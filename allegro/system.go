package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

bool _al_init() {
	return al_init();
}
*/
import "C"
import (
	"errors"
)

func init() {
	if !bool(C._al_init()) {
		panic("failed to initialize allegro!")
	}
}

func Version() uint32 {
	return uint32(C.al_get_allegro_version())
}

func SystemConfig() (*Config, error) {
	cfg := C.al_get_system_config()
	if cfg == nil {
		return nil, errors.New("no system config found")
	}
	return (*Config)(cfg), nil
}

func SetExeName(path string) {
	path_ := C.CString(path)
	defer freeString(path_)
	C.al_set_exe_name(path_)
}

func SetOrgName(name string) {
	name_ := C.CString(name)
	defer freeString(name_)
	C.al_set_org_name(name_)
}

func SetAppName(name string) {
	name_ := C.CString(name)
	defer freeString(name_)
	C.al_set_app_name(name_)
}

func OrgName() string {
	return C.GoString(C.al_get_org_name())
}

func AppName() string {
	return C.GoString(C.al_get_app_name())
}

