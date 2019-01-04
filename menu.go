package webapp

import (
	"github.com/jholder85638/cef/cef"
	"github.com/jholder85638/webapp/keys"
)

// Pre-defined menu IDs. Apps should start their IDs at MenuIDUserBase.
const (
	MenuIDAppMenu = int(cef.MenuIDUserFirst) + iota
	MenuIDFileMenu
	MenuIDEditMenu
	MenuIDWindowMenu
	MenuIDHelpMenu
	MenuIDServicesMenu
	MenuIDAboutItem
	MenuIDDevToolsItem
	MenuIDPreferencesItem
	MenuIDQuitItem
	MenuIDCutItem
	MenuIDCopyItem
	MenuIDPasteItem
	MenuIDDeleteItem
	MenuIDSelectAllItem
	MenuIDMinimizeItem
	MenuIDZoomItem
	MenuIDBringAllWindowsToFrontItem
	MenuIDCloseItem
	MenuIDHideItem
	MenuIDHideOthersItem
	MenuIDShowAllItem
	MenuIDUserBase = MenuIDAppMenu + 250
)

// Menu represents a set of menu items.
type Menu struct {
	PlatformData interface{}
	ID           int
	Title        string
}

// NewMenu creates a new menu.
func NewMenu(id int, title string) *Menu {
	menu := &Menu{
		Title: title,
		ID:    id,
	}
	driver.MenuInit(menu)
	return menu
}

// ItemAtIndex returns the menu item at the specified index within the menu.
func (menu *Menu) ItemAtIndex(index int) *MenuItem {
	return driver.MenuItemAtIndex(menu, index)
}

// Item returns the menu item with the specified id anywhere in the menu and
// and its sub-menus.
func (menu *Menu) Item(id int) *MenuItem {
	for i := menu.Count() - 1; i >= 0; i-- {
		item := menu.ItemAtIndex(i)
		if item.ID == id {
			return item
		}
		if item.SubMenu != nil {
			if item = item.SubMenu.Item(id); item != nil {
				return item
			}
		}
	}
	return nil
}

// SetItemAtIndexTitle sets the title of the menu item at the specified index
// within the menu.
func (menu *Menu) SetItemAtIndexTitle(index int, title string) {
	driver.MenuItemAtIndexSetTitle(menu, index, title)
}

// SetItemTitle sets the title of the menu item with the specified id anywhere
// in the menu and its sub-menus.
func (menu *Menu) SetItemTitle(id int, title string) {
	if item := menu.Item(id); item != nil {
		item.Owner.SetItemAtIndexTitle(item.Index, title)
	}
}

// InsertSeparator inserts a menu separator at the specified item index within
// this menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertSeparator(beforeIndex int) {
	driver.MenuInsertSeparator(menu, beforeIndex)
}

// InsertItem inserts a menu item at the specified item index within this
// menu. Pass in a negative index to append to the end. Both 'validator' and
// 'handler' may be nil for default behavior.
func (menu *Menu) InsertItem(beforeIndex, id int, title string, key *keys.Key, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
	if validator == nil {
		validator = func() bool { return true }
	}
	if handler == nil {
		handler = func() {}
	}
	driver.MenuInsertItem(menu, beforeIndex, id, title, key, keyModifiers, validator, handler)
}

// InsertMenu inserts a new sub-menu at the specified item index within this
// menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertMenu(beforeIndex, id int, title string) *Menu {
	return driver.MenuInsertMenu(menu, beforeIndex, id, title)
}

// Remove the menu item at the specified index from this menu.
func (menu *Menu) Remove(index int) {
	if index >= 0 && index < menu.Count() {
		driver.MenuRemove(menu, index)
	}
}

// Count of menu items in this menu.
func (menu *Menu) Count() int {
	return driver.MenuCount(menu)
}

// Dispose releases any operating system resources associated with this menu.
func (menu *Menu) Dispose() {
	driver.MenuDispose(menu)
}
