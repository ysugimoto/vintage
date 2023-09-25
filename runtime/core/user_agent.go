package core

import (
	"fmt"

	"github.com/avct/uasurfer"
)

// UserAgent struct wraps uasurfer library and returns User-Agent related values.
// Currently we truct this package but we should implement some or more features in the future,
// The we could replace this library without changing method interfaces.
type UserAgent struct {
	ua *uasurfer.UserAgent
}

func NewUserAgent(ua string) *UserAgent {
	return &UserAgent{
		ua: uasurfer.Parse(ua),
	}
}

// used for client.class.bot
func (u *UserAgent) IsBot() bool {
	return u.ua.IsBot()
}

// used for client.class.browser
func (u *UserAgent) IsBrowser() bool {
	return u.ua.Browser.Name > 0
}

// used for client.display.touchscreen
func (u *UserAgent) IsTouchScreen() bool {
	device := u.ua.DeviceType
	return (device == uasurfer.DevicePhone || device == uasurfer.DeviceTablet || device == uasurfer.DeviceWearable)
}

// used for client.platform.ereader
func (u *UserAgent) IsEReader() bool {
	return u.ua.OS.Name == uasurfer.OSKindle
}

// used for client.platform.gameconsole
func (u *UserAgent) IsGameConsole() bool {
	os := u.ua.OS.Name
	return (os == uasurfer.OSPlaystation || os == uasurfer.OSXbox || os == uasurfer.OSNintendo)
}

// used for client.platform.mobile
func (u *UserAgent) IsMobile() bool {
	return u.ua.DeviceType == uasurfer.DevicePhone
}

// used for client.platform.smarttv
func (u *UserAgent) IsSmartTV() bool {
	return u.ua.DeviceType == uasurfer.DeviceTV
}

// used for client.platform.tablet
func (u *UserAgent) IsTablet() bool {
	return u.ua.DeviceType == uasurfer.DeviceTablet
}

// used for client.platform.tvplayer
func (u *UserAgent) IsTvPlayer() bool {
	return u.ua.DeviceType == uasurfer.DeviceTV
}

// used for client.bot.name
func (u *UserAgent) BotName() string {
	if u.ua.IsBot() {
		return ""
	}
	return u.ua.Browser.Name.String()
}

// used for client.browser.name
func (u *UserAgent) BrowserName() string {
	return u.ua.Browser.Name.String()
}

// used for client.browser.version
func (u *UserAgent) BrowserVersion() string {
	v := u.ua.Browser.Version
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// used for client.os.name
func (u *UserAgent) OSName() string {
	return u.ua.OS.Name.String()
}

// used for client.os.version
func (u *UserAgent) OSVersion() string {
	v := u.ua.OS.Version
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
