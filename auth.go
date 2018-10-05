package watson_sdk

import(
    "fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"strings"
	"encoding/json"
)

func GetIAMToken(apikey string) string {
	
	authorizationAPI := "https://iam.bluemix.net/identity/token"
	
	data := url.Values{}
    data.Set("grant_type", "urn:ibm:params:oauth:grant-type:apikey")
    data.Add("apikey", apikey)
	
	client := http.Client{}
    request, _ := http.NewRequest("POST", authorizationAPI, strings.NewReader(data.Encode())) 
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    request.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		
		fmt.Println("----err")
		fmt.Println(err)

		return string("err")

	} else {
        contents, _ := ioutil.ReadAll(resp.Body)
		jsonResp := make(map[string]interface{})
		json.Unmarshal(contents, &jsonResp)
		var token string
		token, _ = jsonResp["access_token"].(string)

		return token
	}

}
