package stdmenu

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

// NewWindowMenu creates a standard 'Window' menu.
func NewWindowMenu() *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuTagWindowMenu, i18n.Text("Window"))
	InsertMinimizeItem(menu, -1)
	InsertZoomItem(menu, -1)
	menu.InsertSeparator(-1)
	menu.InsertItem(-1, webapp.MenuTagBringAllWindowsToFrontItem, i18n.Text("Bring All to Front"), 0, 0, nil, webapp.AllWindowsToFront)
	return menu
}

// InsertMinimizeItem creates the standard "Minimize" menu item that will
// issue the Minimize command to the current key window when chosen.
func InsertMinimizeItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuTagMinimizeItem, i18n.Text("Minimize"), keys.VirtualKeyM, keys.PlatformMenuModifier(), MinimizeValidator, MinimizeHandler)
}

// MinimizeValidator provides the standard validation function for the
// "Minimize" menu item.
func MinimizeValidator() bool {
	w := webapp.KeyWindow()
	return w != nil && w.Minimizable()
}

// MinimizeHandler provides the standard handler function for the "Minimize"
// menu item.
func MinimizeHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		wnd.Minimize()
	}
}

// InsertZoomItem creates the standard "Zoom" menu item that will issue the
// Zoom command to the current key window when chosen.
func InsertZoomItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuTagZoomItem, i18n.Text("Zoom"), keys.VirtualKeyZ, keys.ShiftModifier|keys.PlatformMenuModifier(), ZoomValidator, ZoomHandler)
}

// ZoomValidator provides the standard validation function for the "Zoom" menu
// item.
func ZoomValidator() bool {
	w := webapp.KeyWindow()
	return w != nil && w.Resizable()
}

// ZoomHandler provides the standard handler function for the "Zoom" menu
// item.
func ZoomHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		wnd.Zoom()
	}
}