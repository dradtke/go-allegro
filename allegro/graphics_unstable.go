// +build unstable

package allegro

// #include <allegro5/allegro.h>
import "C"

func SetNewBitmapDepth(depth int) {
	C.al_set_new_bitmap_depth(C.int(depth))
}

func NewBitmapDepth() int {
	return int(C.al_get_new_bitmap_depth())
}
