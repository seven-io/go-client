package sms77api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

type HttpMethod string

const (
	HttpMethodGet  HttpMethod = "GET"
	HttpMethodPost HttpMethod = "POST"
)

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

type resource struct {
	client *Sms77API
}

type Options struct {
	ApiKey   string
	Debug    bool
	SentWith string
}

const senderKey = "sentWith"
const defaultOptionSentWith = "go-client"

func New(options Options) *Sms77API {
	if "" == options.ApiKey {
		panic(errors.New("missing required option ApiKey"))
	}

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

func (api *Sms77API) request(endpoint string, httpMethod string, data interface{}) (string, error) {
	if nil == data {
		data = map[string]interface{}{}
	}

	data, _ = json.Marshal(&data)

	json.Unmarshal(data.([]byte), &data)

	if "POST" == httpMethod {
		return api.post(endpoint, data.(map[string]interface{}))
	}

	return api.get(endpoint, data.(map[string]interface{}))
}

func (api *Sms77API) get(endpoint string, data map[string]interface{}) (string, error) {
	payload := api.createRequestPayload(data)
	uri := buildUri(endpoint) + "?" + payload.Encode()

	if api.Debug {
		fmt.Println("GET", uri)
	}

	req, err := http.NewRequest("GET", uri, nil)

	if err != nil {
		return "", fmt.Errorf("could not execute request! #1 (%s)", err.Error())
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", api.ApiKey))
	req.Header.Add(senderKey, api.SentWith)

	return api.handleResponse(api.client.Do(req))
}

func (api *Sms77API) post(endpoint string, data map[string]interface{}) (string, error) {
	payload := api.createRequestPayload(data)
	payload.Add("p", api.ApiKey)
	payload.Add(senderKey, api.SentWith)

	uri := buildUri(endpoint)

	if api.Debug {
		fmt.Println("POST", uri, payload)
	}

	return api.handleResponse(api.client.PostForm(uri, payload))
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

	str := string(body)

	if api.Debug {
		fmt.Println(str)
	}

	return str, nil
}
