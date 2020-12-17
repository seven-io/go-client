package sms77api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type HttpMethod string
type Options struct {
	ApiKey   string
	Debug    bool
	SentWith string
}
type resource struct {
	client *Sms77API
}
type StatusCode string
type Sms77API struct {
	Options
	client *http.Client
	base   resource // Instead of allocating a struct for each service we reuse a one

	// Resources
	Analytics        *AnalyticsResource
	Balance          *BalanceResource
	Contacts         *ContactsResource
	Hooks            *HooksResource
	Journal          *JournalResource
	Lookup           *LookupResource
	Pricing          *PricingResource
	Sms              *SmsResource
	Status           *StatusResource
	ValidateForVoice *ValidateForVoiceResource
	Voice            *VoiceResource
}

const (
	defaultOptionSentWith = "go-client"

	sentWithKey = "sentWith"

	HttpMethodGet  HttpMethod = "GET"
	HttpMethodPost HttpMethod = "POST"

	StatusCodeErrorCarrierNotAvailable    StatusCode = "11"
	StatusCodeSuccess                     StatusCode = "100"
	StatusCodeSuccessPartial              StatusCode = "101"
	StatusCodeInvalidSender               StatusCode = "201"
	StatusCodeInvalidRecipient            StatusCode = "202"
	StatusCodeMissingParamTo              StatusCode = "301"
	StatusCodeMissingParamText            StatusCode = "305"
	StatusCodeParamTextExceedsLengthLimit StatusCode = "401"
	StatusCodePreventedByReloadLock       StatusCode = "402"
	StatusCodeReachedDailyLimitForNumber  StatusCode = "403"
	StatusCodeInsufficientCredits         StatusCode = "500"
	StatusCodeErrorCarrierDelivery        StatusCode = "600"
	StatusCodeErrorUnknown                StatusCode = "700"
	StatusCodeErrorAuthentication         StatusCode = "900"
	StatusCodeErrorApiDisabledForKey      StatusCode = "902"
	StatusCodeErrorServerIp               StatusCode = "903"
)

var StatusCodes = map[StatusCode]string{
	StatusCodeErrorCarrierNotAvailable:    "ErrorCarrierNotAvailable",
	StatusCodeSuccess:                     "Success",
	StatusCodeSuccessPartial:              "SuccessPartial",
	StatusCodeInvalidSender:               "InvalidSender",
	StatusCodeInvalidRecipient:            "InvalidRecipient",
	StatusCodeMissingParamTo:              "MissingParamTo",
	StatusCodeMissingParamText:            "MissingParamText",
	StatusCodeParamTextExceedsLengthLimit: "ParamTextExceedsLengthLimit",
	StatusCodePreventedByReloadLock:       "PreventedByReloadLock",
	StatusCodeReachedDailyLimitForNumber:  "ReachedDailyLimitForNumber",
	StatusCodeInsufficientCredits:         "InsufficientCredits",
	StatusCodeErrorCarrierDelivery:        "ErrorCarrierDelivery",
	StatusCodeErrorUnknown:                "ErrorUnknown",
	StatusCodeErrorAuthentication:         "ErrorAuthentication",
	StatusCodeErrorApiDisabledForKey:      "ErrorApiDisabledForKey",
	StatusCodeErrorServerIp:               "ErrorServerIp",
}

func New(options Options) *Sms77API {
	if "" == options.SentWith {
		options.SentWith = defaultOptionSentWith
	}

	c := &Sms77API{client: http.DefaultClient}
	c.Options = options
	c.base.client = c

	c.Analytics = (*AnalyticsResource)(&c.base)
	c.Balance = (*BalanceResource)(&c.base)
	c.Contacts = (*ContactsResource)(&c.base)
	c.Hooks = (*HooksResource)(&c.base)
	c.Journal = (*JournalResource)(&c.base)
	c.Lookup = (*LookupResource)(&c.base)
	c.Pricing = (*PricingResource)(&c.base)
	c.Sms = (*SmsResource)(&c.base)
	c.Status = (*StatusResource)(&c.base)
	c.ValidateForVoice = (*ValidateForVoiceResource)(&c.base)
	c.Voice = (*VoiceResource)(&c.base)

	return c
}

func buildUri(endpoint string) string {
	return fmt.Sprintf("https://gateway.sms77.io/api/%s", endpoint)
}

func (api *Sms77API) createRequestPayload(payload map[string]interface{}) url.Values {
	params := url.Values{}

	for key, value := range payload {
		if nil == value {
			continue
		}

		switch reflect.TypeOf(value).Kind() {
		case reflect.Bool:
			if true == value {
				value = "1"
			} else {
				value = "0"
			}
		case reflect.Int64:
			value = strconv.FormatInt(value.(int64), 10)
		}

		params.Add(key, fmt.Sprintf("%v", value))
	}

	return params
}

func (api *Sms77API) get(endpoint string, data map[string]interface{}) (string, error) {
	payload := api.createRequestPayload(data)
	qs := payload.Encode()

	uri := buildUri(endpoint)
	if "" != qs {
		uri += "?" + qs
	}

	if api.Debug {
		log.Println("GET", uri)
	}

	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return "", fmt.Errorf("could not execute request! #1 (%s)", err.Error())
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", api.ApiKey))
	req.Header.Add(sentWithKey, api.SentWith)

	return api.handleResponse(api.client.Do(req))
}

func (api *Sms77API) handleResponse(res *http.Response, err error) (string, error) {
	if err != nil {
		return "", fmt.Errorf("could not execute request! #2 (%s)", err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("could not execute request! #3 (%s)", err.Error())
	}

	str := strings.TrimSpace(string(body))

	if api.Debug {
		log.Println(str)
	}

	length := len(str)

	if 2 == length || 3 == length {
		code, msg := pickMapByKey(str, StatusCodes)
		if nil != code {
			return "", errors.New(fmt.Sprintf("%s: %s", code, msg))
		}
	}

	return str, nil
}

func (api *Sms77API) post(endpoint string, data map[string]interface{}) (string, error) {
	payload := api.createRequestPayload(data)
	payload.Add("p", api.ApiKey)
	payload.Add(sentWithKey, api.SentWith)

	uri := buildUri(endpoint)

	if api.Debug {
		log.Println("POST", uri, payload)
	}

	return api.handleResponse(api.client.PostForm(uri, payload))
}

func (api *Sms77API) request(endpoint string, httpMethod string, data interface{}) (string, error) {
	if "" == api.Options.ApiKey {
		return "", errors.New("missing required option ApiKey")
	}

	if nil == data {
		data = map[string]interface{}{}
	}

	data, _ = json.Marshal(&data)

	json.Unmarshal(data.([]byte), &data)

	m := map[string]func(*Sms77API, string, map[string]interface{}) (string, error){
		"GET":  (*Sms77API).get,
		"POST": (*Sms77API).post,
	}

	return m[httpMethod](api, endpoint, data.(map[string]interface{}))
}
