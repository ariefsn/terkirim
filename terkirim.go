package terkirim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TerkirimClient struct {
	apiKey  string
	baseUrl string
}

func New(apiKey string) *TerkirimClient {
	return &TerkirimClient{
		apiKey:  apiKey,
		baseUrl: DefaultUrl,
	}
}

func (c *TerkirimClient) SetBaseUrl(baseUrl string) {
	c.baseUrl = baseUrl
}

func (c *TerkirimClient) getBaseUrl() string {
	if c.baseUrl == "" {
		return DefaultUrl
	}

	return c.baseUrl
}

func (c *TerkirimClient) Whatsapp(payload WhatsappPayload) (*Response, error) {
	res, err := c.send(Whatsapp, payload)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TerkirimClient) Email(payload EmailPayload) (*Response, error) {
	if len(payload.Attachments) == 0 {
		payload.Attachments = []EmailAttachment{}
	}

	if len(payload.Cc) == 0 {
		payload.Cc = []EmailAccount{}
	}

	if len(payload.Bcc) == 0 {
		payload.Bcc = []EmailAccount{}
	}

	res, err := c.send(Email, payload)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *TerkirimClient) send(service Service, payload interface{}) (*Response, error) {
	payloadM, err := convert[M](payload)
	if err != nil {
		return nil, err
	}

	payloadM["apiKey"] = c.apiKey

	body, err := toJsonBody(payloadM)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.getBaseUrl(), serviceUrlPaths[service])

	resHttp, err := http.Post(url, "application/json", body)
	if err != nil {
		return nil, err
	}

	defer resHttp.Body.Close()

	var res Response
	err = json.NewDecoder(resHttp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
