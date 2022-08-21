package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nhk-news-web-easy/nhk-easy-content-monitor/nhk"
	"time"
)

func HandleRequest(ctx context.Context) (string, error) {
	now := time.Now().UTC()
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	endDate := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)

	_, err := nhk.FetchNews(startDate, endDate)

	if err != nil {
		return "error", err
	}

	return "ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}
