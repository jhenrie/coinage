package coinage

import (
	"testing"
)

func TestCoinPurse_Add(t *testing.T) {
	purse := NewEmptyPurse("test")
	exchange := NewCoinPurse("exchange", 100)

	purse.Add(*exchange)

	if purse.Count != 100 {
		t.Errorf("Add failed purse %s, %d - add %s, %d", purse.Owner, purse.Count, exchange.Owner, exchange.Count)
	}
}

func TestCoinPurse_Remove(t *testing.T) {
	purse := NewCoinPurse("test", 111)
	exchange := NewCoinPurse("remove", 100)

	if err := purse.Remove(*exchange); err != nil {
		t.Errorf("following error was returned")
	}
	if purse.Count != 11 {
		t.Errorf("Remove failed to remove.")
	}
}

func TestCoinPurse_RemoveOverdraft(t *testing.T) {
	purse := NewCoinPurse("test", 111)
	exchange := NewCoinPurse("remove", 1000)

	if err := purse.Remove(*exchange); err == nil {
		t.Errorf("no error was returned to show overdraft")
	}
}
