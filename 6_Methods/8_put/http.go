/*
Complete the updateUser and getUserById functions.

The updateUser function takes a baseURL,
id and apiKey string, and data User as input.
It returns a User and error.

getUserById takes a baseURL,
id and apiKey string. It returns a User and error.

We've included the fullURL creation logic
for you in both functions, we'll be talking
more about URL building in the next chapter.
*/
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id
    jsonData, err := json.Marshal(data)
    if err != nil {
        return User{}, err
    }

    req, err := http.NewRequest(
        "PUT",
        fullURL,
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        return User{}, err
    }
    req.Header.Set("X-API-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return User{}, err
    }
    defer res.Body.Close()

    var user User
    deco := json.NewDecoder(res.Body)
    err = deco.Decode(&user)
    if err != nil {
        return User{}, err
    }

    return user, nil
}

func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

    req, err := http.NewRequest(
        "GET",
        fullURL,
        nil,
    )
    if err != nil {
        return User{}, err
    }
    req.Header.Set("X-API-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        return User{}, err
    }
    defer res.Body.Close()

    var user User
    deco := json.NewDecoder(res.Body)
    err = deco.Decode(&user)
    if err != nil {
        return User{}, err
    }

    return user, nil
}

