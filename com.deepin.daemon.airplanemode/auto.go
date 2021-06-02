package airplanemode

import "context"
import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type AirplaneMode struct {
	airplaneMode // interface com.deepin.daemon.AirplaneMode
	proxy.Object
}

func NewAirplaneMode(conn *dbus.Conn) *AirplaneMode {
	obj := new(AirplaneMode)
	obj.Object.Init_(conn, "com.deepin.daemon.AirplaneMode", "/com/deepin/daemon/AirplaneMode")
	return obj
}

type airplaneMode struct{}

func (v *airplaneMode) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*airplaneMode) GetInterfaceName_() string {
	return "com.deepin.daemon.AirplaneMode"
}

// method DumpState

func (v *airplaneMode) GoDumpState(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DumpState", flags, ch)
}

func (v *airplaneMode) GoDumpStateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DumpState", flags, ch)
}

func (v *airplaneMode) DumpState(flags dbus.Flags) error {
	return (<-v.GoDumpState(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *airplaneMode) DumpStateWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDumpStateWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Enable

func (v *airplaneMode) GoEnable(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Enable", flags, ch, enabled)
}

func (v *airplaneMode) GoEnableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Enable", flags, ch, enabled)
}

func (v *airplaneMode) Enable(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnable(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *airplaneMode) EnableWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnableBluetooth

func (v *airplaneMode) GoEnableBluetooth(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableBluetooth", flags, ch, enabled)
}

func (v *airplaneMode) GoEnableBluetoothWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableBluetooth", flags, ch, enabled)
}

func (v *airplaneMode) EnableBluetooth(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableBluetooth(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *airplaneMode) EnableBluetoothWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableBluetoothWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RestartBluetooth

func (v *airplaneMode) GoRestartBluetooth(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestartBluetooth", flags, ch, enabled)
}

func (v *airplaneMode) GoRestartBluetoothWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RestartBluetooth", flags, ch, enabled)
}

func (v *airplaneMode) RestartBluetooth(flags dbus.Flags, enabled bool) error {
	return (<-v.GoRestartBluetooth(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *airplaneMode) RestartBluetoothWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRestartBluetoothWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnableWifi

func (v *airplaneMode) GoEnableWifi(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableWifi", flags, ch, enabled)
}

func (v *airplaneMode) GoEnableWifiWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableWifi", flags, ch, enabled)
}

func (v *airplaneMode) EnableWifi(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableWifi(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *airplaneMode) EnableWifiWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableWifiWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal AirplaneOn

func (v *airplaneMode) ConnectAirplaneOn(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "AirplaneOn", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".AirplaneOn",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Enabled b

func (v *airplaneMode) Enabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Enabled",
	}
}

// property WifiEnabled b

func (v *airplaneMode) WifiEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WifiEnabled",
	}
}

// property BluetoothEnabled b

func (v *airplaneMode) BluetoothEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "BluetoothEnabled",
	}
}
