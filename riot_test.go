package riot

import (
	"testing"
)

func TestNewApiKeyEmpty(t *testing.T) {
	riot, err := New("", EuropeRegion, HeaderParam, false)

	if riot != nil || err == nil {
		t.Fatalf(`New("", EuropeRegion, HeaderParam, false) = %v, %v, want nil, error`, riot, err)
	}
}

func TestNewRegionEmpty(t *testing.T) {
	riot, err := New("<API key>", "", HeaderParam, false)

	if riot != nil || err == nil {
		t.Fatalf(`New("<API key>", "", HeaderParam, false) = %v, %v, want nil, error`, riot, err)
	}
}

func TestNewWrongRegion(t *testing.T) {
	riot, err := New("<API key>", "WrongRegion", HeaderParam, false)

	if riot != nil || err == nil {
		t.Fatalf(`New("<API key>", "WrongRegion", HeaderParam, false) = %v, %v, want nil, error`, riot, err)
	}
}

func TestNewIncludeApiKeyAsEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, "", false)

	if riot != nil || err == nil {
		t.Fatalf(`New("<API key>", EuropeRegion, "", false) = %v, %v, want nil, error`, riot, err)
	}
}

func TestNewWrongIncludeApiKeyAs(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, "WrongIncludeApiKeyAs", false)

	if riot != nil || err == nil {
		t.Fatalf(`New("<API key>", EuropeRegion, "WrongIncludeApiKeyAs", false) = %v, %v, want nil, error`, riot, err)
	}
}

func TestGetAccountByRiotIdGameNameEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	accountDto, err := riot.GetAccountByRiotId("", "<tagLine>")

	if (AccountDto{} != accountDto) || err == nil {
		t.Fatalf(`GetAccountByRiotId("", "<tagLine>") = %v, %v, want nil, error`, accountDto, err)
	}
}

func TestGetAccountByRiotIdTagLineEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	accountDto, err := riot.GetAccountByRiotId("<gameName>", "")

	if (AccountDto{} != accountDto) || err == nil {
		t.Fatalf(`GetAccountByRiotId("<gameName>", "") = %v, %v, want nil, error`, accountDto, err)
	}
}

func TestGetAccountByRiotIdAllParametersEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	accountDto, err := riot.GetAccountByRiotId("", "")

	if (AccountDto{} != accountDto) || err == nil {
		t.Fatalf(`GetAccountByRiotId("", "") = %v, %v, want nil, error`, accountDto, err)
	}
}

func TestGetAccountByPuuidPuuidEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	accountDto, err := riot.GetAccountByPuuid("")

	if (AccountDto{} != accountDto) || err == nil {
		t.Fatalf(`GetAccountByPuuid("") = %v, %v, want nil, error`, accountDto, err)
	}
}

func TestGetActiveShardPuuidEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	activeShardDto, err := riot.GetActiveShard("", "<game>")

	if (ActiveShardDto{} != activeShardDto) || err == nil {
		t.Fatalf(`GetActiveShard("", "<game>") = %v, %v, want nil, error`, activeShardDto, err)
	}
}

func TestGetActiveShardGameEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	activeShardDto, err := riot.GetActiveShard("<puuid>", "")

	if (ActiveShardDto{} != activeShardDto) || err == nil {
		t.Fatalf(`GetActiveShard("<puuid>", "") = %v, %v, want nil, error`, activeShardDto, err)
	}
}

func TestGetActiveShardAllParametersEmpty(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	activeShardDto, err := riot.GetActiveShard("", "")

	if (ActiveShardDto{} != activeShardDto) || err == nil {
		t.Fatalf(`GetActiveShard("", "") = %v, %v, want nil, error`, activeShardDto, err)
	}
}

func TestGetActiveShardWrongGame(t *testing.T) {
	riot, err := New("<API key>", EuropeRegion, HeaderParam, false)

	activeShardDto, err := riot.GetActiveShard("<puuid>", "wrongGame")

	if (ActiveShardDto{} != activeShardDto) || err == nil {
		t.Fatalf(`GetActiveShard("", "") = %v, %v, want nil, error`, activeShardDto, err)
	}
}
