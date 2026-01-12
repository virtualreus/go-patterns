package main

import (
	"fmt"
	"time"
)

// Паттерн для работы с кластером БД с синхронной репликацией, дабы снизить нагрузку при чтении данных.
// При необходимости прочитать данные - шлем запрос во все реплики, и как только пришел ответ, его обрабатываем, забиваем на остальные

type Database interface {
	Query(query string) string
}

func DistributedQuery(replicas []Database, query string) string {
	responseCh := make(chan string, 1)
	for _, replica := range replicas {
		go func() {
			select {
			case responseCh <- replica.Query(query):
			default:
				return
			}
		}()
	}

	return <-responseCh
}

// Конкретные типы базовые для ебланчиков
type PgSQLDatabase struct {
	Addr string
}

func NewPgSQLDatabase(addr string) *PgSQLDatabase {
	return &PgSQLDatabase{Addr: addr}
}

func (p *PgSQLDatabase) Query(string) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("Success query from addr %s", p.Addr)
}

func main() {
	start := time.Now()
	replicas := []Database{
		NewPgSQLDatabase("127.0.0.1:5432"),
		NewPgSQLDatabase("127.0.0.2:5432"),
		NewPgSQLDatabase("127.0.0.3:5432"),
	}

	response := DistributedQuery(replicas, "SELECT * FROM users")
	fmt.Printf("Ответ: %s\nВремя: %v\n", response, time.Since(start))
}
