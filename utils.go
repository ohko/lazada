package lazada

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// MakeGUID make GUID
func MakeGUID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[8:10], b[6:8], b[4:6], b[10:])
}

// Request 获取http/https内容
func Request(method, url, data string, header map[string]string) ([]byte, error) {
	var client *http.Client

	if strings.HasPrefix(url, "https://") {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client = &http.Client{Transport: tr}
	} else {
		client = &http.Client{}
	}

	req, err := http.NewRequest(method, url, strings.NewReader(data))
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return bs, nil
}
