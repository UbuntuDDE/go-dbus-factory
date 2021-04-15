package power

import "errors"
import "fmt"
import "github.com/godbus/dbus"
import "pkg.deepin.io/lib/dbusutil"
import "pkg.deepin.io/lib/dbusutil/proxy"
import "unsafe"

/* prevent compile error */
var _ = errors.New
var _ dbusutil.SignalHandlerId
var _ = fmt.Sprintf
var _ unsafe.Pointer

type Power struct {
	power // interface com.deepin.system.Power
	proxy.Object
}

func NewPower(conn *dbus.Conn) *Power {
	obj := new(Power)
	obj.Object.Init_(conn, "com.deepin.system.Power", "/com/deepin/system/Power")
	return obj
}

type power struct{}

func (v *power) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*power) GetInterfaceName_() string {
	return "com.deepin.system.Power"
}

// method GetBatteries

func (v *power) GoGetBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetBatteries", flags, ch)
}

func (*power) StoreGetBatteries(call *dbus.Call) (batteries []dbus.ObjectPath, err error) {
	err = call.Store(&batteries)
	return
}

func (v *power) GetBatteries(flags dbus.Flags) (batteries []dbus.ObjectPath, err error) {
	return v.StoreGetBatteries(
		<-v.GoGetBatteries(flags, make(chan *dbus.Call, 1)).Done)
}

// method Refresh

func (v *power) GoRefresh(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Refresh", flags, ch)
}

func (v *power) Refresh(flags dbus.Flags) error {
	return (<-v.GoRefresh(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RefreshBatteries

func (v *power) GoRefreshBatteries(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshBatteries", flags, ch)
}

func (v *power) RefreshBatteries(flags dbus.Flags) error {
	return (<-v.GoRefreshBatteries(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method RefreshMains

func (v *power) GoRefreshMains(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RefreshMains", flags, ch)
}

func (v *power) RefreshMains(flags dbus.Flags) error {
	return (<-v.GoRefreshMains(flags, make(chan *dbus.Call, 1)).Done).Err
}

// method SetCpuGovernor

func (v *power) GoSetCpuGovernor(flags dbus.Flags, ch chan *dbus.Call, governor string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCpuGovernor", flags, ch, governor)
}

func (v *power) SetCpuGovernor(flags dbus.Flags, governor string) error {
	return (<-v.GoSetCpuGovernor(flags, make(chan *dbus.Call, 1), governor).Done).Err
}

// method SetCpuBoost

func (v *power) GoSetCpuBoost(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetCpuBoost", flags, ch, enabled)
}

func (v *power) SetCpuBoost(flags dbus.Flags, enabled bool) error {
	return (<-v.GoSetCpuBoost(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

// method LockCpuFreq

func (v *power) GoLockCpuFreq(flags dbus.Flags, ch chan *dbus.Call, governor string, lockTime int32) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LockCpuFreq", flags, ch, governor, lockTime)
}

func (v *power) LockCpuFreq(flags dbus.Flags, governor string, lockTime int32) error {
	return (<-v.GoLockCpuFreq(flags, make(chan *dbus.Call, 1), governor, lockTime).Done).Err
}

// signal BatteryDisplayUpdate

func (v *power) ConnectBatteryDisplayUpdate(cb func(timestamp int64)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BatteryDisplayUpdate", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BatteryDisplayUpdate",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var timestamp int64
		err := dbus.Store(sig.Body, &timestamp)
		if err == nil {
			cb(timestamp)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BatteryAdded

func (v *power) ConnectBatteryAdded(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BatteryAdded", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BatteryAdded",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal BatteryRemoved

func (v *power) ConnectBatteryRemoved(cb func(path dbus.ObjectPath)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "BatteryRemoved", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".BatteryRemoved",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var path dbus.ObjectPath
		err := dbus.Store(sig.Body, &path)
		if err == nil {
			cb(path)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal LidClosed

func (v *power) ConnectLidClosed(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LidClosed", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LidClosed",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal LidOpened

func (v *power) ConnectLidOpened(cb func()) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "LidOpened", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".LidOpened",
	}
	handlerFunc := func(sig *dbus.Signal) {
		cb()
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal PowerActionCode

func (v *power) ConnectPowerActionCode(cb func(actionCode int32)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "PowerActionCode", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".PowerActionCode",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var actionCode int32
		err := dbus.Store(sig.Body, &actionCode)
		if err == nil {
			cb(actionCode)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property PowerSavingModeAuto b

func (v *power) PowerSavingModeAuto() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PowerSavingModeAuto",
	}
}

// property OnBattery b

func (v *power) OnBattery() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "OnBattery",
	}
}

// property HasLidSwitch b

func (v *power) HasLidSwitch() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasLidSwitch",
	}
}

// property BatteryPercentage d

func (v *power) BatteryPercentage() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "BatteryPercentage",
	}
}

// property BatteryTimeToEmpty t

func (v *power) BatteryTimeToEmpty() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "BatteryTimeToEmpty",
	}
}

// property HasBattery b

func (v *power) HasBattery() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HasBattery",
	}
}

// property BatteryStatus u

func (v *power) BatteryStatus() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "BatteryStatus",
	}
}

// property BatteryTimeToFull t

func (v *power) BatteryTimeToFull() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "BatteryTimeToFull",
	}
}

// property BatteryCapacity d

func (v *power) BatteryCapacity() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "BatteryCapacity",
	}
}

// property PowerSavingModeEnabled b

func (v *power) PowerSavingModeEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PowerSavingModeEnabled",
	}
}

// property PowerSavingModeAutoWhenBatteryLow b

func (v *power) PowerSavingModeAutoWhenBatteryLow() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PowerSavingModeAutoWhenBatteryLow",
	}
}

// property PowerSavingModeBrightnessDropPercent u

func (v *power) PowerSavingModeBrightnessDropPercent() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "PowerSavingModeBrightnessDropPercent",
	}
}

// property CpuGovernor s

func (v *power) CpuGovernor() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "CpuGovernor",
	}
}

// property CpuBoost b

func (v *power) CpuBoost() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CpuBoost",
	}
}

// property IsBoostSupported b

func (v *power) IsBoostSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsBoostSupported",
	}
}

type Battery struct {
	battery // interface com.deepin.system.Power.Battery
	proxy.Object
}

func NewBattery(conn *dbus.Conn, path dbus.ObjectPath) (*Battery, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Battery)
	obj.Object.Init_(conn, "com.deepin.system.Power", path)
	return obj, nil
}

type battery struct{}

func (v *battery) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*battery) GetInterfaceName_() string {
	return "com.deepin.system.Power.Battery"
}

// property EnergyFullDesign d

func (v *battery) EnergyFullDesign() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "EnergyFullDesign",
	}
}

// property Capacity d

func (v *battery) Capacity() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Capacity",
	}
}

// property Technology s

func (v *battery) Technology() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Technology",
	}
}

// property Energy d

func (v *battery) Energy() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Energy",
	}
}

// property EnergyFull d

func (v *battery) EnergyFull() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "EnergyFull",
	}
}

// property Manufacturer s

func (v *battery) Manufacturer() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Manufacturer",
	}
}

// property ModelName s

func (v *battery) ModelName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ModelName",
	}
}

// property TimeToEmpty t

func (v *battery) TimeToEmpty() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeToEmpty",
	}
}

// property IsPresent b

func (v *battery) IsPresent() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsPresent",
	}
}

// property Status u

func (v *battery) Status() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Status",
	}
}

// property EnergyRate d

func (v *battery) EnergyRate() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "EnergyRate",
	}
}

// property Voltage d

func (v *battery) Voltage() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Voltage",
	}
}

// property Percentage d

func (v *battery) Percentage() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Percentage",
	}
}

// property TimeToFull t

func (v *battery) TimeToFull() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeToFull",
	}
}

// property UpdateTime x

func (v *battery) UpdateTime() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "UpdateTime",
	}
}

// property SysfsPath s

func (v *battery) SysfsPath() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SysfsPath",
	}
}

// property SerialNumber s

func (v *battery) SerialNumber() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SerialNumber",
	}
}

// property Name s

func (v *battery) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}
