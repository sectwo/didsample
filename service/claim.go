package service

import "didSample/model"

func CreateClaimsSample() *model.Claims {
	claims := &model.Claims{
		Name:        "Sujan",
		Identifier:  "a123456789",
		Telephone:   "010-0000-0000",
		Address:     "Seoul City",
		Email:       "aaa@example.com",
		Description: "employer",
	}
	return claims
}
