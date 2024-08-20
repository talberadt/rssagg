package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Example: Authorization: ApiKey {api_key}
func GetAPIKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")
	if header == "" {
		return "", errors.New("no authentication info found")
	}
	vals := strings.Split(header, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part auth header")
	}
	return vals[1], nil
}
