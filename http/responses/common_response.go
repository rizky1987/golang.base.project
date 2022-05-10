package responses

type CommonBaseResponse struct {
	Alert AlertResponse          `json:"alert"`
	Data  map[string]interface{} `json:"data"`
}

type AlertResponse struct {
	Code         int    `json:"code"`
	InnerMessage string `json:"inner_message"`
	Message      string `json:"message"`
}

type CommonPagingResponse struct {
	CurrentPage  int  `json:"current_page"`
	Limit        *int `json:"limit"`
	TotalRecords *int `json:"total_records"`
	TotalPages   *int `json:"total_page"`
}
