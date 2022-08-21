package nhk

import "net/http"

func ValidateNews(newsList []News) (bool, error) {
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
		return false, nil
	}

	return true, nil
}
