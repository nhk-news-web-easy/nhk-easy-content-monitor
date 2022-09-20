package nhk

import (
	"errors"
	"fmt"
	"github.com/nhk-news-web-easy/nhk-news-fetcher-go"
	"net/http"
)

func ValidateNews(newsList []nhk_fetcher.News) (bool, error) {
	for _, news := range newsList {
		urls := []string{news.Url, news.M3u8Url}

		for _, url := range urls {
			valid, err := validateUrl(url)

			if err != nil {
				return false, err
			}

			if !valid {
				return false, nil
			}
		}
	}

	return true, nil
}

func validateUrl(url string) (bool, error) {
	request, err := http.NewRequest(http.MethodHead, url, nil)

	if err != nil {
		return false, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return false, err
	}

	if response.StatusCode != http.StatusOK {
		return false, errors.New(fmt.Sprintf("Expected 200, but got %d", response.StatusCode))
	}

	return true, nil
}
