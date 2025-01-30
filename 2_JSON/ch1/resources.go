/*
A junior developer's deadline is looming.
They've been tasked with creating functions
to retrieve and log various resources from Jello's API.
Rather than making multiple structs and
type-safe functions, the junior developer has decided
to make two flexible functions. Help them complete
their task before their boss, Wayne Lagner,
asks them to refactor it.

*/
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func getResources(url string) ([]map[string]any, error) {
	var resources []map[string]any

	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}

	defer res.Body.Close()

    decoder := json.NewDecoder(res.Body)
    if err := decoder.Decode(&resources); err != nil {
        fmt.Println("error decoding response body")
        return nil, err
    }

    return resources, nil
}

func logResources(resources []map[string]any) {
    var formattedStrings []string

    for _, m := range resources {
        for k, v := range m {
            formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", k, v))
        }
    }

	sort.Strings(formattedStrings)

	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

