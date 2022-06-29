package modal

type ClientAddQuery struct {
	Domain string `json:"domain" validate:"required"`
	Secret string `json:"secret" validate:"required"`
	ID     string `json:"id" validate:"required"`
}

type ClientEditQuery struct {
	Domain string `json:"domain"`
	Secret string `json:"secret"`
	ID     string `json:"id" validate:"required"`
}
