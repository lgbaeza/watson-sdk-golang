package watson_sdk

import(
    "fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func ClassifyUrl(url string, apikey string) map[string]interface{} {
	
	visualRecognitionAPI := "https://gateway.watsonplatform.net/visual-recognition/api/v3/classify?version=2018-03-19&Accept-Language=es&url=" + url
		
	token := GetIAMToken(apikey)

	client := http.Client{}
    request, _ := http.NewRequest("GET", visualRecognitionAPI, nil) 
    request.Header.Set("Authorization", "Bearer " + token)

	resp, err := client.Do(request)
	if err != nil {
		
		fmt.Println("----err")
		fmt.Println(err)

		return nil

	} else {
        contents, _ := ioutil.ReadAll(resp.Body)
		jsonResp := make(map[string]interface{})
		json.Unmarshal(contents, &jsonResp)

		return jsonResp
	}

}
