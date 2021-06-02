package sensorproxy

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

type SensorProxy struct {
	sensorProxy // interface net.hadess.SensorProxy
	proxy.Object
}

func NewSensorProxy(conn *dbus.Conn) *SensorProxy {
	obj := new(SensorProxy)
	obj.Object.Init_(conn, "net.hadess.SensorProxy", "/net/hadess/SensorProxy")
	return obj
}

type sensorProxy struct{}

func (v *sensorProxy) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*sensorProxy) GetInterfaceName_() string {
	return "net.hadess.SensorProxy"
}

// method ClaimAccelerometer

func (v *sensorProxy) GoClaimAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClaimAccelerometer", flags, ch)
}

func (v *sensorProxy) GoClaimAccelerometerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ClaimAccelerometer", flags, ch)
}

func (v *sensorProxy) ClaimAccelerometer(flags dbus.Flags) error {
	return (<-v.GoClaimAccelerometer(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sensorProxy) ClaimAccelerometerWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClaimAccelerometerWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ReleaseAccelerometer

func (v *sensorProxy) GoReleaseAccelerometer(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseAccelerometer", flags, ch)
}

func (v *sensorProxy) GoReleaseAccelerometerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReleaseAccelerometer", flags, ch)
}

func (v *sensorProxy) ReleaseAccelerometer(flags dbus.Flags) error {
	return (<-v.GoReleaseAccelerometer(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sensorProxy) ReleaseAccelerometerWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReleaseAccelerometerWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ClaimLight

func (v *sensorProxy) GoClaimLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ClaimLight", flags, ch)
}

func (v *sensorProxy) GoClaimLightWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ClaimLight", flags, ch)
}

func (v *sensorProxy) ClaimLight(flags dbus.Flags) error {
	return (<-v.GoClaimLight(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sensorProxy) ClaimLightWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoClaimLightWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method ReleaseLight

func (v *sensorProxy) GoReleaseLight(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ReleaseLight", flags, ch)
}

func (v *sensorProxy) GoReleaseLightWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ReleaseLight", flags, ch)
}

func (v *sensorProxy) ReleaseLight(flags dbus.Flags) error {
	return (<-v.GoReleaseLight(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *sensorProxy) ReleaseLightWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoReleaseLightWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property HasAccelerometer b

func (v *sensorProxy) HasAccelerometer() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasAccelerometer",
	}
}

// property AccelerometerOrientation s

func (v *sensorProxy) AccelerometerOrientation() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "AccelerometerOrientation",
	}
}

// property HasAmbientLight b

func (v *sensorProxy) HasAmbientLight() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasAmbientLight",
	}
}

// property LightLevelUnit s

func (v *sensorProxy) LightLevelUnit() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "LightLevelUnit",
	}
}

// property LightLevel d

func (v *sensorProxy) LightLevel() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "LightLevel",
	}
}
