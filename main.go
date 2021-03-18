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
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	// "github.com/getyoti/yoti-go-sdk/v2/requests"
	"github.com/getyoti/yoti-go-sdk/v3/requests"
)

func main() {
	if len(os.Args) < 5 {
		checkError(errors.New("provide all arguments (<img_path> <sdk_id> <key_file_path> <endpoint> <url>)"))
	}

	// Read sdk ID and key file path from args.
	imgPath := os.Args[1]
	sdkID := os.Args[2]
	keyFilePath := os.Args[3]
	endpoint := os.Args[4]

	baseURL := "https://api.yoti.com/ai/v1" // Prod URL is the default value.
	if len(os.Args) >= 6 {
		baseURL = os.Args[5]
	}

	// Build request.
	file, err := ioutil.ReadFile(imgPath)
	checkError(err)
	b64Img := base64.StdEncoding.EncodeToString(file)

	keyFile, err := ioutil.ReadFile(keyFilePath)
	checkError(err)

	entity := map[string]string{"img": b64Img}
	reqBody, err := json.Marshal(entity)
	checkError(err)

	req, err := requests.SignedRequest{
		HTTPMethod: http.MethodPost,
		BaseURL:    baseURL,
		Endpoint:   fmt.Sprintf("/%s", endpoint),
		Headers: map[string][]string{
			"Content-Type":   {"application/img"},
			"Accept":         {"application/img"},
			"X-Yoti-Auth-Id": {sdkID},
		},
		Body: reqBody,
	}.WithPemFile(keyFile).Request()
	checkError(err)

	// Request the service.
	res, err := http.DefaultClient.Do(req)
	checkError(err)

	// Print result.
	fmt.Printf("Response status code: %d\n", res.StatusCode)
	fmt.Println("Body:")
	resBody, err := ioutil.ReadAll(res.Body)
	checkError(err)
	fmt.Println(string(resBody))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
