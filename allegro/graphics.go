package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

const rgbaMAX = 0xFFFF

/* Types and Enums */

var (
	BitmapIsNull      = errors.New("bitmap is null")
	BitmapHasNoParent = errors.New("bitmap has no parent")
)

type Color C.ALLEGRO_COLOR
type Bitmap C.ALLEGRO_BITMAP
type LockedRegion C.struct_ALLEGRO_LOCKED_REGION

type DrawFlags int

const (
	FLIP_NONE       DrawFlags = 0
	FLIP_HORIZONTAL           = C.ALLEGRO_FLIP_HORIZONTAL
	FLIP_VERTICAL             = C.ALLEGRO_FLIP_VERTICAL
)

type BitmapFlags int

const (
	ALPHA_TEST             BitmapFlags = C.ALLEGRO_ALPHA_TEST
	FORCE_LOCKING                      = C.ALLEGRO_FORCE_LOCKING
	KEEP_BITMAP_FORMAT                 = C.ALLEGRO_KEEP_BITMAP_FORMAT
	KEEP_INDEX                         = C.ALLEGRO_KEEP_INDEX
	MAG_LINEAR                         = C.ALLEGRO_MAG_LINEAR
	MEMORY_BITMAP                      = C.ALLEGRO_MEMORY_BITMAP
	MIN_LINEAR                         = C.ALLEGRO_MIN_LINEAR
	MIPMAP                             = C.ALLEGRO_MIPMAP
	NO_PREMULTIPLIED_ALPHA             = C.ALLEGRO_NO_PREMULTIPLIED_ALPHA
	NO_PRESERVE_TEXTURE                = C.ALLEGRO_NO_PRESERVE_TEXTURE
	VIDEO_BITMAP                       = C.ALLEGRO_VIDEO_BITMAP
)

type LockFlags int

const (
	LOCK_READONLY  LockFlags = C.ALLEGRO_LOCK_READONLY
	LOCK_WRITEONLY           = C.ALLEGRO_LOCK_WRITEONLY
	LOCK_READWRITE           = C.ALLEGRO_LOCK_READWRITE
)

type PixelFormat int

const (
	PIXEL_FORMAT_ANY               PixelFormat = C.ALLEGRO_PIXEL_FORMAT_ANY
	PIXEL_FORMAT_ANY_NO_ALPHA                  = C.ALLEGRO_PIXEL_FORMAT_ANY_NO_ALPHA
	PIXEL_FORMAT_ANY_WITH_ALPHA                = C.ALLEGRO_PIXEL_FORMAT_ANY_WITH_ALPHA
	PIXEL_FORMAT_ANY_15_NO_ALPHA               = C.ALLEGRO_PIXEL_FORMAT_ANY_15_NO_ALPHA
	PIXEL_FORMAT_ANY_16_NO_ALPHA               = C.ALLEGRO_PIXEL_FORMAT_ANY_16_NO_ALPHA
	PIXEL_FORMAT_ANY_16_WITH_ALPHA             = C.ALLEGRO_PIXEL_FORMAT_ANY_16_WITH_ALPHA
	PIXEL_FORMAT_ANY_24_NO_ALPHA               = C.ALLEGRO_PIXEL_FORMAT_ANY_24_NO_ALPHA
	PIXEL_FORMAT_ANY_32_NO_ALPHA               = C.ALLEGRO_PIXEL_FORMAT_ANY_32_NO_ALPHA
	PIXEL_FORMAT_ANY_32_WITH_ALPHA             = C.ALLEGRO_PIXEL_FORMAT_ANY_32_WITH_ALPHA
	PIXEL_FORMAT_ARGB_8888                     = C.ALLEGRO_PIXEL_FORMAT_ARGB_8888
	PIXEL_FORMAT_RGBA_8888                     = C.ALLEGRO_PIXEL_FORMAT_RGBA_8888
	PIXEL_FORMAT_ARGB_4444                     = C.ALLEGRO_PIXEL_FORMAT_ARGB_4444
	PIXEL_FORMAT_RGB_888                       = C.ALLEGRO_PIXEL_FORMAT_RGB_888
	PIXEL_FORMAT_RGB_565                       = C.ALLEGRO_PIXEL_FORMAT_RGB_565
	PIXEL_FORMAT_RGB_555                       = C.ALLEGRO_PIXEL_FORMAT_RGB_555
	PIXEL_FORMAT_RGBA_5551                     = C.ALLEGRO_PIXEL_FORMAT_RGBA_5551
	PIXEL_FORMAT_ARGB_1555                     = C.ALLEGRO_PIXEL_FORMAT_ARGB_1555
	PIXEL_FORMAT_ABGR_8888                     = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888
	PIXEL_FORMAT_XBGR_8888                     = C.ALLEGRO_PIXEL_FORMAT_XBGR_8888
	PIXEL_FORMAT_BGR_888                       = C.ALLEGRO_PIXEL_FORMAT_BGR_888
	PIXEL_FORMAT_BGR_565                       = C.ALLEGRO_PIXEL_FORMAT_BGR_565
	PIXEL_FORMAT_BGR_555                       = C.ALLEGRO_PIXEL_FORMAT_BGR_555
	PIXEL_FORMAT_RGBX_8888                     = C.ALLEGRO_PIXEL_FORMAT_RGBX_8888
	PIXEL_FORMAT_XRGB_8888                     = C.ALLEGRO_PIXEL_FORMAT_XRGB_8888
	PIXEL_FORMAT_ABGR_F32                      = C.ALLEGRO_PIXEL_FORMAT_ABGR_F32
	PIXEL_FORMAT_ABGR_8888_LE                  = C.ALLEGRO_PIXEL_FORMAT_ABGR_8888_LE
	PIXEL_FORMAT_RGBA_4444                     = C.ALLEGRO_PIXEL_FORMAT_RGBA_4444
)

type BlendingOperation int

const (
	ADD            BlendingOperation = C.ALLEGRO_ADD
	DEST_MINUS_SRC                   = C.ALLEGRO_DEST_MINUS_SRC
	SRC_MINUS_DEST                   = C.ALLEGRO_SRC_MINUS_DEST
)

type BlendingValue int

const (
	ZERO          BlendingValue = C.ALLEGRO_ZERO
	ONE                         = C.ALLEGRO_ONE
	ALPHA                       = C.ALLEGRO_ALPHA
	INVERSE_ALPHA               = C.ALLEGRO_INVERSE_ALPHA
	// These need at least Allegro 5.0.10
	/*
		SRC_COLOR           = C.ALLEGRO_SRC_COLOR
		DEST_COLOR          = C.ALLEGRO_DEST_COLOR
		INVERSE_SRC_COLOR   = C.ALLEGRO_INVERSE_SRC_COLOR
		INVERSE_DEST_COLOR  = C.ALLEGRO_INVERSE_DEST_COLOR
	*/
)

// Static Methods {{{

// Returns the format used for newly created bitmaps.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_new_bitmap_format
func NewBitmapFormat() PixelFormat {
	return PixelFormat(C.al_get_new_bitmap_format())
}

// Sets the pixel format (ALLEGRO_PIXEL_FORMAT) for newly created bitmaps. The
// default format is 0 and means the display driver will choose the best format.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_new_bitmap_format
func SetNewBitmapFormat(format PixelFormat) {
	C.al_set_new_bitmap_format(C.int(format))
}

// Returns the flags used for newly created bitmaps.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_new_bitmap_flags
func NewBitmapFlags() BitmapFlags {
	return BitmapFlags(C.al_get_new_bitmap_flags())
}

// Sets the flags to use for newly created bitmaps. Valid flags are:
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_new_bitmap_flags
func SetNewBitmapFlags(flags BitmapFlags) {
	C.al_set_new_bitmap_flags(C.int(flags))
}

// A convenience function which does the same as
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_add_new_bitmap_flag
func AddNewBitmapFlag(flags BitmapFlags) {
	C.al_add_new_bitmap_flag(C.int(flags))
}

// Creates a new bitmap using the bitmap format and flags for the current
// thread. Blitting between bitmaps of differing formats, or blitting between
// memory bitmaps and display bitmaps may be slow.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_create_bitmap
func CreateBitmap(w, h int) *Bitmap {
	bitmap := (*Bitmap)(C.al_create_bitmap(C.int(w), C.int(h)))
	//runtime.SetFinalizer(bitmap, bitmap.Destroy)
	return bitmap
}

// Clear the complete target bitmap, but confined by the clipping rectangle.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_clear_to_color
func ClearToColor(c Color) {
	C.al_clear_to_color(C.ALLEGRO_COLOR(c))
}

// Loads an image file into a new ALLEGRO_BITMAP. The file type is determined
// by the extension, except if the file has no extension in which case
// al_identify_bitmap is used instead.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_load_bitmap
func LoadBitmap(filename string) (*Bitmap, error) {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	bmp := C.al_load_bitmap(filename_)
	if bmp == nil {
		return nil, fmt.Errorf("failed to load bitmap at '%s'", filename)
	}
	bitmap := (*Bitmap)(bmp)
	//runtime.SetFinalizer(bitmap, bitmap.Destroy)
	return bitmap, nil
}

// Loads an image file into a new ALLEGRO_BITMAP. The file type is determined
// by the extension, except if the file has no extension in which case
// al_identify_bitmap is used instead.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_load_bitmap_flags
func LoadBitmapFlags(filename string, flags BitmapFlags) (*Bitmap, error) {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	bmp := C.al_load_bitmap_flags(filename_, C.int(flags))
	if bmp == nil {
		return nil, fmt.Errorf("failed to load bitmap at '%s'", filename)
	}
	bitmap := (*Bitmap)(bmp)
	//runtime.SetFinalizer(bitmap, bitmap.Destroy)
	return bitmap, nil
}

// Enables or disables deferred bitmap drawing. This allows for efficient
// drawing of many bitmaps that share a parent bitmap, such as sub-bitmaps from
// a tilesheet or simply identical bitmaps. Drawing bitmaps that do not share a
// parent is less efficient, so it is advisable to stagger bitmap drawing calls
// such that the parent bitmap is the same for large number of those calls.
// While deferred bitmap drawing is enabled, the only functions that can be
// used are the bitmap drawing functions and font drawing functions. Changing
// the state such as the blending modes will result in undefined behaviour. One
// exception to this rule are the non-projection transformations. It is
// possible to set a new transformation while the drawing is held.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_hold_bitmap_drawing
func HoldBitmapDrawing(hold bool) {
	C.al_hold_bitmap_drawing(C.bool(hold))
}

// Returns whether the deferred bitmap drawing mode is turned on or off.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_is_bitmap_drawing_held
func IsBitmapDrawingHeld() bool {
	return bool(C.al_is_bitmap_drawing_held())
}

// This function selects the bitmap to which all subsequent drawing operations
// in the calling thread will draw to. To return to drawing to a display, set
// the backbuffer of the display as the target bitmap, using al_get_backbuffer.
// As a convenience, you may also use al_set_target_backbuffer.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_target_bitmap
func SetTargetBitmap(bmp *Bitmap) {
	C.al_set_target_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

// Return the target bitmap of the calling thread.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_target_bitmap
func TargetBitmap() *Bitmap {
	return (*Bitmap)(C.al_get_target_bitmap())
}

// Draw a single pixel on the target bitmap. This operation is slow on
// non-memory bitmaps. Consider locking the bitmap if you are going to use this
// function multiple times on the same bitmap. This function is not affected by
// the transformations or the color blenders.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_put_pixel
func PutPixel(x, y int, color Color) {
	C.al_put_pixel(C.int(x), C.int(y), C.ALLEGRO_COLOR(color))
}

// Like al_put_pixel, but the pixel color is blended using the current blenders
// before being drawn.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_put_blended_pixel
func PutBlendedPixel(x, y int, color Color) {
	C.al_put_blended_pixel(C.int(x), C.int(y), C.ALLEGRO_COLOR(color))
}

// Draws a single pixel at x, y. This function, unlike al_put_pixel, does
// blending and, unlike al_put_blended_pixel, respects the transformations
// (that is, the pixel's position is transformed, but its size is unaffected -
// it remains a pixel). This function can be slow if called often; if you need
// to draw a lot of pixels consider using al_draw_prim with
// ALLEGRO_PRIM_POINT_LIST from the primitives addon.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_pixel
func DrawPixel(x, y float32, color Color) {
	C.al_draw_pixel(C.float(x), C.float(y), C.ALLEGRO_COLOR(color))
}

// Set the region of the target bitmap or display that pixels get clipped to.
// The default is to clip pixels to the entire bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_clipping_rectangle
func SetClippingRectangle(x, y, width, height int) {
	C.al_set_clipping_rectangle(C.int(x), C.int(y), C.int(width), C.int(height))
}

// Equivalent to calling `al_set_clipping_rectangle(0, 0, w, h)' where w and h
// are the width and height of the target bitmap respectively.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_reset_clipping_rectangle
func ResetClippingRectangle() {
	C.al_reset_clipping_rectangle()
}

// Gets the clipping rectangle of the target bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_clipping_rectangle
func ClippingRectangle() (x, y, w, h int) {
	var cx, cy, cw, ch C.int
	C.al_get_clipping_rectangle(&cx, &cy, &cw, &ch)
	return int(cx), int(cy), int(cw), int(ch)
}

// Sets the function to use for blending for the current thread.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_blender
func SetBlender(op BlendingOperation, src, dst BlendingValue) {
	C.al_set_blender(C.int(op), C.int(src), C.int(dst))
}

// Like al_set_blender, but allows specifying a separate blending operation for
// the alpha channel. This is useful if your target bitmap also has an alpha
// channel and the two alpha channels need to be combined in a different way
// than the color components.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_separate_blender
func SetSeparateBlender(op BlendingOperation, src, dst BlendingValue, alpha_op BlendingOperation, alpha_src, alpha_dst BlendingValue) {
	C.al_set_separate_blender(
		C.int(op),
		C.int(src),
		C.int(dst),
		C.int(alpha_op),
		C.int(alpha_src),
		C.int(alpha_dst),
	)
}

// Returns the active blender for the current thread. You can pass NULL for
// values you are not interested in.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_blender
func Blender() (op BlendingOperation, src, dst BlendingValue) {
	var cop, csrc, cdst C.int
	C.al_get_blender(&cop, &csrc, &cdst)
	return BlendingOperation(cop), BlendingValue(csrc), BlendingValue(cdst)
}

// Returns the active blender for the current thread. You can pass NULL for
// values you are not interested in.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_separate_blender
func SeparateBlender() (op BlendingOperation, src, dst BlendingValue, alpha_op BlendingOperation, alpha_src, alpha_dst BlendingValue) {
	var cop, csrc, cdst, calpha_op, calpha_src, calpha_dst C.int
	C.al_get_separate_blender(&cop, &csrc, &cdst, &calpha_op, &calpha_src, &calpha_dst)
	return BlendingOperation(cop), BlendingValue(csrc), BlendingValue(cdst), BlendingOperation(calpha_op), BlendingValue(calpha_src), BlendingValue(calpha_dst)
}

// Sets the color to use for blending when using the ALLEGRO_CONST_COLOR or
// ALLEGRO_INVERSE_CONST_COLOR blend functions. See al_set_blender for more
// information.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_blend_color
func SetBlendColor(c Color) {
	C.al_set_blend_color(C.ALLEGRO_COLOR(c))
}

// Returns the color currently used for constant color blending (white by
// default).
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_blend_color
func BlendColor() Color {
	return Color(C.al_get_blend_color())
}

// Return the display that is "current" for the calling thread, or NULL if
// there is none.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_current_display
func CurrentDisplay() *Display {
	return (*Display)(C.al_get_current_display())
}

// Same as al_set_target_bitmap(al_get_backbuffer(display));
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_target_backbuffer
func SetTargetBackbuffer(d *Display) {
	C.al_set_target_backbuffer((*C.ALLEGRO_DISPLAY)(d))
}

// If you create a bitmap when there is no current display (for example because
// you have not called al_create_display in the current thread) and are using
// the ALLEGRO_CONVERT_BITMAP bitmap flag (which is set by default) then the
// bitmap will be created successfully, but as a memory bitmap. This function
// converts all such bitmaps to proper video bitmaps belonging to the current
// display.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_convert_memory_bitmaps
func ConvertMemoryBitmaps() {
	C.al_convert_memory_bitmaps()
}

//}}}

// Bitmap Instance Methods {{{

// Saves an ALLEGRO_BITMAP to an image file. The file type is determined by the
// extension.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_save_bitmap
func (bmp *Bitmap) Save(filename string) error {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	ok := C.al_save_bitmap(filename_, (*C.ALLEGRO_BITMAP)(bmp))
	if !ok {
		return fmt.Errorf("failed to save bitmap at '%s'", filename)
	}
	return nil
}

// Returns the pixel format of a bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_format
func (bmp *Bitmap) Format() PixelFormat {
	return PixelFormat(C.al_get_bitmap_format((*C.ALLEGRO_BITMAP)(bmp)))
}

// Return the flags used to create the bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_flags
func (bmp *Bitmap) Flags() BitmapFlags {
	return BitmapFlags(C.al_get_bitmap_flags((*C.ALLEGRO_BITMAP)(bmp)))
}

// Destroys the given bitmap, freeing all resources used by it. This function
// does nothing if the bitmap argument is NULL.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_destroy_bitmap
func (bmp *Bitmap) Destroy() {
	C.al_destroy_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

// Returns the width of a bitmap in pixels.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_width
func (bmp *Bitmap) Width() int {
	return (int)(C.al_get_bitmap_width((*C.ALLEGRO_BITMAP)(bmp)))
}

// Returns the height of a bitmap in pixels.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_height
func (bmp *Bitmap) Height() int {
	return (int)(C.al_get_bitmap_height((*C.ALLEGRO_BITMAP)(bmp)))
}

// For a sub-bitmap, return it's x position within the parent.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_x
func (bmp *Bitmap) X() int {
	return (int)(C.al_get_bitmap_x((*C.ALLEGRO_BITMAP)(bmp)))
}

// For a sub-bitmap, return it's y position within the parent.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_y
func (bmp *Bitmap) Y() int {
	return (int)(C.al_get_bitmap_y((*C.ALLEGRO_BITMAP)(bmp)))
}

// For a sub-bitmap, changes the parent, position and size. This is the same as
// destroying the bitmap and re-creating it with al_create_sub_bitmap - except
// the bitmap pointer stays the same. This has many uses, for example an
// animation player could return a single bitmap which can just be re-parented
// to different animation frames without having to re-draw the contents. Or a
// sprite atlas could re-arrange its sprites without having to invalidate all
// existing bitmaps.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_reparent_bitmap
func (bmp *Bitmap) Reparent(parent *Bitmap, x, y, w, h int) {
	C.al_reparent_bitmap(
		(*C.ALLEGRO_BITMAP)(bmp),
		(*C.ALLEGRO_BITMAP)(parent),
		C.int(x),
		C.int(y),
		C.int(w),
		C.int(h),
	)
}

// Converts the bitmap to the current bitmap flags and format. The bitmap will
// be as if it was created anew with al_create_bitmap but retain its contents.
// All of this bitmap's sub-bitmaps are also converted. If the new bitmap type
// is memory, then the bitmap's projection bitmap is reset to be orthographic.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_convert_bitmap
func (bmp *Bitmap) Convert() {
	C.al_convert_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

// Draws an unscaled, unrotated bitmap at the given position to the current
// target bitmap (see al_set_target_bitmap).
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_bitmap
func (bmp *Bitmap) Draw(dx, dy float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
	C.al_draw_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.float(dx),
		C.float(dy),
		C.int(flags),
	)
}

// Draws a region of the given bitmap to the target bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_bitmap_region
func (bmp *Bitmap) DrawRegion(sx, sy, sw, sh, dx, dy float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Draws a scaled version of the given bitmap to the target bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_scaled_bitmap
func (bmp *Bitmap) DrawScaled(sx, sy, sw, sh, dx, dy, dw, dh float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Draws a rotated version of the given bitmap to the target bitmap. The bitmap
// is rotated by 'angle' radians clockwise.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_rotated_bitmap
func (bmp *Bitmap) DrawRotated(cx, cy, dx, dy, angle float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
	C.al_draw_rotated_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.float(cx),
		C.float(cy),
		C.float(dx),
		C.float(dy),
		C.float(angle),
		C.int(flags))
}

// Returns the bitmap this bitmap is a sub-bitmap of. Returns NULL if this
// bitmap is not a sub-bitmap. This function always returns the real bitmap,
// and never a sub-bitmap. This might NOT match what was passed to
// al_create_sub_bitmap. Consider this code, for instance:
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_parent_bitmap
func (bmp *Bitmap) Parent() (*Bitmap, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	parent := C.al_get_parent_bitmap((*C.ALLEGRO_BITMAP)(bmp))
	if parent == nil {
		return nil, BitmapHasNoParent
	}
	return (*Bitmap)(parent), nil
}

// Like al_draw_rotated_bitmap, but can also scale the bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_scaled_rotated_bitmap
func (bmp *Bitmap) DrawScaledRotated(cx, cy, dx, dy, xscale, yscale, angle float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Like al_draw_bitmap but multiplies all colors in the bitmap with the given
// color. For example:
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_tinted_bitmap
func (bmp *Bitmap) DrawTinted(tint Color, dx, dy float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
	C.al_draw_tinted_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.ALLEGRO_COLOR(tint),
		C.float(dx),
		C.float(dy),
		C.int(flags),
	)
}

// Like al_draw_bitmap_region but multiplies all colors in the bitmap with the
// given color.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_tinted_bitmap_region
func (bmp *Bitmap) DrawTintedRegion(tint Color, sx, sy, sw, sh, dx, dy float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Like al_draw_scaled_bitmap but multiplies all colors in the bitmap with the
// given color.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_tinted_scaled_bitmap
func (bmp *Bitmap) DrawTintedScaled(tint Color, sx, sy, sw, sh, dx, dy, dw, dh float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Like al_draw_rotated_bitmap but multiplies all colors in the bitmap with the
// given color.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_tinted_rotated_bitmap
func (bmp *Bitmap) DrawTintedRotated(tint Color, cx, cy, dx, dy, angle float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Like al_draw_scaled_rotated_bitmap but multiplies all colors in the bitmap
// with the given color.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_tinted_scaled_rotated_bitmap
func (bmp *Bitmap) DrawTintedScaledRotated(tint Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
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

// Like al_draw_tinted_scaled_rotated_bitmap but you specify an area within the
// bitmap to be drawn.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_draw_tinted_scaled_rotated_bitmap_region
func (bmp *Bitmap) DrawTintedScaledRotatedRegion(sx, sy, sw, sh float32, tint Color, cx, cy, dx, dy, xscale, yscale, angle float32, flags DrawFlags) {
	if bmp == nil {
		return
	}
	C.al_draw_tinted_scaled_rotated_bitmap_region((*C.ALLEGRO_BITMAP)(bmp),
		C.float(sx),
		C.float(sy),
		C.float(sw),
		C.float(sh),
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

// Convenience method for acting on a locked Bitmap, which will automatically be
// unlocked after the function completes.
func (bmp *Bitmap) WhileLocked(format PixelFormat, flags LockFlags, f func()) error {
	_, err := bmp.Lock(format, flags)
	if err != nil {
		return err
	}
	f()
	bmp.Unlock()
	return nil
}

func (bmp *Bitmap) WithLockedTarget(format PixelFormat, flags LockFlags, f func()) error {
	if _, err := bmp.Lock(format, flags); err != nil {
		return err
	}
	bmp.AsTarget(f)
	bmp.Unlock()
	return nil
}

// Lock an entire bitmap for reading or writing. If the bitmap is a display
// bitmap it will be updated from system memory after the bitmap is unlocked
// (unless locked read only). Returns NULL if the bitmap cannot be locked, e.g.
// the bitmap was locked previously and not unlocked. This function also
// returns NULL if the format is a compressed format.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_lock_bitmap
func (bmp *Bitmap) Lock(format PixelFormat, flags LockFlags) (*LockedRegion, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	reg := C.al_lock_bitmap((*C.ALLEGRO_BITMAP)(bmp), C.int(format), C.int(flags))
	if reg == nil {
		return nil, errors.New("failed to lock bitmap; is it already locked?")
	}
	return (*LockedRegion)(reg), nil
}

// Like al_lock_bitmap, but allows locking bitmaps with a blocked pixel format
// (i.e. a format for which al_get_pixel_block_width or
// al_get_pixel_block_height do not return 1) in that format. To that end, this
// function also does not allow format conversion. For bitmap formats with a
// block size of 1, this function is identical to calling al_lock_bitmap(bmp,
// al_get_bitmap_format(bmp), flags).
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_lock_bitmap_blocked
func (bmp *Bitmap) LockBlocked(flags LockFlags) (*LockedRegion, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	reg := C.al_lock_bitmap_blocked((*C.ALLEGRO_BITMAP)(bmp), C.int(flags))
	if reg == nil {
		return nil, errors.New("failed to lock bitmap; is it already locked?")
	}
	return (*LockedRegion)(reg), nil
}

// Like al_lock_bitmap_blocked, but allows locking a sub-region, for
// performance. Unlike al_lock_bitmap_region the region specified in terms of
// blocks and not pixels.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_lock_bitmap_region_blocked
func (bmp *Bitmap) LockRegionBlocked(x, y, width, height int, flags LockFlags) (*LockedRegion, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	reg := C.al_lock_bitmap_region_blocked(
		(*C.ALLEGRO_BITMAP)(bmp),
		C.int(x),
		C.int(y),
		C.int(width),
		C.int(height),
		C.int(flags),
	)
	if reg == nil {
		return nil, errors.New("failed to lock bitmap; is it already locked?")
	}
	return (*LockedRegion)(reg), nil
}

// Like al_lock_bitmap, but only locks a specific area of the bitmap. If the
// bitmap is a video bitmap, only that area of the texture will be updated when
// it is unlocked. Locking only the region you indend to modify will be faster
// than locking the whole bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_lock_bitmap_region
func (bmp *Bitmap) LockRegion(x, y, width, height int, format PixelFormat, flags LockFlags) (*LockedRegion, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
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

// Returns whether or not a bitmap is already locked.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_is_bitmap_locked
func (bmp *Bitmap) IsLocked() bool {
	if bmp == nil {
		return false
	}
	return bool(C.al_is_bitmap_locked((*C.ALLEGRO_BITMAP)(bmp)))
}

// Unlock a previously locked bitmap or bitmap region. If the bitmap is a video
// bitmap, the texture will be updated to match the system memory copy (unless
// it was locked read only).
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_unlock_bitmap
func (bmp *Bitmap) Unlock() {
	if bmp == nil {
		return
	}
	C.al_unlock_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

// Creates a sub-bitmap of the parent, at the specified coordinates and of the
// specified size. A sub-bitmap is a bitmap that shares drawing memory with a
// pre-existing (parent) bitmap, but possibly with a different size and
// clipping settings.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_create_sub_bitmap
func (bmp *Bitmap) CreateSubBitmap(x, y, w, h int) (*Bitmap, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	sub := C.al_create_sub_bitmap((*C.ALLEGRO_BITMAP)(bmp),
		C.int(x), C.int(y), C.int(w), C.int(h))
	if sub == nil {
		return nil, errors.New("failed to create sub-bitmap")
	}
	return (*Bitmap)(sub), nil
}

// Returns true if the specified bitmap is a sub-bitmap, false otherwise.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_is_sub_bitmap
func (bmp *Bitmap) IsSubBitmap() bool {
	if bmp == nil {
		return false
	}
	return bool(C.al_is_sub_bitmap((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) ParentBitmap() (*Bitmap, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	par := C.al_get_parent_bitmap((*C.ALLEGRO_BITMAP)(bmp))
	if par == nil {
		return nil, errors.New("no parent bitmap")
	}
	return (*Bitmap)(par), nil
}

// Create a new bitmap with al_create_bitmap, and copy the pixel data from the
// old bitmap across. The newly created bitmap will be created with the current
// new bitmap flags, and not the ones that were used to create the original
// bitmap. If the new bitmap is a memory bitmap, its projection bitmap is reset
// to be orthographic.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_clone_bitmap
func (bmp *Bitmap) Clone() (*Bitmap, error) {
	if bmp == nil {
		return nil, BitmapIsNull
	}
	clone := C.al_clone_bitmap((*C.ALLEGRO_BITMAP)(bmp))
	if clone == nil {
		return nil, errors.New("failed to clone bitmap")
	}
	return (*Bitmap)(clone), nil
}

// D3D and OpenGL allow sharing a texture in a way so it can be used for
// multiple windows. Each ALLEGRO_BITMAP created with al_create_bitmap however
// is usually tied to a single ALLEGRO_DISPLAY. This function can be used to
// know if the bitmap is compatible with the given display, even if it is a
// different display to the one it was created with. It returns true if the
// bitmap is compatible (things like a cached texture version can be used) and
// false otherwise (blitting in the current display will be slow).
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_is_compatible_bitmap
func (bmp *Bitmap) IsCompatible() bool {
	if bmp == nil {
		return false
	}
	return bool(C.al_is_compatible_bitmap((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) BitmapFlags() BitmapFlags {
	if bmp == nil {
		return 0
	}
	return BitmapFlags(C.al_get_bitmap_flags((*C.ALLEGRO_BITMAP)(bmp)))
}

func (bmp *Bitmap) BitmapFormat() PixelFormat {
	if bmp == nil {
		return 0
	}
	return PixelFormat(C.al_get_bitmap_format((*C.ALLEGRO_BITMAP)(bmp)))
}

// Get a pixel's color value from the specified bitmap. This operation is slow
// on non-memory bitmaps. Consider locking the bitmap if you are going to use
// this function multiple times on the same bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_pixel
func (bmp *Bitmap) Pixel(x, y int) Color {
	if bmp == nil {
		return *new(Color)
	}
	return (Color)(C.al_get_pixel((*C.ALLEGRO_BITMAP)(bmp), C.int(x), C.int(y)))
}

// Convert the given mask color to an alpha channel in the bitmap. Can be used
// to convert older 4.2-style bitmaps with magic pink to alpha-ready bitmaps.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_convert_mask_to_alpha
func (bmp *Bitmap) ConvertMaskToAlpha(mask_color Color) {
	if bmp == nil {
		return
	}
	C.al_convert_mask_to_alpha((*C.ALLEGRO_BITMAP)(bmp), C.ALLEGRO_COLOR(mask_color))
}

//}}}

// image.Image conformity {{{
func (bmp *Bitmap) ColorModel() color.Model {
	// ???: always use this color model?
	return color.RGBAModel
}

func (bmp *Bitmap) Bounds() image.Rectangle {
	return image.Rect(0, 0, bmp.Width(), bmp.Height())
}

func (bmp *Bitmap) At(x, y int) color.Color {
	return bmp.Pixel(x, y)
}

func (bmp *Bitmap) Set(x, y int, c color.Color) {
	bmp.AsTarget(func() {
		PutPixel(x, y, NewColor(c))
	})
}

func (c Color) RGBA() (r, g, b, a uint32) {
	fr, fg, fb, fa := c.UnmapRGBAf()
	return uint32(fr * rgbaMAX), uint32(fg * rgbaMAX), uint32(fb * rgbaMAX), uint32(fa * rgbaMAX)
}

// NewColor converts a Go image.Color to an Allegro color value. If the parameter
// is itself an Allegro color value, then it is cast and returned; otherwise, a
// new value is constructed using MapRGBAf.
func NewColor(c color.Color) Color {
	if x, ok := c.(Color); ok {
		return x
	}
	r, g, b, a := c.RGBA()
	return MapRGBAf(float32(r)/rgbaMAX, float32(g)/rgbaMAX, float32(b)/rgbaMAX, float32(a)/rgbaMAX)
}

var _ color.Color = (*Color)(nil)
var _ image.Image = (*Bitmap)(nil)
var _ draw.Image = (*Bitmap)(nil)

// }}}

// Color Methods {{{

// Convert r, g, b (ranging from 0-255) into an ALLEGRO_COLOR, using 255 for
// alpha.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_map_rgb
func MapRGB(r, g, b byte) Color {
	return Color(C.al_map_rgb(C.uchar(r), C.uchar(g), C.uchar(b)))
}

// Convert r, g, b, a (ranging from 0-255) into an ALLEGRO_COLOR.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_map_rgba
func MapRGBA(r, g, b, a byte) Color {
	return (Color)(C.al_map_rgba(C.uchar(r), C.uchar(g), C.uchar(b), C.uchar(a)))
}

// Convert r, g, b, (ranging from 0.0f-1.0f) into an ALLEGRO_COLOR, using 1.0f
// for alpha.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_map_rgb_f
func MapRGBf(r, g, b float32) Color {
	return (Color)(C.al_map_rgb_f(C.float(r), C.float(g), C.float(b)))
}

// Convert r, g, b, a (ranging from 0.0f-1.0f) into an ALLEGRO_COLOR.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_map_rgba_f
func MapRGBAf(r, g, b, a float32) Color {
	return (Color)(C.al_map_rgba_f(C.float(r), C.float(g), C.float(b), C.float(a)))
}

// Retrieves components of an ALLEGRO_COLOR, ignoring alpha. Components will
// range from 0-255.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_unmap_rgb
func (c Color) UnmapRGB() (byte, byte, byte) {
	var r, g, b C.uchar
	C.al_unmap_rgb((C.ALLEGRO_COLOR)(c), &r, &g, &b)
	return byte(r), byte(g), byte(b)
}

// Retrieves components of an ALLEGRO_COLOR. Components will range from 0-255.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_unmap_rgba
func (c Color) UnmapRGBA() (byte, byte, byte, byte) {
	var r, g, b, a C.uchar
	C.al_unmap_rgba((C.ALLEGRO_COLOR)(c), &r, &g, &b, &a)
	return byte(r), byte(g), byte(b), byte(a)
}

// Retrieves components of an ALLEGRO_COLOR, ignoring alpha. Components will
// range from 0.0f-1.0f.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_unmap_rgb_f
func (c Color) UnmapRGBf() (float32, float32, float32) {
	var r, g, b C.float
	C.al_unmap_rgb_f((C.ALLEGRO_COLOR)(c), &r, &g, &b)
	return float32(r), float32(g), float32(b)
}

// Retrieves components of an ALLEGRO_COLOR. Components will range from
// 0.0f-1.0f.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_unmap_rgba_f
func (c Color) UnmapRGBAf() (float32, float32, float32, float32) {
	var r, g, b, a C.float
	C.al_unmap_rgba_f((C.ALLEGRO_COLOR)(c), &r, &g, &b, &a)
	return float32(r), float32(g), float32(b), float32(a)
}

//}}}

// Miscellaneous Instance Methods {{{

func (reg *LockedRegion) Data() uintptr {
	return uintptr((*C.struct_ALLEGRO_LOCKED_REGION)(reg).data)
}

func (reg *LockedRegion) Format() PixelFormat {
	return PixelFormat((*C.struct_ALLEGRO_LOCKED_REGION)(reg).format)
}

func (reg *LockedRegion) Pitch() int {
	return int((*C.struct_ALLEGRO_LOCKED_REGION)(reg).pitch)
}

func (reg *LockedRegion) PixelSize() int {
	return int((*C.struct_ALLEGRO_LOCKED_REGION)(reg).pixel_size)
}

// Return the number of bytes that a pixel of the given format occupies. For
// blocked pixel formats (e.g. compressed formats), this returns 0.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_pixel_size
func (format PixelFormat) PixelSize() int {
	return int(C.al_get_pixel_size(C.int(format)))
}

// Return the number of bits that a pixel of the given format occupies. For
// blocked pixel formats (e.g. compressed formats), this returns 0.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_pixel_format_bits
func (format PixelFormat) PixelFormatBits() int {
	return int(C.al_get_pixel_format_bits(C.int(format)))
}

// Loads an image from an ALLEGRO_FILE stream into a new ALLEGRO_BITMAP. The
// file type is determined by the passed 'ident' parameter, which is a file
// name extension including the leading dot. If (and only if) 'ident' is NULL,
// the file type is determined with al_identify_bitmap_f instead.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_load_bitmap_f
func (f *File) LoadBitmap(ident string) (*Bitmap, error) {
	ident_ := C.CString(ident)
	defer freeString(ident_)
	bmp := C.al_load_bitmap_f((*C.ALLEGRO_FILE)(f), ident_)
	if bmp == nil {
		return nil, errors.New("failed to load bitmap from file")
	}
	return (*Bitmap)(bmp), nil
}

// Loads an image from an ALLEGRO_FILE stream into a new ALLEGRO_BITMAP. The
// file type is determined by the passed 'ident' parameter, which is a file
// name extension including the leading dot. If (and only if) 'ident' is NULL,
// the file type is determined with al_identify_bitmap_f instead.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_load_bitmap_flags_f
func (f *File) LoadBitmapFlags(ident string, flags BitmapFlags) (*Bitmap, error) {
	ident_ := C.CString(ident)
	defer freeString(ident_)
	bmp := C.al_load_bitmap_flags_f((*C.ALLEGRO_FILE)(f), ident_, C.int(flags))
	if bmp == nil {
		return nil, errors.New("failed to load bitmap from file")
	}
	return (*Bitmap)(bmp), nil
}

// Saves an ALLEGRO_BITMAP to an ALLEGRO_FILE stream. The file type is
// determined by the passed 'ident' parameter, which is a file name extension
// including the leading dot.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_save_bitmap_f
func (f *File) SaveBitmap(ident string, bmp *Bitmap) error {
	ident_ := C.CString(ident)
	defer freeString(ident_)
	ok := bool(C.al_save_bitmap_f((*C.ALLEGRO_FILE)(f), ident_, (*C.ALLEGRO_BITMAP)(bmp)))
	if !ok {
		return errors.New("failed to save bitmap to file")
	}
	return nil
}

//}}}

// AsTarget() is a utility method for temporarily setting a bitmap as
// the target bitmap, running the provided function, and then undoing
// the change after the function exits. This is very useful for calling
// functions that operate on the target bitmap, e.g. the drawing methods
// provided by the primitives addon.
func (bmp *Bitmap) AsTarget(f func()) *Bitmap {
	old := TargetBitmap()
	SetTargetBitmap(bmp)
	f()
	SetTargetBitmap(old)
	return bmp
}
