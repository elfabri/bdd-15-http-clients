/*
Complete the deleteUser function.

Check the response status code.
If the status code indicates a non-successful response,
return an error. Otherwise return nil.
*/
package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	fullURL := baseURL + "/" + id

    req, err := http.NewRequest(
        "DELETE",
        fullURL,
        nil,
    )
    if err != nil {
        fmt.Println(err)
        return err
    }
    req.Header.Set("X-API-Key", apiKey)

    res, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer res.Body.Close()

    if res.StatusCode > 299 {
        fmt.Println("request to delete location unsuccessful")
        return fmt.Errorf("%w",err)
    }

    fmt.Println("request to delete location successful")
    return nil
}

