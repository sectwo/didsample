package main

import (
	"didSample/service"
	"log"
)

func main() {
	log.Println("=================================== DID SAMPLE ===================================")

	// 0. Issuer DID 생성 및 등록
	issuerDidInfo, _ := service.GenerateDID("issuer")
	log.Println("issuer public key([]byte) : ", string(issuerDidInfo.PubKey))
	log.Println("issuer private key([]byte) : ", string(issuerDidInfo.PrvKey))
	log.Println("issuer did : ", issuerDidInfo.DID)
	log.Println()

	// 1. DID 및 DID Document 생성 및 등록
	holderDidInfo, _ := service.GenerateDID("holder")
	log.Println("holder public key([]byte) : ", string(holderDidInfo.PubKey))
	log.Println("holder private key([]byte) : ", string(holderDidInfo.PrvKey))
	log.Println("holder did : ", holderDidInfo.DID)
	log.Println()

	// 2. Holder와 Issuer DID를 기반으로 VC 생성 및 서명 생성
	//holderDID := holderDidInfo.DID
	claims := service.CreateClaimsSample()
	vc := service.GenerateVC(holderDidInfo, issuerDidInfo, claims)

	// 2-1. vc 검증 테스트
	verifyVCResult := service.VerifyVC(&vc, issuerDidInfo)
	log.Println("Verify VC result : ", verifyVCResult)
	log.Println()

	// 3. Holder 와 Issuer DID 및 VC를 기반으로 VP 생성 및 서명 생성
	// 이때 VP는 사용자 혹은 시스템에 필요한 포맷에 따라 Claims 의 선택 가능
	// VP 생성 시에는 Holder의 Public Key를 사용하여 Sign 생성
	vp := service.GenerateVP(holderDidInfo, issuerDidInfo, vc)

	// 3-1. vp 검증 테스트
	verifyVPResult := service.VerifyVP(&vp, holderDidInfo)
	log.Println("verify VP result : ", verifyVPResult)
}
