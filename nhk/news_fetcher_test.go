package nhk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFetchNews(t *testing.T) {
	startDate := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2022, 1, 1, 23, 59, 59, 0, time.UTC)

	newsList, err := FetchNews(startDate, endDate)

	assert.True(t, len(newsList) > 0)
	assert.Nil(t, err)
}
