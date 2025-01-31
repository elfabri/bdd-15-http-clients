/*
Complete the getUsers function. It should:

 * Call http.Get using its url parameter.

 * Decode the JSON data from the response and return it.

We've done this all before,
but now you're writing it all from scratch!
*/
package main

import (
	"encoding/json"
	"net/http"
)

func getUsers(url string) ([]User, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    usrs := []User{}
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&usrs)
    if err != nil {
        return nil, err
    }

    return usrs, nil
}

