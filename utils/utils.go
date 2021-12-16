package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type ErrMsg struct {
	ErrTxt string `json:"error"`
}

func FromJSON(r io.Reader, data interface{}) error {
	decoder := json.NewDecoder(r)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&data)
	return err
}

func ToJSON(w io.Writer, intf interface{}) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	err := enc.Encode(intf)
	if err == io.EOF {
		err = nil
	}

	return err
}

func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}

func IsJSON(s string) bool {
	var js json.RawMessage
	err := json.Unmarshal([]byte(s), &js)
	if err != nil {
		return false
	}
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		fmt.Println("Is NOT a JSON map")
		return false
	}

	if len(m) == 0 {
		fmt.Println("Empty JSON")
		return false
	}
	return true
}

func HasJsonStructFields(itf interface{}, data string) bool {
	var res bool
	v := reflect.ValueOf(itf)
	typesOf := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if typesOf.Field(i).Name != "R" && typesOf.Field(i).Name != "L" {
			res = HasJsonField(data, typesOf.Field(i).Name)
			if !res {
				fmt.Printf("Field missing: %s\n", typesOf.Field(i).Name)
				return res
			}

		}
	}
	return true

}
func HasJsonField(s, field string) bool {

	var m map[string]interface{}
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		return false
	}

	m2 := make(map[string]interface{})
	for k, v := range m {
		m2[strings.ToLower(k)] = v
	}
	if _, ok := m2[strings.ToLower(field)]; ok {
		return true
	}
	return false

}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func ByteArrayToString(arr []byte) string {
	return bytes.NewBuffer(arr).String()
}

func RespondWithError(w http.ResponseWriter, code int, msgArr []ErrMsg) {
	resp, _ := json.Marshal(msgArr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func AssertResponseStatusCode(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status code: %d but wanted: %d", got, want)
	}
}
func AssertResponseHeader(t *testing.T, response *httptest.ResponseRecorder, header string, headerValue string) {
	t.Helper()
	if response.Result().Header.Get(header) != headerValue {
		t.Errorf("response did not have header %v as expected: %v; but got: %v ", header, headerValue, response.Result().Header)
	}
}
func AssertResponseText(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got response: %s but wanted: %s", got, want)
	}
}
