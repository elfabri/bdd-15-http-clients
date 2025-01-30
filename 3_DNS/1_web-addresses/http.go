/* Domain Names and IP addresses

When we browse the internet, we type in
a human readable domain name. That domain
is converted into an IP address. The IP address
tells our computer where the server is located
on the internet.

I've provided a getIPAddress function that makes a request
to Cloudflare. The function takes a domain name as input
and returns the IP address associated with that domain.

Update the getIPAddress function to return
just the first IP address found within (if it exists).

*/
package main

import (
	"fmt"
	"io"
	"net/http"
    "encoding/json"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

    var ips map[string][]Answer

    json.Unmarshal(body, &ips)

    for k, v := range ips {
        if k == "Answer" {
            // return first ip
            return v[0].Data, nil
        }
    }
	return "", fmt.Errorf("unknown error")
}

