package soundthemeplayer

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

type SoundThemePlayer struct {
	soundThemePlayer // interface com.deepin.api.SoundThemePlayer
	proxy.Object
}

func NewSoundThemePlayer(conn *dbus.Conn) *SoundThemePlayer {
	obj := new(SoundThemePlayer)
	obj.Object.Init_(conn, "com.deepin.api.SoundThemePlayer", "/com/deepin/api/SoundThemePlayer")
	return obj
}

type soundThemePlayer struct{}

func (v *soundThemePlayer) GetObject_() *proxy.Object {
	return (*proxy.Object)(unsafe.Pointer(v))
}

func (*soundThemePlayer) GetInterfaceName_() string {
	return "com.deepin.api.SoundThemePlayer"
}

// method EnableSoundDesktopLogin

func (v *soundThemePlayer) GoEnableSoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".EnableSoundDesktopLogin", flags, ch, enabled)
}

func (v *soundThemePlayer) GoEnableSoundDesktopLoginWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, enabled bool) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".EnableSoundDesktopLogin", flags, ch, enabled)
}

func (v *soundThemePlayer) EnableSoundDesktopLogin(flags dbus.Flags, enabled bool) error {
	return (<-v.GoEnableSoundDesktopLogin(flags, make(chan *dbus.Call, 1), enabled).Done).Err
}

func (v *soundThemePlayer) EnableSoundDesktopLoginWithTimeout(timeout time.Duration, flags dbus.Flags, enabled bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoEnableSoundDesktopLoginWithContext(ctx, flags, make(chan *dbus.Call, 1), enabled).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method Play

func (v *soundThemePlayer) GoPlay(flags dbus.Flags, ch chan *dbus.Call, theme string, event string, device string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".Play", flags, ch, theme, event, device)
}

func (v *soundThemePlayer) GoPlayWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, theme string, event string, device string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".Play", flags, ch, theme, event, device)
}

func (v *soundThemePlayer) Play(flags dbus.Flags, theme string, event string, device string) error {
	return (<-v.GoPlay(flags, make(chan *dbus.Call, 1), theme, event, device).Done).Err
}

func (v *soundThemePlayer) PlayWithTimeout(timeout time.Duration, flags dbus.Flags, theme string, event string, device string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPlayWithContext(ctx, flags, make(chan *dbus.Call, 1), theme, event, device).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method PlaySoundDesktopLogin

func (v *soundThemePlayer) GoPlaySoundDesktopLogin(flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".PlaySoundDesktopLogin", flags, ch)
}

func (v *soundThemePlayer) GoPlaySoundDesktopLoginWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".PlaySoundDesktopLogin", flags, ch)
}

func (v *soundThemePlayer) PlaySoundDesktopLogin(flags dbus.Flags) error {
	return (<-v.GoPlaySoundDesktopLogin(flags, make(chan *dbus.Call, 1)).Done).Err
}

func (v *soundThemePlayer) PlaySoundDesktopLoginWithTimeout(timeout time.Duration, flags dbus.Flags) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoPlaySoundDesktopLoginWithContext(ctx, flags, make(chan *dbus.Call, 1)).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SaveAudioState

func (v *soundThemePlayer) GoSaveAudioState(flags dbus.Flags, ch chan *dbus.Call, activePlayback map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SaveAudioState", flags, ch, activePlayback)
}

func (v *soundThemePlayer) GoSaveAudioStateWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, activePlayback map[string]dbus.Variant) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SaveAudioState", flags, ch, activePlayback)
}

func (v *soundThemePlayer) SaveAudioState(flags dbus.Flags, activePlayback map[string]dbus.Variant) error {
	return (<-v.GoSaveAudioState(flags, make(chan *dbus.Call, 1), activePlayback).Done).Err
}

func (v *soundThemePlayer) SaveAudioStateWithTimeout(timeout time.Duration, flags dbus.Flags, activePlayback map[string]dbus.Variant) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSaveAudioStateWithContext(ctx, flags, make(chan *dbus.Call, 1), activePlayback).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}

// method SetSoundTheme

func (v *soundThemePlayer) GoSetSoundTheme(flags dbus.Flags, ch chan *dbus.Call, theme string) *dbus.Call {
	return v.GetObject_().Go_(v.GetInterfaceName_()+".SetSoundTheme", flags, ch, theme)
}

func (v *soundThemePlayer) GoSetSoundThemeWithContext(ctx context.Context, flags dbus.Flags, ch chan *dbus.Call, theme string) *dbus.Call {
	return v.GetObject_().GoWithContext_(ctx, v.GetInterfaceName_()+".SetSoundTheme", flags, ch, theme)
}

func (v *soundThemePlayer) SetSoundTheme(flags dbus.Flags, theme string) error {
	return (<-v.GoSetSoundTheme(flags, make(chan *dbus.Call, 1), theme).Done).Err
}

func (v *soundThemePlayer) SetSoundThemeWithTimeout(timeout time.Duration, flags dbus.Flags, theme string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	call := <-v.GoSetSoundThemeWithContext(ctx, flags, make(chan *dbus.Call, 1), theme).Done
	if call.Err == nil && ctx.Err() != nil {
		return ctx.Err()
	}

	return call.Err
}
