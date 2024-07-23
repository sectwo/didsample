package service

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"didSample/key"
	"didSample/model"
	"encoding/json"
	"fmt"
	"math/big"
)

func VerifyVC(vc *model.VCInfo, issuer *model.DidInfo) bool {
	fmt.Println("VC(Verifiable Credential) Verifying...")
	vcJson, _ := json.Marshal(vc.Proof)
	// VC의 SHA-256 해시 계산
	hash := sha256.Sum256(vcJson)

	// 서명의 R과 S 값을 big.Int로 변환
	r := new(big.Int).SetBytes(vc.Signature.R)
	s := new(big.Int).SetBytes(vc.Signature.S)

	_, issuerPubKey := key.ByteToECDSA(issuer)

	fmt.Println("VC(Verifiable Credential) Verify Success!!")
	return ecdsa.Verify(issuerPubKey, hash[:], r, s)
}

func VerifyVP(vp *model.VPInfo, holder *model.DidInfo) bool {
	fmt.Println("VP(Verifiable Presentation) Verifying...")
	vcJson, _ := json.Marshal(vp.Proof)
	// VC의 SHA-256 해시 계산
	hash := sha256.Sum256(vcJson)

	// 서명의 R과 S 값을 big.Int로 변환
	r := new(big.Int).SetBytes(vp.Signature.R)
	s := new(big.Int).SetBytes(vp.Signature.S)

	_, issuerPubKey := key.ByteToECDSA(holder)

	fmt.Println("VP(Verifiable Presentation) Verify Success!!")
	return ecdsa.Verify(issuerPubKey, hash[:], r, s)
}
