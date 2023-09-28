package authenticationLibrary

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func RetrieveAuthToken(httpClient http.Client) (string, error) {

	envErrors := initializeEnvironment()
	fmt.Println(envErrors)
	apiURL := environment[authUrl]

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		log.Printf("Log 1 %v", err)
		return "", err

	}
	inputBytes := []byte(environment[clientId] + ":" + environment[clientSecret])
	encodedString := base64.StdEncoding.EncodeToString(inputBytes)
	authHeader := "Basic " + encodedString
	req.Header.Set("Authorization", authHeader)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Log 2 %v", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("status code %d", resp.StatusCode)
		return "", nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// Handle the error
		return "", err
	}
	var jsonData map[string]interface{}

	if err := json.Unmarshal([]byte(body), &jsonData); err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	access_token, found := jsonData["access_token"]
	if found {
		fmt.Print(access_token)
		responseToken := fmt.Sprintf("%v", access_token)
		return responseToken, nil
	}
	fmt.Print(string(body))
	return "", nil
}
