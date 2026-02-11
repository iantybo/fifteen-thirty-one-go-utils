package cards

import "testing"

func TestParseAndString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"AS", "AS"},
		{"10H", "10H"},
		{"KD", "KD"},
		{"2C", "2C"},
		{"JH", "JH"},
		{"QS", "QS"},
	}
	for _, tt := range tests {
		c, err := Parse(tt.input)
		if err != nil {
			t.Errorf("Parse(%q) error: %v", tt.input, err)
			continue
		}
		if got := c.String(); got != tt.want {
			t.Errorf("Parse(%q).String() = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestValue15(t *testing.T) {
	tests := []struct {
		card string
		want int
	}{
		{"AS", 1},
		{"5H", 5},
		{"10D", 10},
		{"JC", 10},
		{"QH", 10},
		{"KS", 10},
	}
	for _, tt := range tests {
		c, _ := Parse(tt.card)
		if got := c.Value15(); got != tt.want {
			t.Errorf("Card(%q).Value15() = %d, want %d", tt.card, got, tt.want)
		}
	}
}

func TestFullDeck(t *testing.T) {
	deck := FullDeck()
	if len(deck) != 52 {
		t.Errorf("FullDeck() has %d cards, want 52", len(deck))
	}
	seen := make(map[string]bool)
	for _, c := range deck {
		key := c.String()
		if seen[key] {
			t.Errorf("duplicate card in deck: %s", key)
		}
		seen[key] = true
	}
}

func TestCombinations(t *testing.T) {
	deck := FullDeck()[:6]
	combos := Combinations(deck, 2)
	// 6 choose 2 = 15
	if len(combos) != 15 {
		t.Errorf("Combinations(6, 2) = %d combos, want 15", len(combos))
	}
}

func TestRemove(t *testing.T) {
	deck := FullDeck()[:6]
	toRemove := deck[:2]
	result := Remove(deck, toRemove)
	if len(result) != 4 {
		t.Errorf("Remove: got %d cards, want 4", len(result))
	}
}
