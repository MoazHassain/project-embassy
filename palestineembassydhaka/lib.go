package main

import (
	"fmt"
	"net/http"
	"strings"
)

//GetBaseURL ..
func GetBaseURL(r *http.Request) string {

	var baseurl, proto string
	if strings.Contains(r.Proto, "HTTP") {
		proto = "http"
	} else {
		proto = "https"
	}
	baseurl = fmt.Sprintf("%s://%s", proto, r.Host)
	return baseurl
}
