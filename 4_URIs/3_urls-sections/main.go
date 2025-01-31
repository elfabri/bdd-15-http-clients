/* Complete the newParsedURL function.
It should return a ParsedURL
with all the parts of a URL.
For example, given this URL:

http://testuser:testpass@testdomain.com:8080/testpath?testsearch=testvalue#testhash

newParsedURL should return a struct with these fields:

return ParsedURL{
	protocol: "http",
	username: "testuser",
	password: "testpass",
	hostname: "testdomain.com",
	port:     "8080",
	pathname: "/testpath",
	search:   "testsearch=testvalue",
	hash:     "testhash",
}

*/
package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

    pass, _ := parsedUrl.User.Password()

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: parsedUrl.User.Username(),
		password: pass,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}

