package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

type Headers map[string]interface{}
type QueryParams map[string]interface{}
type PathParams []interface{}

type Params struct {
	Method       string
	URL          string
	Body         interface{}
	Headers      Headers
	Timeout      int
	PathParams   PathParams
	QueryParams  QueryParams
	BasicAuth    *BasicAuth
	HandleErrors *bool
}

type BasicAuth struct {
	Username string
	Password string
}

type Response struct {
	StatusCode int
	Headers    Headers
	Body       map[string]interface{}
	RawBody    []byte
}

func New(params Params) (*Response, error) {
	var body *bytes.Reader

	if params.Body != nil {
		data, err := json.Marshal(params.Body)

		if err != nil {
			return &Response{}, err
		}

		body = bytes.NewReader(data)
	}

	if len(params.PathParams) > 0 {
		params.URL = strings.TrimSuffix(params.URL, "/")

		for _, v := range params.PathParams {
			params.URL += "/" + toString(v)
		}

	}

	if len(params.QueryParams) > 0 {
		query := url.Values{}

		for k, v := range params.QueryParams {
			query.Add(k, toString(v))
		}

		params.URL += "?" + query.Encode()
	}

	var request *http.Request
	var err error

	if body == nil {
		request, err = http.NewRequest(params.Method, params.URL, nil)
	} else {
		request, err = http.NewRequest(params.Method, params.URL, body)
	}

	if err != nil {
		return &Response{}, err
	}

	if params.BasicAuth != nil {
		request.SetBasicAuth(params.BasicAuth.Username, params.BasicAuth.Password)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	if len(params.Headers) > 0 {

		for header, value := range params.Headers {
			request.Header.Set(header, toString(value))
		}

	}

	client := &http.Client{}

	if params.Timeout == 0 {
		params.Timeout = 40
	}

	client.Timeout = time.Duration(params.Timeout) * time.Second

	res, err := client.Do(request)

	if err != nil {
		return &Response{}, err
	}

	defer res.Body.Close()

	headers := Headers{}
	for name, values := range res.Header {
		headers[name] = values[0]
	}

	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return &Response{StatusCode: res.StatusCode, Headers: headers, RawBody: rawBody}, err
	}

	var untypedResponseBody interface{}
	err = json.Unmarshal(rawBody, &untypedResponseBody)

	var responseBody map[string]interface{}
	if jsonObject, isJsonObject := untypedResponseBody.(map[string]interface{}); isJsonObject {
		responseBody = jsonObject
	} else {
		responseBody = map[string]interface{}{"data": untypedResponseBody}
	}

	if params.HandleErrors == nil || *params.HandleErrors {

		if res.StatusCode > 299 {
			return &Response{StatusCode: res.StatusCode, Headers: headers, RawBody: rawBody, Body: responseBody}, getError(responseBody, rawBody)
		}

	}

	if err != nil {

		if res.ContentLength == 0 && res.StatusCode < 300 {
			return &Response{StatusCode: res.StatusCode, Headers: headers, Body: map[string]interface{}{}, RawBody: []byte{}}, nil
		} else {
			return &Response{StatusCode: res.StatusCode, Headers: headers, RawBody: rawBody, Body: responseBody}, err
		}

	}

	return &Response{StatusCode: res.StatusCode, RawBody: rawBody, Body: responseBody, Headers: headers}, nil
}

func getError(body map[string]interface{}, rawBody ...[]byte) error {

	if body != nil {
		if body["error"] != nil {
			return errors.New(toString(body["error"]))
		} else if body["errors"] != nil {
			return errors.New(stringSlice(body["errors"])[0])
		} else if body["authErrors"] != nil {
			return errors.New(stringSlice(body["authErrors"])[0])
		} else if body["message"] != nil {
			return errors.New(toString(body["message"]))
		}
	}

	if len(rawBody) > 0 && len(rawBody[0]) > 0 {
		return errors.New(string(rawBody[0]))
	}

	return errors.New("Ocorreu uma falha ao realizar operação") //lint:ignore ST1005 ignore
}

func stringSlice(itface interface{}) []string {
	s, ok := itface.([]interface{})
	if !ok {
		return []string{toString(itface)}
	}
	str := make([]string, len(s))
	for i, value := range s {
		str[i] = toString(value)
	}
	return str
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {

		if rv.IsNil() {
			return ""
		} else {
			return fmt.Sprintf("%v", rv.Elem())
		}

	}

	return fmt.Sprintf("%v", rv)
}
