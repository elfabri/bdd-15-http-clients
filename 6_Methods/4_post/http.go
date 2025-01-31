/*
Complete the createUser function. It should:

 1 - Take a URL and apiKey string, and User data as parameters

 2 - Encode the user data using json.Marshal
 
 3 - Create a new POST request using http.NewRequest.
     Use a bytes.NewBuffer to create a io.Reader
     from the JSON data.

 4 - Modify the request headers:

   * - Set the Content-Type header, with application/json as its value

   * - Set the X-API-Key header, with apiKey as its value

 5 - Make the request using the http.Client's Do method.
     You can create a new http.Client or use the http.DefaultClient

 6 - Decode and return the response's JSON body (which is also a User)

Don't copy paste from the code above.
Type it out and understand what each line does.
*/
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func createUser(url, apiKey string, data User) (User, error) {
    // 2
    encodedData, err := json.Marshal(data)
    if err != nil {
		return User{}, err
    }

    // 3
    req, err := http.NewRequest(
        "POST",
        url,
        bytes.NewBuffer(encodedData),
    )
    if err != nil {
		return User{}, err
    }

    // 4
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    // 5
    res, err := http.DefaultClient.Do(req)
    if err != nil {
		return User{}, err
    }
    defer res.Body.Close()

    // 6
    var user User
    decoder := json.NewDecoder(res.Body)
    err = decoder.Decode(&user)
	if err != nil {
		return User{}, err
	}

    return user, nil
}

