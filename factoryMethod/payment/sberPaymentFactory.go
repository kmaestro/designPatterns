package payment

import "designPatterns/factoryMethod/factory"

type sberPaymentFactory struct {
}

func NewSberPaymentFactory() *sberPaymentFactory {
	return &sberPaymentFactory{}
}

func (s *sberPaymentFactory) CteatePayment() factory.Payment {
	return NewSberPayment()
}
