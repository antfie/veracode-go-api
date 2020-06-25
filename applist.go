package api

import (
	"encoding/xml"
	"net/http"
)

// GetApplicationList returns a list of applications.
func GetApplicationList(apiKeyID, apiKeySecret string) (Application, error) {
	var apiUrl = "https://analysiscenter.veracode.com/api/5.0/getapplist.do"
	response, err := makeApiRequest(apiKeyID, apiKeySecret, apiUrl, http.MethodGet)
	var result Application

	if err != nil {
		return result, err
	}

	if err := xml.Unmarshal([]byte(response), &result); err != nil {
		return result, err
	}

	return result, nil
}

type Application struct {
	XMLName xml.Name `xml:"app"`
	Id      int      `xml:"app_id"`
	Name    string   `xml:"app_name"`
}

type ApplicationList struct {
	XMLName      xml.Name      `xml:applist"`
	Applications []Application `xml:"app"`
}
