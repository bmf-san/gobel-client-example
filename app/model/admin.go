package model

// A Admins represents the plural of admin.
type Admins []Admin

// A Admin represents the singular of admin.
type Admin struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
