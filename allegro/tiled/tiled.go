// Package tiled implements support for the Allegro Tiled addon.
package tiled

// #include <allegro5/allegro_tiled.h>
// #include "../util.c"
import "C"
import (
	"fmt"
	"path/filepath"
)

type Map *C.ALLEGRO_MAP

func OpenMap(filename string) (*Map, error) {
	base := filepath.Base(filename)
	base_ := C.CString(base)
	defer C.free_string(base_)
	dir := filename[:base]
	dir_ := C.CString(dir)
	defer C.free_string(dir_)
	m := C.al_open_map(dir_, base_)
	if m == nil {
		return nil, fmt.Errorf("failed to load map file: %s", filename)
	}
	return (*Map)(m), nil
}

// TODO: add a bunch of drawing methods

func (m *Map) Width() int {
	return int(C.al_get_map_width((*C.ALLEGRO_MAP)(m)))
}

func (m *Map) Height() int {
	return int(C.al_get_map_height((*C.ALLEGRO_MAP)(m)))
}

type Tile *C.ALLEGRO_MAP_TILE

type (t *Tile) Prop(name, def string) string {
    name_ := C.CString(name)
    defer C.free_string(name_)
    def_ := C.CString(def)
    defer C.free_string(def_)
    p := C.al_get_tile_property((*C.ALLEGRO_MAP_TILE)(t), name_, def_)
    defer C.free_string(p)
    return C.GoString(p)
}

type Object *C.ALLEGRO_MAP_OBJECT

type (o *Object) Prop(name, def string) string {
    name_ := C.CString(name)
    defer C.free_string(name_)
    def_ := C.CString(def)
    defer C.free_string(def_)
    p := C.al_get_tile_property((*C.ALLEGRO_MAP_OBJECT)(o), name_, def_)
    defer C.free_string(p)
    return C.GoString(p)
}
