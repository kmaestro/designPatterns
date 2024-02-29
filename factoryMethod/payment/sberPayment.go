package payment

import (
	"designPatterns/factoryMethod/payment/dto"
	"fmt"
)

type sberPayment struct {
}

func NewSberPayment() *sberPayment {
	return &sberPayment{}
}

func (s *sberPayment) Pay(order dto.Order) {
	fmt.Printf("Sber pay success %d rub\n", order.Sum)
}
