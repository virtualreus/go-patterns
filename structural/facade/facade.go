package facade

import "fmt"

// Фасад - паттерн который предоставляет простой интерфейс к сложной системе классов, библиотеке или фреймворку

// AvailabilityChecker Подсистема: Модуль проверки доступности
type AvailabilityChecker struct{}

func (a *AvailabilityChecker) CheckFlight(date string) bool {
	fmt.Printf("Checking flight availability for %s...\n", date)
	// Сложная логика проверки
	return true
}

func (a *AvailabilityChecker) CheckHotel(city string, date string) bool {
	fmt.Printf("Checking hotel availability in %s for %s...\n", city, date)
	// Сложная логика проверки
	return true
}

// BookingSystem Подсистема: Модуль бронирования
type BookingSystem struct{}

func (b *BookingSystem) BookFlight(date string, passenger string) string {
	fmt.Printf("Booking flight for %s on %s...\n", passenger, date)
	return fmt.Sprintf("FLIGHT-%s-%s", passenger, date)
}

func (b *BookingSystem) BookHotel(city string, date string, guest string) string {
	fmt.Printf("Booking hotel in %s for %s on %s...\n", city, guest, date)
	return fmt.Sprintf("HOTEL-%s-%s-%s", city, guest, date)
}

// PaymentProcessor Подсистема: Модуль оплаты
type PaymentProcessor struct{}

func (p *PaymentProcessor) ProcessPayment(amount float64, method string) bool {
	fmt.Printf("Processing payment of $%.2f via %s...\n", amount, method)
	// Сложная логика оплаты
	return true
}

func (p *PaymentProcessor) GenerateInvoice(bookingID string, amount float64) string {
	return fmt.Sprintf("INVOICE for %s: $%.2f", bookingID, amount)
}

// NotificationService Подсистема: Модуль уведомлений
type NotificationService struct{}

func (n *NotificationService) SendEmail(to, subject, body string) {
	fmt.Printf("Sending email to %s: %s\n", to, subject)
	// Сложная логика отправки email
}

func (n *NotificationService) SendSMS(to, message string) {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
	// Сложная логика отправки SMS
}

// Фасад: Упрощенный интерфейс для системы путешествий
type TravelFacade struct {
	availability *AvailabilityChecker
	booking      *BookingSystem
	payment      *PaymentProcessor
	notification *NotificationService
}

func NewTravelFacade() *TravelFacade {
	return &TravelFacade{
		availability: &AvailabilityChecker{},
		booking:      &BookingSystem{},
		payment:      &PaymentProcessor{},
		notification: &NotificationService{},
	}
}

// Упрощенный метод для бронирования всей поездки
func (t *TravelFacade) BookCompleteTrip(
	passenger string,
	email string,
	phone string,
	destination string,
	travelDate string,
	nights int,
	paymentMethod string,
) (*TripDetails, error) {

	fmt.Println("=== Starting trip booking process ===")

	// 1. Проверка доступности
	if !t.availability.CheckFlight(travelDate) {
		return nil, fmt.Errorf("no flights available")
	}

	if !t.availability.CheckHotel(destination, travelDate) {
		return nil, fmt.Errorf("no hotels available")
	}

	// 2. Бронирование
	flightID := t.booking.BookFlight(travelDate, passenger)
	hotelID := t.booking.BookHotel(destination, travelDate, passenger)

	// 3. Расчет стоимости
	totalAmount := t.calculateTotal(500, 200, nights) // цены условные

	// 4. Оплата
	if !t.payment.ProcessPayment(totalAmount, paymentMethod) {
		return nil, fmt.Errorf("payment failed")
	}

	// 5. Генерация документов
	invoice := t.payment.GenerateInvoice(flightID, totalAmount)

	// 6. Отправка уведомлений
	t.notification.SendEmail(
		email,
		"Your trip is confirmed!",
		fmt.Sprintf("Flight: %s\nHotel: %s\nInvoice: %s", flightID, hotelID, invoice),
	)

	t.notification.SendSMS(
		phone,
		fmt.Sprintf("Trip booked! Total: $%.2f", totalAmount),
	)

	fmt.Println("=== Trip booking completed successfully ===")

	return &TripDetails{
		FlightID:    flightID,
		HotelID:     hotelID,
		Invoice:     invoice,
		TotalAmount: totalAmount,
	}, nil
}

func (t *TravelFacade) calculateTotal(flightPrice, hotelPrice float64, nights int) float64 {
	return flightPrice + (hotelPrice * float64(nights))
}

type TripDetails struct {
	FlightID    string
	HotelID     string
	Invoice     string
	TotalAmount float64
}
