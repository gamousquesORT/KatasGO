//http://www.tddbuddy.com/katas/Heavy%20Metal%20Bake%20Sale.pdf
/*
There are four items they would like to sell on this sale with specific prices and quantities of each:
Brownie - 65c (48)
Muffin - $1.00 (36)
Cake Pop - $1.35 (24)
Water - $1.50 (30)

The application must calculate the smallest amount of change to give a person if they overpay.
If you do not have stock of an item, you cannot make the sale.

Purchases are input as comma delimited string. The first letter of each item is its input code, B for Brownie, M for
Muffin, C for Cake Pop and W for Water.
The application then responds with a total if there are enough of each item.
Otherwise responds with “Not enough stock.”
If the total appears the user, then inputs the amount paid. If the amount is equal to or greater than the amount due
the application, then figures out the amount of change to give.
If the amount is less than the amount due it responds with “Not enough money.”
Examples
Items to Purchase > B, C, W / Total > $3.50 /Amount Paid > $4.00 Change > $0.50

Items to Purchase > B
Total > $0.65
Amount Paid > $0.75
Change > $0.10
Items to Purchase > C,M
Total > $2.35
Amount Paid > $2.00
Change > Not enough money
Items to Purchase > W
Total > Not enough stock
Hint
Looks to use Stubs or Mocks or both when solving input and output problems.

*/
package BakeSale

import (

	"testing"
)


func TestBakeSale(t *testing.T) {
	t.Run("Should_Total$3.50_for_BCW",func(t *testing.T) {
		got, change := BakeSale("B,C,W", 4.00)
		wantTotal := "$3.50"
		wantChange := 0.50

		if got != wantTotal {
			t.Errorf("got %s want %s",got,wantTotal)

		}
		if change != wantChange {
			t.Errorf("got %f want %f",change,wantChange)
		}
	})

	t.Run("Should_Total$0.65_for_B",func(t *testing.T) {
		got, change := BakeSale("B", 0.75)
		wantTotal := "$0.65"
		wantChange := 0.10

		if got != wantTotal {
			t.Errorf("got %s want %s",got,wantTotal)

		}
		if change != wantChange {
			t.Errorf("got %f want %f",change,wantChange)
		}
	})


}