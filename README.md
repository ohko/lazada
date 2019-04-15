# lazada api

# demo
```go
import "github.com/ohko/lazada"

api, err := lazada.New(EndpointMalaysia, os.Getenv("LAZADA_APPKEY"), os.Getenv("LAZADA_APPSECRET"))
if err!=nil{
   ...
}

resp1, err := api.GenerateAccessToken("0_109531_Wj1639BSGHkK6Z2T3zPo62Im760", "123456")
if err != nil {
   ...
}

resp2, err := api.GetBrands("0", "10")
if err != nil {
   ...
}
```