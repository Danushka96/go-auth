package types

type User struct {
	Id        string `bson:"_id,omitempty"`
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
}
