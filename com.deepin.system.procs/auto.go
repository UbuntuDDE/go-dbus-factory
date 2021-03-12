package procs

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

type Procs struct {
	procs // interface com.deepin.system.Procs
	proxy.Object
}

func NewProcs(conn *dbus.Conn) *Procs {
	obj := new(Procs)
	obj.Object.Init_(conn, "com.deepin.system.Procs", "/com/deepin/system/Procs")
	return obj
}

type procs struct{}

func (v *procs) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*procs) GetInterfaceName_() string {
	return "com.deepin.system.Procs"
}

// signal ExecProc

func (v *procs) ConnectExecProc(cb func(ExecPath string, CGroupPath string, Pid string, PPid string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ExecProc", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ExecProc",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ExecPath string
		var CGroupPath string
		var Pid string
		var PPid string
		err := dbus.Store(sig.Body, &ExecPath, &CGroupPath, &Pid, &PPid)
		if err == nil {
			cb(ExecPath, CGroupPath, Pid, PPid)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// signal ExitProc

func (v *procs) ConnectExitProc(cb func(ExecPath string, CGroupPath string, Pid string, PPid string)) (dbusutil.SignalHandlerId, error) {
	if cb == nil {
		return 0, errors.New("nil callback")
	}
	obj := v.GetObject_()
	rule := fmt.Sprintf(
		"type='signal',interface='%s',member='%s',path='%s',sender='%s'",
		v.GetInterfaceName_(), "ExitProc", obj.Path_(), obj.ServiceName_())

	sigRule := &dbusutil.SignalRule{
		Path: obj.Path_(),
		Name: v.GetInterfaceName_() + ".ExitProc",
	}
	handlerFunc := func(sig *dbus.Signal) {
		var ExecPath string
		var CGroupPath string
		var Pid string
		var PPid string
		err := dbus.Store(sig.Body, &ExecPath, &CGroupPath, &Pid, &PPid)
		if err == nil {
			cb(ExecPath, CGroupPath, Pid, PPid)
		}
	}

	return obj.ConnectSignal_(rule, sigRule, handlerFunc)
}

// property Procs a{s(ssss)}

func (v *procs) Procs() PropProcsMap {
	return PropProcsMap{
		Impl: v,
		Name: "Procs",
	}
}

type PropProcsMap struct {
	Impl proxy.Implementer
	Name string
}

func (p PropProcsMap) Get(flags dbus.Flags) (value map[string]ProcMessage, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(),
		p.Name, &value)
	return
}

func (p PropProcsMap) Set(flags dbus.Flags, value map[string]ProcMessage) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropProcsMap) ConnectChanged(cb func(hasValue bool, value map[string]ProcMessage)) error {
	if cb == nil {
		return errors.New("nil callback")
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			var v map[string]ProcMessage
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