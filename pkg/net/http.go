package net

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
)

type defaultHTTPClient struct {
	client  *http.Client
	apiKey  string
	baseUrl string
}

func NewHTTPClient(openApiKey, baseUrl string, proxy *string) (HTTPClient, error) {
	client := defaultHTTPClient{
		client:  &http.Client{},
		apiKey:  openApiKey,
		baseUrl: baseUrl,
	}
	if proxy != nil {
		url, err := url.Parse(*proxy)
		if err != nil {
			return nil, err
		}
		fmt.Printf("Using proxy: %s\n", url)
		client.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(url),
		}
	}
	return &client, nil
}

func (c *defaultHTTPClient) Do(ctx context.Context, req *OpenAIRequest) (*OpenAIResponse, error) {
	httpReq, err := mapOpenAIToHttp(req, c.baseUrl, c.apiKey)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			// Log the error if closing the body fails
			log.Printf("error closing response body: %v", err)
		}
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error code: %d", resp.StatusCode)
	}

	headers := make(map[string]string, len(resp.Header))
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0] // Take the first value if multiple exist
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &OpenAIResponse{
		Headers:      headers,
		Body:         body,
		ResponseCode: resp.StatusCode,
	}, nil
}

func mapOpenAIToHttp(req *OpenAIRequest, baseUrl, apiKey string) (*http.Request, error) {
	var bodyReader io.Reader
	headers := make(map[string]string, len(req.Headers))
	for k, v := range req.Headers {
		headers[k] = v
	}
	if len(req.FormValues) > 0 || len(req.FormFiles) > 0 {
		if len(req.Body) > 0 {
			return nil, fmt.Errorf("body and form data cannot both be set")
		}
		buf := &bytes.Buffer{}
		writer := multipart.NewWriter(buf)
		for key, value := range req.FormValues {
			if err := writer.WriteField(key, value); err != nil {
				return nil, err
			}
		}
		for _, formFile := range req.FormFiles {
			if formFile.FieldName == "" {
				return nil, fmt.Errorf("form file field name is required")
			}
			if formFile.FileName == "" {
				return nil, fmt.Errorf("form file name is required")
			}
			if formFile.File == nil {
				return nil, fmt.Errorf("form file data is required")
			}
			header := textproto.MIMEHeader{}
			header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, formFile.FieldName, formFile.FileName))
			if mimeType := formFile.File.MimeType; mimeType != "" {
				header.Set("Content-Type", mimeType)
			}
			part, err := writer.CreatePart(header)
			if err != nil {
				return nil, err
			}
			if _, err := part.Write(formFile.File.Data); err != nil {
				return nil, err
			}
		}
		if err := writer.Close(); err != nil {
			return nil, err
		}
		bodyReader = buf
		headers["Content-Type"] = writer.FormDataContentType()
	} else {
		bodyReader = bytes.NewReader(req.Body)
	}

	httpReq, err := http.NewRequest(req.Method, baseUrl+req.Path, bodyReader)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	contentTypeSet := false
	for k, v := range headers {
		httpReq.Header.Set(k, v)
		if strings.EqualFold(k, "Content-Type") {
			contentTypeSet = true
		}
	}
	if !contentTypeSet {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	return httpReq, nil
}
