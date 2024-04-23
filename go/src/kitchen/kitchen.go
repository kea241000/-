package kitchen

import "fmt"

// Order структура представляет заказ на кухне
type Order struct {
    ID         int
    Status     string
    Courier    string
}

// Kitchen структура представляет кухню
type Kitchen struct{}

// NewKitchen создает новую кухню
func NewKitchen() *Kitchen {
    return &Kitchen{}
}

// ViewOrders отображает список заказов, поступивших на приготовление
func (k *Kitchen) ViewOrders(orders []Order) {
    for _, order := range orders {
        fmt.Printf("Order ID: %d, Status: %s, Courier: %s\n", order.ID, order.Status, order.Courier)
    }
}

// StartPreparation подтверждает начало приготовления заказа
func (k *Kitchen) StartPreparation(orderID int) {
    // Реализация подтверждения начала приготовления заказа
}

// ChangeStatus изменяет статус заказа на "ожидает курьера"
func (k *Kitchen) ChangeStatus(orderID int) {
    // Реализация изменения статуса заказа
}

// ConfirmHandover подтверждает передачу заказа курьеру
func (k *Kitchen) ConfirmHandover(orderID int, courierName string) {
    // Реализация подтверждения передачи заказа курьеру
}
