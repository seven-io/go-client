package sms77api

type AnalyticsParams struct {
	Start       string `json:"start,omitempty"`
	End         string `json:"end,omitempty"`
	GroupBy     string `json:"group_by,omitempty"`
	Label       string `json:"label,omitempty"`
	Subaccounts string `json:"subaccounts,omitempty"`
}

type Analytics struct {
	Account *string `json:"account"`
	Country *string `json:"country"`
	Date    *string `json:"date"`
	Label   *string `json:"label"`

	Direct   int     `json:"direct"`
	Economy  int     `json:"economy"`
	Hlr      int     `json:"hlr"`
	Inbound  int     `json:"inbound"`
	Mnp      int     `json:"mnp"`
	Voice    int     `json:"voice"`
	UsageEur float64 `json:"usage_eur"`
}

type AnalyticsResponse []Analytics

type Carrier struct {
	Country     string `json:"country"`
	Name        string `json:"name"`
	NetworkCode string `json:"network_code"`
	NetworkType string `json:"network_type"`
}

type Contact struct {
	Id    int64  `json:"id"`
	Nick  string `json:"nick"`
	Phone string `json:"empfaenger"`
	EMail string `json:"email"`
}

type ContactsParams struct {
	Action string `json:"action"`

	Json  *bool   `json:"json,omitempty"`
	Id    *int64  `json:"id,omitempty"`
	Nick  *string `json:"nick,omitempty"`
	Phone *string `json:"empfaenger,omitempty"`
	EMail *string `json:"email,omitempty"`
}

type CountryNetwork struct {
	Comment     string   `json:"comment,omitempty"`
	Features    []string `json:"features,omitempty"`
	Mcc         string   `json:"mcc,omitempty"`
	Mncs        []string `json:"mncs,omitempty"`
	NetworkName string   `json:"networkName,omitempty"`
	Price       float64  `json:"price,omitempty"`
}

type CountryPricing struct {
	CountryCode   string           `json:"countryCode,omitempty"`
	CountryName   string           `json:"countryName,omitempty"`
	CountryPrefix string           `json:"countryPrefix,omitempty"`
	Networks      []CountryNetwork `json:"networks,omitempty"`
}

type LookupParams struct {
	Type   string `json:"type"`
	Number string `json:"number,omitempty"`
	Json   bool   `json:"json,omitempty"`
}

type LookupCnamResponse struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Number  string `json:"number"`
	Success string `json:"success"`
}

type LookupFormatResponse struct {
	National               string `json:"national"`
	Carrier                string `json:"carrier"`
	CountryCode            string `json:"country_code"`
	CountryIso             string `json:"country_iso"`
	CountryName            string `json:"country_name"`
	International          string `json:"international"`
	InternationalFormatted string `json:"international_formatted"`
	NetworkType            string `json:"network_type"`
	Success                bool   `json:"success"`
}

type LookupHlrResponse struct {
	CountryCode               string  `json:"country_code"`
	CountryCodeIso3           *string `json:"country_code_iso3,omitempty"`
	CountryName               string  `json:"country_name"`
	CountryPrefix             string  `json:"country_prefix"`
	CurrentCarrier            Carrier `json:"current_carrier"`
	GsmCode                   string  `json:"gsm_code"`
	GsmMessage                string  `json:"gsm_message"`
	InternationalFormatNumber string  `json:"international_format_number"`
	InternationalFormatted    string  `json:"international_formatted"`
	LookupOutcome             bool    `json:"lookup_outcome"`
	LookupOutcomeMessage      string  `json:"lookup_outcome_message"`
	NationalFormatNumber      string  `json:"national_format_number"`
	OriginalCarrier           Carrier `json:"original_carrier"`
	Ported                    string  `json:"ported"`
	Reachable                 string  `json:"reachable"`
	Roaming                   string  `json:"roaming"`
	Status                    bool    `json:"status"`
	StatusMessage             string  `json:"status_message"`
	ValidNumber               string  `json:"valid_number"`
}

type LookupMnpResponse struct {
	Code    int64   `json:"code"`
	Mnp     Mnp     `json:"mnp"`
	Price   float64 `json:"price"`
	Success bool    `json:"success"`
}

type Mnp struct {
	Country                string `json:"country"`
	InternationalFormatted string `json:"international_formatted"`
	IsPorted               bool   `json:"isPorted"`
	Mccmnc                 string `json:"mccmnc"`
	NationalFormat         string `json:"national_format"`
	Network                string `json:"network"`
	Number                 string `json:"number"`
}

type PricingParams struct {
	Country string `json:"country,omitempty"`
	Format  string `json:"format,omitempty"`
}

type PricingResponse struct {
	CountCountries int64            `json:"countCountries"`
	CountNetworks  int64            `json:"countNetworks"`
	Countries      []CountryPricing `json:"countries"`
}

type SmsParams struct {
	Debug               bool   `json:"debug,omitempty"`
	Delay               string `json:"delay,omitempty"`
	Details             bool   `json:"details,omitempty"`
	Flash               bool   `json:"flash,omitempty"`
	ForeignId           string `json:"foreign_id,omitempty"`
	From                string `json:"from,omitempty"`
	Label               string `json:"label,omitempty"`
	Json                bool   `json:"json,omitempty"`
	NoReload            bool   `json:"no_reload,omitempty"`
	PerformanceTracking bool   `json:"performance_tracking,omitempty"`
	ReturnMessageId     bool   `json:"return_msg_id,omitempty"`
	Text                string `json:"text"`
	To                  string `json:"to"`
	Ttl                 int64  `json:"ttl,omitempty"`
	Udh                 string `json:"udh,omitempty"`
	Unicode             bool   `json:"unicode,omitempty"`
	Utf8                bool   `json:"utf8,omitempty"`
}

type SmsResponse struct {
	Debug    string  `json:"debug"`
	Balance  float64 `json:"balance"`
	Messages []struct {
		Encoding  string    `json:"encoding"`
		Error     string    `json:"error"`
		ErrorText string    `json:"error_text"`
		Id        string    `json:"id"`
		Messages  *[]string `json:"messages,omitempty"`
		Parts     int64     `json:"parts"`
		Price     float64   `json:"price"`
		Recipient string    `json:"recipient"`
		Sender    string    `json:"sender"`
		Success   bool      `json:"success"`
		Text      string    `json:"text"`
	} `json:"messages"`
	SmsType    string  `json:"sms_type"`
	Success    string  `json:"success"`
	TotalPrice float64 `json:"total_price"`
}

type StatusParams struct {
	MessageId int64 `json:"msg_id"`
}

type ValidateForVoiceParams struct {
	Number   string `json:"number"`
	Networks string `json:"callback,omitempty"`
}

type ValidateForVoiceResponse struct {
	Code            string  `json:"code,omitempty"`
	Error           *string `json:"error"`
	FormattedOutput *string `json:"formatted_output,omitempty"`
	Id              *int64  `json:"id,omitempty"`
	Sender          string  `json:"sender,omitempty"`
	Success         bool    `json:"success"`
	Voice           bool    `json:"voice,omitempty"`
}

type VoiceParams struct {
	To   string `json:"to"`
	Text string `json:"text"`
	Xml  bool   `json:"xml,omitempty"`
	From string `json:"from,omitempty"`
}
