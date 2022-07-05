package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// isJSONArray returns true if a JSON payload is a collection of JSON
func isJSONArray(jsonBytes []byte) bool {
	trimmedJSON := bytes.TrimLeft(jsonBytes, " \t\r\n")
	isCollection := len(trimmedJSON) > 0 && string(trimmedJSON[0]) == "["
	return isCollection
}

// convertToJSONArray alters a JSON object to be contained within an array
func convertToJSONArray(jsonBytes []byte) []byte {
	jsonBytes = append(jsonBytes, []byte{0, 0}...)
	copy(jsonBytes[1:], jsonBytes)
	jsonBytes[0] = byte('[')
	jsonBytes[len(jsonBytes)-1] = byte(']')
	return jsonBytes
}

// SerializeJSONPayload arranges our JSON into an array if it's valid
func SerializeJSONPayload(ctx *gin.Context) ([]byte, error) {
	reader := ctx.Request.Body
	json, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return nil, readErr // TODO better error handling
	}

	if !isJSONArray(json) {
		json = convertToJSONArray(json)
	}

	return json, nil
}
