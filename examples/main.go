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
	riot := riot.New("<API key>", riot.EuropeRegion, false)

	response, err := riot.GetAccountByRiotId("<gameName>", "<tagLine>")

	if err != nil {
		fmt.Println(response)
	} else {
		PrintResult(response)
	}
}
