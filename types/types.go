package types

type LoginReq struct {
	Name string `json:"name" binding:"required"`
	// name 이라는 값은 해당 api 에서 필수적으로 들어와야함, 그렇기에 binding 에 required 을 줘서 필수값을 받도록 했음.

}
