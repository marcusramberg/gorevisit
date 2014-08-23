package gorevisit

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// BytesToDataURI returns a data URI encoded string given a byte array and a content type
// See RFC2397 - http://tools.ietf.org/html/rfc2397
func BytesToDataURI(data []byte, contentType string) string {
	return fmt.Sprintf("data:%s;base64,%s",
		contentType, base64.StdEncoding.EncodeToString(data))
}

// DataURIToBytes returns a content type string and an array of bytes
// given a data URI encoded string.
// See RFC2397 - http://tools.ietf.org/html/rfc2397
func DataURIToBytes(dataURI string) (*DecodedContent, error) {
	parts := strings.Split(dataURI, ",")
	contentType := parts[0]
	contentBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return &DecodedContent{}, err
	} else {
		return &DecodedContent{Type: contentType, Data: contentBytes}, nil
	}
}
