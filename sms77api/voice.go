package sms77api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

type Voice struct {
	Code int
	Cost float64
	Id   int
}

type VoiceParams struct {
	// Debug is for internal use only.
	Debug bool `json:"debug,omitempty"`

	To        string `json:"to"`
	Text      string `json:"text"`
	From      string `json:"from,omitempty"`
	Ringtime  int    `json:"ringtime,omitempty"`
	ForeignId string `json:"foreign_id,omitempty"`
	// Deprecated: Xml exists for historical compatibility and should not be used.
	// Please remove all uses of this option.
	Xml bool `json:"xml,omitempty"`
}

type VoiceResource resource

func makeVoice(res string) Voice {
	lines := strings.Split(res, "\n")

	code, _ := strconv.Atoi(lines[0])
	id, _ := strconv.Atoi(lines[1])
	cost, _ := strconv.ParseFloat(lines[2], 64)

	return Voice{
		Code: code,
		Cost: cost,
		Id:   id,
	}
}

func (api *VoiceResource) Text(p VoiceParams) (*string, error) {
	return api.TextContext(context.Background(), p)
}

func (api *VoiceResource) TextContext(ctx context.Context, p VoiceParams) (*string, error) {
	res, err := api.client.request(ctx, "voice", http.MethodPost, p)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (api *VoiceResource) Json(p VoiceParams) (o Voice, e error) {
	r, e := api.Text(p)

	if nil != e {
		return
	}

	return makeVoice(*r), nil
}
