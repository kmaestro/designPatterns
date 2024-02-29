package factory

import "designPatterns/factoryMethod/payment/dto"

type Payment interface {
	Pay(dto.Order)
}
