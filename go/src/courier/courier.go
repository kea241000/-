package courier

import "fmt"

// Order структура представляет заказ курьера
type Order struct {
    ID     int
    Status string
}

// Courier структура представляет курьера
type Courier struct{}

// NewCourier создает нового курьера
func NewCourier() *Courier {
    return &Courier{}
}

// ViewOrders отображает список заказов, которые находятся на кухне
func (c *Courier) ViewOrders(orders []Order) {
    for _, order := range orders {
        fmt.Printf("Order ID: %d, Status: %s\n", order.ID, order.Status)
    }
}

// PickupOrder выбирает заказ на доставку
func (c *Courier) PickupOrder(orderID int) {
    // Реализация выбора заказа на доставку
}

// ConfirmDelivery подтверждает доставку заказа
func (c *Courier) ConfirmDelivery(orderID int) {
    // Реализация подтверждения доставки заказа
}

// ReportProblem сообщает менеджеру о проблеме с заказом
func (c *Courier) ReportProblem(orderID int, problem string) {
    // Реализация сообщения о проблеме с заказом
}
