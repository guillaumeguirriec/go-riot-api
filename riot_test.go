package riot

import (
	"testing"
)

func TestNewApiKeyEmpty(t *testing.T) {
	riot, err := New("", EuropeRegion, false)

	if riot != nil || err == nil {
		t.Fatalf(`New("", EuropeRegion, false) = %v+, %v, want nil, error`, riot, err)
	}
}

func TestNewRegionEmpty(t *testing.T) {
	riot, err := New("<API key>", "", false)

	if riot != nil || err == nil {
		t.Fatalf(`New("<API key>", "", false) = %v+, %v, want nil, error`, riot, err)
	}
}

func TestNewWrongRegion(t *testing.T) {
	riot, err := New("<API key>", "WrongRegion", false)

	if riot != nil || err == nil {
		t.Fatalf(`New("<API key>", "WrongRegion", false) = %v+, %v, want nil, error`, riot, err)
	}
}
