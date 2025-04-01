# go-aws-cli

> `aws-cli` clone in Go, this is a toy project, DO NOT USE IN PRODUCTION

## Supported Commands
- `aws sts get-caller-identity`: `app`

## How to Use This Program

```
# export AWS credentials
# - AWS_ACCESS_KEY_ID
# - AWS_SECRET_ACCESS_KEY
# - AWS_SESSION_TOKEN (optional)
# - AWS_DEFAULT_REGION (optional)
# run: go run main.go

$ make run
```

## Dependencies

- [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2)