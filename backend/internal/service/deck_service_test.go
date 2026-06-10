package service

import "testing"

func TestDeckMergeRequiresTwoDecks(t *testing.T) {
	_, err := (&DeckService{}).Merge("d", nil, "Merged")
	if err == nil {
		t.Fatal("expected error")
	}
}
