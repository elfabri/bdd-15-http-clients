/*
In previous lessons, we've converted response
into slices of bytes, and then strings. Now,
decode the response data directly
into a slice of issues and return that instead.

If any error occurs return a nil slice and the error

*/
package main

import (
	"fmt"
	"net/http"
    "encoding/json"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

    var issues []Issue
    decoder := json.NewDecoder(res.Body)
    if err := decoder.Decode(&issues); err != nil {
        fmt.Println("error decoding response body")
        return nil, err
    }

    return issues, nil
}
