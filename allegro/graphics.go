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
	"fmt"
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

type BlendingOperation int

const (
	ADD            BlendingOperation = C.ALLEGRO_ADD
	DEST_MINUS_SRC BlendingOperation = C.ALLEGRO_DEST_MINUS_SRC
	SRC_MINUS_DEST BlendingOperation = C.ALLEGRO_SRC_MINUS_DEST
)

type BlendingValue int

const (
	ZERO               BlendingValue = C.ALLEGRO_ZERO
	ONE                BlendingValue = C.ALLEGRO_ONE
	ALPHA              BlendingValue = C.ALLEGRO_ALPHA
	INVERSE_ALPHA      BlendingValue = C.ALLEGRO_INVERSE_ALPHA
	// These need at least Allegro 5.0.10
	/*
	SRC_COLOR          BlendingValue = C.ALLEGRO_SRC_COLOR
	DEST_COLOR         BlendingValue = C.ALLEGRO_DEST_COLOR
	INVERSE_SRC_COLOR  BlendingValue = C.ALLEGRO_INVERSE_SRC_COLOR
	INVERSE_DEST_COLOR BlendingValue = C.ALLEGRO_INVERSE_DEST_COLOR
	*/
)

// Static Methods {{{

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

func ClearToColor(c Color) {
	C.al_clear_to_color(C.ALLEGRO_COLOR(c))
}

func LoadBitmap(filename string) (*Bitmap, error) {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	bmp := C.al_load_bitmap(filename_)
	if bmp == nil {
		return nil, fmt.Errorf("failed to load bitmap at '%s'", filename)
	}
	return (*Bitmap)(bmp), nil
}

func HoldBitmapDrawing(hold bool) {
	C.al_hold_bitmap_drawing(C.bool(hold))
}

func IsBitmapDrawingHeld() bool {
	return bool(C.al_is_bitmap_drawing_held())
}

func SetTargetBitmap(bmp *Bitmap) {
	C.al_set_target_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

func TargetBitmap() *Bitmap {
	return (*Bitmap)(C.al_get_target_bitmap())
}

func PutPixel(x, y int, color Color) {
	C.al_put_pixel(C.int(x), C.int(y), C.ALLEGRO_COLOR(color))
}

func PutBlendedPixel(x, y int, color Color) {
	C.al_put_blended_pixel(C.int(x), C.int(y), C.ALLEGRO_COLOR(color))
}

func DrawPixel(x, y float32, color Color) {
	C.al_draw_pixel(C.float(x), C.float(y), C.ALLEGRO_COLOR(color))
}

func SetClippingRectangle(x, y, width, height int) {
	C.al_set_clipping_rectangle(C.int(x), C.int(y), C.int(width), C.int(height))
}

func ResetClippingRectangle() {
	C.al_reset_clipping_rectangle()
}

func ClippingRectangle() (x, y, w, h int) {
	var cx, cy, cw, ch C.int
	C.al_get_clipping_rectangle(&cx, &cy, &cw, &ch)
	return int(cx), int(cy), int(cw), int(ch)
}

func SetBlender(op BlendingOperation, src, dst BlendingValue) {
	C.al_set_blender(C.int(op), C.int(src), C.int(dst))
}

func Blender() (op BlendingOperation, src, dst BlendingValue) {
	var cop, csrc, cdst C.int
	C.al_get_blender(&cop, &csrc, &cdst)
	return BlendingOperation(cop), BlendingValue(csrc), BlendingValue(cdst)
}

func CurrentDisplay() *Display {
	return (*Display)(C.al_get_current_display())
}

func SetTargetBackbuffer(d *Display) {
	C.al_set_target_backbuffer((*C.ALLEGRO_DISPLAY)(d))
}

//}}}

// Bitmap Instance Methods {{{

func (bmp *Bitmap) Save(filename string) error {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	ok := C.al_save_bitmap(filename_, (*C.ALLEGRO_BITMAP)(bmp))
	if !ok {
		return fmt.Errorf("failed to save bitmap at '%s'", filename)
	}
	return nil
}

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
		C.int(flags),
	)
}

func (bmp *Bitmap) DrawRegion(sx, sy, sw, sh, dx, dy float32, flags DrawFlags) {
	C.al_draw_bitmap_region((*C.ALLEGRO_BITMAP)(bmp),
		C.float(sx),
		C.float(sy),
		C.float(sw),
		C.float(sh),
		C.float(dx),
		C.float(dy),
		C.int(flags),
	)
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
		C.int(flags),
	)
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

func (bmp *Bitmap) Parent() (*Bitmap, error) {
	parent := C.al_get_parent_bitmap((*C.ALLEGRO_BITMAP)(bmp))
	if parent == nil {
		return nil, errors.New("bitmap has no parent")
	}
	return (*Bitmap)(parent), nil
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
		C.int(flags),
	)
}

func (bmp *Bitmap) DrawTinted(tint Color, dx, dy float32, flags DrawFlags) {
	C.al_draw_tinted_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(dx),
		C.float(dy),
		C.int(flags),
	)
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
		C.int(flags),
	)
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
		C.int(flags),
	)
}

func (bmp *Bitmap) DrawTintedRotated(tint Color, cx, cy, dx, dy, angle float32, flags DrawFlags) {
	C.al_draw_tinted_rotated_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(cx),
		C.float(cy),
		C.float(dx),
		C.float(dy),
		C.float(angle),
		C.int(flags),
	)
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
		C.int(flags),
	)
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
		C.int(flags),
	)
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

func (bmp *Bitmap) CreateSubBitmap(x, y, w, h int) (*Bitmap, error) {
	sub := C.al_create_sub_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.int(x), C.int(y), C.int(w), C.int(h))
	if sub == nil {
		return nil, errors.New("failed to create sub-bitmap")
	}
	return (*Bitmap)(sub), nil
}

func (bmp *Bitmap) IsSubBitmap() bool {
	return bool(C.al_is_sub_bitmap((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) ParentBitmap() (*Bitmap, error) {
	par := C.al_get_parent_bitmap((*C.ALLEGRO_BITMAP)(bmp))
	if par == nil {
		return nil, errors.New("no parent bitmap")
	}
	return (*Bitmap)(par), nil
}

func (bmp *Bitmap) Clone() (*Bitmap, error) {
	clone := C.al_clone_bitmap((*C.ALLEGRO_BITMAP)(bmp))
	if clone == nil {
		return nil, errors.New("failed to clone bitmap")
	}
	return (*Bitmap)(clone), nil
}

func (bmp *Bitmap) IsCompatible() bool {
	return bool(C.al_is_compatible_bitmap((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) BitmapFlags() BitmapFlags {
	return BitmapFlags(C.al_get_bitmap_flags((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) BitmapFormat() PixelFormat {
	return PixelFormat(C.al_get_bitmap_format((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) Pixel(x, y int) Color {
	return (Color)(C.al_get_pixel((*C.ALLEGRO_BITMAP)(bmp), C.int(x), C.int(y)))
}

func (bmp *Bitmap) ConvertMaskToAlpha(mask_color Color) {
	C.al_convert_mask_to_alpha((*C.ALLEGRO_BITMAP)(bmp), C.ALLEGRO_COLOR(mask_color))
}

//}}}

// Color Methods {{{

func MapRGB(r, g, b byte) Color {
	return Color(C.al_map_rgb(C.uchar(r), C.uchar(g), C.uchar(b)))
}

func MapRGBA(r, g, b, a byte) Color {
	return (Color)(C.al_map_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a)))
}

// ??? name?
func MapRGBf(r, g, b float32) Color {
	return (Color)(C.al_map_rgb_f(C.float(r), C.float(g), C.float(b)))
}

func MapRGBAf(r, g, b, a float32) Color {
	return (Color)(C.al_map_rgba_f(C.float(r), C.float(g), C.float(b), C.float(a)))
}

func (c Color) UnmapRGB() (byte, byte, byte) {
	var r, g, b C.uchar
	C.al_unmap_rgb((C.ALLEGRO_COLOR)(c), &r, &g, &b)
	return byte(r), byte(g), byte(b)
}

func (c Color) UnmapRGBA() (byte, byte, byte, byte) {
	var r, g, b, a C.uchar
	C.al_unmap_rgba((C.ALLEGRO_COLOR)(c), &r, &g, &b, &a)
	return byte(r), byte(g), byte(b), byte(a)
}

func (c Color) UnmapRGBf() (float32, float32, float32) {
	var r, g, b C.float
	C.al_unmap_rgb_f((C.ALLEGRO_COLOR)(c), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

func (c Color) UnmapRGBAf() (float32, float32, float32, float32) {
	var r, g, b, a C.float
	C.al_unmap_rgba_f((C.ALLEGRO_COLOR)(c), &r, &g, &b, &a)
	return float32(r), float32(g), float32(b), float32(a)
}

//}}}

// Miscellaneous Instance Methods {{{

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

func (format PixelFormat) PixelSize() int {
	return int(C.al_get_pixel_size(C.int(format)))
}

func (format PixelFormat) PixelFormatBits() int {
	return int(C.al_get_pixel_format_bits(C.int(format)))
}

//}}}
