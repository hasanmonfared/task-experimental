package estimate

import (
	"gameapp/entity/estimateentity"
	"time"
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
func (c Client) GetEstimate(orderID uint) (estimateentity.Estimate, error) {
	const op = "estimate.GetEstimate"

	//Because not work webservice

	//response, err := http.Get(c.address)
	//if err != nil {
	//	return estimateentity.Estimate{}, richerror.New(op).WithErr(err)
	//}
	//defer response.Body.Close()
	//
	//body, err := io.ReadAll(response.Body)
	//if err != nil {
	//	return estimateentity.Estimate{}, richerror.New(op).WithErr(err)
	//}
	//var apiResponse ApiResponse
	//
	//err = json.Unmarshal(body, &apiResponse)
	//if err != nil {
	//	return estimateentity.Estimate{}, richerror.New(op).WithErr(err)
	//}
	timeNow := time.Now()
	return estimateentity.Estimate{NewEstimate: timeNow.Add(time.Second * 50)}, nil
}
