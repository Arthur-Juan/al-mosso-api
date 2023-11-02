package types

type AddFoodToAppointmentInput struct {
	PIN    string `json:"PIN"`
	FoodId uint64 `json:"food_id"`
}
