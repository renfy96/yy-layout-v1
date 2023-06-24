package params

type LoginReq struct {
	Account  string `json:"account" binding:"required" form:"account"`
	Password string `json:"password" binding:"required" form:"password"`
}
