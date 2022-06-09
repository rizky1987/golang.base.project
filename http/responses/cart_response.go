package responses

import (
	"example/commonHelpers"
	mongoEntity "example/databases/entities/mongo"
)

type CartResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	Title        string `json:"title"`
	CreatedBy    string `json:"created_by"`
	CreatedDate  string `json:"created_date"`
	ModifiedBy   string `json:"modified_by"`
	ModifiedDate string `json:"modified_date"`
}
type CartSuccessResponse struct {
	CommonBaseResponse
	Data CartResponse `json:"data"`
}

type CartFailedResponse struct {
	CommonBaseResponse
}

func ConvertListCartMongoEntityToCartResponseResponse(carts []*mongoEntity.Cart) []*CartResponse {

	results := []*CartResponse{}

	if len(carts) > 0 {

		for _, cartData := range carts {
			results = append(results,
				&CartResponse{
					Id:          cartData.Id.String(),
					Name:        cartData.Name,
					Content:     cartData.Content,
					Title:       cartData.Title,
					CreatedBy:   cartData.Title,
					CreatedDate: commonHelpers.ConvertDateToStringFormatYYYYMMDD(&cartData.CreatedDate),
				},
			)
		}
	}

	return results
}
