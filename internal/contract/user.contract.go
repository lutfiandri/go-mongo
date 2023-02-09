package contract

type InsertOneUserReqBody struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type FindOneUserReqParams struct {
	ID string `params:"id" validate:"required"`
}
