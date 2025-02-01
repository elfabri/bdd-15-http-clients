/*
Complete the fetchData function.
It needs to handle network errors
and non-OK responses. It should
return the status code, an int, and an error.
*/
package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
    res, err := http.Get(url)
    if err != nil {
        return 0, fmt.Errorf("network error: %v", err)
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        return res.StatusCode, fmt.Errorf("non-OK HTTP status: %s", res.Status)
    }

    return res.StatusCode, nil
}

