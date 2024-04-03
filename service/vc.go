package service

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"didSample/key"
	"didSample/model"
	"didSample/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateVCJson(data interface{}) {

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", " ")

	err = ioutil.WriteFile("./VCredential.json", out.Bytes(), os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
	}
}

func GenerateVC(holderDID *model.DidInfo, issuerDID *model.DidInfo, claims *model.Claims) model.VCInfo {
	/*
		블록체인에서 DID와 DID Document 를 호출하였다고 가정하에 VC 생성 프로세스 진행
	*/
	vc := &model.VCInfo{}

	vc.Context = []string{"https://www.w3.org/2018/credentials/v1"}
	vc.Type = []string{"VerifiableCredential"}
	vc.Issuer = issuerDID.DID
	vc.Version = "1.0"
	vc.IssuanceDate = util.GetCurrentDate()
	vc.ExpirationDate = util.GetExpireDate() // 220518 수정 필요 expiredate 임의설정 중
	// Document의 Claim을 VC의 CREDENTIAL 형태로 변경
	vc.CredentialSubject.Id = holderDID.DID
	vc.CredentialSubject.Name = claims.Name
	vc.CredentialSubject.Photo = "Photo Link"
	vc.CredentialSubject.Identifier = claims.Identifier
	vc.CredentialSubject.Company = "BU"
	vc.CredentialSubject.Department = "Blockchain Center"
	vc.CredentialSubject.Position = "Researcher"
	vc.CredentialSubject.Status = "Working"
	vc.CredentialSubject.Email = claims.Email
	vc.CredentialSubject.Telephone = claims.Telephone
	vc.CredentialSubject.Address = claims.Address
	vc.CredentialSubject.Gender = "Female"
	vc.CredentialSubject.Created = util.GetCurrentDate()
	vc.CredentialSubject.Updated = "None"

	// proof 를 위한 별도의 서명 기능 상세 설계 필요
	vc.Proof.ProofValue = "sign for proof"
	vc.Proof.Created = util.GetCurrentDate()
	vc.Proof.Proofpurpose = "authentication"
	vc.Proof.Type = "ECDSA"
	vc.Proof.VerificationMethod = issuerDID.DID + "#key-1"

	// 서명 (VC Proof의 내용을 사용하여)
	r, s, _ := singVC(issuerDID, vc)

	vc.Signature.R = r
	vc.Signature.S = s

	CreateVCJson(vc)
	return *vc
}

func singVC(didInfo *model.DidInfo, vc *model.VCInfo) ([]byte, []byte, error) {
	vcProofBytes, err := json.Marshal(vc.Proof)
	if err != nil {
		return nil, nil, err
	}

	// VC의 SHA-256 해시 계산
	hash := sha256.Sum256(vcProofBytes)

	// 개인키로 VC의 해시에 서명
	prvKey, _ := key.ByteToECDSA(didInfo)
	r, s, err := ecdsa.Sign(rand.Reader, prvKey, hash[:])
	if err != nil {
		return nil, nil, err
	}
	// 서명된 데이터 반환
	rBytes, sBytes := r.Bytes(), s.Bytes()
	return rBytes, sBytes, nil
}
