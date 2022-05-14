package httphelpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

var noresponse = map[string]interface{}{
	"": "",
}

// Converts map[string]interface{} to JSON
func MapToJSON(mapData interface{}) []byte {
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

// Checks if all required parameters are present
func CheckRequiredParameters(params map[string]interface{}, keys []string) error {
	for _, key := range keys {

		if _, ok := params[key]; !ok {
			pc := make([]uintptr, 10)
			runtime.Callers(2, pc)
			f := runtime.FuncForPC(pc[0]).Name()
			details := strings.Split(f, ".")
			funcName := details[len(details)-1]
			return fmt.Errorf("%s is a required parameter for %s\n", key, funcName)
		}
	}

	return nil
}

// Makes a post request to rave api
func MakePostRequest(data interface{}, url string) (error error, response map[string]interface{}) {
	postData := MapToJSON(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return err, noresponse
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return nil, result

}

// Makes a get request to rave api
func MakeGetRequest(url string, params map[string]string) (error error, response map[string]interface{}) {
	var addToUrl string = "?"
	for k, v := range params {
		addToUrl += fmt.Sprintf("%s=%s&", k, v)
	}
	url += addToUrl
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return err, noresponse
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return nil, result

}

// ParseDynamicQuery - parse url query to possible keys
func ParseDynamicQuery(queries url.Values, keys map[string]interface{}) map[string]interface{} {
	for key := range keys {
		query := queries.Get(key)
		if query != "" {
			keys[key] = query
		} else {
			delete(keys, key)
		}
	}
	return keys
}

// ParseSingleQuery - parse url query to possible keys
func ParseSingleQuery(queries url.Values, keys map[string]interface{}) map[string]interface{} {
	for key := range keys {
		query := queries.Get(key)
		if query != "" {
			keys[key] = query
			break
		} else {
			delete(keys, key)
		}
	}
	return keys
}
