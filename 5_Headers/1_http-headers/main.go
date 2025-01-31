/*
Complete the getContentType function. It takes
a pointer to a http.Response as input
and should return the Content-Type header.

Use the .Get() method on the
Response struct's Header field to get what you need.
*/
package main

import (
	"net/http"
)

func getContentType(res *http.Response) string {
    
    // reading a header from the response
    header := res.Header.Get("content-type")
    return header
}

