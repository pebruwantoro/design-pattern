package donut

// Abstract Product
type IDonut interface {
	SetName(name string)
	GetName() string
	SetPrice(price float64)
	GetPrice() float64
}

type Donut struct {
	Name  string
	Price float64
}

func (d *Donut) SetName(name string) {
	d.Name = name
}

func (d *Donut) GetName() string {
	return d.Name
}

func (d *Donut) SetPrice(price float64) {
	d.Price = price
}

func (d *Donut) GetPrice() float64 {
	return d.Price
}
