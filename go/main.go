package main

import (
    "fmt"
    "net/http"

    "goi/src/courier"
    "goi/src/customer"
    "goi/src/kitchen"
    "goi/src/manager"
)


func main() {

    // Обработчик для статических файлов (HTML и CSS)
    http.Handle("/", http.FileServer(http.Dir("public")))

    // Запуск HTTP-сервера на порту 8080
    http.ListenAndServe(":8080", nil)




    // Создаем экземпляры менеджера, покупателя, курьера и кухни
    manager := manager.NewManager()
    customer := customer.NewCustomer()
    courier := courier.NewCourier()
    kitchen := kitchen.NewKitchen()

    // Пример использования функциональности менеджера
    orders := []manager.Order{
        {ID: 1, Items: []string{"Pizza", "Salad"}, Status: "Pending"},
        {ID: 2, Items: []string{"Burger", "Fries"}, Status: "In Progress"},
    }
    manager.ViewOrders(orders)

    // Пример использования функциональности покупателя
    items := []customer.Item{
        {Name: "Pizza", Ingredients: map[string]string{"Cheese": "Yes", "Pepperoni": "Yes"}},
        {Name: "Salad", Ingredients: map[string]string{"Lettuce": "Yes", "Tomato": "Yes"}},
    }
    order := customer.Order{
        Items:           items,
        DeliveryAddress: "123 Main St",
        DeliveryTime:    "ASAP",
        Comment:         "Extra napkins please",
        Status:          "Pending",
    }
    customer.PlaceOrder(order)

    // Пример использования функциональности курьера
    courierOrders := []courier.Order{
        {ID: 1, Status: "Pending"},
        {ID: 2, Status: "In Progress"},
    }
    courier.ViewOrders(courierOrders)

    // Пример использования функциональности кухни
    kitchenOrders := []kitchen.Order{
        {ID: 1, Status: "Pending", Courier: "John"},
        {ID: 2, Status: "In Progress", Courier: "Alice"},
    }
    kitchen.ViewOrders(kitchenOrders)
}
