package main

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

// don't touch above this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

func newUser(name string, membershipType membershipType) User {
	if membershipType == "standard" {
		var user = User{
			Name: name,
			Membership: Membership{
				Type:             TypeStandard,
				MessageCharLimit: 100,
			},
		}
		return user
	} else {
		var user = User{
			Name: name,
			Membership: Membership{
				Type:             TypePremium,
				MessageCharLimit: 1000,
			},
		}
		return user
	}
}
