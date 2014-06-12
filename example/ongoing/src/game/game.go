package game

import (
	al "github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"os"
)

// Fatal() shows an error message box, then quits the
// game when the user clicks 'Close'.
func Fatal(err error) {
	dialog.ShowNativeMessageBoxWithButtons(display, "Application Error", "", err.Error(), []string{"Close"}, dialog.MESSAGEBOX_ERROR)
	Exit(1)
}

// Exit() causes the game to quit with the provided
// error code.
func Exit(code int) {
	Cleanup()
	al.Uninstall()
	os.Exit(code)
}
