package cards

// Combinations returns all k-element subsets of the given cards.
func Combinations(cards []Card, k int) [][]Card {
	if k <= 0 || k > len(cards) {
		return nil
	}
	var result [][]Card
	combo := make([]Card, k)
	var generate func(start, depth int)
	generate = func(start, depth int) {
		if depth == k {
			tmp := make([]Card, k)
			copy(tmp, combo)
			result = append(result, tmp)
			return
		}
		for i := start; i <= len(cards)-(k-depth); i++ {
			combo[depth] = cards[i]
			generate(i+1, depth+1)
		}
	}
	generate(0, 0)
	return result
}

// Remove returns a new slice with the specified cards removed from the source.
// Each card in toRemove is removed at most once.
func Remove(source []Card, toRemove []Card) []Card {
	removed := make(map[int]bool)
	for _, tr := range toRemove {
		for i, c := range source {
			if !removed[i] && c.Rank == tr.Rank && c.Suit == tr.Suit {
				removed[i] = true
				break
			}
		}
	}
	result := make([]Card, 0, len(source)-len(removed))
	for i, c := range source {
		if !removed[i] {
			result = append(result, c)
		}
	}
	return result
}
