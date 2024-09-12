package widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"runtime"
)

type EntryEx struct {
	Entry

	changePrevFn    func()
	changeNextFn    func()
	createItemFn    func()
	formatContentFn func(string) string
}

func NewEntryEx() *EntryEx {
	e := new(EntryEx)
	e.ExtendBaseWidget(e)
	return e
}

// NewMultiLineEntryEx creates a new entry that allows multiple lines
func NewMultiLineEntryEx(rightFn, leftFn, createItemFn func(), fcFn func(string) string) *EntryEx {
	e := &EntryEx{Entry: Entry{MultiLine: true, Wrapping: fyne.TextWrap(fyne.TextTruncateClip)}, changeNextFn: rightFn, changePrevFn: leftFn, createItemFn: createItemFn, formatContentFn: fcFn}
	e.ExtendBaseWidget(e)
	return e
}

// ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.
func (e *EntryEx) ExtendBaseWidget(wid fyne.Widget) {
	impl := e.super()
	if impl != nil {
		return
	}

	e.impl.Store(&wid)

	e.propertyLock.Lock()
	e.registerShortcut()
	e.propertyLock.Unlock()
}

func (e *EntryEx) registerShortcut() {
	moveWordModifier := fyne.KeyModifierShortcutDefault
	if runtime.GOOS == "darwin" {
		moveWordModifier = fyne.KeyModifierAlt
	}

	e.Entry.registerShortcut()

	e.shortcut.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyDown, Modifier: moveWordModifier},
		func(fyne.Shortcut) {
			if e.changeNextFn != nil {
				e.changeNextFn()
			}
		})
	e.shortcut.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyUp, Modifier: moveWordModifier},
		func(fyne.Shortcut) {
			if e.changePrevFn != nil {
				e.changePrevFn()
			}
		})
	e.shortcut.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyEqual, Modifier: moveWordModifier},
		func(fyne.Shortcut) {
			if e.createItemFn != nil {
				e.createItemFn()
			}
		})
	e.shortcut.AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyJ, Modifier: moveWordModifier},
		func(fyne.Shortcut) {
			if e.formatContentFn != nil {
				e.Text = e.formatContentFn(e.Text)
				e.Refresh()
			}
		})
}
