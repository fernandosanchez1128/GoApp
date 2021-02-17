package Model

type User struct {
	Id       int     `json:"id"`
	Document string  `json:"document"`
	Name     string  `json:"name"`
	Lastname string  `json:"lastname"`
	Phones   []Phone `json:"phones"`
}

func NewUser(id int, document string, name string, lastname string) *User {
	return &User{Id: id, Document: document, Name: name, Lastname: lastname}
}

func (u *User) SetPhones(phones []Phone) {
	u.Phones = phones
}
