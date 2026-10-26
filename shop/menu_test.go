package shop

import (
	"math"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPriceCalculatorAddMenu(t *testing.T) {
	t.Run("shop adds Black set to the menu, customer can order it and gets pair discount, pays 114", func(t *testing.T) {
		calculator := NewPriceCalculator()
		assert.Equal(t, nil, calculator.AddMenu("Black set", 60.0))

		got, err := calculator.CalculatePrice(Order{orders: []OrderItem{{name: "Black set", quantity: 2}}})
		assert.Equal(t, nil, err)
		assert.Equal(t, 114.0, math.Round(got*100)/100) // (60 x 2) - 5%
	})

	t.Run("shop cannot add a set that is already on the menu", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.AddMenu("Red set", 10.0)
		assert.Equal(t, `menu "Red set" already exists`, err.Error())
	})

	t.Run("shop cannot add a set with no price", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.AddMenu("Free set", 0)
		assert.Equal(t, `menu "Free set" price must be greater than 0`, err.Error())
	})
}

func TestPriceCalculatorRemoveMenu(t *testing.T) {
	t.Run("shop removes Red set from the menu, customer can no longer order it", func(t *testing.T) {
		calculator := NewPriceCalculator()
		assert.Equal(t, nil, calculator.RemoveMenu("Red set"))

		_, err := calculator.CalculatePrice(Order{orders: []OrderItem{{name: "Red set", quantity: 1}}})
		assert.Equal(t, "the following items are not in the menu: [Red set]", err.Error())
	})

	t.Run("shop cannot remove a set that is not on the menu", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.RemoveMenu("Black set")
		assert.Equal(t, `menu "Black set" not found`, err.Error())
	})
}

func TestPriceCalculatorUpdateMenuPrice(t *testing.T) {
	t.Run("shop changes Red set price to 60, customer pays the new price", func(t *testing.T) {
		calculator := NewPriceCalculator()
		assert.Equal(t, nil, calculator.UpdateMenuPrice("Red set", 60.0))

		got, err := calculator.CalculatePrice(Order{orders: []OrderItem{{name: "Red set", quantity: 1}}})
		assert.Equal(t, nil, err)
		assert.Equal(t, 60.0, got)
	})

	t.Run("shop cannot change the price of a set that is not on the menu", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.UpdateMenuPrice("Black set", 60.0)
		assert.Equal(t, `menu "Black set" not found`, err.Error())
	})

	t.Run("shop cannot change a set price to 0", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.UpdateMenuPrice("Red set", 0)
		assert.Equal(t, `menu "Red set" price must be greater than 0`, err.Error())
	})
}

func TestPriceCalculatorRenameMenu(t *testing.T) {
	t.Run("shop renames Red set to Crimson set, customer orders by the new name at the same price", func(t *testing.T) {
		calculator := NewPriceCalculator()
		assert.Equal(t, nil, calculator.RenameMenu("Red set", "Crimson set"))

		got, err := calculator.CalculatePrice(Order{orders: []OrderItem{{name: "Crimson set", quantity: 1}}})
		assert.Equal(t, nil, err)
		assert.Equal(t, 50.0, got)

		_, err = calculator.CalculatePrice(Order{orders: []OrderItem{{name: "Red set", quantity: 1}}})
		assert.Equal(t, "the following items are not in the menu: [Red set]", err.Error())
	})

	t.Run("shop cannot rename a set that is not on the menu", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.RenameMenu("Black set", "White set")
		assert.Equal(t, `menu "Black set" not found`, err.Error())
	})

	t.Run("shop cannot rename a set to a name that is already on the menu", func(t *testing.T) {
		calculator := NewPriceCalculator()
		err := calculator.RenameMenu("Red set", "Green set")
		assert.Equal(t, `menu "Green set" already exists`, err.Error())
	})
}
