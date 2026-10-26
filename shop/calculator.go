// Package shop calculates order totals for a set-menu shop, applying pair and member discounts.
package shop

import (
	"fmt"
	"sort"
)

type PriceCalculator struct {
	menuList []OrderMenu
}

const (
	pairDiscountRate   = 0.05
	memberDiscountRate = 0.10
)

func NewPriceCalculator() *PriceCalculator {
	return &PriceCalculator{menuList: MenuList()}
}

func (c *PriceCalculator) CalculatePrice(order Order) (float64, error) {
	orderMap := order.MapOrder()
	if missingItems := c.checkOrderItemExists(orderMap); len(missingItems) > 0 {
		return 0.0, fmt.Errorf("the following items are not in the menu: %v", missingItems)
	}

	total := c.calculateDiscountPair(orderMap)
	if order.memberCard {
		total -= total * memberDiscountRate
	}
	return total, nil
}

func (c *PriceCalculator) checkOrderItemExists(orderMap map[string]int) []string {
	menuNames := make(map[string]bool, len(c.menuList))
	for _, menuItem := range c.menuList {
		menuNames[menuItem.name] = true
	}

	orderNotExists := []string{}
	for name := range orderMap {
		if !menuNames[name] {
			orderNotExists = append(orderNotExists, name)
		}
	}
	sort.Strings(orderNotExists)
	return orderNotExists
}

func (c *PriceCalculator) calculateDiscountPair(orderMap map[string]int) float64 {
	price := 0.0
	for _, menuItem := range c.menuList {
		quantity := orderMap[menuItem.name]
		pairs := quantity / 2
		singles := quantity % 2
		pairPrice := menuItem.price * 2 * (1 - pairDiscountRate)
		price += float64(pairs)*pairPrice + float64(singles)*menuItem.price
	}
	return price
}
