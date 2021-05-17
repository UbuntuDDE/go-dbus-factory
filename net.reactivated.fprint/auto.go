package fprint

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

type Manager struct {
	manager // interface net.reactivated.Fprint.Manager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "net.reactivated.Fprint", "/net/reactivated/Fprint/Manager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "net.reactivated.Fprint.Manager"
}

// method GetDefaultDevice

func (v *manager) GoGetDefaultDevice(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (v *manager) GoGetDefaultDeviceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDefaultDevice", flags, ch)
}

func (*manager) StoreGetDefaultDevice(call *dbus.Call) (device dbus.ObjectPath, err error) {
	err = call.Store(&device)
	return
}

func (v *manager) GetDefaultDevice(flags dbus.Flags) (device dbus.ObjectPath, err error) {
	return v.StoreGetDefaultDevice(
		<-v.GoGetDefaultDevice(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) GetDefaultDeviceWithTimeout(timeout time.Duration, flags dbus.Flags) (device dbus.ObjectPath, err error) {
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

func (v *manager) GoGetDevices(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (v *manager) GoGetDevicesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetDevices", flags, ch)
}

func (*manager) StoreGetDevices(call *dbus.Call) (devices []dbus.ObjectPath, err error) {
	err = call.Store(&devices)
	return
}

func (v *manager) GetDevices(flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
	return v.StoreGetDevices(
		<-v.GoGetDevices(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) GetDevicesWithTimeout(timeout time.Duration, flags dbus.Flags) (devices []dbus.ObjectPath, err error) {
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

type Device struct {
	device // interface net.reactivated.Fprint.Device
	proxy.Object
}

func NewDevice(conn *dbus.Conn, path dbus.ObjectPath) (*Device, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Device)
	obj.Object.Init_(conn, "net.reactivated.Fprint", path)
	return obj, nil
}

type device struct{}

func (v *device) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*device) GetInterfaceName_() string {
	return "net.reactivated.Fprint.Device"
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

// method EnrollStart

func (v *device) GoEnrollStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnrollStart", flags, ch, finger_name)
}

func (v *device) GoEnrollStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnrollStart", flags, ch, finger_name)
}

func (v *device) EnrollStart(flags dbus.Flags, finger_name string) error {
	return (<-v.GoEnrollStart(flags, make(chan *dbus.Call, 1), finger_name).Done).Err
}

func (v *device) EnrollStartWithTimeout(timeout time.Duration, flags dbus.Flags, finger_name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnrollStartWithContext(ctx, flags, make(chan *dbus.Call, 1), finger_name).Done
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

// method VerifyStart

func (v *device) GoVerifyStart(flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".VerifyStart", flags, ch, finger_name)
}

func (v *device) GoVerifyStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, finger_name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".VerifyStart", flags, ch, finger_name)
}

func (v *device) VerifyStart(flags dbus.Flags, finger_name string) error {
	return (<-v.GoVerifyStart(flags, make(chan *dbus.Call, 1), finger_name).Done).Err
}

func (v *device) VerifyStartWithTimeout(timeout time.Duration, flags dbus.Flags, finger_name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoVerifyStartWithContext(ctx, flags, make(chan *dbus.Call, 1), finger_name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
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

// method DeleteEnrolledFinger

func (v *device) GoDeleteEnrolledFinger(flags dbus.Flags, ch chan *dbus.Call, username string, finger_name string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DeleteEnrolledFinger", flags, ch, username, finger_name)
}

func (v *device) GoDeleteEnrolledFingerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string, finger_name string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DeleteEnrolledFinger", flags, ch, username, finger_name)
}

func (v *device) DeleteEnrolledFinger(flags dbus.Flags, username string, finger_name string) error {
	return (<-v.GoDeleteEnrolledFinger(flags, make(chan *dbus.Call, 1), username, finger_name).Done).Err
}

func (v *device) DeleteEnrolledFingerWithTimeout(timeout time.Duration, flags dbus.Flags, username string, finger_name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteEnrolledFingerWithContext(ctx, flags, make(chan *dbus.Call, 1), username, finger_name).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ListEnrolledFingers

func (v *device) GoListEnrolledFingers(flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListEnrolledFingers", flags, ch, username)
}

func (v *device) GoListEnrolledFingersWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, username string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListEnrolledFingers", flags, ch, username)
}

func (*device) StoreListEnrolledFingers(call *dbus.Call) (enrolled_fingers []string, err error) {
	err = call.Store(&enrolled_fingers)
	return
}

func (v *device) ListEnrolledFingers(flags dbus.Flags, username string) (enrolled_fingers []string, err error) {
	return v.StoreListEnrolledFingers(
		<-v.GoListEnrolledFingers(flags, make(chan *dbus.Call, 1), username).Done)
}

func (v *device) ListEnrolledFingersWithTimeout(timeout time.Duration, flags dbus.Flags, username string) (enrolled_fingers []string, err error) {
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

// signal EnrollStatus

func (v *device) ConnectEnrollStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		var arg1 bool
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyStatus

func (v *device) ConnectVerifyStatus(cb func(arg0 string, arg1 bool)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		var arg1 bool
		err := dbus.Store(sig.Body, &arg0, &arg1)
		if err == nil {
			cb(arg0, arg1)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal VerifyFingerSelected

func (v *device) ConnectVerifyFingerSelected(cb func(arg0 string)) (dbusutil.SignalHandlerId, error) {
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
		var arg0 string
		err := dbus.Store(sig.Body, &arg0)
		if err == nil {
			cb(arg0)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property scan-type s

func (v *device) ScanType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "scan-type",
	}
}

// property num-enroll-stages i

func (v *device) NumEnrollStages() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "num-enroll-stages",
	}
}

// property name s

func (v *device) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "name",
	}
}
