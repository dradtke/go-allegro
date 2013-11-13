package physfs

/*
#cgo pkg-config: allegro_physfs-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_physfs.h>
*/
import "C"

func UseFileInterface() {
	C.al_set_physfs_file_interface()
}

func Version() uint32 {
	return uint32(C.al_get_allegro_physfs_version())
}
