package utils

import (
	"fmt"
	"net/url"
)

// GetUris
func GetUris(uris []string, port string) []string {
	muris := make([]string, 0)
	for _, uri := range uris {
		u, err := url.Parse(uri)
		if err != nil {
			// TODO:
		}
		var mp string
		if u.Port() == "443" || u.Port() == "26657" {
			mp = u.Port()
		} else {
			mp = port
		}
		var path string
		if u.Path == "/" {
			path = ""
		} else {
			path = u.Path
		}
		muris = append(muris, fmt.Sprintf("%s://%s:%s%s", u.Scheme, u.Hostname(), mp, path))
	}
	return muris
}
