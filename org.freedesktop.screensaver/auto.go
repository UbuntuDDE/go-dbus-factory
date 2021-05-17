package screensaver

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

type ScreenSaver struct {
	screenSaver // interface org.freedesktop.ScreenSaver
	proxy.Object
}

func NewScreenSaver(conn *dbus.Conn) *ScreenSaver {
	obj := new(ScreenSaver)
	obj.Object.Init_(conn, "org.freedesktop.ScreenSaver", "/org/freedesktop/ScreenSaver")
	return obj
}

type screenSaver struct{}

func (v *screenSaver) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*screenSaver) GetInterfaceName_() string {
	return "org.freedesktop.ScreenSaver"
}

// method Inhibit

func (v *screenSaver) GoInhibit(flags dbus.Flags, ch chan *dbus.Call, name string, reason string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Inhibit", flags, ch, name, reason)
}

func (v *screenSaver) GoInhibitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, reason string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Inhibit", flags, ch, name, reason)
}

func (*screenSaver) StoreInhibit(call *dbus.Call) (cookie uint32, err error) {
	err = call.Store(&cookie)
	return
}

func (v *screenSaver) Inhibit(flags dbus.Flags, name string, reason string) (cookie uint32, err error) {
	return v.StoreInhibit(
		<-v.GoInhibit(flags, make(chan *dbus.Call, 1), name, reason).Done)
}

func (v *screenSaver) InhibitWithTimeout(timeout time.Duration, flags dbus.Flags, name string, reason string) (cookie uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoInhibitWithContext(ctx, flags, make(chan *dbus.Call, 1), name, reason).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreInhibit(call)
}

// method SetTimeout

func (v *screenSaver) GoSetTimeout(flags dbus.Flags, ch chan *dbus.Call, seconds uint32, interval uint32, blank bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetTimeout", flags, ch, seconds, interval, blank)
}

func (v *screenSaver) GoSetTimeoutWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, seconds uint32, interval uint32, blank bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetTimeout", flags, ch, seconds, interval, blank)
}

func (v *screenSaver) SetTimeout(flags dbus.Flags, seconds uint32, interval uint32, blank bool) error {
	return (<-v.GoSetTimeout(flags, make(chan *dbus.Call, 1), seconds, interval, blank).Done).Err
}

func (v *screenSaver) SetTimeoutWithTimeout(timeout time.Duration, flags dbus.Flags, seconds uint32, interval uint32, blank bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTimeoutWithContext(ctx, flags, make(chan *dbus.Call, 1), seconds, interval, blank).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SimulateUserActivity

func (v *screenSaver) GoSimulateUserActivity(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SimulateUserActivity", flags, ch)
}

func (v *screenSaver) GoSimulateUserActivityWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SimulateUserActivity", flags, ch)
}

func (v *screenSaver) SimulateUserActivity(flags dbus.Flags) error {
	return (<-v.GoSimulateUserActivity(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *screenSaver) SimulateUserActivityWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSimulateUserActivityWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UnInhibit

func (v *screenSaver) GoUnInhibit(flags dbus.Flags, ch chan *dbus.Call, cookie uint32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UnInhibit", flags, ch, cookie)
}

func (v *screenSaver) GoUnInhibitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, cookie uint32) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UnInhibit", flags, ch, cookie)
}

func (v *screenSaver) UnInhibit(flags dbus.Flags, cookie uint32) error {
	return (<-v.GoUnInhibit(flags, make(chan *dbus.Call, 1), cookie).Done).Err
}

func (v *screenSaver) UnInhibitWithTimeout(timeout time.Duration, flags dbus.Flags, cookie uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnInhibitWithContext(ctx, flags, make(chan *dbus.Call, 1), cookie).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal IdleOn

func (v *screenSaver) ConnectIdleOn(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IdleOn", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IdleOn",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal CycleActive

func (v *screenSaver) ConnectCycleActive(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "CycleActive", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".CycleActive",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal IdleOff

func (v *screenSaver) ConnectIdleOff(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "IdleOff", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".IdleOff",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
