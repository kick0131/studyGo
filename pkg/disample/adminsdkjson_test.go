package disample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type DummySampleA struct{}

// SampleA インタフェースの実装
func (DummySampleA) SampleA() error {
	fmt.Println("this is SampleA")
	return nil
}

// InjectionServiceInterfaceA は注入用のインタフェースを生成するメソッドです
func InjectionServiceInterfaceA() ServiceInterfaceA {
	return DummySampleA{}
}

type DummySampleB struct{}

// SampleB インタフェースの実装(ファイルパスから構造体にマッピング)
func (DummySampleB) SampleB(filepath string) (*AdminsdkjsonInfo, error) {
	fmt.Println("this is SampleB")
	raw, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var jsonData AdminsdkjsonInfo

	json.Unmarshal(raw, &jsonData)
	return &jsonData, nil

}

// InjectionServiceInterfaceB は注入用のインタフェースを生成するメソッドです
func InjectionServiceInterfaceB() ServiceInterfaceB {
	return DummySampleB{}
}

func TestDiSample(t *testing.T) {
	fmt.Println("=== CheckPoint WriteFsSubCollection 1")

	// 外部で定義したインタフェース実装を注入
	service := BuildAdminsdkjsonService(
		InjectionServiceInterfaceA(),
		InjectionServiceInterfaceB(),
	)
	obj, err := service.MappingToStructFromJSON("../../data/adminsdk.json")
	if err != nil {
		t.Errorf("sub-collection Create error\n")
	}

	fmt.Printf("get json data is : %s\n", obj)

}
