package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_UpdateQuality(t *testing.T) {
	t.Run("given a regular product sold after sellin update quality twice as fast", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"OldProd", 0, 10},
		}

		gildedrose.UpdateQuality(items)
		var expected = 8

		if items[0].Quality != expected {
			t.Errorf("Name: Expected %d but got %d ", expected, items[0].Quality)
		}

	})
	t.Run("given a regular product quality should never be negative", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"NegQuality", 0, 0},
		}

		gildedrose.UpdateQuality(items)
		var expected = 0

		if items[0].Quality != expected {
			t.Errorf("Name: Expected %d but got %d ", expected, items[0].Quality)
		}

	})
	t.Run("given a regular product quality and sellin should decrease", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"RegularDecrease", 1, 1},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 0
		var expectedQuality = 0

		if items[0].Quality != expectedQuality {
			t.Errorf("Name: Expected %d but got %d ", expectedQuality, items[0].Quality)
		}
		if items[0].SellIn != expectedSellin {
			t.Errorf("Name: Expected %d but got %d ", expectedSellin, items[0].SellIn)
		}
	})
	t.Run("given a Brie quality increases and sellin decrease", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Aged Brie", 2, 1},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 1
		var expectedQuality = 2

		if items[0].Quality != expectedQuality {
			t.Errorf("Name: Expected %d but got %d ", expectedQuality, items[0].Quality)
		}
		if items[0].SellIn != expectedSellin {
			t.Errorf("Name: Expected %d but got %d ", expectedSellin, items[0].SellIn)
		}
	})

	t.Run("given a product quality should ve more than 50", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Aged Brie", 1, 51},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 0
		var expectedQuality = 50

		if items[0].Quality != expectedQuality {
			t.Errorf("Name: Expected %d but got %d ", expectedQuality, items[0].Quality)
		}
		if items[0].SellIn != expectedSellin {
			t.Errorf("Name: Expected %d but got %d ", expectedSellin, items[0].SellIn)
		}
	})

	t.Run("given a Sulfuras never sold and quality intact", func(t *testing.T) {

		var items = []*gildedrose.Item{
			{"Sulfuras", 1, 15},
		}

		gildedrose.UpdateQuality(items)
		var expectedSellin = 1
		var expectedQuality = 15

		if items[0].Quality != expectedQuality {
			t.Errorf("Name: Expected %d but got %d ", expectedQuality, items[0].Quality)
		}
		if items[0].SellIn != expectedSellin {
			t.Errorf("Name: Expected %d but got %d ", expectedSellin, items[0].SellIn)
		}
	})
}
