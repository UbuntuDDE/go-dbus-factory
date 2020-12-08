package keyevent

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

type KeyEvent struct {
	keyEvent // interface com.deepin.daemon.KeyEvent
	proxy.Object
}

func NewKeyEvent(conn *dbus.Conn) *KeyEvent {
	obj := new(KeyEvent)
	obj.Object.Init_(conn, "com.deepin.daemon.KeyEvent", "/com/deepin/daemon/KeyEvent")
	return obj
}

type keyEvent struct{}

func (v *keyEvent) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*keyEvent) GetInterfaceName_() string {
	return "com.deepin.daemon.KeyEvent"
}

// signal KeyEvent

func (v *keyEvent) ConnectKeyEvent(cb func(keycode uint32, pressed bool, ctrlPressed bool, shiftPressed bool, altPressed bool, superPressed bool)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "KeyEvent", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".KeyEvent",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var keycode uint32
		var pressed bool
		var ctrlPressed bool
		var shiftPressed bool
		var altPressed bool
		var superPressed bool
		err := dbus.Store(sig.Body, &keycode, &pressed, &ctrlPressed, &shiftPressed, &altPressed, &superPressed)
		if err == nil {
			cb(keycode, pressed, ctrlPressed, shiftPressed, altPressed, superPressed)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}
