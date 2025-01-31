/*
Complete the getDomainNameFromURL function.
Given a full URL, it should return
the domain (or host) name.
Simply return any potential errors.
*/
package main

import (
	"fmt"
	"net/url"
)

func getDomainNameFromURL(rawURL string) (string, error) {
    domain, err := url.Parse(rawURL)
    if err != nil {
        return "", fmt.Errorf("error parsing url: %w", err)
    }
    domainStr := domain.Hostname();
    return domainStr, nil
}

