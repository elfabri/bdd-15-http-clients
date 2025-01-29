/* Throughout this course, we'll be building parts
 * of an online issue tracking app called "Jello".
 * It's a typical issue tracker
 * with one key difference: this time it's good, actually.
 * 
 * Take a look at the getIssueData function
 * It retrieves information about issues from
 * Jello's servers via HTTP as a slice of bytes []byte.
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	// Don't edit above this line

    fmt.Printf("%s", string(issues))
}

