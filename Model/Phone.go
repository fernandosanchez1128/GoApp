package Model

type Phone struct {
	PhoneNumber string `json:"phoneNumber"`
}

func NewPhone(phoneNumber string) *Phone {
	return &Phone{PhoneNumber: phoneNumber}
}
