package invitational

type Invitation struct {
	Email string `json:"email" bson:"email"`
	Code  string `json:"code" bson:"code"`
}
