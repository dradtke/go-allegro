// +build unstable

package allegro

// #include <allegro5/allegro.h>
import "C"

// Sets the depthbuffer depth used by newly created bitmaps (on the current
// thread) if they are used with al_set_target_bitmap. 0 means no depth-buffer
// will be created when drawing into the bitmap, which is the default.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_new_bitmap_depth
func SetNewBitmapDepth(depth int) {
	C.al_set_new_bitmap_depth(C.int(depth))
}

// Returns the value currently set with al_set_new_bitmap_depth on the current
// thread or 0 if none was set.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_new_bitmap_depth
func NewBitmapDepth() int {
	return int(C.al_get_new_bitmap_depth())
}

// Sets the multi-sampling samples used by newly created bitmaps (on the
// current thread) if they are used with al_set_target_bitmap. 0 means
// multi-sampling will not be used when drawing into the bitmap, which is the
// default. 1 means multi-sampling will be used but only using a single sample
// per pixel (so usually there will be no visual difference to not using
// multi-sampling at all).
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_new_bitmap_samples
func SetNewBitmapSamples(samples int) {
	C.al_set_new_bitmap_samples(C.int(samples))
}

// Returns the value currently set with al_set_new_bitmap_samples on the
// current thread or 0 if none was set.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_new_bitmap_samples
func NewBitmapSamples() int {
	return int(C.al_get_new_bitmap_samples())
}

// Return the depthbuffer depth used by this bitmap if it is used with
// al_set_target_bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_depth
func (bmp *Bitmap) Depth() int {
	if bmp == nil {
		return 0
	}
	return int(C.al_get_bitmap_depth((*C.ALLEGRO_BITMAP)(bmp)))
}

// Return the multi-sampling samples used by this bitmap if it is used with
// al_set_target_bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_samples
func (bmp *Bitmap) Samples() int {
	if bmp == nil {
		return 0
	}
	return int(C.al_get_bitmap_samples((*C.ALLEGRO_BITMAP)(bmp)))
}

// On some platforms, notably Windows Direct3D and Android, textures may be
// lost at any time for events such as display resize or switching out of the
// app. On those platforms, bitmaps created without the
// ALLEGRO_NO_PRESERVE_TEXTURE flag automatically get backed up to system
// memory every time al_flip_display is called.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_backup_dirty_bitmap
func (bmp *Bitmap) BackupDirty() {
	C.al_backup_dirty_bitmap((*C.ALLEGRO_BITMAP)(bmp))
}

// Sets the color to use for ALLEGRO_CONST_COLOR or ALLEGRO_INVERSE_CONST_COLOR
// blend operations.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_set_bitmap_blend_color
func SetBitmapBlendColor(c Color) {
	C.al_set_bitmap_blend_color(C.ALLEGRO_COLOR(c))
}

// Returns the color currently used for constant color blending on the target
// bitmap.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_blend_color
func BitmapBlendColor() Color {
	return Color(C.al_get_bitmap_blend_color())
}

// Returns the current blender being used by the target bitmap. You can pass
// NULL for values you are not interested in.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_bitmap_blender
func BitmapBlender() (op BlendingOperation, src, dst BlendingValue) {
	var op_, src_, dst_ C.int
	C.al_get_bitmap_blender(&op_, &src_, &dst_)
	return BlendingOperation(op_), BlendingValue(src_), BlendingValue(dst_)
}

// Returns the current blender being used by the target bitmap. You can pass
// NULL for values you are not interested in.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_get_separate_bitmap_blender
func SeparateBitmapBlender() (op BlendingOperation, src, dst BlendingValue, alphaOp BlendingOperation, alphaSrc, alphaDst BlendingValue) {
	var op_, src_, dst_, alphaOp_, alphaSrc_, alphaDst_ C.int
	C.al_get_separate_bitmap_blender(&op_, &src_, &dst_, &alphaOp_, &alphaSrc_, &alphaDst_)
	return BlendingOperation(op_), BlendingValue(src_), BlendingValue(dst_), BlendingOperation(alphaOp_), BlendingValue(alphaSrc_), BlendingValue(alphaDst_)
}

// Resets the blender for this bitmap to the default. After resetting the
// bitmap blender, the values set for
// al_set_bitmap_blender/al_set_separate_bitmap_blender will be used instead.
//
// See https://liballeg.org/a5docs/5.2.6/graphics.html#al_reset_bitmap_blender
func ResetBitmapBlender() {
	C.al_reset_bitmap_blender()
}
