package yandexgpt

import (
	"net/http"
)

const (
	BaseURL = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type YandexGPTClient struct {
	config     *YandexGPTClientConfig
	httpClient *http.Client
}

func NewYandexGPTClient(
	apiKey,
	catalogID string,
	model yandexGPTModel,
) *YandexGPTClient {
	client := &http.Client{}
	config := NewYandexGPTClientConfig(apiKey, catalogID, model)

	return &YandexGPTClient{
		config:     config,
		httpClient: client,
	}
}

func (c *YandexGPTClient) CreateRequest(
	request YandexGPTRequest,
) (response YandexGPTResponse, err error) {

	//1. Validate Request
	//2. Create Request via c.newRequest(...)
	//3. Send request via c.sendRequest(...)
	//P.s. Use pointers
	return
}
