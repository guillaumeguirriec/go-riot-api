// TODO: deal with rate limit
// TODO: make use of debugging boolean

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
	AmericasRegion     = "americas"
	EuropeRegion       = "europe"
	AsiaRegion         = "asia"
	Valorant           = "val"
	LegendsOfRuneterra = "lor"
	HeaderParam        = "headerParam"
	QueryParam         = "queryParam"
)

type (
	AccountDto struct {
		Puuid, GameName, TagLine string
	}

	ActiveShardDto struct {
		Puuid, Game, ActiveShard string
	}

	Riot struct {
		apiKey, region, includeApiKeyAs string
		debugging                       bool
	}
)

func New(apiKey, region, includeApiKeyAs string, debugging bool) (*Riot, error) {
	if apiKey == "" {
		errorMessage := fmt.Sprintf("API key not given")
		return nil, errors.New(errorMessage)
	}

	if region == "" {
		errorMessage := fmt.Sprintf("Region not given")
		return nil, errors.New(errorMessage)
	}

	if region != AmericasRegion && region != EuropeRegion && region != AsiaRegion {
		errorMessage := fmt.Sprintf("Given region to execute against is not correct (given: %v, wanted: %v, %v or %v).", region, AmericasRegion, EuropeRegion, AsiaRegion)
		return nil, errors.New(errorMessage)
	}

	if includeApiKeyAs == "" {
		errorMessage := fmt.Sprintf("Way to include API key not given")
		return nil, errors.New(errorMessage)
	}

	if includeApiKeyAs != HeaderParam && includeApiKeyAs != QueryParam {
		errorMessage := fmt.Sprintf("Given way to include API key is not correct (given: %v, wanted: %v or %v).", includeApiKeyAs, HeaderParam, QueryParam)
		return nil, errors.New(errorMessage)
	}

	riot := new(Riot)

	riot.apiKey = apiKey
	riot.region = region
	riot.includeApiKeyAs = includeApiKeyAs
	riot.debugging = debugging

	return riot, nil
}

func (riot Riot) sendGetRequest(endpoint string) []byte {
	httpClient := &http.Client{}
	request, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(nil))

	if err != nil {
		// TODO: deal with error
	}

	if riot.includeApiKeyAs == HeaderParam {
		request.Header.Set("X-Riot-Token", riot.apiKey)
	} else if riot.includeApiKeyAs == QueryParam {
		request.Header.Set("X-Riot-Token", riot.apiKey)
		endpoint = endpoint + "?api_key=" + riot.apiKey
	}

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
	var accountDto AccountDto

	if gameName == "" || tagLine == "" {
		errorMessage := fmt.Sprintf("Required parameters not given (gameName: %v, tagLine: %v).", gameName, tagLine)
		return accountDto, errors.New(errorMessage)
	}

	responseBody := riot.sendGetRequest("https://" + riot.region + ".api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + gameName + "/" + tagLine)

	err := json.Unmarshal(responseBody, &accountDto)

	if err != nil {
		// TODO: deal with error
	}

	return accountDto, nil
}

func (riot Riot) GetAccountByPuuid(puuid string) (AccountDto, error) {
	var accountDto AccountDto

	if puuid == "" {
		errorMessage := fmt.Sprintf("Required parameters not given (puuid: %v).", puuid)
		return accountDto, errors.New(errorMessage)
	}

	responseBody := riot.sendGetRequest("https://" + riot.region + ".api.riotgames.com/riot/account/v1/accounts/by-puuid/" + puuid)

	err := json.Unmarshal(responseBody, &accountDto)

	if err != nil {
		// TODO: deal with error
	}

	return accountDto, nil
}

func (riot Riot) GetActiveShard(puuid, game string) (ActiveShardDto, error) {
	var activeShardDto ActiveShardDto

	if puuid == "" || game == "" {
		errorMessage := fmt.Sprintf("Required parameters not given (puuid: %v, game: %v).", puuid, game)
		return activeShardDto, errors.New(errorMessage)
	}

	if game != Valorant && game != LegendsOfRuneterra {
		errorMessage := fmt.Sprintf("Given game is not correct (given: %v, wanted: %v or %v).", game, Valorant, LegendsOfRuneterra)
		return activeShardDto, errors.New(errorMessage)
	}

	responseBody := riot.sendGetRequest("https://" + riot.region + ".api.riotgames.com/riot/account/v1/active-shards/by-game/" + game + "/by-puuid/" + puuid)

	err := json.Unmarshal(responseBody, &activeShardDto)

	if err != nil {
		// TODO: deal with error
	}

	return activeShardDto, nil
}
