/*
The Jello server uses URL paths to denote
different "resources". For example,

/issues
/projects

The entire server is hosted at the path
/v1/courses_rest_api/learn-http, however,
so all requests must be prefixed with that path.

Update the URL in the code to fetch "project" data
instead of "issue" data.

Use the /v1/courses_rest_api/learn-http/projects path.
*/
package main

const url = "https://api.boot.dev/v1/courses_rest_api/learn-http/projects"

