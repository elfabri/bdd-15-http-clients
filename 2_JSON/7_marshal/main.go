/* If there is a way to unmarshal JSON data,
there must be a way to marshal it as well.
The json.Marshal function converts a Go struct
into a slice of bytes representing JSON data.

Complete the marshalAll function. The expectation is
that they are structs of various forms.

 * It should return a slice of slices of bytes (I didn't stutter),
where each resulting slice of bytes represents
the JSON representation of an item in the input slice.

 * If an item cannot be marshaled,
the function should immediately return an error.

 * Return the marshalled data in the same order as the input items. 
*/
package main

import (
	"encoding/json"
)

func marshalAll[T any](items []T) ([][]byte, error) {
    data := make([][]byte, len(items))
    var err error
    for idx, item := range items {
        data[idx], err = json.Marshal(item)
        if err != nil {
            return nil, err
        }
    }
    return data, nil
}
