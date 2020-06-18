package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/antfie/veracode-go-hmac-authentication/hmac"
)

// Returns a list of applications.
func GetApplicationList(apiKeyID, apiKeySecret string) (string, error) {
	var apiUrl = "https://analysiscenter.veracode.com/api/5.0/getapplist.do"
	return makeApiRequest(apiKeyID, apiKeySecret, apiUrl, http.MethodGet)
}

func makeApiRequest(apiKeyID, apiKeySecret, apiUrl, httpMethod string) (string, error) {
	parsedUrl, err := url.Parse(apiUrl)

	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, parsedUrl.String(), nil)

	if err != nil {
		return "", err
	}

	authorizationHeader, err := hmac.CalculateAuthorizationHeader(parsedUrl, httpMethod, apiKeyID, apiKeySecret)

	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", authorizationHeader)

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Expected status 200. Status was: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body[:]), nil
}
