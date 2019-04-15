package lazada

import (
	"encoding/json"
	"errors"
)

// ResponseGetBrands ...
type ResponseGetBrands struct {
	Code    string `json:"code"`              // "0"
	Type    string `json:"type,omitempty"`    // "ISV"
	Message string `json:"message,omitempty"` // "missing required parameter: access_token"
	Detail  string `json:"detail,omitempty"`

	Data []GetBrandsItem `json:"data"`
}

// GetBrandsItem ...
type GetBrandsItem struct {
	Name             string `json:"name"`              // "Apple"
	BrandID          int    `json:"brand_id"`          // 1001
	GlobalIdentifier string `json:"global_identifier"` // "Apple"
	NameEN           string `json:"name_en"`           // "Apple"
}

// GetBrands Use this API to retrieve all product brands in the system.
func (o *Lazada) GetBrands(offset, limit string) (*ResponseGetBrands, error) {

	api := "/brands/get"
	req := o.mkParams(api, map[string]string{
		"offset": offset,
		"limit":  limit,
	})
	bs, err := o.request("GET", api, req)
	if err != nil {
		return nil, err
	}

	var resp ResponseGetBrands
	if err := json.Unmarshal(bs, &resp); err != nil {
		return nil, err
	}

	if resp.Code != "0" {
		return nil, errors.New(string(bs))
	}

	return &resp, nil
}
