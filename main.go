package main

import (
	"didSample/service"
	"fmt"
)

func main() {
	fmt.Println("=================================== DID SAMPLE ===================================")

	// 0. Issuer DID 생성 및 등록
	issuerDidInfo, _ := service.GenerateDID("issuer")
	fmt.Println("issuer public key([]byte) : ", string(issuerDidInfo.PubKey))
	fmt.Println("issuer private key([]byte) : ", string(issuerDidInfo.PrvKey))
	fmt.Println("issuer did : ", issuerDidInfo.DID)
	fmt.Println()

	// 1. DID 및 DID Document 생성 및 등록
	holderDidInfo, _ := service.GenerateDID("holder")
	fmt.Println("holder public key([]byte) : ", string(holderDidInfo.PubKey))
	fmt.Println("holder private key([]byte) : ", string(holderDidInfo.PrvKey))
	fmt.Println("holder did : ", holderDidInfo.DID)
	fmt.Println()

	// 2. Holder와 Issuer DID를 기반으로 VC 생성 및 서명 생성
	//holderDID := holderDidInfo.DID
	claims := service.CreateClaimsSample()
	vc := service.GenerateVC(holderDidInfo, issuerDidInfo, claims)

	// 2-1. vc 검증 테스트
	verifyVCResult := service.VerifyVC(&vc, issuerDidInfo)
	fmt.Println("Verify VC result : ", verifyVCResult)
	fmt.Println()

	// 3. Holder 와 Issuer DID 및 VC를 기반으로 VP 생성 및 서명 생성
	// 이때 VP는 사용자 혹은 시스템에 필요한 포맷에 따라 Claims 의 선택 가능
	// VP 생성 시에는 Holder의 Public Key를 사용하여 Sign 생성
	vp := service.GenerateVP(holderDidInfo, issuerDidInfo, vc)

	// 3-1. vp 검증 테스트
	verifyVPResult := service.VerifyVP(&vp, holderDidInfo)
	fmt.Println("verify VP result : ", verifyVPResult)
}
