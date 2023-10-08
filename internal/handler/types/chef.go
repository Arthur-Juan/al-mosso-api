package types

type ChefOutput struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}

type InsertChefInput struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	Description string `json:"description"`
	Photo       *TFile
}
