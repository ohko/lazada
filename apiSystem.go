package lazada

import (
	"encoding/json"
	"errors"
)

// ResponseGenerateAccessToken ...
type ResponseGenerateAccessToken struct {
	Code    string `json:"code"`              // "0"
	Type    string `json:"type,omitempty"`    // "ISV"
	Message string `json:"message,omitempty"` // "missing required parameter: access_token"
	Detail  string `json:"detail,omitempty"`

	AccessToken      string            `json:"access_token"`       // "50000601c30atpedfgu3LVvik87Ixlsvle3mSoB7701ceb156fPunYZ43GBg"
	Country          string            `json:"country"`            // "sg"
	RefreshToken     string            `json:"refresh_token"`      // "500016000300bwa2WteaQyfwBMnPxurcA0mXGhQdTt18356663CfcDTYpWoi"
	AccountID        string            `json:"account_id"`         // "7063844"
	AccountPlatform  string            `json:"account_platform"`   // "seller_center"
	RefreshExpiresIn int               `json:"refresh_expires_in"` // 60
	CountryUserInfo  []CountryUserInfo `json:"country_user_info"`  // [...]
	ExpiresIn        int               `json:"expires_in"`         // "10"
	RequestID        string            `json:"request_id"`         // "0ba2887315178178017221014"
	Account          string            `json:"account"`            // "xxx@126.com
}

// CountryUserInfo ...
type CountryUserInfo struct {
	Country   string `json:"country"`    // "my"
	UserID    string `json:"user_id"`    // "100322142"
	SellerID  string `json:"seller_id"`  // "1000087545"
	ShortCode string `json:"short_code"` // "MYJ2FMZV"
}

// GenerateAccessToken generate access_token for call api, the endpoint is https://auth.lazada.com/rest
func (o *Lazada) GenerateAccessToken(code, uuid string) (*ResponseGenerateAccessToken, error) {

	api := "/auth/token/create"
	req := o.mkParams(api, map[string]string{
		"code": code,
		"uuid": uuid,
	})
	bs, err := o.request("POST", api, req)
	if err != nil {
		return nil, err
	}

	var resp ResponseGenerateAccessToken
	if err := json.Unmarshal(bs, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "0" {
		return nil, errors.New(string(bs))
	}

	o.accessToken = resp.AccessToken

	return &resp, nil
}
