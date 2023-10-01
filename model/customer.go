package model

type Customer struct {
	Id        int
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool // has customer been contacted
}
