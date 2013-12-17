go-allegro
==========

This repository contains bindings for writing [Allegro 5](http://alleg.sourceforge.net) games in Go.

I assume that you have a working Allegro 5 development environment set up. If not, go do that first.

Screenshot of one of the included examples (open the `example` folder for more details on the included examples):

![screenshot](https://github.com/dradtke/go-allegro/raw/5.0/example/img/screenshot.png)

Function documentation is included in the source, but it's pulled directly from Allegro's C API documentation, so not everything will line up as far as parameters and return values. However, the C API maps pretty well to the Go API, so if you're familiar with the patterns (e.g. `error`'s instead of boolean success values, multiple return values instead of output parameters, object functions as instance methods on structs), then it shouldn't be hard to figure out what's going on.

The following functions still need to be implemented or blacklisted (in addition to a few more modules). This list is generated using the included `coverage_test.go`:

```
--- FAIL: TestCoverage (0.25 seconds)
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_prim' [int al_draw_prim(const void* vtxs, const ALLEGRO_VERTEX_DECL* decl, ALLEGRO_BITMAP* texture, int start, int end, int type)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_indexed_prim' [int al_draw_indexed_prim(const void* vtxs, const ALLEGRO_VERTEX_DECL* decl, ALLEGRO_BITMAP* texture, const int* indices, int num_vtx, int type)]
	coverage_test.go:239: Module 'primitives' missing function 'al_create_vertex_decl' [ALLEGRO_VERTEX_DECL* al_create_vertex_decl(const ALLEGRO_VERTEX_ELEMENT* elements, int stride)]
	coverage_test.go:239: Module 'primitives' missing function 'al_destroy_vertex_decl' [void al_destroy_vertex_decl(ALLEGRO_VERTEX_DECL* decl)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_soft_triangle' [void al_draw_soft_triangle(ALLEGRO_VERTEX* v1, ALLEGRO_VERTEX* v2, ALLEGRO_VERTEX* v3, uintptr_t state,void (*init)(uintptr_t, ALLEGRO_VERTEX*, ALLEGRO_VERTEX*, ALLEGRO_VERTEX*),void (*first)(uintptr_t, int, int, int, int),void (*step)(uintptr_t, int),void (*draw)(uintptr_t, int, int, int))]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_soft_line' [void al_draw_soft_line(ALLEGRO_VERTEX* v1, ALLEGRO_VERTEX* v2, uintptr_t state,void (*first)(uintptr_t, int, int, ALLEGRO_VERTEX*, ALLEGRO_VERTEX*),void (*step)(uintptr_t, int),void (*draw)(uintptr_t, int, int))]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_circle' [void al_draw_circle(float cx, float cy, float r, ALLEGRO_COLOR color, float thickness)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_arc' [void al_draw_arc(float cx, float cy, float r, float start_theta, float delta_theta, ALLEGRO_COLOR color, float thickness)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_elliptical_arc' [void al_draw_elliptical_arc(float cx, float cy, float rx, float ry, float start_theta, float delta_theta, ALLEGRO_COLOR color, float thickness)]
	coverage_test.go:239: Module 'primitives' missing function 'al_calculate_spline' [void al_calculate_spline(float* dest, int stride, float points[8], float thickness, int num_segments)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_spline' [void al_draw_spline(float points[8], ALLEGRO_COLOR color, float thickness)]
	coverage_test.go:239: Module 'primitives' missing function 'al_calculate_ribbon' [void al_calculate_ribbon(float* dest, int dest_stride, const float *points, int points_stride, float thickness, int num_segments)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_ribbon' [void al_draw_ribbon(const float *points, int points_stride, ALLEGRO_COLOR color, float thickness, int num_segments)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_filled_ellipse' [void al_draw_filled_ellipse(float cx, float cy, float rx, float ry, ALLEGRO_COLOR color)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_filled_circle' [void al_draw_filled_circle(float cx, float cy, float r, ALLEGRO_COLOR color)]
	coverage_test.go:239: Module 'primitives' missing function 'al_draw_filled_pieslice' [void al_draw_filled_pieslice(float cx, float cy, float r, float start_theta, float delta_theta, ALLEGRO_COLOR color)]
FAIL
exit status 1
FAIL	github.com/dradtke/go-allegro	0.299s
```
