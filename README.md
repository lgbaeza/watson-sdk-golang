
# IBM Watson SDK - GOLANG

This is a starting project for IBM Watson SDK on golang. Currently supports the following services and API Methods:

- Authentication
	- IAM token auth
- Visual Recognition
	- GET Classify

Version 0.1

## Here is a sample code snippet to see it working
```go
package main

import (
	"fmt"
	"strconv"
	"github.com/lgbaeza/watson_sdk"
)

func main() {

	apikey := "" //your apikey
	imgUrl := "https://alkaway.com/wp-content/uploads/runner.jpg" //a sample image
	res := watson_sdk.ClassifyUrl(imgUrl, apikey)

	class := getFirstClass(res)
	class_name := class["class"].(string)
	confidence := class["score"].(float64) * 100

	conf := strconv.FormatFloat(confidence, 'f', 0, 64)

	fmt.Println("Watson está " + conf + "% seguro que ve " + class_name + " en la imagen ")
}

func getFirstClass(watsonClassification map[string]interface {}) map[string]interface {}{
	images := watsonClassification["images"].([]interface {})
	classifiers := images[0].(map[string]interface {})["classifiers"].([]interface {})
	classifier := classifiers[0].(map[string]interface {})
	classes := classifier["classes"].([]interface {})
	class := classes[0].(map[string]interface {})

	return class
}
```

## Contributing
If you want to contribute to this project, feel free to fork the repo and start building. You can always reach me at @lgbaeza over twitter or linkedin