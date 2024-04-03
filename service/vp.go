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

func CreateVPJson(data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", " ")

	err = ioutil.WriteFile("./VPresentation.json", out.Bytes(), os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
	}
}

func generateProof(holderDID string) model.Proof {
	proof := model.Proof{
		ProofValue:         "sing for VP using Holder Public Key",
		Created:            util.GetCurrentDate(),
		Proofpurpose:       "authentication",
		Type:               "ECDSA",
		VerificationMethod: holderDID + "key-1",
	}

	return proof
}

func createSamplePresentation(vc model.VCInfo) model.VCInfo {
	credential := &model.Credential{
		Id: vc.CredentialSubject.Id,
	}
	presentation := &model.VCInfo{
		CredentialSubject: *credential,
		Proof:             vc.Proof,
	}

	return *presentation
}

func GenerateVP(holderDID *model.DidInfo, issuerDID *model.DidInfo, vc model.VCInfo) model.VPInfo {
	presentation := createSamplePresentation(vc)

	vp := &model.VPInfo{
		Context:              vc.Context,
		Created:              util.GetCurrentDate(),
		Type:                 []string{"VerifiablePresentation"},
		Presenter:            vc.CredentialSubject.Id,
		VerifiableCredential: presentation,
		Proof:                generateProof(holderDID.DID),
	}

	// 서명 (VC Proof의 내용을 사용하여)
	r, s, _ := singVP(holderDID, vp)

	fmt.Println("vp sign")
	fmt.Println("r : ", r)
	fmt.Println("s : ", s)

	vp.Signature.R = r
	vp.Signature.S = s

	CreateVPJson(vp)

	return *vp
}

func singVP(didInfo *model.DidInfo, vp *model.VPInfo) ([]byte, []byte, error) {
	vpProofBytes, err := json.Marshal(vp.Proof)
	if err != nil {
		return nil, nil, err
	}

	// VC의 SHA-256 해시 계산
	hash := sha256.Sum256(vpProofBytes)

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
