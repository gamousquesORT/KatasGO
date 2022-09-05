package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const maxQuality = 50

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if items[i].Name != "Aged Brie" && items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					NormalQualityDecrease(items, i)
				}
			}
		} else {
			if items[i].Quality < maxQuality {
				IncrementQuality(items, i)
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						NormalQualityIncrease(items, i)
					}
				}
				if items[i].SellIn < 6 {
					NormalQualityIncrease(items, i)
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							NormalQualityDecrease(items, i)
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				NormalQualityIncrease(items, i)
			}
		}
	}

}

func NormalQualityIncrease(items []*Item, i int) {
	if items[i].Quality < maxQuality {
		IncrementQuality(items, i)
	}
}

func IncrementQuality(items []*Item, i int) {
	items[i].Quality = items[i].Quality + 1
}

func NormalQualityDecrease(items []*Item, i int) {
	items[i].Quality = items[i].Quality - 1
}
