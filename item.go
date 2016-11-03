package main

// NewItemFromASCII returns a new Item with values initialized from a
// given ascii code
func NewItemFromASCII(val int32) Item {
	// Scores
	if val >= 49 && val <= 57 {
		return Item{
			Score: int(val % 48),
		}
	}

	if val == 112 {
		return Item{
			IsPickaxe: true,
		}
	}

	if val == 115 {
		return Item{
			IsStart: true,
		}
	}
	// Everything else counts as an empty square.
	return Item{}
}

type Item struct {
	Point

	IsPickaxe bool
	IsStart   bool
	Score     int
}

func (i Item) IsEmpty() bool {
	return !i.IsPickaxe && i.Score == 0
}
