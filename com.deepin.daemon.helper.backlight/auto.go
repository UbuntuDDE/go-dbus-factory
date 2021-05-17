package backlight

import "context"
import "errors"
import "fmt"
import dbus "pkg.deepin.io/lib/dbus1"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "time"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Backlight struct {
	backlight // interface com.deepin.daemon.helper.Backlight
	proxy.Object
}

func NewBacklight(conn *dbus.Conn) *Backlight {
	obj := new(Backlight)
	obj.Object.Init_(conn, "com.deepin.daemon.helper.Backlight", "/com/deepin/daemon/helper/Backlight")
	return obj
}

type backlight struct{}

func (v *backlight) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*backlight) GetInterfaceName_() string {
	return "com.deepin.daemon.helper.Backlight"
}

// method SetBrightness

func (v *backlight) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, type0 uint8, name string, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, type0, name, value)
}

func (v *backlight) GoSetBrightnessWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, type0 uint8, name string, value int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetBrightness", flags, ch, type0, name, value)
}

func (v *backlight) SetBrightness(flags dbus.Flags, type0 uint8, name string, value int32) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), type0, name, value).Done).Err
}

func (v *backlight) SetBrightnessWithTimeout(timeout time.Duration, flags dbus.Flags, type0 uint8, name string, value int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetBrightnessWithContext(ctx, flags, make(chan *dbus.Call, 1), type0, name, value).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

type DDCCI struct {
	ddcci // interface com.deepin.daemon.helper.Backlight.DDCCI
	proxy.Object
}

func NewDDCCI(conn *dbus.Conn) *DDCCI {
	obj := new(DDCCI)
	obj.Object.Init_(conn, "com.deepin.daemon.helper.Backlight", "/com/deepin/daemon/helper/Backlight/DDCCI")
	return obj
}

type ddcci struct{}

func (v *ddcci) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*ddcci) GetInterfaceName_() string {
	return "com.deepin.daemon.helper.Backlight.DDCCI"
}

// method CheckSupport

func (v *ddcci) GoCheckSupport(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CheckSupport", flags, ch, edidChecksum)
}

func (v *ddcci) GoCheckSupportWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CheckSupport", flags, ch, edidChecksum)
}

func (*ddcci) StoreCheckSupport(call *dbus.Call) (support bool, err error) {
	err = call.Store(&support)
	return
}

func (v *ddcci) CheckSupport(flags dbus.Flags, edidChecksum string) (support bool, err error) {
	return v.StoreCheckSupport(
		<-v.GoCheckSupport(flags, make(chan *dbus.Call, 1), edidChecksum).Done)
}

func (v *ddcci) CheckSupportWithTimeout(timeout time.Duration, flags dbus.Flags, edidChecksum string) (support bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCheckSupportWithContext(ctx, flags, make(chan *dbus.Call, 1), edidChecksum).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCheckSupport(call)
}

// method GetBrightness

func (v *ddcci) GoGetBrightness(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBrightness", flags, ch, edidChecksum)
}

func (v *ddcci) GoGetBrightnessWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, edidChecksum string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetBrightness", flags, ch, edidChecksum)
}

func (*ddcci) StoreGetBrightness(call *dbus.Call) (value int32, err error) {
	err = call.Store(&value)
	return
}

func (v *ddcci) GetBrightness(flags dbus.Flags, edidChecksum string) (value int32, err error) {
	return v.StoreGetBrightness(
		<-v.GoGetBrightness(flags, make(chan *dbus.Call, 1), edidChecksum).Done)
}

func (v *ddcci) GetBrightnessWithTimeout(timeout time.Duration, flags dbus.Flags, edidChecksum string) (value int32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetBrightnessWithContext(ctx, flags, make(chan *dbus.Call, 1), edidChecksum).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetBrightness(call)
}

// method SetBrightness

func (v *ddcci) GoSetBrightness(flags dbus.Flags, ch chan *dbus.Call, edidChecksum string, value int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetBrightness", flags, ch, edidChecksum, value)
}

func (v *ddcci) GoSetBrightnessWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, edidChecksum string, value int32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetBrightness", flags, ch, edidChecksum, value)
}

func (v *ddcci) SetBrightness(flags dbus.Flags, edidChecksum string, value int32) error {
	return (<-v.GoSetBrightness(flags, make(chan *dbus.Call, 1), edidChecksum, value).Done).Err
}

func (v *ddcci) SetBrightnessWithTimeout(timeout time.Duration, flags dbus.Flags, edidChecksum string, value int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetBrightnessWithContext(ctx, flags, make(chan *dbus.Call, 1), edidChecksum, value).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RefreshDisplays

func (v *ddcci) GoRefreshDisplays(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshDisplays", flags, ch)
}

func (v *ddcci) GoRefreshDisplaysWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RefreshDisplays", flags, ch)
}

func (v *ddcci) RefreshDisplays(flags dbus.Flags) error {
	return (<-v.GoRefreshDisplays(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *ddcci) RefreshDisplaysWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRefreshDisplaysWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
