package lazada

import (
	"log"
	"os"
	"strings"
	"testing"
)

func newLazada(t *testing.T) *Lazada {
	log.SetFlags(log.Lshortfile)
	api, err := New(EndpointMalaysia, os.Getenv("LAZADA_APPKEY"), os.Getenv("LAZADA_APPSECRET"))
	if err != nil {
		t.Fatal(err)
	}
	return api
}

func TestLazada_mkParams(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	api, err := New(EndpointMalaysia, "123456", "helloworld")
	if err != nil {
		t.Fatal(err)
	}
	sign := api.mkParams("/order/get", map[string]string{
		"access_token": "test",
		"timestamp":    "1517820392000",
		"sign_method":  "sha256",
		"order_id":     "1234",
	})
	log.Println(sign)
	if strings.Index(sign, "4190D32361CFB9581350222F345CB77F3B19F0E31D162316848A2C1FFD5FAB4A") <= 0 {
		t.Fatal(sign)
	}
}
