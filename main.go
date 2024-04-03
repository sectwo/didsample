package main

import (
	"didSample/model"
	"didSample/service"
	"fmt"
)

func createClaimsSample() *model.Claims {
	claims := &model.Claims{
		Name:        "Sujan",
		Identifier:  "a123456789",
		Telephone:   "010-0000-0000",
		Address:     "Seoul City",
		Email:       "aaa@bu.com",
		Description: "employer",
	}
	return claims
}

func main() {
	fmt.Println("DID SAMPLE")

	// 0. Issuer DID 생성 및 등록
	issuerDidInfo, _ := service.GenerateDID("issuer")

	fmt.Println("issuer did : ", issuerDidInfo.DID)
	fmt.Println("issuer public key([]byte) : ", issuerDidInfo.PubKey)
	fmt.Println("issuer private key([]byte) : ", issuerDidInfo.PrvKey)

	// 1. DID 및 DID Document 생성 및 등록
	holderDidInfo, _ := service.GenerateDID("holder")
	fmt.Println("holder did : ", holderDidInfo.DID)
	fmt.Println("holder public key([]byte) : ", holderDidInfo.PubKey)
	fmt.Println("holder private key([]byte) : ", holderDidInfo.PrvKey)

	// 2. Holder와 Issuer DID를 기반으로 VC 생성 및 서명 생성
	//holderDID := holderDidInfo.DID
	claims := createClaimsSample()
	vc := service.GenerateVC(holderDidInfo, issuerDidInfo, claims)

	// 2-1. vc 검증 테스트
	verifyVCResult := service.VerifyVC(&vc, issuerDidInfo)
	fmt.Println("verify VC result : ", verifyVCResult)

	// 3. Holder 와 Issuer DID 및 VC를 기반으로 VP 생성 및 서명 생성
	// 이때 VP는 사용자 혹은 시스템에 필요한 포맷에 따라 Claims 의 선택 가능
	// VP 생성 시에는 Holder의 Public Key를 사용하여 Sign 생성
	vp := service.GenerateVP(holderDidInfo, issuerDidInfo, vc)

	// 3-1. vp 검증 테스트
	verifyVPResult := service.VerifyVP(&vp, holderDidInfo)
	fmt.Println("verify VP result : ", verifyVPResult)
}
