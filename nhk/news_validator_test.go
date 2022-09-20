package nhk

import (
	nhk_fetcher "github.com/nhk-news-web-easy/nhk-news-fetcher-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateNews(t *testing.T) {
	startDate := time.Date(2022, 8, 8, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 8, 8, 23, 59, 59, 0, time.UTC)
	newsList, err := nhk_fetcher.FetchNews(startDate, endDate)

	assert.True(t, len(newsList) > 0)
	assert.Nil(t, err)

	valid, err := ValidateNews(newsList)

	assert.True(t, valid)
	assert.Nil(t, err)
}
