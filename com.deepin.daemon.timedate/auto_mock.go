// Code generated by "./generator ./com.deepin.daemon.timedate"; DO NOT EDIT.

package timedate

import (
	"fmt"

	"github.com/stretchr/testify/mock"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/dbusutil/proxy"
)

type MockTimedate struct {
	MockInterfaceTimedate // interface com.deepin.daemon.Timedate
	proxy.MockObject
}

type MockInterfaceTimedate struct {
	mock.Mock
}

// signal TimeUpdate

func (v *MockInterfaceTimedate) ConnectTimeUpdate(cb func()) (dbusutil.SignalHandlerId, error) {
	mockArgs := v.Called(cb)

	ret0, ok := mockArgs.Get(0).(dbusutil.SignalHandlerId)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, mockArgs.Get(0)))
	}

	return ret0, mockArgs.Error(1)
}
