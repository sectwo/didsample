package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"didSample/model"
	"fmt"
)

type KeyPair struct {
	PrvKey []byte `json:"prvkey"`
	PubKey []byte `json:"pubkey"`
}

func GenerateKeyPair(uid string) *KeyPair {
	// 1. ECDSA 키 쌍 생성
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		fmt.Println("키 생성 중 오류 발생:", err)
		return nil
	}

	derPrivateKey, _ := x509.MarshalECPrivateKey(privateKey)
	derPublicKey, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)

	keyPair := &KeyPair{
		PrvKey: derPrivateKey,
		PubKey: derPublicKey,
	}

	return keyPair
}

func ByteToECDSA(didInfo *model.DidInfo) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	derPrvKey := didInfo.PrvKey
	derPubKey := didInfo.PubKey

	prvKey, _ := x509.ParseECPrivateKey(derPrvKey)
	pubKeyInterface, _ := x509.ParsePKIXPublicKey(derPubKey)
	pubKey := pubKeyInterface.(*ecdsa.PublicKey)

	return prvKey, pubKey
}
