package entity

import "fmt"

type Auto struct {
	brand   string
	model   string
	price   int
	status  string
	mileage int
}

const (
	inWay     = "in the way"
	inStock   = "in stock"
	sold      = "sold"
	withdrawn = "withdrawn from sale",
)

var statusList = map[string]string{
	inWay:     inWay,
	inStock:   inStock,
	sold:      sold,
	withdrawn: withdrawn,
}

func NewAuto(brand string, model string, mileage int) *Auto {
	return &Auto{brand: brand, model: model, mileage: mileage}
}

func (a *Auto) Mileage() int {
	return a.mileage
}

func (a *Auto) Status() string {
	return a.status
}

func (a *Auto) Model() string {
	return a.model
}

func (a *Auto) Brand() string {
	return a.brand
}

func (a *Auto) Price() int {
	return a.price
}

func (a *Auto) SetPrice(price int) error {
	if price < 0 {
		return fmt.Errorf("price: `%d` is less than zero", price)
	}
	a.price = price

	return nil
}

func (a *Auto) SetStatus(status string) error {
	_, exist := statusList[status]
	if !exist {
		return fmt.Errorf("status: `%s` not found", status)
	}
	a.status = status

	return nil
}
