package sms77api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const BaseUri = "https://gateway.sms77.io/api/"
var _, isDev = os.LookupEnv("SMS77_DEBUG")

type Sms77Api = Sms77API

type Sms77API struct {
	apiKey string
	client *http.Client
}

func New(apiKey string) Sms77API {
	return Sms77API{
		apiKey: apiKey,
		client: http.DefaultClient,
	}
}

func (api *Sms77API) Analytics(p *AnalyticsParams) (AnalyticsResponse, error) {
	res, err := api.request("analytics", "GET", p)

	if err != nil {
		return nil, err
	}

	var js = AnalyticsResponse{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return js, nil
}

func (api *Sms77API) Balance() (*float64, error) {
	res, err := api.request("balance", "GET", nil)

	if err != nil {
		return nil, err
	}

	float, err := strconv.ParseFloat(res, 64)
	if err != nil {
		return nil, err
	}

	return &float, nil
}

func (api *Sms77API) Contacts(p ContactsParams) (*string, error) {
	method := "POST"
	if p.Action == "read" {
		method = "GET"
	}

	res, err := api.request("contacts", method, p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *Sms77API) Lookup(p LookupParams) (interface{}, error) {
	res, err := api.request("lookup", "GET", p)

	if err != nil {
		return nil, err
	}

	var js interface{}

	switch p.Type {
	case "mnp":
		if !p.Json {
			return res, nil
		}

		js = &LookupMnpResponse{}
	case "cnam":
		js = &LookupCnamResponse{}
	case "format":
		js = &LookupFormatResponse{}
	case "hlr":
		js = &LookupHlrResponse{}
	}

	if err := json.Unmarshal([]byte(res), js); err != nil {
		return nil, err
	}

	return js, nil
}

func (api *Sms77API) Pricing(p PricingParams) (interface{}, error) {
	res, err := api.request("pricing", "GET", p)

	if err != nil {
		return nil, err
	}

	if p.Format == "csv" {
		return res, nil
	}

	var js = PricingResponse{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return js, nil
}

func (api *Sms77API) Sms(p SmsParams) (interface{}, error) {
	res, err := api.request("sms", "POST", p)

	if err != nil {
		return nil, err
	}

	if p.Json == false {
		return res, nil
	}

	var js = SmsResponse{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return js, nil
}

func (api *Sms77API) Status(p StatusParams) (*string, error) {
	res, err := api.request("status", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *Sms77API) ValidateForVoice(p ValidateForVoiceParams) (*ValidateForVoiceResponse, error) {
	res, err := api.request("validate_for_voice", "GET", p)

	if err != nil {
		return nil, err
	}

	var js = ValidateForVoiceResponse{}

	if err := json.Unmarshal([]byte(res), &js); err != nil {
		return nil, err
	}

	return &js, nil
}

func (api *Sms77API) Voice(p VoiceParams) (*string, error) {
	res, err := api.request("voice", "POST", p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func createRequestPayload(i interface{}) url.Values {
	params := url.Values{}

	if i != nil {
		fields := reflect.TypeOf(i)
		values := reflect.ValueOf(i)
		num := fields.NumField()

		for i := 0; i < num; i++ {
			value := values.Field(i)

			if !value.IsZero() {
				jsonTag := fields.Field(i).Tag.Get("json")
				commaIdx := strings.Index(jsonTag, ",")
				name := jsonTag
				if commaIdx >= 0 {
					name = jsonTag[:commaIdx]
				}

				var str = ""
				switch value.Type().Kind() {
				case reflect.Bool:
					if value.Bool() == true {
						str = "1"
					} else {
						str = "0"
					}
				case reflect.Int64:
					str = strconv.FormatInt(value.Int(), 10)
				}
				if str != "" {
					value = reflect.ValueOf(str)
				}

				params.Add(name, value.String())
			}
		}
	}

	return params
}

func (api *Sms77API) request(endpoint string, method string, data interface{}) (string, error) {
	const senderKey = "sentWith"
	const senderValue = "go"
	uri := fmt.Sprintf("%s%s", BaseUri, endpoint)
	payload := createRequestPayload(data)

	if isDev {
		fmt.Println(method + "@" + uri)
	}

	var res *http.Response
	var err error

	if method == "POST" {
		payload.Add("p", api.apiKey)
		payload.Add(senderKey, senderValue)

		if isDev {
			fmt.Println(payload)
		}

		res, err = http.PostForm(uri, payload)
	} else {
		req, err := http.NewRequest(method, uri+"?"+payload.Encode(), nil)

		if err != nil {
			return "", fmt.Errorf("could not execute request! #1 (%s)", err.Error())
		}

		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", api.apiKey))
		req.Header.Add(senderKey, senderValue)

		res, err = api.client.Do(req)
	}

	if err != nil {
		return "", fmt.Errorf("could not execute request! #2 (%s)", err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("could not execute request! #3 (%s)", err.Error())
	}

	str := string(body)

	if isDev {
		fmt.Println(str)
	}

	return str, nil
}
