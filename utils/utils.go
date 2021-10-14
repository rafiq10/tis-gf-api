package utils

import (
	"encoding/json"
	"io"
)

func FromJSON(r io.Reader, fs interface{}) error {
	return json.NewDecoder(r).Decode(&fs)
}

func ToJSON(w io.Writer, intf interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(intf)
}

func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}
