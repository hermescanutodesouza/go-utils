package json

import (
	"bytes"
	stdjson "encoding/json"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

var (
	JSON = jsoniter.Config{
		EscapeHTML:             false,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
)

var (
	arrayPrefix  = []byte("[")
	objectPrefix = []byte("{")
)

type RawMessage = stdjson.RawMessage

// UseNumber solve very big int64 digits loss.
func UseNumber() {
	JSON = jsoniter.Config{
		UseNumber:              true,
		EscapeHTML:             false,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
}

// Marshal returns the JSON encoding of v.
func Marshal(v interface{}) ([]byte, error) {
	return JSON.Marshal(v)
}

// MustMarshal must returns the JSON encoding of v.
func MustMarshal(v interface{}) []byte {
	data, _ := JSON.Marshal(v)
	return data
}

// MarshalToString returns the JSON encoding to string of v.
func MarshalToString(v interface{}) (string, error) {
	return JSON.MarshalToString(v)
}

// MustMarshalToString must returns the JSON encoding to string of v.
func MustMarshalToString(v interface{}) string {
	str, _ := JSON.MarshalToString(v)
	return str
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	return JSON.Unmarshal(data, v)
}

// UnmarshalFromString unmarshal string to v.
func UnmarshalFromString(str string, v interface{}) error {
	return JSON.UnmarshalFromString(str, v)
}

// Valid check JSON data.
func Valid(data []byte) bool {
	return JSON.Valid(data)
}

// ValidFromString check JSON string.
func ValidFromString(str string) bool {
	return Valid([]byte(str))
}

// Get get value from JSON data by path.
// https://pkg.go.dev/github.com/tidwall/gjson?tab=doc#Get
func Get(data []byte, path ...string) gjson.Result {
	if len(path) == 0 || len(path[0]) == 0 {
		return gjson.ParseBytes(data)
	}

	return gjson.GetBytes(data, path[0])
}

// GetFromString get value from JSON string by path.
// https://pkg.go.dev/github.com/tidwall/gjson?tab=doc#Get
func GetFromString(str string, path ...string) gjson.Result {
	if len(path) == 0 || len(path[0]) == 0 {
		return gjson.Parse(str)
	}

	return gjson.Get(str, path[0])
}

func Format(raw []byte) []byte {
	raw = bytes.TrimSpace(raw)
	if bytes.HasPrefix(raw, arrayPrefix) {
		val := []interface{}{}
		if err := Unmarshal(raw, &val); err != nil {
			return raw
		}

		data, err := Marshal(val)
		if err != nil {
			return raw
		}

		return data
	}

	if bytes.HasPrefix(raw, objectPrefix) {
		val := map[string]interface{}{}
		if err := Unmarshal(raw, &val); err != nil {
			return raw
		}

		data, err := Marshal(val)
		if err != nil {
			return raw
		}

		return data
	}

	return raw
}

func FormatFromString(raw string) string {
	return string(Format([]byte(raw)))
}
