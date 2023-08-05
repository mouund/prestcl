package controllers

import (
	b64 "encoding/base64"
)

func Encode64(stringToEncode string) string {
	encodedString := b64.StdEncoding.EncodeToString([]byte(stringToEncode))
	return encodedString
}