package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>

void *locked_region_get_data(ALLEGRO_LOCKED_REGION *region) {
	return region->data;
}

int locked_region_get_format(ALLEGRO_LOCKED_REGION *region) {
	return region->format;
}

int locked_region_get_pitch(ALLEGRO_LOCKED_REGION *region) {
	return region->pitch;
}

int locked_region_get_pixel_size(ALLEGRO_LOCKED_REGION *region) {
	return region->pixel_size;
}
*/
import "C"
import (
	"errors"
)

/* Types and Enums */

type Color C.ALLEGRO_COLOR
type Bitmap C.ALLEGRO_BITMAP
type LockedRegion C.ALLEGRO_LOCKED_REGION

type DrawFlags int

const (
	FLIP_NONE       DrawFlags = 0
	FLIP_HORIZONTAL DrawFlags = C.ALLEGRO_FLIP_HORIZONTAL
	FLIP_VERTICAL   DrawFlags = C.ALLEGRO_FLIP_VERTICAL
)

type BitmapFlags int

const (
	VIDEO_BITMAP           BitmapFlags = C.ALLEGRO_VIDEO_BITMAP
	MEMORY_BITMAP          BitmapFlags = C.ALLEGRO_MEMORY_BITMAP
	KEEP_BITMAP_FORMAT     BitmapFlags = C.ALLEGRO_KEEP_BITMAP_FORMAT
	FORCE_LOCKING          BitmapFlags = C.ALLEGRO_FORCE_LOCKING
	NO_PRESERVE_TEXTURE    BitmapFlags = C.ALLEGRO_NO_PRESERVE_TEXTURE
	ALPHA_TEST             BitmapFlags = C.ALLEGRO_ALPHA_TEST
	MIN_LINEAR             BitmapFlags = C.ALLEGRO_MIN_LINEAR
	MAG_LINEAR             BitmapFlags = C.ALLEGRO_MAG_LINEAR
	MIPMAP                 BitmapFlags = C.ALLEGRO_MIPMAP
	NO_PREMULTIPLIED_ALPHA BitmapFlags = C.ALLEGRO_NO_PREMULTIPLIED_ALPHA
)

type LockFlags int

const (
	LOCK_READONLY  LockFlags = C.ALLEGRO_LOCK_READONLY
	LOCK_WRITEONLY LockFlags = C.ALLEGRO_LOCK_WRITEONLY
	LOCK_READWRITE LockFlags = C.ALLEGRO_LOCK_READWRITE
)

type PixelFormat int

const (
	PIXEL_FORMAT_ANY               PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY
	PIXEL_FORMAT_ANY_NO_ALPHA      PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_NO_ALPHA
	PIXEL_FORMAT_ANY_WITH_ALPHA    PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_WITH_ALPHA
	PIXEL_FORMAT_ANY_15_NO_ALPHA   PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_15_NO_ALPHA
	PIXEL_FORMAT_ANY_16_NO_ALPHA   PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_16_NO_ALPHA
	PIXEL_FORMAT_ANY_16_WITH_ALPHA PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_16_WITH_ALPHA
	PIXEL_FORMAT_ANY_24_NO_ALPHA   PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_24_NO_ALPHA
	PIXEL_FORMAT_ANY_32_NO_ALPHA   PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_32_NO_ALPHA
	PIXEL_FORMAT_ANY_32_WITH_ALPHA PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY_32_WITH_ALPHA
	PIXEL_FORMAT_ARGB_8888         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ARGB_8888
	PIXEL_FORMAT_RGBA_8888         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGBA_8888
	PIXEL_FORMAT_ARGB_4444         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ARGB_4444
	PIXEL_FORMAT_RGB_888           PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGB_888
	PIXEL_FORMAT_RGB_565           PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGB_565
	PIXEL_FORMAT_RGB_555           PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGB_555
	PIXEL_FORMAT_RGBA_5551         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGBA_5551
	PIXEL_FORMAT_ARGB_1555         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ARGB_1555
	PIXEL_FORMAT_ABGR_8888         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888
	PIXEL_FORMAT_XBGR_8888         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_XBGR_8888
	PIXEL_FORMAT_BGR_888           PixelFormat = C.ALLEGRO_PIXEL_FORMAT_BGR_888
	PIXEL_FORMAT_BGR_565           PixelFormat = C.ALLEGRO_PIXEL_FORMAT_BGR_565
	PIXEL_FORMAT_BGR_555           PixelFormat = C.ALLEGRO_PIXEL_FORMAT_BGR_555
	PIXEL_FORMAT_RGBX_8888         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGBX_8888
	PIXEL_FORMAT_XRGB_8888         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_XRGB_8888
	PIXEL_FORMAT_ABGR_F32          PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ABGR_F32
	PIXEL_FORMAT_ABGR_8888_LE      PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888_LE
	PIXEL_FORMAT_RGBA_4444         PixelFormat = C.ALLEGRO_PIXEL_FORMAT_RGBA_4444
)

/* Static Methods */

func NewBitmapFormat() PixelFormat {
	return PixelFormat(C.al_get_new_bitmap_format())
}

func SetNewBitmapFormat(format PixelFormat) {
	C.al_set_new_bitmap_format(C.int(format))
}

func NewBitmapFlags() BitmapFlags {
	return BitmapFlags(C.al_get_new_bitmap_flags())
}

func SetNewBitmapFlags(flags BitmapFlags) {
	C.al_set_new_bitmap_flags(C.int(flags))
}

func AddNewBitmapFlag(flags BitmapFlags) {
	C.al_add_new_bitmap_flag(C.int(flags))
}

func CreateBitmap(w, h int) *Bitmap {
	return (*Bitmap)(C.al_create_bitmap(C.int(w), C.int(h)))
}

func MapRGB(r, g, b byte) Color {
	return Color(C.al_map_rgb(C.uchar(r), C.uchar(g), C.uchar(b)))
}

func ClearToColor(c Color) {
	C.al_clear_to_color(C.ALLEGRO_COLOR(c))
}

func LoadBitmap(filename string) (*Bitmap, error) {
	filename_ := C.CString(filename)
	defer FreeString(filename_)
	bmp := C.al_load_bitmap(filename_)
	if bmp == nil {
		return nil, errors.New("failed to load bitmap at '" + filename + "'")
	}
	return (*Bitmap)(bmp), nil
}

func HoldDrawing(hold bool) {
	C.al_hold_bitmap_drawing(C.bool(hold))
}

func IsDrawingHeld() bool {
	return bool(C.al_is_bitmap_drawing_held())
}

/* Bitmap Instance Methods */

func (bmp *Bitmap) Format() PixelFormat {
	return PixelFormat(C.al_get_bitmap_format((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Flags() BitmapFlags {
	return BitmapFlags(C.al_get_bitmap_flags((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Destroy() {
	C.al_destroy_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

func (bmp *Bitmap) Width() int {
	return (int)(C.al_get_bitmap_width((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Height() int {
	return (int)(C.al_get_bitmap_height((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Draw(dx, dy float32, flags DrawFlags) {
	C.al_draw_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.float(dx),
		C.float(dy),
		C.int(flags))
}

func (bmp *Bitmap) DrawRegion(sx, sy, sw, sh, dx, dy float32, flags DrawFlags) {
	C.al_draw_bitmap_region((*C.ALLEGRO_BITMAP)(bmp),
		C.float(sx),
		C.float(sy),
		C.float(sw),
		C.float(sh),
		C.float(dx),
		C.float(dy),
		C.int(flags))
}

func (bmp *Bitmap) DrawScaled(sx, sy, sw, sh, dx, dy, dw, dh float32, flags DrawFlags) {
	C.al_draw_scaled_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.float(sx),
		C.float(sy),
		C.float(sw),
		C.float(sh),
		C.float(dx),
		C.float(dy),
		C.float(dw),
		C.float(dh),
		C.int(flags))
}

func (bmp *Bitmap) DrawRotated(cx, cy, dx, dy, angle float32, flags DrawFlags) {
	C.al_draw_rotated_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.float(cx),
		C.float(cy),
		C.float(dx),
		C.float(dy),
		C.float(angle),
		C.int(flags))
}

func (bmp *Bitmap) DrawScaledRotated(cx, cy, dx, dy, xscale, yscale, angle float32, flags DrawFlags) {
	C.al_draw_scaled_rotated_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.float(cx),
		C.float(cy),
		C.float(dx),
		C.float(dy),
		C.float(xscale),
		C.float(yscale),
		C.float(angle),
		C.int(flags))
}

func (bmp *Bitmap) DrawTinted(tint Color, dx, dy float32, flags DrawFlags) {
	C.al_draw_tinted_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(dx),
		C.float(dy),
		C.int(flags))
}

func (bmp *Bitmap) DrawTintedRegion(tint Color, sx, sy, sw, sh, dx, dy float32, flags DrawFlags) {
	C.al_draw_tinted_bitmap_region((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(sx),
		C.float(sy),
		C.float(sw),
		C.float(sh),
		C.float(dx),
		C.float(dy),
		C.int(flags))
}

func (bmp *Bitmap) DrawTintedScaled(tint Color, sx, sy, sw, sh, dx, dy, dw, dh float32, flags DrawFlags) {
	C.al_draw_tinted_scaled_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(sx),
		C.float(sy),
		C.float(sw),
		C.float(sh),
		C.float(dx),
		C.float(dy),
		C.float(dw),
		C.float(dh),
		C.int(flags))
}

func (bmp *Bitmap) DrawTintedRotated(tint Color, cx, cy, dx, dy, angle float32, flags DrawFlags) {
	C.al_draw_tinted_rotated_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(cx),
		C.float(cy),
		C.float(dx),
		C.float(dy),
		C.float(angle),
		C.int(flags))
}

func (bmp *Bitmap) DrawTintedScaledRotated(tint Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags DrawFlags) {
	C.al_draw_tinted_scaled_rotated_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(cx),
		C.float(cy),
		C.float(dx),
		C.float(dy),
		C.float(xscale),
		C.float(yscale),
		C.float(angle),
		C.int(flags))
}

func (bmp *Bitmap) Lock(format PixelFormat, flags LockFlags) (*LockedRegion, error) {
	reg := C.al_lock_bitmap((*C.ALLEGRO_BITMAP)(bmp), C.int(format), C.int(flags))
	if reg == nil {
		return nil, errors.New("failed to lock bitmap; is it already locked?")
	}
	return (*LockedRegion)(reg), nil
}

func (bmp *Bitmap) LockRegion(x, y, width, height int, format PixelFormat, flags LockFlags) (*LockedRegion, error) {
	reg := C.al_lock_bitmap_region((*C.ALLEGRO_BITMAP)(bmp),
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		C.int(format),
		C.int(flags))
	if reg == nil {
		return nil, errors.New("failed to lock bitmap region; is it already locked?")
	}
	return (*LockedRegion)(reg), nil
}

func (bmp *Bitmap) IsLocked() bool {
	return bool(C.al_is_bitmap_locked((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Unlock() {
	C.al_unlock_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

/* Locked Region Instance Methods */

func (reg *LockedRegion) Data() uintptr {
	return uintptr(C.locked_region_get_data((*C.ALLEGRO_LOCKED_REGION)(reg)))
}

func (reg *LockedRegion) Format() PixelFormat {
	return PixelFormat(C.locked_region_get_format((*C.ALLEGRO_LOCKED_REGION)(reg)))
}

func (reg *LockedRegion) Pitch() int {
	return int(C.locked_region_get_pitch((*C.ALLEGRO_LOCKED_REGION)(reg)))
}

func (reg *LockedRegion) PixelSize() int {
	return int(C.locked_region_get_pixel_size((*C.ALLEGRO_LOCKED_REGION)(reg)))
}
