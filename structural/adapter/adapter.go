package adapter

import "fmt"

// Паттерн позволяет обьектам с несовместимыми интерфейсами работать вместе
// Используется для работы с либами или легаси кодом

// PaymentProcess Целевой интерфейс который ожидает клиент
type PaymentProcess interface {
	ProcessPayment(amount float64) string
}

// LegacyPaymentSystem Легаси система со старым интерфейсом
type LegacyPaymentSystem struct {
}

func (l *LegacyPaymentSystem) ProcessPayment(amount float64, currency string) string {
	return fmt.Sprintf("Paid %.2f %s via legacy system", amount, currency)
}

type LegacyPaymentAdapter struct {
	legacyPaymentSystem *LegacyPaymentSystem
}

func NewLegacyPaymentAdapter() *LegacyPaymentAdapter {
	return &LegacyPaymentAdapter{
		legacyPaymentSystem: &LegacyPaymentSystem{},
	}
}

func (a *LegacyPaymentAdapter) ProcessPayment(amount float64) string {
	return a.legacyPaymentSystem.ProcessPayment(amount, "USD")
}
