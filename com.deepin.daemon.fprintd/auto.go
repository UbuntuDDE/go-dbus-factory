package fprintd

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

type Fprintd struct {
	fprintd // interface com.deepin.daemon.Fprintd
	proxy.Object
}

func NewFprintd(conn *dbus.Conn) *Fprintd {
	obj := new(Fprintd)
	obj.Object.Init_(conn, "com.deepin.daemon.Fprintd", "/com/deepin/daemon/Fprintd")
	return obj
}

type fprintd struct{}

func (v *fprintd) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*fprintd) GetInterfaceName_() string {
	return "com.deepin.daemon.Fprintd"
}

// method GetDefaultDevice

func (v *fprintd) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (v *fprintd) GoGetDefaultDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (*fprintd) StoreGetDefaultDevice(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *fprintd) GetDefaultDevice(flags dbus.Flags) (device dbus.ObjectPath, err error) {
	return v.StoreGetDefaultDevice(
		<-v.GoGetDefaultDevice(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *fprintd) GetDefaultDeviceWithTimeout(timeout time.Duration, flags dbus.Flags) (device dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetDefaultDeviceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetDefaultDevice(call)
}

// method GetDevices

func (v *fprintd) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (v *fprintd) GoGetDevicesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*fprintd) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *fprintd) GetDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *fprintd) GetDevicesWithTimeout(timeout time.Duration, flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetDevicesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetDevices(call)
}

// method TriggerUDevEvent

func (v *fprintd) GoTriggerUDevEvent(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".TriggerUDevEvent", flags, ch)
}

func (v *fprintd) GoTriggerUDevEventWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".TriggerUDevEvent", flags, ch)
}

func (v *fprintd) TriggerUDevEvent(flags dbus.Flags) error {
	return (<-v.GoTriggerUDevEvent(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *fprintd) TriggerUDevEventWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoTriggerUDevEventWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Devices ao

func (v *fprintd) Devices() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "Devices",
	}
}

type Device struct {
	device // interface com.deepin.daemon.Fprintd.Device
	proxy.Object
}

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (*Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Device)
	obj.Object.Init_(conn, "com.deepin.daemon.Fprintd", path)
	return obj, nil
}

type device struct{}

func (v *device) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*device) GetInterfaceName_() string {
	return "com.deepin.daemon.Fprintd.Device"
}

// method Claim

func (v *device) GoClaim(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Claim", flags, ch, username)
}

func (v *device) GoClaimWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Claim", flags, ch, username)
}

func (v *device) Claim(flags dbus.Flags, username string) error {
	return (<-v.GoClaim(flags, make(chan *dbus.Call, 1), username).Done).Err
}

func (v *device) ClaimWithTimeout(timeout time.Duration, flags dbus.Flags, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClaimWithContext(ctx, flags, make(chan *dbus.Call, 1), username).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ClaimForce

func (v *device) GoClaimForce(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClaimForce", flags, ch, username)
}

func (v *device) GoClaimForceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ClaimForce", flags, ch, username)
}

func (v *device) ClaimForce(flags dbus.Flags, username string) error {
	return (<-v.GoClaimForce(flags, make(chan *dbus.Call, 1), username).Done).Err
}

func (v *device) ClaimForceWithTimeout(timeout time.Duration, flags dbus.Flags, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClaimForceWithContext(ctx, flags, make(chan *dbus.Call, 1), username).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteEnrolledFinger

func (v *device) GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFinger", flags, ch, username, finger)
}

func (v *device) GoDeleteEnrolledFingerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string, finger string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteEnrolledFinger", flags, ch, username, finger)
}

func (v *device) DeleteEnrolledFinger(flags dbus.Flags, username string, finger string) error {
	return (<-v.GoDeleteEnrolledFinger(flags, make(chan *dbus.Call, 1), username, finger).Done).Err
}

func (v *device) DeleteEnrolledFingerWithTimeout(timeout time.Duration, flags dbus.Flags, username string, finger string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteEnrolledFingerWithContext(ctx, flags, make(chan *dbus.Call, 1), username, finger).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DeleteEnrolledFingers

func (v *device) GoDeleteEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFingers", flags, ch, username)
}

func (v *device) GoDeleteEnrolledFingersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteEnrolledFingers", flags, ch, username)
}

func (v *device) DeleteEnrolledFingers(flags dbus.Flags, username string) error {
	return (<-v.GoDeleteEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done).Err
}

func (v *device) DeleteEnrolledFingersWithTimeout(timeout time.Duration, flags dbus.Flags, username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteEnrolledFingersWithContext(ctx, flags, make(chan *dbus.Call, 1), username).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnrollStart

func (v *device) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStart", flags, ch, finger)
}

func (v *device) GoEnrollStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnrollStart", flags, ch, finger)
}

func (v *device) EnrollStart(flags dbus.Flags, finger string) error {
	return (<-v.GoEnrollStart(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

func (v *device) EnrollStartWithTimeout(timeout time.Duration, flags dbus.Flags, finger string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnrollStartWithContext(ctx, flags, make(chan *dbus.Call, 1), finger).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method EnrollStop

func (v *device) GoEnrollStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStop", flags, ch)
}

func (v *device) GoEnrollStopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnrollStop", flags, ch)
}

func (v *device) EnrollStop(flags dbus.Flags) error {
	return (<-v.GoEnrollStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) EnrollStopWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnrollStopWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetCapabilities

func (v *device) GoGetCapabilities(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetCapabilities", flags, ch)
}

func (v *device) GoGetCapabilitiesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetCapabilities", flags, ch)
}

func (*device) StoreGetCapabilities(call *dbus.Call) (caps []string, err error) {
	err = call.Store(&caps)
	return
}

func (v *device) GetCapabilities(flags dbus.Flags) (caps []string, err error) {
	return v.StoreGetCapabilities(
		<-v.GoGetCapabilities(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *device) GetCapabilitiesWithTimeout(timeout time.Duration, flags dbus.Flags) (caps []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetCapabilitiesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetCapabilities(call)
}

// method ListEnrolledFingers

func (v *device) GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListEnrolledFingers", flags, ch, username)
}

func (v *device) GoListEnrolledFingersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListEnrolledFingers", flags, ch, username)
}

func (*device) StoreListEnrolledFingers(call *dbus.Call) (fingers []string, err error) {
	err = call.Store(&fingers)
	return
}

func (v *device) ListEnrolledFingers(flags dbus.Flags, username string) (fingers []string, err error) {
	return v.StoreListEnrolledFingers(
		<-v.GoListEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done)
}

func (v *device) ListEnrolledFingersWithTimeout(timeout time.Duration, flags dbus.Flags, username string) (fingers []string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListEnrolledFingersWithContext(ctx, flags, make(chan *dbus.Call, 1), username).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListEnrolledFingers(call)
}

// method Release

func (v *device) GoRelease(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Release", flags, ch)
}

func (v *device) GoReleaseWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Release", flags, ch)
}

func (v *device) Release(flags dbus.Flags) error {
	return (<-v.GoRelease(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) ReleaseWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReleaseWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method VerifyStart

func (v *device) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStart", flags, ch, finger)
}

func (v *device) GoVerifyStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, finger string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".VerifyStart", flags, ch, finger)
}

func (v *device) VerifyStart(flags dbus.Flags, finger string) error {
	return (<-v.GoVerifyStart(flags, make(chan *dbus.Call, 1), finger).Done).Err
}

func (v *device) VerifyStartWithTimeout(timeout time.Duration, flags dbus.Flags, finger string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoVerifyStartWithContext(ctx, flags, make(chan *dbus.Call, 1), finger).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method VerifyStop

func (v *device) GoVerifyStop(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStop", flags, ch)
}

func (v *device) GoVerifyStopWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".VerifyStop", flags, ch)
}

func (v *device) VerifyStop(flags dbus.Flags) error {
	return (<-v.GoVerifyStop(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *device) VerifyStopWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoVerifyStopWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// signal EnrollStatus

func (v *device) ConnectEnrollStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "EnrollStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".EnrollStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var done bool
		err := dbus.Store(sig.Body, &status, &done)
		if err == nil {
			cb(status, done)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *device) ConnectVerifyStatus(cb func(status string, done bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyStatus", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyStatus",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var status string
		var done bool
		err := dbus.Store(sig.Body, &status, &done)
		if err == nil {
			cb(status, done)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyFingerSelected

func (v *device) ConnectVerifyFingerSelected(cb func(finger string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "VerifyFingerSelected", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".VerifyFingerSelected",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var finger string
		err := dbus.Store(sig.Body, &finger)
		if err == nil {
			cb(finger)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property ScanType s

func (v *device) ScanType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ScanType",
	}
}
