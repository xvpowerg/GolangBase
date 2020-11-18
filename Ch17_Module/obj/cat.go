package obj

import (
	"errors"
)

type Cat struct {
	age  int
	Name string
}

func (c *Cat) SetAge(age int) error {
	if age >= 0 && age <= 100 {
		c.age = age
	} else {
		return errors.New("錯誤的年齡")
	}
	return nil
}
func (c *Cat) GetAge() int {
	return c.age
}
