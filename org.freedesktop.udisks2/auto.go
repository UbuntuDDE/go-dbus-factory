package udisks2

import "context"
import "errors"
import "fmt"
import "github.com/linuxdeepin/go-dbus-factory/object_manager"
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

type UDisks struct {
	object_manager.ObjectManager // interface org.freedesktop.DBus.ObjectManager
	proxy.Object
}

func NewUDisks(conn *dbus.Conn) *UDisks {
	obj := new(UDisks)
	obj.Object.Init_(conn, "org.freedesktop.UDisks2", "/org/freedesktop/UDisks2")
	return obj
}

type Manager struct {
	manager // interface org.freedesktop.UDisks2.Manager
	proxy.Object
}

func NewManager(conn *dbus.Conn) *Manager {
	obj := new(Manager)
	obj.Object.Init_(conn, "org.freedesktop.UDisks2", "/org/freedesktop/UDisks2/Manager")
	return obj
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.Manager"
}

// method LoopSetup

func (v *manager) GoLoopSetup(flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".LoopSetup", flags, ch, fd, options)
}

func (v *manager) GoLoopSetupWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, fd dbus.UnixFD, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".LoopSetup", flags, ch, fd, options)
}

func (*manager) StoreLoopSetup(call *dbus.Call) (resulting_device dbus.ObjectPath, err error) {
	err = call.Store(&resulting_device)
	return
}

func (v *manager) LoopSetup(flags dbus.Flags, fd dbus.UnixFD, options map[string]dbus.Variant) (resulting_device dbus.ObjectPath, err error) {
	return v.StoreLoopSetup(
		<-v.GoLoopSetup(flags, make(chan *dbus.Call, 1), fd, options).Done)
}

func (v *manager) LoopSetupWithTimeout(timeout time.Duration, flags dbus.Flags, fd dbus.UnixFD, options map[string]dbus.Variant) (resulting_device dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoLoopSetupWithContext(ctx, flags, make(chan *dbus.Call, 1), fd, options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreLoopSetup(call)
}

// method MDRaidCreate

func (v *manager) GoMDRaidCreate(flags dbus.Flags, ch chan *dbus.Call, blocks []dbus.ObjectPath, level string, name string, chunk uint64, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".MDRaidCreate", flags, ch, blocks, level, name, chunk, options)
}

func (v *manager) GoMDRaidCreateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, blocks []dbus.ObjectPath, level string, name string, chunk uint64, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".MDRaidCreate", flags, ch, blocks, level, name, chunk, options)
}

func (*manager) StoreMDRaidCreate(call *dbus.Call) (resulting_array dbus.ObjectPath, err error) {
	err = call.Store(&resulting_array)
	return
}

func (v *manager) MDRaidCreate(flags dbus.Flags, blocks []dbus.ObjectPath, level string, name string, chunk uint64, options map[string]dbus.Variant) (resulting_array dbus.ObjectPath, err error) {
	return v.StoreMDRaidCreate(
		<-v.GoMDRaidCreate(flags, make(chan *dbus.Call, 1), blocks, level, name, chunk, options).Done)
}

func (v *manager) MDRaidCreateWithTimeout(timeout time.Duration, flags dbus.Flags, blocks []dbus.ObjectPath, level string, name string, chunk uint64, options map[string]dbus.Variant) (resulting_array dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMDRaidCreateWithContext(ctx, flags, make(chan *dbus.Call, 1), blocks, level, name, chunk, options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreMDRaidCreate(call)
}

// property Version s

func (v *manager) Version() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Version",
	}
}

type Drive struct {
	drive    // interface org.freedesktop.UDisks2.Drive
	driveAta // interface org.freedesktop.UDisks2.Drive.Ata
	proxy.Object
}

func NewDrive(conn *dbus.Conn, path dbus.ObjectPath) (*Drive, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Drive)
	obj.Object.Init_(conn, "org.freedesktop.UDisks2", path)
	return obj, nil
}

func (obj *Drive) Drive() *drive {
	return &obj.drive
}

type drive struct{}

func (v *drive) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*drive) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.Drive"
}

// method Eject

func (v *drive) GoEject(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Eject", flags, ch, options)
}

func (v *drive) GoEjectWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Eject", flags, ch, options)
}

func (v *drive) Eject(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoEject(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *drive) EjectWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEjectWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetConfiguration

func (v *drive) GoSetConfiguration(flags dbus.Flags, ch chan *dbus.Call, value map[string]dbus.Variant, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetConfiguration", flags, ch, value, options)
}

func (v *drive) GoSetConfigurationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, value map[string]dbus.Variant, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetConfiguration", flags, ch, value, options)
}

func (v *drive) SetConfiguration(flags dbus.Flags, value map[string]dbus.Variant, options map[string]dbus.Variant) error {
	return (<-v.GoSetConfiguration(flags, make(chan *dbus.Call, 1), value, options).Done).Err
}

func (v *drive) SetConfigurationWithTimeout(timeout time.Duration, flags dbus.Flags, value map[string]dbus.Variant, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetConfigurationWithContext(ctx, flags, make(chan *dbus.Call, 1), value, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PowerOff

func (v *drive) GoPowerOff(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PowerOff", flags, ch, options)
}

func (v *drive) GoPowerOffWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PowerOff", flags, ch, options)
}

func (v *drive) PowerOff(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoPowerOff(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *drive) PowerOffWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPowerOffWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Vendor s

func (v *drive) Vendor() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Vendor",
	}
}

// property Model s

func (v *drive) Model() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Model",
	}
}

// property Revision s

func (v *drive) Revision() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Revision",
	}
}

// property Serial s

func (v *drive) Serial() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Serial",
	}
}

// property WWN s

func (v *drive) WWN() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "WWN",
	}
}

// property Id s

func (v *drive) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Configuration a{sv}

func (v *drive) Configuration() PropDriveConfiguration {
	return PropDriveConfiguration{
		Impl: v,
	}
}

type PropDriveConfiguration struct {
	Impl proxy.Implementer
}

func (p PropDriveConfiguration) Get(flags dbus.Flags) (value map[string]dbus.Variant, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Configuration", &value)
	return
}

func (p PropDriveConfiguration) ConnectChanged(cb func(hasValue bool, value map[string]dbus.Variant)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]dbus.Variant
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Configuration", cb0)
}

// property Media s

func (v *drive) Media() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Media",
	}
}

// property MediaCompatibility as

func (v *drive) MediaCompatibility() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "MediaCompatibility",
	}
}

// property MediaRemovable b

func (v *drive) MediaRemovable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MediaRemovable",
	}
}

// property MediaAvailable b

func (v *drive) MediaAvailable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MediaAvailable",
	}
}

// property MediaChangeDetected b

func (v *drive) MediaChangeDetected() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "MediaChangeDetected",
	}
}

// property Size t

func (v *drive) Size() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Size",
	}
}

// property TimeDetected t

func (v *drive) TimeDetected() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeDetected",
	}
}

// property TimeMediaDetected t

func (v *drive) TimeMediaDetected() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "TimeMediaDetected",
	}
}

// property Optical b

func (v *drive) Optical() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Optical",
	}
}

// property OpticalBlank b

func (v *drive) OpticalBlank() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "OpticalBlank",
	}
}

// property OpticalNumTracks u

func (v *drive) OpticalNumTracks() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "OpticalNumTracks",
	}
}

// property OpticalNumAudioTracks u

func (v *drive) OpticalNumAudioTracks() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "OpticalNumAudioTracks",
	}
}

// property OpticalNumDataTracks u

func (v *drive) OpticalNumDataTracks() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "OpticalNumDataTracks",
	}
}

// property OpticalNumSessions u

func (v *drive) OpticalNumSessions() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "OpticalNumSessions",
	}
}

// property RotationRate i

func (v *drive) RotationRate() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "RotationRate",
	}
}

// property ConnectionBus s

func (v *drive) ConnectionBus() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "ConnectionBus",
	}
}

// property Seat s

func (v *drive) Seat() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Seat",
	}
}

// property Removable b

func (v *drive) Removable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Removable",
	}
}

// property Ejectable b

func (v *drive) Ejectable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Ejectable",
	}
}

// property SortKey s

func (v *drive) SortKey() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SortKey",
	}
}

// property CanPowerOff b

func (v *drive) CanPowerOff() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "CanPowerOff",
	}
}

// property SiblingId s

func (v *drive) SiblingId() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SiblingId",
	}
}

func (obj *Drive) DriveAta() *driveAta {
	return &obj.driveAta
}

type driveAta struct{}

func (v *driveAta) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*driveAta) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.Drive.Ata"
}

// method SmartUpdate

func (v *driveAta) GoSmartUpdate(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SmartUpdate", flags, ch, options)
}

func (v *driveAta) GoSmartUpdateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SmartUpdate", flags, ch, options)
}

func (v *driveAta) SmartUpdate(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoSmartUpdate(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *driveAta) SmartUpdateWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSmartUpdateWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SmartGetAttributes

func (v *driveAta) GoSmartGetAttributes(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SmartGetAttributes", flags, ch, options)
}

func (v *driveAta) GoSmartGetAttributesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SmartGetAttributes", flags, ch, options)
}

func (*driveAta) StoreSmartGetAttributes(call *dbus.Call) (attributes []Attribute, err error) {
	err = call.Store(&attributes)
	return
}

func (v *driveAta) SmartGetAttributes(flags dbus.Flags, options map[string]dbus.Variant) (attributes []Attribute, err error) {
	return v.StoreSmartGetAttributes(
		<-v.GoSmartGetAttributes(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *driveAta) SmartGetAttributesWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (attributes []Attribute, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSmartGetAttributesWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreSmartGetAttributes(call)
}

// method SmartSelftestStart

func (v *driveAta) GoSmartSelftestStart(flags dbus.Flags, ch chan *dbus.Call, type0 string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SmartSelftestStart", flags, ch, type0, options)
}

func (v *driveAta) GoSmartSelftestStartWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, type0 string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SmartSelftestStart", flags, ch, type0, options)
}

func (v *driveAta) SmartSelftestStart(flags dbus.Flags, type0 string, options map[string]dbus.Variant) error {
	return (<-v.GoSmartSelftestStart(flags, make(chan *dbus.Call, 1), type0, options).Done).Err
}

func (v *driveAta) SmartSelftestStartWithTimeout(timeout time.Duration, flags dbus.Flags, type0 string, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSmartSelftestStartWithContext(ctx, flags, make(chan *dbus.Call, 1), type0, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SmartSelftestAbort

func (v *driveAta) GoSmartSelftestAbort(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SmartSelftestAbort", flags, ch, options)
}

func (v *driveAta) GoSmartSelftestAbortWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SmartSelftestAbort", flags, ch, options)
}

func (v *driveAta) SmartSelftestAbort(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoSmartSelftestAbort(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *driveAta) SmartSelftestAbortWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSmartSelftestAbortWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SmartSetEnabled

func (v *driveAta) GoSmartSetEnabled(flags dbus.Flags, ch chan *dbus.Call, value bool, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SmartSetEnabled", flags, ch, value, options)
}

func (v *driveAta) GoSmartSetEnabledWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, value bool, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SmartSetEnabled", flags, ch, value, options)
}

func (v *driveAta) SmartSetEnabled(flags dbus.Flags, value bool, options map[string]dbus.Variant) error {
	return (<-v.GoSmartSetEnabled(flags, make(chan *dbus.Call, 1), value, options).Done).Err
}

func (v *driveAta) SmartSetEnabledWithTimeout(timeout time.Duration, flags dbus.Flags, value bool, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSmartSetEnabledWithContext(ctx, flags, make(chan *dbus.Call, 1), value, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PmGetState

func (v *driveAta) GoPmGetState(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PmGetState", flags, ch, options)
}

func (v *driveAta) GoPmGetStateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PmGetState", flags, ch, options)
}

func (*driveAta) StorePmGetState(call *dbus.Call) (state uint8, err error) {
	err = call.Store(&state)
	return
}

func (v *driveAta) PmGetState(flags dbus.Flags, options map[string]dbus.Variant) (state uint8, err error) {
	return v.StorePmGetState(
		<-v.GoPmGetState(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *driveAta) PmGetStateWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (state uint8, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPmGetStateWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePmGetState(call)
}

// method PmStandby

func (v *driveAta) GoPmStandby(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PmStandby", flags, ch, options)
}

func (v *driveAta) GoPmStandbyWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PmStandby", flags, ch, options)
}

func (v *driveAta) PmStandby(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoPmStandby(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *driveAta) PmStandbyWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPmStandbyWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PmWakeup

func (v *driveAta) GoPmWakeup(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PmWakeup", flags, ch, options)
}

func (v *driveAta) GoPmWakeupWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PmWakeup", flags, ch, options)
}

func (v *driveAta) PmWakeup(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoPmWakeup(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *driveAta) PmWakeupWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPmWakeupWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SecurityEraseUnit

func (v *driveAta) GoSecurityEraseUnit(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SecurityEraseUnit", flags, ch, options)
}

func (v *driveAta) GoSecurityEraseUnitWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SecurityEraseUnit", flags, ch, options)
}

func (v *driveAta) SecurityEraseUnit(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoSecurityEraseUnit(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *driveAta) SecurityEraseUnitWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSecurityEraseUnitWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property SmartSupported b

func (v *driveAta) SmartSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SmartSupported",
	}
}

// property SmartEnabled b

func (v *driveAta) SmartEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SmartEnabled",
	}
}

// property SmartUpdated t

func (v *driveAta) SmartUpdated() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SmartUpdated",
	}
}

// property SmartFailing b

func (v *driveAta) SmartFailing() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SmartFailing",
	}
}

// property SmartPowerOnSeconds t

func (v *driveAta) SmartPowerOnSeconds() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "SmartPowerOnSeconds",
	}
}

// property SmartTemperature d

func (v *driveAta) SmartTemperature() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "SmartTemperature",
	}
}

// property SmartNumAttributesFailing i

func (v *driveAta) SmartNumAttributesFailing() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SmartNumAttributesFailing",
	}
}

// property SmartNumAttributesFailedInThePast i

func (v *driveAta) SmartNumAttributesFailedInThePast() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SmartNumAttributesFailedInThePast",
	}
}

// property SmartNumBadSectors x

func (v *driveAta) SmartNumBadSectors() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "SmartNumBadSectors",
	}
}

// property SmartSelftestStatus s

func (v *driveAta) SmartSelftestStatus() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "SmartSelftestStatus",
	}
}

// property SmartSelftestPercentRemaining i

func (v *driveAta) SmartSelftestPercentRemaining() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SmartSelftestPercentRemaining",
	}
}

// property PmSupported b

func (v *driveAta) PmSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PmSupported",
	}
}

// property PmEnabled b

func (v *driveAta) PmEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "PmEnabled",
	}
}

// property ApmSupported b

func (v *driveAta) ApmSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ApmSupported",
	}
}

// property ApmEnabled b

func (v *driveAta) ApmEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ApmEnabled",
	}
}

// property AamSupported b

func (v *driveAta) AamSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AamSupported",
	}
}

// property AamEnabled b

func (v *driveAta) AamEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AamEnabled",
	}
}

// property AamVendorRecommendedValue i

func (v *driveAta) AamVendorRecommendedValue() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "AamVendorRecommendedValue",
	}
}

// property WriteCacheSupported b

func (v *driveAta) WriteCacheSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WriteCacheSupported",
	}
}

// property WriteCacheEnabled b

func (v *driveAta) WriteCacheEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "WriteCacheEnabled",
	}
}

// property ReadLookaheadSupported b

func (v *driveAta) ReadLookaheadSupported() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ReadLookaheadSupported",
	}
}

// property ReadLookaheadEnabled b

func (v *driveAta) ReadLookaheadEnabled() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ReadLookaheadEnabled",
	}
}

// property SecurityEraseUnitMinutes i

func (v *driveAta) SecurityEraseUnitMinutes() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SecurityEraseUnitMinutes",
	}
}

// property SecurityEnhancedEraseUnitMinutes i

func (v *driveAta) SecurityEnhancedEraseUnitMinutes() proxy.PropInt32 {
	return proxy.PropInt32{
		Impl: v,
		Name: "SecurityEnhancedEraseUnitMinutes",
	}
}

// property SecurityFrozen b

func (v *driveAta) SecurityFrozen() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SecurityFrozen",
	}
}

type Block struct {
	block          // interface org.freedesktop.UDisks2.Block
	partitionTable // interface org.freedesktop.UDisks2.PartitionTable
	partition      // interface org.freedesktop.UDisks2.Partition
	filesystem     // interface org.freedesktop.UDisks2.Filesystem
	proxy.Object
}

func NewBlock(conn *dbus.Conn, path dbus.ObjectPath) (*Block, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Block)
	obj.Object.Init_(conn, "org.freedesktop.UDisks2", path)
	return obj, nil
}

func (obj *Block) Block() *block {
	return &obj.block
}

type block struct{}

func (v *block) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*block) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.Block"
}

// method AddConfigurationItem

func (v *block) GoAddConfigurationItem(flags dbus.Flags, ch chan *dbus.Call, item ConfigurationItem, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".AddConfigurationItem", flags, ch, item, options)
}

func (v *block) GoAddConfigurationItemWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, item ConfigurationItem, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".AddConfigurationItem", flags, ch, item, options)
}

func (v *block) AddConfigurationItem(flags dbus.Flags, item ConfigurationItem, options map[string]dbus.Variant) error {
	return (<-v.GoAddConfigurationItem(flags, make(chan *dbus.Call, 1), item, options).Done).Err
}

func (v *block) AddConfigurationItemWithTimeout(timeout time.Duration, flags dbus.Flags, item ConfigurationItem, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoAddConfigurationItemWithContext(ctx, flags, make(chan *dbus.Call, 1), item, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method RemoveConfigurationItem

func (v *block) GoRemoveConfigurationItem(flags dbus.Flags, ch chan *dbus.Call, item ConfigurationItem, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemoveConfigurationItem", flags, ch, item, options)
}

func (v *block) GoRemoveConfigurationItemWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, item ConfigurationItem, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemoveConfigurationItem", flags, ch, item, options)
}

func (v *block) RemoveConfigurationItem(flags dbus.Flags, item ConfigurationItem, options map[string]dbus.Variant) error {
	return (<-v.GoRemoveConfigurationItem(flags, make(chan *dbus.Call, 1), item, options).Done).Err
}

func (v *block) RemoveConfigurationItemWithTimeout(timeout time.Duration, flags dbus.Flags, item ConfigurationItem, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemoveConfigurationItemWithContext(ctx, flags, make(chan *dbus.Call, 1), item, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UpdateConfigurationItem

func (v *block) GoUpdateConfigurationItem(flags dbus.Flags, ch chan *dbus.Call, old_item ConfigurationItem, new_item ConfigurationItem, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateConfigurationItem", flags, ch, old_item, new_item, options)
}

func (v *block) GoUpdateConfigurationItemWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, old_item ConfigurationItem, new_item ConfigurationItem, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UpdateConfigurationItem", flags, ch, old_item, new_item, options)
}

func (v *block) UpdateConfigurationItem(flags dbus.Flags, old_item ConfigurationItem, new_item ConfigurationItem, options map[string]dbus.Variant) error {
	return (<-v.GoUpdateConfigurationItem(flags, make(chan *dbus.Call, 1), old_item, new_item, options).Done).Err
}

func (v *block) UpdateConfigurationItemWithTimeout(timeout time.Duration, flags dbus.Flags, old_item ConfigurationItem, new_item ConfigurationItem, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUpdateConfigurationItemWithContext(ctx, flags, make(chan *dbus.Call, 1), old_item, new_item, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method GetSecretConfiguration

func (v *block) GoGetSecretConfiguration(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".GetSecretConfiguration", flags, ch, options)
}

func (v *block) GoGetSecretConfigurationWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".GetSecretConfiguration", flags, ch, options)
}

func (*block) StoreGetSecretConfiguration(call *dbus.Call) (configuration []ConfigurationItem, err error) {
	err = call.Store(&configuration)
	return
}

func (v *block) GetSecretConfiguration(flags dbus.Flags, options map[string]dbus.Variant) (configuration []ConfigurationItem, err error) {
	return v.StoreGetSecretConfiguration(
		<-v.GoGetSecretConfiguration(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *block) GetSecretConfigurationWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (configuration []ConfigurationItem, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoGetSecretConfigurationWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreGetSecretConfiguration(call)
}

// method Format

func (v *block) GoFormat(flags dbus.Flags, ch chan *dbus.Call, type0 string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Format", flags, ch, type0, options)
}

func (v *block) GoFormatWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, type0 string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Format", flags, ch, type0, options)
}

func (v *block) Format(flags dbus.Flags, type0 string, options map[string]dbus.Variant) error {
	return (<-v.GoFormat(flags, make(chan *dbus.Call, 1), type0, options).Done).Err
}

func (v *block) FormatWithTimeout(timeout time.Duration, flags dbus.Flags, type0 string, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoFormatWithContext(ctx, flags, make(chan *dbus.Call, 1), type0, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method OpenForBackup

func (v *block) GoOpenForBackup(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OpenForBackup", flags, ch, options)
}

func (v *block) GoOpenForBackupWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".OpenForBackup", flags, ch, options)
}

func (*block) StoreOpenForBackup(call *dbus.Call) (fd dbus.UnixFD, err error) {
	err = call.Store(&fd)
	return
}

func (v *block) OpenForBackup(flags dbus.Flags, options map[string]dbus.Variant) (fd dbus.UnixFD, err error) {
	return v.StoreOpenForBackup(
		<-v.GoOpenForBackup(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *block) OpenForBackupWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (fd dbus.UnixFD, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoOpenForBackupWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreOpenForBackup(call)
}

// method OpenForRestore

func (v *block) GoOpenForRestore(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OpenForRestore", flags, ch, options)
}

func (v *block) GoOpenForRestoreWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".OpenForRestore", flags, ch, options)
}

func (*block) StoreOpenForRestore(call *dbus.Call) (fd dbus.UnixFD, err error) {
	err = call.Store(&fd)
	return
}

func (v *block) OpenForRestore(flags dbus.Flags, options map[string]dbus.Variant) (fd dbus.UnixFD, err error) {
	return v.StoreOpenForRestore(
		<-v.GoOpenForRestore(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *block) OpenForRestoreWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (fd dbus.UnixFD, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoOpenForRestoreWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreOpenForRestore(call)
}

// method OpenForBenchmark

func (v *block) GoOpenForBenchmark(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".OpenForBenchmark", flags, ch, options)
}

func (v *block) GoOpenForBenchmarkWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".OpenForBenchmark", flags, ch, options)
}

func (*block) StoreOpenForBenchmark(call *dbus.Call) (fd dbus.UnixFD, err error) {
	err = call.Store(&fd)
	return
}

func (v *block) OpenForBenchmark(flags dbus.Flags, options map[string]dbus.Variant) (fd dbus.UnixFD, err error) {
	return v.StoreOpenForBenchmark(
		<-v.GoOpenForBenchmark(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *block) OpenForBenchmarkWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (fd dbus.UnixFD, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoOpenForBenchmarkWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreOpenForBenchmark(call)
}

// method Rescan

func (v *block) GoRescan(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Rescan", flags, ch, options)
}

func (v *block) GoRescanWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Rescan", flags, ch, options)
}

func (v *block) Rescan(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoRescan(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *block) RescanWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRescanWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Device ay

func (v *block) Device() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "Device",
	}
}

// property PreferredDevice ay

func (v *block) PreferredDevice() proxy.PropByteArray {
	return proxy.PropByteArray{
		Impl: v,
		Name: "PreferredDevice",
	}
}

// property Symlinks aay

func (v *block) Symlinks() PropByteSliceSlice {
	return PropByteSliceSlice{
		Impl: v,
		Name: "Symlinks",
	}
}

// property DeviceNumber t

func (v *block) DeviceNumber() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "DeviceNumber",
	}
}

// property Id s

func (v *block) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Size t

func (v *block) Size() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Size",
	}
}

// property ReadOnly b

func (v *block) ReadOnly() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "ReadOnly",
	}
}

// property Drive o

func (v *block) Drive() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Drive",
	}
}

// property MDRaid o

func (v *block) MDRaid() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "MDRaid",
	}
}

// property MDRaidMember o

func (v *block) MDRaidMember() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "MDRaidMember",
	}
}

// property IdUsage s

func (v *block) IdUsage() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IdUsage",
	}
}

// property IdType s

func (v *block) IdType() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IdType",
	}
}

// property IdVersion s

func (v *block) IdVersion() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IdVersion",
	}
}

// property IdLabel s

func (v *block) IdLabel() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IdLabel",
	}
}

// property IdUUID s

func (v *block) IdUUID() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "IdUUID",
	}
}

// property Configuration a(sa{sv})

func (v *block) Configuration() PropBlockConfiguration {
	return PropBlockConfiguration{
		Impl: v,
	}
}

type PropBlockConfiguration struct {
	Impl proxy.Implementer
}

func (p PropBlockConfiguration) Get(flags dbus.Flags) (value []ConfigurationItem, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		"Configuration", &value)
	return
}

func (p PropBlockConfiguration) ConnectChanged(cb func(hasValue bool, value []ConfigurationItem)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v []ConfigurationItem
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		"Configuration", cb0)
}

// property CryptoBackingDevice o

func (v *block) CryptoBackingDevice() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "CryptoBackingDevice",
	}
}

// property HintPartitionable b

func (v *block) HintPartitionable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HintPartitionable",
	}
}

// property HintSystem b

func (v *block) HintSystem() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HintSystem",
	}
}

// property HintIgnore b

func (v *block) HintIgnore() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HintIgnore",
	}
}

// property HintAuto b

func (v *block) HintAuto() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "HintAuto",
	}
}

// property HintName s

func (v *block) HintName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HintName",
	}
}

// property HintIconName s

func (v *block) HintIconName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HintIconName",
	}
}

// property HintSymbolicIconName s

func (v *block) HintSymbolicIconName() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "HintSymbolicIconName",
	}
}

func (obj *Block) PartitionTable() *partitionTable {
	return &obj.partitionTable
}

type partitionTable struct{}

func (v *partitionTable) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*partitionTable) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.PartitionTable"
}

// method CreatePartition

func (v *partitionTable) GoCreatePartition(flags dbus.Flags, ch chan *dbus.Call, offset uint64, size uint64, type0 string, name string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CreatePartition", flags, ch, offset, size, type0, name, options)
}

func (v *partitionTable) GoCreatePartitionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, offset uint64, size uint64, type0 string, name string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CreatePartition", flags, ch, offset, size, type0, name, options)
}

func (*partitionTable) StoreCreatePartition(call *dbus.Call) (created_partition dbus.ObjectPath, err error) {
	err = call.Store(&created_partition)
	return
}

func (v *partitionTable) CreatePartition(flags dbus.Flags, offset uint64, size uint64, type0 string, name string, options map[string]dbus.Variant) (created_partition dbus.ObjectPath, err error) {
	return v.StoreCreatePartition(
		<-v.GoCreatePartition(flags, make(chan *dbus.Call, 1), offset, size, type0, name, options).Done)
}

func (v *partitionTable) CreatePartitionWithTimeout(timeout time.Duration, flags dbus.Flags, offset uint64, size uint64, type0 string, name string, options map[string]dbus.Variant) (created_partition dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCreatePartitionWithContext(ctx, flags, make(chan *dbus.Call, 1), offset, size, type0, name, options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCreatePartition(call)
}

// property Type s

func (v *partitionTable) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

func (obj *Block) Partition() *partition {
	return &obj.partition
}

type partition struct{}

func (v *partition) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*partition) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.Partition"
}

// method SetType

func (v *partition) GoSetType(flags dbus.Flags, ch chan *dbus.Call, type0 string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetType", flags, ch, type0, options)
}

func (v *partition) GoSetTypeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, type0 string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetType", flags, ch, type0, options)
}

func (v *partition) SetType(flags dbus.Flags, type0 string, options map[string]dbus.Variant) error {
	return (<-v.GoSetType(flags, make(chan *dbus.Call, 1), type0, options).Done).Err
}

func (v *partition) SetTypeWithTimeout(timeout time.Duration, flags dbus.Flags, type0 string, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetTypeWithContext(ctx, flags, make(chan *dbus.Call, 1), type0, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetName

func (v *partition) GoSetName(flags dbus.Flags, ch chan *dbus.Call, name string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetName", flags, ch, name, options)
}

func (v *partition) GoSetNameWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, name string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetName", flags, ch, name, options)
}

func (v *partition) SetName(flags dbus.Flags, name string, options map[string]dbus.Variant) error {
	return (<-v.GoSetName(flags, make(chan *dbus.Call, 1), name, options).Done).Err
}

func (v *partition) SetNameWithTimeout(timeout time.Duration, flags dbus.Flags, name string, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetNameWithContext(ctx, flags, make(chan *dbus.Call, 1), name, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetFlags

func (v *partition) GoSetFlags(flags dbus.Flags, ch chan *dbus.Call, flags0 uint64, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetFlags", flags, ch, flags0, options)
}

func (v *partition) GoSetFlagsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, flags0 uint64, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetFlags", flags, ch, flags0, options)
}

func (v *partition) SetFlags(flags dbus.Flags, flags0 uint64, options map[string]dbus.Variant) error {
	return (<-v.GoSetFlags(flags, make(chan *dbus.Call, 1), flags0, options).Done).Err
}

func (v *partition) SetFlagsWithTimeout(timeout time.Duration, flags dbus.Flags, flags0 uint64, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetFlagsWithContext(ctx, flags, make(chan *dbus.Call, 1), flags0, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Delete

func (v *partition) GoDelete(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Delete", flags, ch, options)
}

func (v *partition) GoDeleteWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Delete", flags, ch, options)
}

func (v *partition) Delete(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoDelete(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *partition) DeleteWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDeleteWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property Number u

func (v *partition) Number() proxy.PropUint32 {
	return proxy.PropUint32{
		Impl: v,
		Name: "Number",
	}
}

// property Type s

func (v *partition) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Flags t

func (v *partition) Flags() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Flags",
	}
}

// property Offset t

func (v *partition) Offset() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Offset",
	}
}

// property Size t

func (v *partition) Size() proxy.PropUint64 {
	return proxy.PropUint64{
		Impl: v,
		Name: "Size",
	}
}

// property Name s

func (v *partition) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property UUID s

func (v *partition) UUID() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "UUID",
	}
}

// property Table o

func (v *partition) Table() proxy.PropObjectPath {
	return proxy.PropObjectPath{
		Impl: v,
		Name: "Table",
	}
}

// property IsContainer b

func (v *partition) IsContainer() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsContainer",
	}
}

// property IsContained b

func (v *partition) IsContained() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "IsContained",
	}
}

func (obj *Block) Filesystem() *filesystem {
	return &obj.filesystem
}

type filesystem struct{}

func (v *filesystem) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*filesystem) GetInterfaceName_() string {
	return "org.freedesktop.UDisks2.Filesystem"
}

// method SetLabel

func (v *filesystem) GoSetLabel(flags dbus.Flags, ch chan *dbus.Call, label string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLabel", flags, ch, label, options)
}

func (v *filesystem) GoSetLabelWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, label string, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLabel", flags, ch, label, options)
}

func (v *filesystem) SetLabel(flags dbus.Flags, label string, options map[string]dbus.Variant) error {
	return (<-v.GoSetLabel(flags, make(chan *dbus.Call, 1), label, options).Done).Err
}

func (v *filesystem) SetLabelWithTimeout(timeout time.Duration, flags dbus.Flags, label string, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLabelWithContext(ctx, flags, make(chan *dbus.Call, 1), label, options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Mount

func (v *filesystem) GoMount(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Mount", flags, ch, options)
}

func (v *filesystem) GoMountWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Mount", flags, ch, options)
}

func (*filesystem) StoreMount(call *dbus.Call) (mount_path string, err error) {
	err = call.Store(&mount_path)
	return
}

func (v *filesystem) Mount(flags dbus.Flags, options map[string]dbus.Variant) (mount_path string, err error) {
	return v.StoreMount(
		<-v.GoMount(flags, make(chan *dbus.Call, 1), options).Done)
}

func (v *filesystem) MountWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) (mount_path string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoMountWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreMount(call)
}

// method Unmount

func (v *filesystem) GoUnmount(flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Unmount", flags, ch, options)
}

func (v *filesystem) GoUnmountWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, options map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Unmount", flags, ch, options)
}

func (v *filesystem) Unmount(flags dbus.Flags, options map[string]dbus.Variant) error {
	return (<-v.GoUnmount(flags, make(chan *dbus.Call, 1), options).Done).Err
}

func (v *filesystem) UnmountWithTimeout(timeout time.Duration, flags dbus.Flags, options map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUnmountWithContext(ctx, flags, make(chan *dbus.Call, 1), options).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property MountPoints aay

func (v *filesystem) MountPoints() PropByteSliceSlice {
	return PropByteSliceSlice{
		Impl: v,
		Name: "MountPoints",
	}
}

type PropByteSliceSlice struct {
	Impl proxy.Implementer
	Name string
}

func (p PropByteSliceSlice) Get(flags dbus.Flags) (value [][]byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropByteSliceSlice) Set(flags dbus.Flags, value [][]byte) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropByteSliceSlice) ConnectChanged(cb func(hasValue bool, value [][]byte)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v [][]byte
			err := dbus.Store([]interface{}{value}, &v)
			if err != nil {
				return
			}
			cb(true, v)
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(),
		p.Name, cb0)
}
