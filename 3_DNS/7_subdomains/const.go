/*
I've updated the getIssues function from before.
Now it accepts just a domain as input.
It's convenient this way because it means
if the API we're using ever changes its domain,
we only need to update one variable.

Problem is, there is a bug. The API isn't hosted
on boot.dev, it's hosted on the "api" subdomain!
Fix the bug.
*/

package main

// const domain = "boot.dev"
const domain = "api.boot.dev"

