package entity

import (
	uiza "api-wrapper-go"
	"net/http"
)

// Client is used to invoke /Entity and entity-related APIs.
type Client struct {
	B   uiza.Backend
	Key string
}

const baseURL = "api/public/v3/media/entity"

// Retrieve Entity API
func Retrieve(params *uiza.EntityRetrieveParams) (string, error) {
	return getC().Retrieve(params)
}

// Retrieve Entity API
func (c Client) Retrieve(params *uiza.EntityRetrieveParams) (string, error) {
	var entity string
	path := uiza.FormatURLPath(baseURL)
	err := c.B.Call(http.MethodGet, path, c.Key, params, &entity)

	return entity, err
}

// Create Entity API
func Create(params *uiza.EntityCreateParams) (string, error) {
	return getC().Create(params)
}

// Create Entity API
func (c Client) Create(params *uiza.EntityCreateParams) (string, error) {
	var entity string

	if err := CheckCreateParamsIsValid(params); err != nil {
		return "", err
	}

	err := c.B.Call(http.MethodPost, baseURL, c.Key, params, &entity)
	return entity, err
}

// Get Backend Client
func getC() Client {
	return Client{uiza.GetBackend(uiza.APIBackend), uiza.Key}
}
