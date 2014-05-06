#include <allegro5/allegro.h>

static void free_string(char *str) {
	al_free(str);
}

static void *_al_malloc(unsigned int size) {
    return al_malloc(size);
}

