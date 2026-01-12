package errgroup_for_shards

import (
	"context"
	"sync"
)

// Паттерн для паралеллельного выполнения задач с автоматической отменой всех операций при первой ошибке

type Group struct {
	cancel  func(error) // для отмены контекста при ошибке
	wg      sync.WaitGroup
	errOnce sync.Once // гарантирует сохранение только первой ошибки
	err     error     // хранит первую ошибку
}

func NewErrGroup(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancelCause(ctx)
	return &Group{cancel: cancel}, ctx
}

type Database interface {
	Query(query string) (string, error)
}

func DistQueryWithErrGroup(shards []Database, query string) ([]string, error) {
	g, ctx := NewErrGroup(context.Background())
	responseCh := make(chan string)

	for _, shard := range shards {
		g.wg.Add(1)
		go func() {
			defer g.wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			response, err := shard.Query(query)
			if err != nil {
				g.cancel(err)
				return
			}
			select {
			case responseCh <- response:
				return
			}
		}()
	}
	go func() {
		g.wg.Wait()
		close(responseCh)
	}()
	var responses []string
	for response := range responseCh {
		responses = append(responses, response)
	}
	return responses, g.err
}
