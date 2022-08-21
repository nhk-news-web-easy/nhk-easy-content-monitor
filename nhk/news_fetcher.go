package nhk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const requestUrl = "https://nhk.dekiru.app/news"
const isoFormat = "2006-01-02T15:04:05.000Z07:00"

type News struct {
	NewsId          string
	Title           string
	TitleWithRuby   string
	OutlineWithRuby string
	Body            string
	BodyWithoutRuby string
	Url             string
	M3u8Url         string
	ImageUrl        string
	PublishedAtUtc  time.Time
}

func FetchNews(startDate time.Time, endDate time.Time) ([]News, error) {
	url := fmt.Sprintf(requestUrl+"?startDate=%s&endDate=%s", startDate.Format(isoFormat), endDate.Format(isoFormat))
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, nil
	}

	body, err := ioutil.ReadAll(response.Body)

	var newsList []News

	err = json.Unmarshal(body, &newsList)

	if err != nil {
		return nil, err
	}

	return newsList, nil
}
