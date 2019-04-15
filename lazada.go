package lazada

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"
)

// ...
const (
	EndpointSingapore   = "https://api.lazada.sg/rest"
	EndpointThailand    = "https://api.lazada.co.th/rest"
	EndpointMalaysia    = "https://api.lazada.com.my/rest"
	EndpointVietnam     = "https://api.lazada.vn/rest"
	EndpointPhilippines = "https://api.lazada.com.ph/rest"
	EndpointIndonesia   = "https://api.lazada.co.id/rest"
)

// Lazada ...
type Lazada struct {
	endpoint     string
	appKey       string
	appSecret    string
	accessToken  string
	refreshToken string

	debug bool
}

// New create new lazada
func New(endpoint, appKey, appSecret string) (*Lazada, error) {
	if appKey == "" || appSecret == "" {
		return nil, errors.New("appKey/appSecret is empty")
	}
	return &Lazada{
		endpoint:  endpoint,
		appKey:    appKey,
		appSecret: appSecret,
	}, nil
}

// SetDebug open/close debug log
func (o *Lazada) SetDebug(debug bool) {
	o.debug = debug
}

func (o *Lazada) request(method, api, data string) (resp []byte, err error) {

	var bs []byte

	if method == "POST" {
		bs, err = Request(method, o.endpoint+api, data, nil)
	} else {
		bs, err = Request(method, o.endpoint+api+"?"+data, "", nil)
	}
	if err != nil {
		return nil, err
	}

	if o.debug {
		if method == "POST" {
			log.Printf("curl -X POST '%s%s' -H 'Content-Type:application/x-www-form-urlencoded;charset=utf-8' -d '%s'\n", o.endpoint, api, data)
		} else {
			log.Printf("curl '%s%s?%s'\n", o.endpoint, api, data)
		}
		log.Println(string(bs))
	}
	return bs, nil
}

func (o *Lazada) mkParams(api string, params map[string]string) string {
	common := map[string]string{
		"app_key":     o.appKey,
		"timestamp":   fmt.Sprintf("%d", time.Now().Unix()*1000),
		"sign_method": "sha256",
	}
	if o.accessToken != "" {
		common["access_token"] = o.accessToken
	}

	for k, v := range params {
		common[k] = v
	}

	var arr []string
	for k, v := range common {
		arr = append(arr, k+v)
	}

	sort.Strings(arr)
	mac := hmac.New(sha256.New, []byte(o.appSecret))
	mac.Write([]byte(api + strings.Join(arr, "")))
	common["sign"] = strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))

	rqs := url.Values{}
	for k, v := range common {
		rqs.Set(k, v)
	}

	return rqs.Encode()
}
