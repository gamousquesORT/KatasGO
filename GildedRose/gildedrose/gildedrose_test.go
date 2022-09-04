package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestUpdateQuality(t *testing.T) {

	checkValues := func(t testing.TB, item *gildedrose.Item, expectedSellin int, expectedQuality int) {
		t.Helper()
		if item.Quality != expectedQuality {
			t.Errorf("Name: Expected %d but got %d ", expectedQuality, item.Quality)
		}
		if item.SellIn != expectedSellin {
			t.Errorf("Name: Expected %d but got %d ", expectedSellin, item.SellIn)
		}
	}

	t.Run("given a regular product quality and sellin should decrease", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"RegularDecrease", 1, 1},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 0
		var expectedQuality = 0
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a regular product quality should never be negative", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"NegQuality", 10, 0},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 9
		var expectedQuality = 0
		checkValues(t, items[0], expectedSellin, expectedQuality)
	})

	t.Run("given a regular product sold after sellin update quality twice as fast", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"OldProd", 0, 10},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = -1
		var expectedQuality = 8

		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a Brie quality increases and sellin decrease", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Aged Brie", 10, 10},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 9
		var expectedQuality = 11
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a product quality should be more than 50", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Aged Brie", 10, 50},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 9
		var expectedQuality = 50
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a Sulfuras never sold and quality intact", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Sulfuras, Hand of Ragnaros", 10, 40},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 10
		var expectedQuality = 40
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a Backstage passes to a TAFKAL80ETC concert after 10 days", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Backstage passes to a TAFKAL80ETC concert", 10, 40},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 9
		var expectedQuality = 42
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a Backstage passes to a TAFKAL80ETC concert after 5 days", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Backstage passes to a TAFKAL80ETC concert", 5, 40},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 4
		var expectedQuality = 43
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})

	t.Run("given a Backstage passes to a TAFKAL80ETC concert after 0 days", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Backstage passes to a TAFKAL80ETC concert", 0, 40},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = -1
		var expectedQuality = 0
		checkValues(t, items[0], expectedSellin, expectedQuality)

	})
}
