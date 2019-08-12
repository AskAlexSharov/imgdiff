package resource

import "net/http"

//--
// Data model objects and persistence mocks:
//--

// User data model
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// User fixture data
var users = []*User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

//--
// Request and Response payloads for the REST api.
//
// The payloads embed the data model objects an
//
// In a real-world project, it would make sense to put these payloads
// in another file, or another sub-package.
//--

type UserPayload struct {
	*User
	Role string `json:"role"`
}

func NewUserPayloadResponse(user *User) *UserPayload {
	return &UserPayload{User: user}
}

// Bind on UserPayload will run after the unmarshalling is complete, its
// a good time to focus some post-processing after a decoding.
func (u *UserPayload) Bind(r *http.Request) error {
	return nil
}

func (u *UserPayload) Render(w http.ResponseWriter, r *http.Request) error {
	u.Role = "collaborator"
	return nil
}
