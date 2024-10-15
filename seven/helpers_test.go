package seven

import (
	b64 "encoding/base64"
	a "github.com/stretchr/testify/assert"
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"
)

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

func GetClient() (*API, bool) {
	var dummy = true
	var apiKey = os.Getenv("SEVEN_API_KEY_SANDBOX")

	if "" == apiKey {
		apiKey = os.Getenv("SEVEN_API_KEY")

		dummy = false
	}

	testOptions.ApiKey = apiKey

	return New(testOptions), dummy
}

func stringToBase64(text string) string {
	return b64.StdEncoding.EncodeToString([]byte(text))
}

func parseFloat(line string) (float64, error) {
	return strconv.ParseFloat(line, 32)
}

func parseURL(text string) *url.URL {
	u, err := url.Parse(text)
	if err != nil {
		panic(err)
	}

	return u
}

func timestampString() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
