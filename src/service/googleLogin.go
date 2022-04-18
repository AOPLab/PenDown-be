package service

import (
	"net/http"

	"google.golang.org/api/oauth2/v2"
)

var httpClient = &http.Client{}

func VerifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(httpClient)
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
