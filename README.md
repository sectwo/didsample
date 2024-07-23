DID/DID Document Example

1. DID와 Document 저장을 위한 별도의 블록체인은 존재한다고 가정한다.
2. DID와 키 관리를 위한 별도의 DB와 KMS가 존재한다고 가정한다.
3. DID, Document, VC, VP 그리고 Verfier 의 내용과 과정은 간단한 Usecase를 기반으로 하며, 세부 상세 내용들은 Legacy system에 따라 언제든지 변경 가능하다.
4. W3C 기반의 표준화를 따른다.


실행 과정
1. Holder 와 Issuer 의 DID와 DID Document 생성 및 저장
   1) ECDSA 기반의 Key pair 생성
   2) Key 구조체에 적합하도록 X509 마샬링
   3) 저장
2. VC 생성
   1) Credential 생성을 위한 Sample Credential 생성(Using claim)
   2) Issuer의 Private key를 사용하여 Credential 서명 및 VC 생성
   3) 서명과 검증 확인을 위하여 VC 검증 실행 및 확인
3. VP 생성
   1) Presentation 생성을 위하여 selective disclosure(선택적 정보 공개)와 같은 방법을 통하여 자신 혹은 정책에서 필요한 Claim의 선택적 공개 기능 실행
   2) 1) 에서의 예제를 위하여 CredentialSubject의 사용자 ID 만을 선택하여 정보 공개
   3) 이후 Presentation 생성 및 Holder의 Private key를 사용하여 서명을 통한 VP 생성 
   2) 서명과 검증 확인을 위하여 VP 검증 실행 및 확인

추가 
- 서명 테스트를 위하여 Proof의 내용만을 기준으로 서명을 진행
- Proof JSON 마샬링 이후 SHA-256을 사용하여 해싱 진행
- vC의 경우 Holder의 개인키를 통하여 생성
- Issuer의 개인키를 통하여 서명을 진행하였으며, ECDSA 알고리즘을 사용하여 서명과 검증에 필요한 r, s 값 생성
- VC와 VP는 서명시 사용된 키쌍중 공개키를 사용하여 검증 진행 
