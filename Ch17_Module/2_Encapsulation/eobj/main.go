package eobj

type Fruit struct {
	name  string
	price int
}

func (f *Fruit) SetName(name string) {
	f.name = name
}
func (f *Fruit) SetPrice(price int) {
	f.price = price
}
func (f *Fruit) GetName() string {
	return f.name
}
func (f *Fruit) GetPrice() int {
	return f.price
}
