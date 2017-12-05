package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Error: You must provide the word whose definition you wish to retrieve.\n")
		os.Exit(1)
	}

	word := os.Args[1]
	url := fmt.Sprintf("https://od-api.oxforddictionaries.com/api/v1/entries/en/%s", word)

	client := &http.Client{}
	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("app_id", os.Getenv("ODICTAPIID"))
	request.Header.Add("app_key", os.Getenv("ODICTAPIKEY"))

	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	code := resp.StatusCode
	bytes, err := ioutil.ReadAll(resp.Body)

	dictResp := &DictResponse{}
	json.Unmarshal(bytes, dictResp)

	header := `
===================================================

                |  +              |
  | | |         |      | | |  | | | | |
|       |   | | |  |  |           |
|       |  |    |  |  |           |
  | | |     | | |  |   | | |      | | |

===================================================

`
	fmt.Printf(header)

	if code == 404 {
		fmt.Printf("Unable to find %s\n", word)
		os.Exit(1)
	}

	entries := dictResp.Results[0].LexicalEntries[0].Entries

	if len(entries) <= 0 {
		fmt.Printf("No entries found.\n")
		os.Exit(0)
	}

	senses := entries[0].Senses

	if len(senses) <= 0 {
		fmt.Printf("No definitions for %s\n", word)
		os.Exit(0)
	}

	definitions := senses[0].Definitions

	if len(definitions) <= 0 {
		fmt.Printf("No definitions for %s\n", word)
		os.Exit(0)
	}

	fmt.Printf("Word:\n\t%s\n\n", word)
	fmt.Printf("Definition:\n\t%s\n\n", definitions[0])

	// Print all of the other definitions.
	if len(senses[0].Subsenses) > 0 {
		fmt.Printf("Other Meanings:\n")

		for index, item := range senses[0].Subsenses {
			fmt.Printf("\t%d. %s\n", index+1, item.Definitions[0])
		}
	}
}
