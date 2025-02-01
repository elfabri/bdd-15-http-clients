/*
Developers at Jello were prone to making insecure requests,
so management has mandated the use of
a errIfNotHTTPS function to validate all url inputs.

Somehow a developer has still managed
to try to use the wrong URL! Fix the bug.
*/

package main

const URL = "https://api.boot.dev/v1/courses_rest_api/learn-http/users"

