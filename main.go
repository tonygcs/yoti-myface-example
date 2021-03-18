// Steps to run the example:
//
//   1. go mod download
//   2. go run main.go <img_path> <sdk_id> <key_file_path>
//
// Swith between v2 and v3.

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	// "github.com/getyoti/yoti-go-sdk/v2/requests"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
)

func main() {
	// Read sdk ID and key file path from args.
	imgPath := os.Args[1]
	sdkID := os.Args[2]
	keyFilePath := os.Args[3]
	baseURL := "https://api.yoti.com/ai/v1"
	if len(os.Args) >= 5 {
		baseURL = os.Args[4]
	}

	// Build request.
	file, _ := ioutil.ReadFile(imgPath)
	b64Img := base64.StdEncoding.EncodeToString(file)

	keyFile, _ := ioutil.ReadFile(keyFilePath)

	entity := map[string]string{"img": b64Img}
	reqBody, _ := json.Marshal(entity)

	req, _ := requests.SignedRequest{
		HTTPMethod: http.MethodPost,
		BaseURL:    baseURL,
		Endpoint:   "/age-antispoofing",
		Headers: map[string][]string{
			"Content-Type":   {"application/img"},
			"Accept":         {"application/img"},
			"X-Yoti-Auth-Id": {sdkID},
		},
		Body: reqBody,
	}.WithPemFile(keyFile).Request()

	// Request the service.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// Print result.
	fmt.Printf("Response status code: %d\n", res.StatusCode)
	fmt.Println("Body:")
	resBody, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(resBody))
}
