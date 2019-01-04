package webapp

import (
	"unsafe"

	"github.com/jholder85638/cef/cef"
	"github.com/jholder85638/toolbox/xmath/geom"
	"github.com/jholder85638/webapp/keys"
)

// Driver defines the required functions each platform driver must provide.
type Driver interface {
	PrepareForStart() error
	PrepareForEventLoop()
	RunEventLoop()
	OnPreKeyEvent(event *cef.KeyEvent, isKeyboardShortcut *int32) int32
	OnKeyEvent(event *cef.KeyEvent) int32

	AttemptQuit()
	MayQuitNow(quit bool)

	MenuBarForWindow(wnd *Window) (bar *MenuBar, isGlobal, isFirst bool)
	MenuBarMenuAtIndex(bar *MenuBar, index int) *Menu
	MenuBarInsert(bar *MenuBar, beforeIndex int, menu *Menu)
	MenuBarRemove(bar *MenuBar, index int)
	MenuBarCount(bar *MenuBar) int
	MenuBarHeightInWindow() float64

	MenuInit(menu *Menu)
	MenuItemAtIndex(menu *Menu, index int) *MenuItem
	MenuItemAtIndexSetTitle(menu *Menu, index int, title string)
	MenuInsertSeparator(menu *Menu, beforeIndex int)
	MenuInsertItem(menu *Menu, beforeIndex, id int, title string, key *keys.Key, keyModifiers keys.Modifiers, validator func() bool, handler func())
	MenuInsertMenu(menu *Menu, beforeIndex, id int, title string) *Menu
	MenuRemove(menu *Menu, index int)
	MenuCount(menu *Menu) int
	MenuDispose(menu *Menu)

	Displays() []*Display
	KeyWindow() *Window
	BringAllWindowsToFront()

	WindowInit(wnd *Window, style StyleMask, bounds geom.Rect, title string) error
	WindowBrowserParent(wnd *Window) unsafe.Pointer
	WindowClose(wnd *Window)
	WindowSetTitle(wnd *Window, title string)
	WindowBounds(wnd *Window) geom.Rect
	WindowSetBounds(wnd *Window, bounds geom.Rect)
	WindowContentSize(wnd *Window) geom.Size
	WindowToFront(wnd *Window)
	WindowMinimize(wnd *Window)
	WindowZoom(wnd *Window)
}

// AppVisibilityController defines optional APIs a platform can provide for
// manipulating application visibility.
type AppVisibilityController interface {
	HideApp()
	HideOtherApps()
	ShowAllApps()
}
