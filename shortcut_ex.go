package fyne

// ShortcutNext describes a shortcut redo action.
//
// Since: 2.5
type ShortcutNext struct{}

var _ KeyboardShortcut = (*ShortcutNext)(nil)

// Key returns the KeyName for this shortcut.
//
// Implements: KeyboardShortcut
func (se *ShortcutNext) Key() KeyName {
	return KeyRight
}

// Mod returns the KeyModifier for this shortcut.
//
// Implements: KeyboardShortcut
func (se *ShortcutNext) Mod() KeyModifier {
	return KeyModifierShortcutDefault
}

// ShortcutName returns the shortcut name
func (se *ShortcutNext) ShortcutName() string {
	return "Next"
}

// ShortcutPrev describes a shortcut redo action.
//
// Since: 2.5
type ShortcutPrev struct{}

var _ KeyboardShortcut = (*ShortcutPrev)(nil)

// Key returns the KeyName for this shortcut.
//
// Implements: KeyboardShortcut
func (se *ShortcutPrev) Key() KeyName {
	return KeyLeft
}

// Mod returns the KeyModifier for this shortcut.
//
// Implements: KeyboardShortcut
func (se *ShortcutPrev) Mod() KeyModifier {
	return KeyModifierShortcutDefault
}

// ShortcutName returns the shortcut name
func (se *ShortcutPrev) ShortcutName() string {
	return "Prev"
}
