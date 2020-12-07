package main

import (
	"encoding/json"
	"fmt"
	"github.com/guillaumeguirriec/go-riot-api"
)

func PrintResult(value interface{}) (err error) {
	bytes, err := json.MarshalIndent(value, "", "  ")

	if err == nil {
		fmt.Println(string(bytes))
	}

	return err
}

func main() {
	riot, err := riot.New("<API key>", riot.EuropeRegion, false)

	if err != nil {
		fmt.Println(err)
	}

	response, err := riot.GetAccountByRiotId("<gameName>", "<tagLine>")
	// response, err := riot.GetAccountByPuuid("<puuid>")
	// response, err := riot.GetActiveShard("<puuid>", "<game>")

	if err != nil {
		fmt.Println(err)
	} else {
		PrintResult(response)
	}
}
