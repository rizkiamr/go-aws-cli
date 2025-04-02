# go-aws-cli

> `aws-cli` clone in Go, this is a toy project, DO NOT USE IN PRODUCTION

## Supported Commands
- `aws-cli sts get-caller-identity`

## How to Use This Program

```
# export AWS credentials
# - AWS_ACCESS_KEY_ID
# - AWS_SECRET_ACCESS_KEY
# - AWS_SESSION_TOKEN (optional)
# - AWS_DEFAULT_REGION (optional)

$ make build
$ ./build/aws-cli <command> <subcommand> [parameters]
```

## Dependencies

- [AWS SDK for Go v2](https://github.com/aws/aws-sdk-go-v2)
- [Cobra](https://github.com/spf13/cobra)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)