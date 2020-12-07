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
