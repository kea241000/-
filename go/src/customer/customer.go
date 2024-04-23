package customer

import "fmt"

// Item структура представляет товар в заказе
type Item struct {
    Name       string
    Ingredients map[string]string
}

// Order структура представляет заказ покупателя
type Order struct {
    Items      []Item
    DeliveryAddress string
    DeliveryTime string
    Comment string
    Status     string
}

// Customer структура представляет покупателя
type Customer struct{}

// NewCustomer создает нового покупателя
func NewCustomer() *Customer {
    return &Customer{}
}

// ViewCart отображает содержимое корзины покупателя
func (c *Customer) ViewCart(cart []Item) {
    for _, item := range cart {
        fmt.Printf("Item: %s, Ingredients: %v\n", item.Name, item.Ingredients)
    }
}

// PlaceOrder размещает новый заказ
func (c *Customer) PlaceOrder(order Order) {
    // Реализация размещения заказа
}

// ModifyOrder модифицирует существующий заказ
func (c *Customer) ModifyOrder(orderID int, newItems []Item, newAddress string, newTime string, newComment string) {
    // Реализация модификации заказа
}

// TrackOrder отслеживает статус заказа
func (c *Customer) TrackOrder(orderID int) string {
    // Реализация отслеживания статуса заказа
}
