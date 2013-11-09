package allegro

// WARNING: this file was written on Linux and therefore not yet tested.

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
)

type Direct3DDevice C.LPDIRECT3DDEVICE9
type Direct3DTexture C.LPDIRECT3DTEXTURE9

func (d *Display) D3DDevice() (Direct3DDevice, error) {
	device := C.al_get_d3d_device()
	if device == nil {
		return nil, errors.New("failed to get D3D device; did you forget the Direct3D display flag?")
	}
	return Direct3DDevice(device), nil
}

func (d *Display) IsD3DDeviceLost() bool {
	return bool(C.al_is_d3d_device_lost((*C.ALLEGRO_DISPLAY)(d))
}

func (bmp *Bitmap) D3DSystemTexture() (Direct3DTexture, error) {
	texture := C.al_get_d3d_system_texture((*C.ALLEGRO_BITMAP)(bmp))
	if texture == nil {
		return nil, errors.New("failed to get D3D texture")
	}
	return Direct3DTexture(texture), nil
}

func (bmp *Bitmap) TexturePosition() (int, int) {
	var u, v C.int
	C.al_get_d3d_texture_position((*C.ALLEGRO_BITMAP)(bmp), &u, &v)
	return int(u), int(v)
}

func HaveD3DNonSquareTextureSupport() bool {
	return bool(C.al_have_d3d_non_square_texture_support())
}
