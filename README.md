# MyFace GO example

This repository contains a small example how to run a request over Yoti MyFace
service. Run the following commands:

```sh
go mod download
go run main.go <img_path> <sdk_id> <key_file_path>
```

This is the output right now:

```sh
Response status code: 401
Body:
{"error_message": "Failed to verify signature.", "error_code": "INVALID_SIGNATURE"}
```

We can get a result if we provide the same keys in
[this](https://github.com/lampkicking/web-fcm-demo) other example.

Tested version 2 and 3 of `yoti-go-sdk` package
([Guide](https://developers.yoti.com/yoti-myface/integration-guide)).
