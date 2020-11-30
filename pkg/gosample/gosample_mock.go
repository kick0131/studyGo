package gosample

import "github.com/stretchr/testify/mock"

// MockGoSample gosample.goのモックインターフェース定義です
type MockGoSample struct {
	mock.Mock
}

// SubFuncHello はモック化メソッドです
func (m *MockGoSample) SubFuncHello() string {
	ret := m.Called()
	return ret.Get(0).(string)
}
