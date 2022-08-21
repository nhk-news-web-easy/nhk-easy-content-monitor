# nhk-easy-content-monitor [![Go](https://github.com/nhk-news-web-easy/nhk-easy-content-monitor/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/nhk-news-web-easy/nhk-easy-content-monitor/actions/workflows/build.yml) [![codecov](https://codecov.io/gh/nhk-news-web-easy/nhk-easy-content-monitor/branch/main/graph/badge.svg?token=2JDP3YUP8K)](https://codecov.io/gh/nhk-news-web-easy/nhk-easy-content-monitor) [![Go Report Card](https://goreportcard.com/badge/github.com/nhk-news-web-easy/nhk-easy-content-monitor)](https://goreportcard.com/report/github.com/nhk-news-web-easy/nhk-easy-content-monitor)
An AWS Lambda that validates the content links of [nhk-easy-api](https://github.com/nhk-news-web-easy/nhk-easy-api).

## Usage
Create an AWS Lambda function, and run `sh ./package.sh` to package the function, then upload `lambda.zip` to your Lambda function.

## License
[MIT](LICENSE)
