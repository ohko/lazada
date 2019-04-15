package lazada

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLazada_GetBrands(t *testing.T) {
	api := newLazada(t)
	resp, err := api.GetBrands("0", "10")
	if err != nil {
		t.Fatal(err)
	}

	bs, _ := json.Marshal(resp)
	fmt.Println(string(bs))
	// output:
	// {"code":"0","data":[{"name":"test carmen-edit3","brand_id":1,"global_identifier":"test_carmen","name_en":"test carmen"}, ... }
}
