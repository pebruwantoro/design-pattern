package croissant

// Abstract Product
type ICroissant interface {
	SetName(name string)
	GetName() string
	SetPrice(price float64)
	GetPrice() float64
}

type Croissant struct {
	Name  string
	Price float64
}

func (cr *Croissant) SetName(name string) {
	cr.Name = name
}

func (cr *Croissant) GetName() string {
	return cr.Name
}

func (cr *Croissant) SetPrice(price float64) {
	cr.Price = price
}

func (cr *Croissant) GetPrice() float64 {
	return cr.Price
}
