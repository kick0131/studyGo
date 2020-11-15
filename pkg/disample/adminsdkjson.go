package disample

import (
	"fmt"
)

// AdminsdkjsonInfo はGCPのFirebaseAdminSDKのjsonファイルを表す構造体です
type AdminsdkjsonInfo struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

// Adminsdkjson はサービス本体です
type Adminsdkjson struct {
	servicea ServiceInterfaceA
	serviceb ServiceInterfaceB
}

// AdminsdkService はサービス利用者に公開するインタフェースです
type AdminsdkService interface {
	MappingToStructFromJSON(filepath string) (*AdminsdkjsonInfo, error)
}

// MappingToStructFromJSON は外部から注入されたインタフェース実装を使い、XXXを実現します。
func (internal Adminsdkjson) MappingToStructFromJSON(path string) (*AdminsdkjsonInfo, error) {
	// ToDo サービスインタフェースを使った処理
	fmt.Printf("arg is : %s\n", path)
	internal.servicea.SampleA()
	object, err := internal.serviceb.SampleB(path)

	return object, err
}

// ServiceInterfaceA は外部から注入されるインタフェース定義です
type ServiceInterfaceA interface {
	SampleA() error
}

// ServiceInterfaceB は外部から注入されるインタフェース定義です
type ServiceInterfaceB interface {
	SampleB(string) (*AdminsdkjsonInfo, error)
}

// BuildAdminsdkjsonService はサービスビルダーメソッドです
func BuildAdminsdkjsonService(a ServiceInterfaceA, b ServiceInterfaceB) AdminsdkService {
	return Adminsdkjson{
		a,
		b,
	}
}
