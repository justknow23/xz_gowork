package json

import (
	"bytes"
	encodingjson "encoding/json"
	jsoniter "github.com/json-iterator/go"
	"io"
)

// Number -
var Number encodingjson.Number

// JSON -
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an InvalidUnmarshalError.
func Unmarshal(data []byte, v interface{}) error { return JSON.Unmarshal(data, v) }

// Marshal returns the JSON encoding of v.
func Marshal(v interface{}) ([]byte, error) { return JSON.Marshal(v) }

// MarshalIndent is like Marshal but applies Indent to format the output.
// Each JSON element in the output will begin on a new line beginning with prefix
// followed by one or more copies of indent according to the indentation nesting.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return JSON.MarshalIndent(v, prefix, indent)
}

// UnmarshalDecoder By Decoder
func UnmarshalDecoder(data []byte, v interface{}) error {
	decoder := encodingjson.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	return decoder.Decode(v)
}

// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
func HTMLEscape(dst *bytes.Buffer, src []byte) {
	encodingjson.HTMLEscape(dst, src)
}

// NewDecoder returns a new decoder that reads from r.
func NewDecoder(r io.Reader) *encodingjson.Decoder {
	return encodingjson.NewDecoder(r)
}

// NewEncoder returns a new encoder that writes to w.
func NewEncoder(w io.Writer) *encodingjson.Encoder {
	return encodingjson.NewEncoder(w)
}

// ToJSONStringPretty encodingJson MarshalIndent Pretty
func ToJSONStringPretty(v interface{}) string {
	jsonStr, _ := JSON.MarshalIndent(v, "", "    ")
	return string(jsonStr)
}

// FromJSONString encodingJson Unmarshal string
func FromJSONString(data string, v interface{}) error {
	return JSON.Unmarshal([]byte(data), v)
}
