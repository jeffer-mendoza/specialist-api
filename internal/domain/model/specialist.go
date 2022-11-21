package model

type Specialist struct {
	ID       string `json:"id"`
	FistName string `json:"first_name"`
	LastName string `json:"last_name"`
	Type     string `json:"type"`
}

type Type struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
