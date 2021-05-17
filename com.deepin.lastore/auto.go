package lastore

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

type Lastore struct {
	manager // interface com.deepin.lastore.Manager
	updater // interface com.deepin.lastore.Updater
	proxy.Object
}

func NewLastore(conn *dbus.Conn) *Lastore {
	obj := new(Lastore)
	obj.Object.Init_(conn, "com.deepin.lastore", "/com/deepin/lastore")
	return obj
}

func (obj *Lastore) Manager() *manager {
	return &obj.manager
}

type manager struct{}

func (v *manager) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*manager) GetInterfaceName_() string {
	return "com.deepin.lastore.Manager"
}

// method CleanArchives

func (v *manager) GoCleanArchives(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CleanArchives", flags, ch)
}

func (v *manager) GoCleanArchivesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CleanArchives", flags, ch)
}

func (*manager) StoreCleanArchives(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) CleanArchives(flags dbus.Flags) (job dbus.ObjectPath, err error) {
	return v.StoreCleanArchives(
		<-v.GoCleanArchives(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) CleanArchivesWithTimeout(timeout time.Duration, flags dbus.Flags) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCleanArchivesWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreCleanArchives(call)
}

// method CleanJob

func (v *manager) GoCleanJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".CleanJob", flags, ch, jobId)
}

func (v *manager) GoCleanJobWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".CleanJob", flags, ch, jobId)
}

func (v *manager) CleanJob(flags dbus.Flags, jobId string) error {
	return (<-v.GoCleanJob(flags, make(chan *dbus.Call, 1), jobId).Done).Err
}

func (v *manager) CleanJobWithTimeout(timeout time.Duration, flags dbus.Flags, jobId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoCleanJobWithContext(ctx, flags, make(chan *dbus.Call, 1), jobId).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method DistUpgrade

func (v *manager) GoDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".DistUpgrade", flags, ch)
}

func (v *manager) GoDistUpgradeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".DistUpgrade", flags, ch)
}

func (*manager) StoreDistUpgrade(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) DistUpgrade(flags dbus.Flags) (job dbus.ObjectPath, err error) {
	return v.StoreDistUpgrade(
		<-v.GoDistUpgrade(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) DistUpgradeWithTimeout(timeout time.Duration, flags dbus.Flags) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoDistUpgradeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreDistUpgrade(call)
}

// method FixError

func (v *manager) GoFixError(flags dbus.Flags, ch chan *dbus.Call, errType string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".FixError", flags, ch, errType)
}

func (v *manager) GoFixErrorWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, errType string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".FixError", flags, ch, errType)
}

func (*manager) StoreFixError(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) FixError(flags dbus.Flags, errType string) (job dbus.ObjectPath, err error) {
	return v.StoreFixError(
		<-v.GoFixError(flags, make(chan *dbus.Call, 1), errType).Done)
}

func (v *manager) FixErrorWithTimeout(timeout time.Duration, flags dbus.Flags, errType string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoFixErrorWithContext(ctx, flags, make(chan *dbus.Call, 1), errType).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreFixError(call)
}

// method InstallPackage

func (v *manager) GoInstallPackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".InstallPackage", flags, ch, jobName, packages)
}

func (v *manager) GoInstallPackageWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".InstallPackage", flags, ch, jobName, packages)
}

func (*manager) StoreInstallPackage(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) InstallPackage(flags dbus.Flags, jobName string, packages string) (job dbus.ObjectPath, err error) {
	return v.StoreInstallPackage(
		<-v.GoInstallPackage(flags, make(chan *dbus.Call, 1), jobName, packages).Done)
}

func (v *manager) InstallPackageWithTimeout(timeout time.Duration, flags dbus.Flags, jobName string, packages string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoInstallPackageWithContext(ctx, flags, make(chan *dbus.Call, 1), jobName, packages).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreInstallPackage(call)
}

// method PackageDesktopPath

func (v *manager) GoPackageDesktopPath(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageDesktopPath", flags, ch, pkgId)
}

func (v *manager) GoPackageDesktopPathWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PackageDesktopPath", flags, ch, pkgId)
}

func (*manager) StorePackageDesktopPath(call *dbus.Call) (desktopPath string, err error) {
	err = call.Store(&desktopPath)
	return
}

func (v *manager) PackageDesktopPath(flags dbus.Flags, pkgId string) (desktopPath string, err error) {
	return v.StorePackageDesktopPath(
		<-v.GoPackageDesktopPath(flags, make(chan *dbus.Call, 1), pkgId).Done)
}

func (v *manager) PackageDesktopPathWithTimeout(timeout time.Duration, flags dbus.Flags, pkgId string) (desktopPath string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPackageDesktopPathWithContext(ctx, flags, make(chan *dbus.Call, 1), pkgId).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePackageDesktopPath(call)
}

// method PackageExists

func (v *manager) GoPackageExists(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageExists", flags, ch, pkgId)
}

func (v *manager) GoPackageExistsWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PackageExists", flags, ch, pkgId)
}

func (*manager) StorePackageExists(call *dbus.Call) (exist bool, err error) {
	err = call.Store(&exist)
	return
}

func (v *manager) PackageExists(flags dbus.Flags, pkgId string) (exist bool, err error) {
	return v.StorePackageExists(
		<-v.GoPackageExists(flags, make(chan *dbus.Call, 1), pkgId).Done)
}

func (v *manager) PackageExistsWithTimeout(timeout time.Duration, flags dbus.Flags, pkgId string) (exist bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPackageExistsWithContext(ctx, flags, make(chan *dbus.Call, 1), pkgId).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePackageExists(call)
}

// method PackageInstallable

func (v *manager) GoPackageInstallable(flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackageInstallable", flags, ch, pkgId)
}

func (v *manager) GoPackageInstallableWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, pkgId string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PackageInstallable", flags, ch, pkgId)
}

func (*manager) StorePackageInstallable(call *dbus.Call) (installable bool, err error) {
	err = call.Store(&installable)
	return
}

func (v *manager) PackageInstallable(flags dbus.Flags, pkgId string) (installable bool, err error) {
	return v.StorePackageInstallable(
		<-v.GoPackageInstallable(flags, make(chan *dbus.Call, 1), pkgId).Done)
}

func (v *manager) PackageInstallableWithTimeout(timeout time.Duration, flags dbus.Flags, pkgId string) (installable bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPackageInstallableWithContext(ctx, flags, make(chan *dbus.Call, 1), pkgId).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePackageInstallable(call)
}

// method PackagesDownloadSize

func (v *manager) GoPackagesDownloadSize(flags dbus.Flags, ch chan *dbus.Call, packages []string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PackagesDownloadSize", flags, ch, packages)
}

func (v *manager) GoPackagesDownloadSizeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, packages []string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PackagesDownloadSize", flags, ch, packages)
}

func (*manager) StorePackagesDownloadSize(call *dbus.Call) (size int64, err error) {
	err = call.Store(&size)
	return
}

func (v *manager) PackagesDownloadSize(flags dbus.Flags, packages []string) (size int64, err error) {
	return v.StorePackagesDownloadSize(
		<-v.GoPackagesDownloadSize(flags, make(chan *dbus.Call, 1), packages).Done)
}

func (v *manager) PackagesDownloadSizeWithTimeout(timeout time.Duration, flags dbus.Flags, packages []string) (size int64, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPackagesDownloadSizeWithContext(ctx, flags, make(chan *dbus.Call, 1), packages).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePackagesDownloadSize(call)
}

// method PauseJob

func (v *manager) GoPauseJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PauseJob", flags, ch, jobId)
}

func (v *manager) GoPauseJobWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PauseJob", flags, ch, jobId)
}

func (v *manager) PauseJob(flags dbus.Flags, jobId string) error {
	return (<-v.GoPauseJob(flags, make(chan *dbus.Call, 1), jobId).Done).Err
}

func (v *manager) PauseJobWithTimeout(timeout time.Duration, flags dbus.Flags, jobId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPauseJobWithContext(ctx, flags, make(chan *dbus.Call, 1), jobId).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PrepareDistUpgrade

func (v *manager) GoPrepareDistUpgrade(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PrepareDistUpgrade", flags, ch)
}

func (v *manager) GoPrepareDistUpgradeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PrepareDistUpgrade", flags, ch)
}

func (*manager) StorePrepareDistUpgrade(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) PrepareDistUpgrade(flags dbus.Flags) (job dbus.ObjectPath, err error) {
	return v.StorePrepareDistUpgrade(
		<-v.GoPrepareDistUpgrade(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) PrepareDistUpgradeWithTimeout(timeout time.Duration, flags dbus.Flags) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPrepareDistUpgradeWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StorePrepareDistUpgrade(call)
}

// method RemovePackage

func (v *manager) GoRemovePackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RemovePackage", flags, ch, jobName, packages)
}

func (v *manager) GoRemovePackageWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RemovePackage", flags, ch, jobName, packages)
}

func (*manager) StoreRemovePackage(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) RemovePackage(flags dbus.Flags, jobName string, packages string) (job dbus.ObjectPath, err error) {
	return v.StoreRemovePackage(
		<-v.GoRemovePackage(flags, make(chan *dbus.Call, 1), jobName, packages).Done)
}

func (v *manager) RemovePackageWithTimeout(timeout time.Duration, flags dbus.Flags, jobName string, packages string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRemovePackageWithContext(ctx, flags, make(chan *dbus.Call, 1), jobName, packages).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreRemovePackage(call)
}

// method SetAutoClean

func (v *manager) GoSetAutoClean(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoClean", flags, ch, enable)
}

func (v *manager) GoSetAutoCleanWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetAutoClean", flags, ch, enable)
}

func (v *manager) SetAutoClean(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetAutoClean(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

func (v *manager) SetAutoCleanWithTimeout(timeout time.Duration, flags dbus.Flags, enable bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetAutoCleanWithContext(ctx, flags, make(chan *dbus.Call, 1), enable).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetLogger

func (v *manager) GoSetLogger(flags dbus.Flags, ch chan *dbus.Call, levels string, format string, output string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetLogger", flags, ch, levels, format, output)
}

func (v *manager) GoSetLoggerWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, levels string, format string, output string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetLogger", flags, ch, levels, format, output)
}

func (v *manager) SetLogger(flags dbus.Flags, levels string, format string, output string) error {
	return (<-v.GoSetLogger(flags, make(chan *dbus.Call, 1), levels, format, output).Done).Err
}

func (v *manager) SetLoggerWithTimeout(timeout time.Duration, flags dbus.Flags, levels string, format string, output string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetLoggerWithContext(ctx, flags, make(chan *dbus.Call, 1), levels, format, output).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetRegion

func (v *manager) GoSetRegion(flags dbus.Flags, ch chan *dbus.Call, region string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetRegion", flags, ch, region)
}

func (v *manager) GoSetRegionWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, region string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetRegion", flags, ch, region)
}

func (v *manager) SetRegion(flags dbus.Flags, region string) error {
	return (<-v.GoSetRegion(flags, make(chan *dbus.Call, 1), region).Done).Err
}

func (v *manager) SetRegionWithTimeout(timeout time.Duration, flags dbus.Flags, region string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetRegionWithContext(ctx, flags, make(chan *dbus.Call, 1), region).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method StartJob

func (v *manager) GoStartJob(flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".StartJob", flags, ch, jobId)
}

func (v *manager) GoStartJobWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, jobId string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".StartJob", flags, ch, jobId)
}

func (v *manager) StartJob(flags dbus.Flags, jobId string) error {
	return (<-v.GoStartJob(flags, make(chan *dbus.Call, 1), jobId).Done).Err
}

func (v *manager) StartJobWithTimeout(timeout time.Duration, flags dbus.Flags, jobId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStartJobWithContext(ctx, flags, make(chan *dbus.Call, 1), jobId).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method UpdatePackage

func (v *manager) GoUpdatePackage(flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdatePackage", flags, ch, jobName, packages)
}

func (v *manager) GoUpdatePackageWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, jobName string, packages string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UpdatePackage", flags, ch, jobName, packages)
}

func (*manager) StoreUpdatePackage(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) UpdatePackage(flags dbus.Flags, jobName string, packages string) (job dbus.ObjectPath, err error) {
	return v.StoreUpdatePackage(
		<-v.GoUpdatePackage(flags, make(chan *dbus.Call, 1), jobName, packages).Done)
}

func (v *manager) UpdatePackageWithTimeout(timeout time.Duration, flags dbus.Flags, jobName string, packages string) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUpdatePackageWithContext(ctx, flags, make(chan *dbus.Call, 1), jobName, packages).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreUpdatePackage(call)
}

// method UpdateSource

func (v *manager) GoUpdateSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".UpdateSource", flags, ch)
}

func (v *manager) GoUpdateSourceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".UpdateSource", flags, ch)
}

func (*manager) StoreUpdateSource(call *dbus.Call) (job dbus.ObjectPath, err error) {
	err = call.Store(&job)
	return
}

func (v *manager) UpdateSource(flags dbus.Flags) (job dbus.ObjectPath, err error) {
	return v.StoreUpdateSource(
		<-v.GoUpdateSource(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *manager) UpdateSourceWithTimeout(timeout time.Duration, flags dbus.Flags) (job dbus.ObjectPath, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoUpdateSourceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreUpdateSource(call)
}

// property JobList ao

func (v *manager) JobList() proxy.PropObjectPathArray {
	return proxy.PropObjectPathArray{
		Impl: v,
		Name: "JobList",
	}
}

// property SystemArchitectures as

func (v *manager) SystemArchitectures() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "SystemArchitectures",
	}
}

// property UpgradableApps as

func (v *manager) UpgradableApps() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UpgradableApps",
	}
}

// property SystemOnChanging b

func (v *manager) SystemOnChanging() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "SystemOnChanging",
	}
}

// property AutoClean b

func (v *manager) AutoClean() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutoClean",
	}
}

func (obj *Lastore) Updater() *updater {
	return &obj.updater
}

type updater struct{}

func (v *updater) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*updater) GetInterfaceName_() string {
	return "com.deepin.lastore.Updater"
}

// method ApplicationUpdateInfos

func (v *updater) GoApplicationUpdateInfos(flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ApplicationUpdateInfos", flags, ch, lang)
}

func (v *updater) GoApplicationUpdateInfosWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ApplicationUpdateInfos", flags, ch, lang)
}

func (*updater) StoreApplicationUpdateInfos(call *dbus.Call) (updateInfos [][]interface{}, err error) {
	err = call.Store(&updateInfos)
	return
}

func (v *updater) ApplicationUpdateInfos(flags dbus.Flags, lang string) (updateInfos [][]interface{}, err error) {
	return v.StoreApplicationUpdateInfos(
		<-v.GoApplicationUpdateInfos(flags, make(chan *dbus.Call, 1), lang).Done)
}

func (v *updater) ApplicationUpdateInfosWithTimeout(timeout time.Duration, flags dbus.Flags, lang string) (updateInfos [][]interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoApplicationUpdateInfosWithContext(ctx, flags, make(chan *dbus.Call, 1), lang).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreApplicationUpdateInfos(call)
}

// method ListMirrorSources

func (v *updater) GoListMirrorSources(flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".ListMirrorSources", flags, ch, lang)
}

func (v *updater) GoListMirrorSourcesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, lang string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".ListMirrorSources", flags, ch, lang)
}

func (*updater) StoreListMirrorSources(call *dbus.Call) (mirrorSources [][]interface{}, err error) {
	err = call.Store(&mirrorSources)
	return
}

func (v *updater) ListMirrorSources(flags dbus.Flags, lang string) (mirrorSources [][]interface{}, err error) {
	return v.StoreListMirrorSources(
		<-v.GoListMirrorSources(flags, make(chan *dbus.Call, 1), lang).Done)
}

func (v *updater) ListMirrorSourcesWithTimeout(timeout time.Duration, flags dbus.Flags, lang string) (mirrorSources [][]interface{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoListMirrorSourcesWithContext(ctx, flags, make(chan *dbus.Call, 1), lang).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreListMirrorSources(call)
}

// method RestoreSystemSource

func (v *updater) GoRestoreSystemSource(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".RestoreSystemSource", flags, ch)
}

func (v *updater) GoRestoreSystemSourceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".RestoreSystemSource", flags, ch)
}

func (v *updater) RestoreSystemSource(flags dbus.Flags) error {
	return (<-v.GoRestoreSystemSource(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *updater) RestoreSystemSourceWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoRestoreSystemSourceWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetAutoCheckUpdates

func (v *updater) GoSetAutoCheckUpdates(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoCheckUpdates", flags, ch, enable)
}

func (v *updater) GoSetAutoCheckUpdatesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetAutoCheckUpdates", flags, ch, enable)
}

func (v *updater) SetAutoCheckUpdates(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetAutoCheckUpdates(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

func (v *updater) SetAutoCheckUpdatesWithTimeout(timeout time.Duration, flags dbus.Flags, enable bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetAutoCheckUpdatesWithContext(ctx, flags, make(chan *dbus.Call, 1), enable).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetAutoDownloadUpdates

func (v *updater) GoSetAutoDownloadUpdates(flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetAutoDownloadUpdates", flags, ch, enable)
}

func (v *updater) GoSetAutoDownloadUpdatesWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enable bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetAutoDownloadUpdates", flags, ch, enable)
}

func (v *updater) SetAutoDownloadUpdates(flags dbus.Flags, enable bool) error {
	return (<-v.GoSetAutoDownloadUpdates(flags, make(chan *dbus.Call, 1), enable).Done).Err
}

func (v *updater) SetAutoDownloadUpdatesWithTimeout(timeout time.Duration, flags dbus.Flags, enable bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetAutoDownloadUpdatesWithContext(ctx, flags, make(chan *dbus.Call, 1), enable).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetMirrorSource

func (v *updater) GoSetMirrorSource(flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetMirrorSource", flags, ch, id)
}

func (v *updater) GoSetMirrorSourceWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, id string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetMirrorSource", flags, ch, id)
}

func (v *updater) SetMirrorSource(flags dbus.Flags, id string) error {
	return (<-v.GoSetMirrorSource(flags, make(chan *dbus.Call, 1), id).Done).Err
}

func (v *updater) SetMirrorSourceWithTimeout(timeout time.Duration, flags dbus.Flags, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetMirrorSourceWithContext(ctx, flags, make(chan *dbus.Call, 1), id).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// property AutoCheckUpdates b

func (v *updater) AutoCheckUpdates() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutoCheckUpdates",
	}
}

// property AutoDownloadUpdates b

func (v *updater) AutoDownloadUpdates() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "AutoDownloadUpdates",
	}
}

// property MirrorSource s

func (v *updater) MirrorSource() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "MirrorSource",
	}
}

// property UpdatableApps as

func (v *updater) UpdatableApps() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UpdatableApps",
	}
}

// property UpdatablePackages as

func (v *updater) UpdatablePackages() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "UpdatablePackages",
	}
}

type Job struct {
	job // interface com.deepin.lastore.Job
	proxy.Object
}

func NewJob(conn *dbus.Conn, path dbus.ObjectPath) (*Job, error) {
	if !path.IsValid() {
		return nil, errors.New("path is invalid")
	}
	obj := new(Job)
	obj.Object.Init_(conn, "com.deepin.lastore", path)
	return obj, nil
}

type job struct{}

func (v *job) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*job) GetInterfaceName_() string {
	return "com.deepin.lastore.Job"
}

// method String

func (v *job) GoString(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".String", flags, ch)
}

func (v *job) GoStringWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".String", flags, ch)
}

func (*job) StoreString(call *dbus.Call) (arg0 string, err error) {
	err = call.Store(&arg0)
	return
}

func (v *job) String(flags dbus.Flags) (arg0 string, err error) {
	return v.StoreString(
		<-v.GoString(flags, make(chan *dbus.Call, 1)).Done)
}

func (v *job) StringWithTimeout(timeout time.Duration, flags dbus.Flags) (arg0 string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoStringWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		err = ctx.Err()
		return
	} else if call.Err != nil {
		err = call.Err
		return
	}

	return v.StoreString(call)
}

// property Id s

func (v *job) Id() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Id",
	}
}

// property Name s

func (v *job) Name() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Name",
	}
}

// property Packages as

func (v *job) Packages() proxy.PropStringArray {
	return proxy.PropStringArray{
		Impl: v,
		Name: "Packages",
	}
}

// property CreateTime x

func (v *job) CreateTime() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "CreateTime",
	}
}

// property Type s

func (v *job) Type() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Type",
	}
}

// property Status s

func (v *job) Status() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Status",
	}
}

// property Progress d

func (v *job) Progress() proxy.PropDouble {
	return proxy.PropDouble{
		Impl: v,
		Name: "Progress",
	}
}

// property Description s

func (v *job) Description() proxy.PropString {
	return proxy.PropString{
		Impl: v,
		Name: "Description",
	}
}

// property Speed x

func (v *job) Speed() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "Speed",
	}
}

// property DownloadSize x

func (v *job) DownloadSize() proxy.PropInt64 {
	return proxy.PropInt64{
		Impl: v,
		Name: "DownloadSize",
	}
}

// property Cancelable b

func (v *job) Cancelable() proxy.PropBool {
	return proxy.PropBool{
		Impl: v,
		Name: "Cancelable",
	}
}
