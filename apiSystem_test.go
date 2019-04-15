package lazada

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLazada_GenerateAccessToken(t *testing.T) {
	api := newLazada(t)
	resp, err := api.GenerateAccessToken("0_109531_Wj1639BSGHkK6Z2T3zPo62Im760", "123456")
	if err != nil {
		t.Fatal(err)
	}

	bs, _ := json.Marshal(resp)
	fmt.Println(string(bs))
	// output:
	// {"code":"0","access_token":"50000201117x1laaYjUiJrDxsgWda1740d4b2ydnRxiHIOG2GvUDlWjRmhx5N3", ... }
}
