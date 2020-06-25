package api

import (
	"encoding/xml"
	"net/http"
)

// GetApplicationList returns a list of applications.
func GetApplicationList(apiKeyID, apiKeySecret string) (ApplicationList, error) {
	var apiUrl = "https://analysiscenter.veracode.com/api/5.0/getapplist.do"
	var result ApplicationList
	return result, makeApiRequest(apiKeyID, apiKeySecret, apiUrl, http.MethodGet, &result)
}

type ApplicationList struct {
	XMLName      xml.Name      `xml:"applist"`
	Applications []Application `xml:"app"`
}

type Application struct {
	XMLName xml.Name `xml:"app"`
	Id      int      `xml:"app_id"`
	Name    string   `xml:"app_name"`
}
