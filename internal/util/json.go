package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

// isJSONArray returns true if a JSON payload is a collection of JSON
func isJSONArray(jsonBytes []byte) bool {
	trimmedJSON := bytes.TrimLeft(jsonBytes, " \t\r\n")
	isCollection := len(trimmedJSON) > 0 && string(trimmedJSON[0]) == "["
	return isCollection
}

// castToJSONArray alters a JSON object to be contained within an array
func castToJSONArray(jsonBytes []byte) []byte {
	jsonBytes = append(jsonBytes, []byte{0, 0}...)
	copy(jsonBytes[1:], jsonBytes)
	jsonBytes[0] = byte('[')
	jsonBytes[len(jsonBytes)-1] = byte(']')
	return jsonBytes
}

// SerializeJSONPayload arranges our JSON into an array if it's valid
func SerializeJSONPayload(ctx *gin.Context) (*[]byte, error) {
	if !strings.Contains(ctx.GetHeader("Content-Type"), "application/json") {
		return nil, errors.New("missing required content-type json header")
	}

	reader := ctx.Request.Body
	jsonBytes, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return nil, errors.New("failed to read json payload")
	}

	if !json.Valid(jsonBytes) {
		return nil, errors.New("invalid json")
	}

	if !isJSONArray(jsonBytes) {
		jsonBytes = castToJSONArray(jsonBytes)
	}

	return &jsonBytes, nil
}
