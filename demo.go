package main

func main() {
	var apiKeyID = "YOUR_VERACODE_API_KEY_ID"
	var apiKeySecret = "YOUR_VERACODE_API_KEY_SECRET"

	response, err := GetApplicationList(apiKeyID, apiKeySecret)

	if err != nil {
		panic(err)
	}

	print(response)
}
