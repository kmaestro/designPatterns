package payment

import "designPatterns/factoryMethod/factory"

type cashPaymentFactory struct {

}

func NewCashPaymentFactory() (* cashPaymentFactory) {
	return &cashPaymentFactory{}
}

func (c *cashPaymentFactory) CteatePayment() factory.Payment {
	return NewCashPayment()
}