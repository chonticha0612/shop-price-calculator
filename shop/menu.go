package shop

import "fmt"

type OrderMenu struct {
	name  string
	price float64
}

func (c *PriceCalculator) AddMenu(name string, price float64) error {
	if _, found := c.findMenu(name); found {
		return fmt.Errorf("menu %q already exists", name)
	}
	if price <= 0 {
		return fmt.Errorf("menu %q price must be greater than 0", name)
	}
	c.menuList = append(c.menuList, OrderMenu{name: name, price: price})
	return nil
}

func (c *PriceCalculator) RemoveMenu(name string) error {
	i, found := c.findMenu(name)
	if !found {
		return fmt.Errorf("menu %q not found", name)
	}
	c.menuList = append(c.menuList[:i], c.menuList[i+1:]...)
	return nil
}

func (c *PriceCalculator) UpdateMenuPrice(name string, price float64) error {
	i, found := c.findMenu(name)
	if !found {
		return fmt.Errorf("menu %q not found", name)
	}
	if price <= 0 {
		return fmt.Errorf("menu %q price must be greater than 0", name)
	}
	c.menuList[i].price = price
	return nil
}

func (c *PriceCalculator) RenameMenu(oldName, newName string) error {
	i, found := c.findMenu(oldName)
	if !found {
		return fmt.Errorf("menu %q not found", oldName)
	}
	if _, exists := c.findMenu(newName); exists {
		return fmt.Errorf("menu %q already exists", newName)
	}
	c.menuList[i].name = newName
	return nil
}

func (c *PriceCalculator) findMenu(name string) (int, bool) {
	for i, menuItem := range c.menuList {
		if menuItem.name == name {
			return i, true
		}
	}
	return 0, false
}

func MenuList() []OrderMenu {
	return []OrderMenu{
		{
			name:  "Red set",
			price: 50.00,
		},
		{
			name:  "Green set",
			price: 40.00,
		},
		{
			name:  "Blue set",
			price: 30.00,
		},
		{
			name:  "Yellow set",
			price: 50.00,
		}, {
			name:  "Pink set",
			price: 80.00,
		}, {
			name:  "Purple set",
			price: 90.00,
		}, {
			name:  "Orange set",
			price: 120.00,
		},
	}
}
