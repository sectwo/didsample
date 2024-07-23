package service

import (
	"didSample/model"
	"log"
)

func CreateClaimsSample() *model.Claims {
	log.Println("Sample Claims Generating...")
	claims := &model.Claims{
		Name:        "Sujan",
		Identifier:  "a123456789",
		Telephone:   "010-0000-0000",
		Address:     "Seoul City",
		Email:       "aaa@example.com",
		Description: "employer",
	}
	log.Println("Sample Claims Generate Success!!")
	return claims
}
