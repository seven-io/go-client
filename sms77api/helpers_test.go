package sms77api

import (
	b64 "encoding/base64"
	a "github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"testing"
)

const VinTelekom = "4915126716517"

var client, testIsDummy = GetClient()
var testBadClient = New(Options{})
var testOptions = Options{
	Debug:    true,
	SentWith: "go-client-test",
}

func AssertKeylessCall(t *testing.T, res interface{}, err error) {
	a.Error(t, err)
	a.Nil(t, res)
}

func GetClient() (*Sms77API, bool) {
	var dummy = true
	var apiKey = os.Getenv("SMS77_DUMMY_API_KEY")

	if "" == apiKey {
		apiKey = os.Getenv("SMS77_API_KEY")

		dummy = false
	}

	testOptions.ApiKey = apiKey

	return New(testOptions), dummy
}

func stringToBase64(text string) string {
	return b64.StdEncoding.EncodeToString([]byte(text))
}

func parseURL(text string) *url.URL {
	u, err := url.Parse(text)
	if err != nil {
		panic(err)
	}

	return u
}
