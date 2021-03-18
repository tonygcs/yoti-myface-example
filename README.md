# MyFace GO example

This repository contains a small example how to run a request over Yoti MyFace
service. Run the following commands:

```sh
go mod download
go run main.go <img_path> <sdk_id> <key_file_path> <endpoint> <url>
```

## Arguments

- img_path: the path where the image to send is located.
- sdk_id: the SDK id to request the AI service.
- key_file_path: the path where the pem file with the key is located.
- endpoint: the prediction that the client will request (`age`,
  `age-antispoofing` or `antispoofing`).
- url: the url where the service is located (the default value is the production
  environment: `https://api.yoti.com/ai/v1`).
