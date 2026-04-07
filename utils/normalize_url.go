package utils

import (
	"net/url"
	"path"
	"strings"
)

func lower(url *url.URL) {
	url.Host = strings.ToLower(url.Host)
	url.Scheme = strings.ToLower(url.Scheme)
}
func cleanup(url *url.URL) {
	url.Fragment = ""
	url.Path = path.Clean(url.Path)
}
func stripDefaultPorts(url *url.URL) {
	host := url.Hostname()
	port := url.Port()
	if (url.Scheme == "http" && port == "80") ||
		(url.Scheme == "https" && port == "443") {
		url.Host = host
	}
}
func stripTrailingSlash(url *url.URL) {
	if url.Path != "/" {
		url.Path = strings.TrimSuffix(url.Path, "/")
		if url.Path == "" {
			url.Path = "/"
		}
	}
}
func stripTrackingParams(url *url.URL) {
	q := url.Query()
	for key := range q {
		k := strings.ToLower(key)
		if (strings.HasPrefix(k, "utm")) ||
		 (k == "ref") ||
		 (k == "fbclid") ||
		 (k == "gclid") {
			q.Del(key)
		}
	}
	url.RawQuery = q.Encode()
}

func normalizeUrl(baseUrl string, href string) (string, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return "", err
	}
	ref, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	// resolve wrt base
	u := base.ResolveReference(ref)
	lower(u)
	cleanup(u)
	stripDefaultPorts(u)
	stripTrailingSlash(u)
	stripTrackingParams(u)
	return u.String(), nil
}