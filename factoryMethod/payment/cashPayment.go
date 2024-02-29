package payment

import (
	"designPatterns/factoryMethod/payment/dto"
	"fmt"
)

type cashPayment struct {
}

func NewCashPayment() *cashPayment {
	return &cashPayment{}
}

func (c *cashPayment) Pay(order dto.Order) {

	fmt.Printf("Cash pay success %d rub\n", order.Sum)
}
