/*
Jello's CEO has a new idea. They want us to
implement a feature that allows users to see
the issues with the lowest time estimate,
with the amount of issues based on their availability.

Complete the fetchTasks function.
It should add two query parameters
to the URL passed to getIssues:

 * sort
 * limit

Sort by the estimate property.

Set the limit based on the availability:

 * Low availability = 1 issue
 * Medium availability = 3 issues
 * High availability = 5 issues
*/
package main

func fetchTasks(baseURL, availability string) []Issue {
    avaNum := "1"
    switch availability {
    case "Low":
        avaNum = "1"
    case "Medium":
        avaNum = "3"
    case "High":
        avaNum = "5"
    }
    queryParams := "?sort=estimate&limit=" + avaNum

	fullURL := baseURL + queryParams
	return getIssues(fullURL)
}

