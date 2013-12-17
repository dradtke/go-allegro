package physfs

/*
#cgo pkg-config: allegro_physfs-5.0
#include <allegro5/allegro.h>
#include <allegro5/allegro_physfs.h>
*/
import "C"

// After calling this, subsequent calls to al_fopen will be handled by
// PHYSFS_open(). Operations on the files returned by al_fopen will then be
// performed through PhysicsFS.
func UseFileInterface() {
	C.al_set_physfs_file_interface()
}

// Returns the (compiled) version of the addon, in the same format as
// al_get_allegro_version.
func Version() uint32 {
	return uint32(C.al_get_allegro_physfs_version())
}

