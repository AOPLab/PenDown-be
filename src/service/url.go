package service

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/AOPLab/PenDown-be/src/model"
	"github.com/AOPLab/PenDown-be/src/persistence"
)

func AddUrl(Original_url string, expireAt time.Time) (*model.Url, error) {
	str := generateRandomString()
	fmt.Println(str)
	url := &model.Url{
		Original_url: Original_url,
		Expired_date: expireAt,
		Url_id:       str,
	}
	err := persistence.DB.Model(&model.Url{}).Create(&url).Error
	if err != nil {
		return nil, err
	} else {
		return url, nil
	}
}

func GetUrl(url_id string) (*model.Url, error) {
	url := &model.Url{}
	err := persistence.DB.Select("original_url", "expired_date").Where("url_id = ?", url_id).First(&url).Error
	if err != nil {
		return nil, err
	} else {
		return url, nil
	}
}

func CheckUrl(Original_url string) error {
	_, err := url.ParseRequestURI(Original_url)
	if err != nil {
		return err
	}
	_, httpErr := http.Get(Original_url)
	if httpErr != nil {
		return httpErr
	}
	return nil
}

func generateRandomString() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
