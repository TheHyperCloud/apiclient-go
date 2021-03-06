package hypercloud

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"

	"github.com/franela/goreq"
)

type ApiClient struct {
	applicationID     string
	applicationSecret string
	baseURL           string
	token             string
	tokenExpiresAt    int64
	apiPath           string
	ConsoleSession    ConsoleSession
	Disk              Disk
	Instance          Instance
	IpAddress         IpAddress
	Network           Network
	PerformanceTier   PerformanceTier
	PublicKey         PublicKey
	Region            Region
	Template          Template
}

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func setup(client *ApiClient) {
	client.apiPath = "/api/v1"
	client.ConsoleSession = ConsoleSession{client: client}
	client.Disk = Disk{client: client}
	client.Instance = Instance{client: client}
	client.IpAddress = IpAddress{client: client}
	client.Network = Network{client: client}
	client.PerformanceTier = PerformanceTier{client: client}
	client.PublicKey = PublicKey{client: client}
	client.Region = Region{client: client}
	client.Template = Template{client: client}
}

func NewApplicationClient(base_url string, app_id string, app_secret string) ApiClient {
	client := ApiClient{
		applicationID:     app_id,
		applicationSecret: app_secret,
		baseURL:           base_url,
	}
	setup(&client)
	return client
}

func NewAccessTokenClient(base_url string, access_token string) ApiClient {
	client := ApiClient{
		baseURL:        base_url,
		token:          access_token,
		tokenExpiresAt: time.Now().AddDate(10, 0, 0).Unix(),
	}
	setup(&client)
	return client
}

func (api *ApiClient) Auth() error {
	request := goreq.Request{
		Uri:         api.baseURL + "/oauth/token",
		Method:      "POST",
		Accept:      "application/json",
		ContentType: "application/json",
		Timeout:     5 * time.Second,
		Body: AuthRequest{
			GrantType:    "client_credentials",
			ClientId:     api.applicationID,
			ClientSecret: api.applicationSecret,
		},
	}
	resp, err := request.Do()

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New("Authentication failed against " + api.baseURL)
	}

	var auth_response AuthResponse
	resp.Body.FromJsonTo(&auth_response)
	api.token = auth_response.AccessToken
	api.tokenExpiresAt = time.Now().Unix() + auth_response.ExpiresIn
	return nil
}

func (api *ApiClient) MapRequest(url string, method string, data interface{}) (status int, rbody map[string]interface{}, err error) {
	status, bytes, err := api.Request(url, method, data)
	if err != nil {
		return -1, nil, err
	}

	var blob map[string]interface{}
	if err := json.Unmarshal(bytes, &blob); err != nil {
		return -1, nil, err
	}
	return status, blob, nil
}

func (api *ApiClient) ArrayRequest(url string, method string, data interface{}) (status int, rbody []map[string]interface{}, err error) {
	status, bytes, err := api.Request(url, method, data)
	if err != nil {
		return -1, nil, err
	}

	var blob []map[string]interface{}
	if err := json.Unmarshal(bytes, &blob); err != nil {
		return -1, nil, err
	}
	return status, blob, nil
}

func (api *ApiClient) Request(url string, method string, data interface{}) (status int, response []byte, err error) {
	if (time.Now().Unix() + 60) > api.tokenExpiresAt {
		if err := api.Auth(); err != nil {
			return -1, nil, err
		}
	}

	request := goreq.Request{
		Uri:         api.baseURL + url,
		Method:      method,
		Accept:      "application/json",
		ContentType: "application/json",
		Body:        data,
		Timeout:     5 * time.Second,
	}
	request.AddHeader("Authorization", "Bearer "+api.token)
	resp, err := request.Do()
	if err != nil {
		return -1, nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return resp.StatusCode, buf.Bytes(), nil
}
