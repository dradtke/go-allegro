go-allegro
==========

This repository contains bindings for writing [Allegro 5](http://alleg.sourceforge.net) games in Go.

I assume that you have a working Allegro 5 development environment set up. If not, go do that first.

Screenshot of one of the included examples (open the `example` folder for more details on the included examples):

![screenshot](https://github.com/dradtke/go-allegro/raw/5.0/example/img/screenshot.png)

The following functions still need to be implemented or blacklisted (in addition to a few more modules):

```
--- FAIL: TestCoverage (0.14 seconds)
	coverage_test.go:347: Missing allegro function 'al_draw_tinted_scaled_rotated_bitmap_region' in file '/usr/include/allegro5/bitmap.h'
	coverage_test.go:347: Missing allegro function 'al_set_separate_blender' in file '/usr/include/allegro5/bitmap.h'
	coverage_test.go:347: Missing allegro function 'al_get_separate_blender' in file '/usr/include/allegro5/bitmap.h'
	coverage_test.go:347: Missing allegro function 'al_register_assert_handler' in file '/usr/include/allegro5/debug.h'
	coverage_test.go:347: Missing allegro function 'al_emit_user_event' in file '/usr/include/allegro5/events.h'
	coverage_test.go:347: Missing allegro function 'al_wait_for_event_timed' in file '/usr/include/allegro5/events.h'
	coverage_test.go:347: Missing allegro function 'al_make_temp_file' in file '/usr/include/allegro5/file.h'
	coverage_test.go:347: Missing allegro function 'al_set_new_file_interface' in file '/usr/include/allegro5/file.h'
	coverage_test.go:347: Missing allegro function 'al_malloc_with_context' in file '/usr/include/allegro5/memory.h'
	coverage_test.go:347: Missing allegro function 'al_free_with_context' in file '/usr/include/allegro5/memory.h'
	coverage_test.go:347: Missing allegro function 'al_realloc_with_context' in file '/usr/include/allegro5/memory.h'
	coverage_test.go:347: Missing allegro function 'al_calloc_with_context' in file '/usr/include/allegro5/memory.h'
	coverage_test.go:347: Missing allegro function 'al_wait_cond_until' in file '/usr/include/allegro5/threads.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_dup_substr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ref_buffer' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ref_ustr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_insert' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_insert_cstr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_vappendf' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_remove_range' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_assign_substr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_replace_range' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_chr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_rfind_chr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_set' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_set_cstr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_cset' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_cset_cstr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_str' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_cstr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_rfind_str' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_rfind_cstr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_replace' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_find_replace_cstr' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:347: Missing allegro function 'al_ustr_ncompare' in file '/usr/include/allegro5/utf8.h'
	coverage_test.go:349: Module 'color' missing function 'al_color_hsv_to_rgb' [void al_color_hsv_to_rgb(float hue, float saturation,float value, float *red, float *green, float *blue)]
	coverage_test.go:349: Module 'color' missing function 'al_color_rgb_to_hsl' [void al_color_rgb_to_hsl(float red, float green, float blue,float *hue, float *saturation, float *lightness)]
	coverage_test.go:349: Module 'color' missing function 'al_color_rgb_to_hsv' [void al_color_rgb_to_hsv(float red, float green, float blue,float *hue, float *saturation, float *value)]
	coverage_test.go:349: Module 'color' missing function 'al_color_name_to_rgb' [bool al_color_name_to_rgb(char const *name, float *r, float *g,float *b)]
	coverage_test.go:349: Module 'color' missing function 'al_color_rgb_to_name' [const char* al_color_rgb_to_name(float r, float g, float b)]
	coverage_test.go:349: Module 'color' missing function 'al_color_rgb_to_cmyk' [void al_color_rgb_to_cmyk(float red, float green, float blue,float *cyan, float *magenta, float *yellow, float *key)]
	coverage_test.go:349: Module 'color' missing function 'al_color_yuv_to_rgb' [void al_color_yuv_to_rgb(float y, float u, float v,float *red, float *green, float *blue)]
	coverage_test.go:349: Module 'color' missing function 'al_color_rgb_to_yuv' [void al_color_rgb_to_yuv(float red, float green, float blue,float *y, float *u, float *v)]
	coverage_test.go:349: Module 'color' missing function 'al_color_rgb_to_html' [void al_color_rgb_to_html(float red, float green, float blue,char *string)]
	coverage_test.go:349: Module 'color' missing function 'al_color_html_to_rgb' [void al_color_html_to_rgb(char const *string,float *red, float *green, float *blue)]
	coverage_test.go:349: Module 'color' missing function 'al_color_yuv' [ALLEGRO_COLOR al_color_yuv(float y, float u, float v)]
	coverage_test.go:349: Module 'color' missing function 'al_color_hsv' [ALLEGRO_COLOR al_color_hsv(float h, float s, float v)]
	coverage_test.go:349: Module 'color' missing function 'al_color_name' [ALLEGRO_COLOR al_color_name(char const *name)]
	coverage_test.go:349: Module 'color' missing function 'al_color_html' [ALLEGRO_COLOR al_color_html(char const *string)]
FAIL
exit status 1
FAIL	github.com/dradtke/go-allegro	0.153s
```
