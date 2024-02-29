package main

import (
	"designPatterns/factoryMethod/factory"
	"designPatterns/factoryMethod/payment"
	"designPatterns/factoryMethod/payment/dto"
)

type orderData struct {
	order       dto.Order
	paymentType string
}

var orders []orderData

func main() {
	orders = append(orders, orderData{
		order:       dto.Order{Sum: 500},
		paymentType: "cash",
	})
	orders = append(orders, orderData{
		order:       dto.Order{Sum: 3000},
		paymentType: "sber",
	})
	for _, order := range orders {
		getPaymentFactory(order.paymentType).CteatePayment().Pay(order.order)
	}

}

func getPaymentFactory(paymentType string) factory.PaymentFactory {
	switch paymentType {
	case "cash":
		return payment.NewCashPaymentFactory()
	case "sber":
		return payment.NewSberPaymentFactory()
	default:
		panic("error")
	}
}
