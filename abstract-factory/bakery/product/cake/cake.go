package cake

// Abstract Product
type ICake interface {
	SetName(name string)
	GetName() string
	SetPrice(price float64)
	GetPrice() float64
}

type Cake struct {
	Name  string
	Price float64
}

func (c *Cake) SetName(name string) {
	c.Name = name
}

func (c *Cake) GetName() string {
	return c.Name
}

func (c *Cake) SetPrice(price float64) {
	c.Price = price
}

func (c *Cake) GetPrice() float64 {
	return c.Price
}
