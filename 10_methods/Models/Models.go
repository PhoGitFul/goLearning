package Models

type Employee struct {
	iD        int    `json:"id"`
	firstName string `json:"first_name"`
	lastName  string `json:"last_name"`
}

func (e *Employee) SetID(newId int) {
	e.iD = newId
}

func (e Employee) GetID() int {
	return e.iD
}
