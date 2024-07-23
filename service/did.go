package service

import (
	"bytes"
	"crypto/sha256"
	"didSample/key"
	"didSample/model"
	"didSample/util"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func GenerateDID(uid string) (*model.DidInfo, *model.DocumentInfo) {
	log.Println("DID Generating...")
	// 1. KeyPair 생성
	didStr := "did:method:spec"
	keyPair := key.GenerateKeyPair(uid)
	//fmt.Println(keyPair)

	holder := &model.DidInfo{
		PrvKey:  keyPair.PrvKey,
		PubKey:  keyPair.PubKey,
		Created: util.GetCurrentDate(),
	}
	// 공개 키의 SHA-256 해시 계산
	hash := sha256.Sum256(keyPair.PubKey)

	// 해시를 hex 문자열로 인코딩
	hashHex := hex.EncodeToString(hash[:])

	didSpecific := hashHex
	didStr += didSpecific
	holder.DID = didStr

	document := GenerateDIDDocument(*holder)
	/*
		이후 블록체인에 등록과정 필요
	*/
	log.Println("DID Generate Success!!")
	return holder, document
}

func GenerateDIDDocument(holder model.DidInfo) *model.DocumentInfo {
	log.Println("DID Document Generating...")
	services := serviceSample(holder)
	newDocument := &model.DocumentInfo{
		Context:        []string{"https://www.w3.org/ns/did/v1"},
		Id:             holder.DID,
		Created:        util.GetCurrentDate(),
		Authentication: []string{fmt.Sprintf("%s#keys-1", holder.DID)},
		Service:        services.Service,
	}

	CreateDocumentJson(newDocument)
	log.Println("DID Document Generate Success!!")
	return newDocument
}

func serviceSample(holder model.DidInfo) *model.DocumentInfo {
	result := &model.DocumentInfo{}

	Service1 := &model.Service{}
	Service1.Id = fmt.Sprintf("%s#EmployeeAuth", holder.DID)
	Service1.Type = "AuthenticateEmployee"
	Service1.ServiceEndpoint = "https://did.sample.com/employeeAuth/"

	Service2 := &model.Service{}
	Service2.Id = fmt.Sprintf("%s#EmployeeList", holder.DID)
	Service2.Type = "employeeList"
	Service2.ServiceEndpoint = "https://did.sample.com/employeeList/"

	result.Service = append(result.Service, *Service1)
	result.Service = append(result.Service, *Service2)

	return result
}

// func CreateDocumentJson(docu *model.DocumentInfo) {
func CreateDocumentJson(data interface{}) {

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", " ")

	err = ioutil.WriteFile("./DIDDocument.json", out.Bytes(), os.FileMode(0644)) // articles.json 파일에 JSON 문서 저장
	if err != nil {
		fmt.Println(err)
	}
}
