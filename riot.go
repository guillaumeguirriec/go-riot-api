// TODO: deal with rate limit
// TODO: make use of debugging boolean
// TODO: GetAccountByPuuid
// TODO: GetActiveShards

package riot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	AmericasRegion = "americas"
	EuropeRegion   = "europe"
	AsiaRegion     = "asia"
)

type (
	AccountDto struct {
		puuid, gameName, tagLine string
	}

	ActiveShardDto struct {
		puuid, game, activeShard string
	}

	Riot struct {
		apiKey, region string
		debugging      bool
	}
)

func New(apiKey, region string, debugging bool) (*Riot, error) {

	errorMessage := ""

	if apiKey == "" {
		errorMessage = fmt.Sprintf("API key not given")
	}

	if region != AmericasRegion && region != EuropeRegion && region != AsiaRegion {
		errorMessage = fmt.Sprintf("Given region to execute against is not correct (given: %v, wanted: %v, %v or %v).", region, AmericasRegion, EuropeRegion, AsiaRegion)
	}

	if errorMessage != "" {
		return nil, errors.New(errorMessage)
	}

	riot := new(Riot)

	riot.apiKey = apiKey
	riot.region = region
	riot.debugging = debugging

	return riot, nil
}

func (riot Riot) sendGetRequest(endpoint string) []byte {
	httpClient := &http.Client{}
	request, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(nil))

	if err != nil {
		// TODO: deal with error
	}

	request.Header.Set("X-Riot-Token", riot.apiKey)

	response, err := httpClient.Do(request)

	if err != nil {
		// TODO: deal with error
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		// TODO: deal with error
	}

	return responseBody
}

func (riot Riot) GetAccountByRiotId(gameName, tagLine string) (AccountDto, error) {
	responseJsonBody := riot.sendGetRequest("https://" + riot.region + ".api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + gameName + "/" + tagLine)

	var accountDto AccountDto
	err := json.Unmarshal(responseJsonBody, &accountDto)

	if err != nil {
		// TODO: deal with error
	}

	return accountDto, nil
}