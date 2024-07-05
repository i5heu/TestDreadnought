package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// PostRequest performs a POST request with the given headers and data.
func PostRequest(baseUrl, url string, headers map[string]string, data interface{}) (string, string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(data).
		Post(baseUrl + url)

	if err != nil {
		return "", "", fmt.Errorf("Error making POST request: %w", err)
	}

	return resp.String(), string(resp.Body()), nil
}

// GetRequest performs a GET request with the given headers.
func GetRequest(baseUrl, url string, headers map[string]string) (string, string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		Get(baseUrl + url)

	if err != nil {
		return "", "", fmt.Errorf("Error making GET request: %w", err)
	}

	return resp.String(), string(resp.Body()), nil
}

// PatchRequest performs a PATCH request with the given headers and data.
func PatchRequest(baseUrl, url string, headers map[string]string, data interface{}) (string, string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		SetBody(data).
		Patch(baseUrl + url)

	if err != nil {
		return "", "", fmt.Errorf("Error making PATCH request: %w", err)
	}

	return resp.String(), string(resp.Body()), nil
}

// DeleteRequest performs a DELETE request with the given headers.
func DeleteRequest(baseUrl, url string, headers map[string]string) (string, string, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeaders(headers).
		Delete(baseUrl + url)

	if err != nil {
		return "", "", fmt.Errorf("Error making DELETE request: %w", err)
	}

	return resp.String(), string(resp.Body()), nil
}
