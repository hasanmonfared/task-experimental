package estimate

import (
	"encoding/json"
	"gameapp/entity/estimateentity"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

type Client struct {
	address string
}
type ApiResponse struct {
}

func New(address string) Client {
	return Client{
		address: address,
	}
}
func (c Client) GetEstimate(ctx context.Context, orderID uint) (estimateentity.Estimate, error) {
	response, err := http.Get(c.address)
	if err != nil {
		return estimateentity.Estimate{}, nil
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return estimateentity.Estimate{}, nil
	}
	var apiResponse ApiResponse

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return estimateentity.Estimate{}, nil
	}
	return estimateentity.Estimate{NewEstimate: 5}, nil
}
