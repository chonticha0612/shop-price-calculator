package shop

import (
	"math"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCalculatePrice(t *testing.T) {
	tests := []struct {
		name  string
		order Order
		want  float64
	}{
		{
			name:  "customer buys 1 Red set, pays full price 50",
			order: Order{orders: []OrderItem{{name: "Red set", quantity: 1}}},
			want:  50.0,
		},
		{
			name:  "customer buys 2 Red sets as a pair, gets 5% off and pays 95",
			order: Order{orders: []OrderItem{{name: "Red set", quantity: 2}}},
			want:  95.0,
		},
		{
			name:  "customer buys 3 Red sets, 1 pair gets 5% off and the leftover pays full price, total 145",
			order: Order{orders: []OrderItem{{name: "Red set", quantity: 3}}},
			want:  145.0,
		},
		{
			name:  "customer buys 4 Red sets, gets 2 pairs with 5% off each, total 190",
			order: Order{orders: []OrderItem{{name: "Red set", quantity: 4}}},
			want:  190.0,
		},
		{
			name: "customer orders Orange set twice in separate lines, still counted as 1 pair, total 268",
			order: Order{orders: []OrderItem{
				{name: "Orange set", quantity: 1},
				{name: "Green set", quantity: 1},
				{name: "Orange set", quantity: 1},
			}},
			want: 268.0,
		},
		{
			name: "customer buys different sets with no pair, pays full price for all, total 120",
			order: Order{orders: []OrderItem{
				{name: "Red set", quantity: 1},
				{name: "Green set", quantity: 1},
				{name: "Blue set", quantity: 1},
			}},
			want: 120.0,
		},
		{
			name: "member buys 2 Red sets and 1 Blue set, gets pair discount then 10% off total, pays 112.5",
			order: Order{
				orders:     []OrderItem{{name: "Red set", quantity: 2}, {name: "Blue set", quantity: 1}},
				memberCard: true,
			},
			want: 112.5,
		},
		{
			name: "member buys 1 Red set with no pair, still gets 10% off total, pays 45",
			order: Order{
				orders:     []OrderItem{{name: "Red set", quantity: 1}},
				memberCard: true,
			},
			want: 45.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPriceCalculator().CalculatePrice(tt.order)
			assert.Equal(t, nil, err)
			assert.Equal(t, tt.want, math.Round(got*100)/100)
		})
	}
}

func TestCalculatePriceMenuNotFound(t *testing.T) {
	t.Run("customer orders sets that are not on the menu, order is rejected and tells which sets are unavailable", func(t *testing.T) {
		order := Order{orders: []OrderItem{
			{name: "Red set", quantity: 1},
			{name: "Black set", quantity: 1},
			{name: "White set", quantity: 2},
		}}

		got, err := NewPriceCalculator().CalculatePrice(order)
		assert.Equal(t, 0.0, got)
		assert.NotEqual(t, nil, err)
		assert.Equal(t, "the following items are not in the menu: [Black set White set]", err.Error())
	})
}
