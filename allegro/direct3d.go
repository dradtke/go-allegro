// +build windows

package allegro

// #include <allegro5/allegro.h>
// #include <allegro5/allegro_direct3d.h>
import "C"
import (
	"errors"
)

const (
	DIRECT3D DisplayFlags = C.ALLEGRO_DIRECT3D_INTERNAL
)

type Direct3DDevice C.LPDIRECT3DDEVICE9
type Direct3DTexture C.LPDIRECT3DTEXTURE9

// Returns the Direct3D device of the display. The return value is undefined if
// the display was not created with the Direct3D flag.
//
// See https://liballeg.org/a5docs/5.2.6/direct3d.html#al_get_d3d_device
func (d *Display) D3DDevice() (Direct3DDevice, error) {
	device := C.al_get_d3d_device((*C.ALLEGRO_DISPLAY)(d))
	if device == nil {
		return nil, errors.New("failed to get D3D device; did you forget the Direct3D display flag?")
	}
	return Direct3DDevice(device), nil
}

// Returns a boolean indicating whether or not the Direct3D device belonging to
// the given display is in a lost state.
//
// See https://liballeg.org/a5docs/5.2.6/direct3d.html#al_is_d3d_device_lost
func (d *Display) IsD3DDeviceLost() bool {
	return bool(C.al_is_d3d_device_lost((*C.ALLEGRO_DISPLAY)(d)))
}

// Returns the system texture (stored with the D3DPOOL_SYSTEMMEM flags). This
// texture is used for the render-to-texture feature set.
//
// See https://liballeg.org/a5docs/5.2.6/direct3d.html#al_get_d3d_system_texture
func (bmp *Bitmap) D3DSystemTexture() (Direct3DTexture, error) {
	texture := C.al_get_d3d_system_texture((*C.ALLEGRO_BITMAP)(bmp))
	if texture == nil {
		return nil, errors.New("failed to get D3D texture")
	}
	return Direct3DTexture(texture), nil
}

// Returns the u/v coordinates for the top/left corner of the bitmap within the
// used texture, in pixels.
//
// See https://liballeg.org/a5docs/5.2.6/direct3d.html#al_get_d3d_texture_position
func (bmp *Bitmap) TexturePosition() (int, int) {
	var u, v C.int
	C.al_get_d3d_texture_position((*C.ALLEGRO_BITMAP)(bmp), &u, &v)
	return int(u), int(v)
}

// Returns whether the Direct3D device supports textures that are not square.
//
// See https://liballeg.org/a5docs/5.2.6/direct3d.html#al_have_d3d_non_square_texture_support
func HaveD3DNonSquareTextureSupport() bool {
	return bool(C.al_have_d3d_non_square_texture_support())
}
