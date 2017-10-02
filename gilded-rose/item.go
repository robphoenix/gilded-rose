package gildedrose

import "fmt"

// Updater is the interface that wraps the Update method.
type Updater interface {
	// Update updates the quality and days remaining.
	Update()
}

// Item describes an item sold by the Gilded Rose Inn.
type Item struct {
	name    string
	days    int
	quality int
}

// DaysRemaining returns the days left to sell the item in.
func (i *Item) DaysRemaining() int {
	return i.days
}

// Quality returns the quality of the item.
func (i *Item) Quality() int {
	return i.quality
}

// Update updates the quality and days remaining.
func (i Item) Update() {
	// noop
}

// String outputs a string representation of an Item.
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.days, i.quality)
}

// Normal is a normal item.
type Normal struct {
	*Item
}

// Update updates the quality and days remaining.
func (i Normal) Update() {
	i.days--
	if i.quality == 0 {
		return
	}
	i.quality--
	if i.days <= 0 {
		i.quality--
	}
}

// Brie is an "Aged Brie".
type Brie struct {
	*Item
}

// Update updates the quality and days remaining.
func (i Brie) Update() {
	i.days--
	if i.quality >= 50 {
		return
	}
	if i.days <= 0 {
		i.quality++
	}
	if i.quality < 50 {
		i.quality++
	}
}

// Backstage is "Backstage passes to a TAFKAL80ETC concert".
type Backstage struct {
	*Item
}

// Update updates the quality and days remaining.
func (i Backstage) Update() {
	i.days--
	if i.quality >= 50 {
		return
	}
	if i.days < 0 {
		i.quality = 0
		return
	}
	i.quality++
	if i.days < 10 {
		i.quality++
	}
	if i.days < 5 {
		i.quality++
	}
}

// New creates a new Item.
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		days:    days,
		quality: quality,
	}
}

// UpdateQuality ages the item by a day, and updates the quality of the item.
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		var t Updater
		switch item.name {
		case "normal":
			t = Normal{item}
		case "Aged Brie":
			t = Brie{item}
		case "Backstage passes to a TAFKAL80ETC concert":
			t = Backstage{item}
		default:
			t = item
		}
		t.Update()
	}
}
