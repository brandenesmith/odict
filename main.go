package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"encoding/json"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Error: You must provide the word whose definition you wish to retrieve.")
		os.Exit(1)
	}
	word := os.Args[1]
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v1/entries/en/%s", word)

	// Create the request.
	client := &http.Client{}
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	// Add headers for api id and api key.
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("app_id", os.Getenv("ODICTAPIID"))
	request.Header.Add("app_key", os.Getenv("ODICTAPIKEY"))

	// Make the request.
	resp, err := client.Do(request)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	dictResp := &DictResponse{}

	json.Unmarshal(bytes, dictResp)

	str := `
===================================================

| | | |     +                |
|      |         | | |   | | | | |
|       |   |   |            |
|      |    |   |            |
| | | |     |    | | |       | | |

===================================================

Word:
	%s

Definition:
	%s
`

	fmt.Printf(
		str,
		word,
		dictResp.Results[0].LexicalEntries[0].Entries[0].Senses[0].Definitions[0],
	)
}
