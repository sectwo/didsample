package model

type DidInfo struct {
	PrvKey  []byte
	PubKey  []byte
	DID     string
	Created string
}

type DocumentInfo struct {
	Context        []string  `json:"@context"` //22.05.19_sujin : 추가
	Id             string    `json:"id"`
	Created        string    `json:"created"`        //(Unix Timestamp string)
	Authentication []string  `json:"authentication"` //(limit 지정해줘야될듯)
	Service        []Service `json:"service"`        //(structure)
}

type Service struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

type Signature struct {
	R []byte `json:"r"`
	S []byte `json:"s"`
}

type VCInfo struct {
	Context           []string   `json:"@context"`
	Type              []string   `json:"type"`
	Version           string     `json:"version"`
	Issuer            string     `json:"issuer"`
	IssuanceDate      string     `json:"issuanceDate"`        //(Unix Timestamp string)
	ExpirationDate    string     `json:"expirationDate"`      //(Unix Timestamp string)
	CredentialSubject Credential `json:"credentialSubject"`   //(structure)
	Signature         Signature  `json:"signature,omitempty"` // 서명은 옵셔널 필드
	Proof             Proof      `json:"proof"`               //(structure)
}

type VPInfo struct {
	Context              []string  `json:"@context"`
	Type                 []string  `json:"type"`
	Presenter            string    `json:"issuer"`
	Created              string    `json:"created"` //(Unix Timestamp string)
	VerifiableCredential VCInfo    `json:"verifiableCredential"`
	Signature            Signature `json:"signature,omitempty"` // 서명은 옵셔널 필드
	Proof                Proof     `json:"proof"`               //(structure)
}

type Credential struct {
	Id         string `json:"id"`         //(DID)
	Name       string `json:"name"`       //(이름)
	Photo      string `json:"photo"`      //(증명사진)
	Identifier string `json:"identifier"` //(사원번호)
	Company    string `json:"company"`    //(회사명)
	Department string `json:"department"` //(부서)
	Position   string `json:"position"`   //(직급)
	Status     string `json:"status"`     //(재직상태)
	Email      string `json:"email"`      //(이메일주소)
	Telephone  string `json:"telephone"`  //(전화번호)
	Address    string `json:"address"`    //(주소)
	Gender     string `json:"gender"`     //(성별)
	Created    string `json:"created"`    //(VC 생성시간)
	Updated    string `json:"updated"`    //(VC 업데이트시간)
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	Proofpurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"` //(소유자DID#keys-1 ...)
	ProofValue         string `json:"proofValue"`         //(서명값)
}

type Claims struct {
	Name        string `json:"name"`        //(이름)
	Identifier  string `json:"identifier"`  //(사번)
	Telephone   string `json:"telephone"`   //(전화번호)
	Address     string `json:"address"`     //(주소)
	Email       string `json:"email"`       //(이메일주소)
	Description string `json:"description"` //(주민등록번호)
}
