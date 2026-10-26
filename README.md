# Set Menu Price Calculator

## Where to start

The entry point is `CalculatePrice` in [shop/calculator.go](shop/calculator.go).
Read it first, then the tests in [shop/calculator_test.go](shop/calculator_test.go) show each pricing rule as an example.

## Running the tests

Requires Go 1.24 or later.

```sh
go test ./...
```

## Menu

| Set        | Price |
|------------|------:|
| Red set    |    50 |
| Green set  |    40 |
| Blue set   |    30 |
| Yellow set |    50 |
| Pink set   |    80 |
| Purple set |    90 |
| Orange set |   120 |

## Pricing rules

1. **Pair discount (5%)**: every 2 of the same set get 5% off that pair. A leftover single pays full price.
   - 3 × Red set = (50 × 2 × 0.95) + 50 = **145**
2. **Same set on separate order lines is combined** before pairs are counted.
   - Orange set, Green set, Orange set = (120 × 2 × 0.95) + 40 = **268**
3. **Member discount (10%)**: a customer with a member card gets 10% off the total, applied after the pair discount.
   - Member, 2 × Red set + 1 × Blue set = (95 + 30) × 0.9 = **112.5**


## Menu management

The shop can change the menu on a `PriceCalculator`:

| Method            | Fails when                                         |
|-------------------|----------------------------------------------------|
| `AddMenu`         | the set already exists, or the price is 0 or less  |
| `RemoveMenu`      | the set is not on the menu                         |
| `UpdateMenuPrice` | the set is not on the menu, or the price is 0 or less |
| `RenameMenu`      | the old name is not on the menu, or the new name already exists |

## Project layout

```
shop/
├── calculator.go        Price calculation and discounts (start here)
├── calculator_test.go   Tests for pricing rules
├── menu.go              Default menu and menu management
├── menu_test.go         Tests for menu management
└── order.go             Order and OrderItem; merges repeated items
```
