package coinage

import (
	"testing"
)

func TestParseCoinage_simple(t *testing.T) {
	swap := "1c"

	purse, err := ParseCoinage(swap)
	if err != nil {
		t.Errorf("recieved the following error %s", err)
	}

	if purse.Count != 1 {
		t.Errorf("expected 1 got %d", purse.Count)
	}
}

func TestParseCoinage_complex(t *testing.T) {
	swap := "1gp1SP1c"

	purse, err := ParseCoinage(swap)
	if err != nil {
		t.Errorf("recieved the following error %s", err)
	}

	if purse.Count != 111 {
		t.Errorf("expected 1 got %d", purse.Count)
	}
}

func TestParseCoinage_BadChar(t *testing.T) {
	swap := "11r"

	_, err := ParseCoinage(swap)
	if err == nil {
		t.Errorf("did not invalidate swap string")
	}
}

func TestParseCoinage_NoChar(t *testing.T) {
	swap := "11"

	_, err := ParseCoinage(swap)
	if err == nil {
		t.Errorf("did not invalidate swap string")
	}
}

func TestParseCoinage_NoNum(t *testing.T) {
	swap := "c"

	_, err := ParseCoinage(swap)
	if err == nil {
		t.Errorf("did not invalidate swap string")
	}
}
