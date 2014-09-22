#include <allegro5/allegro.h>

extern void go_main();

static int c_main(int argc, char **argv) {
    go_main();
    return 0;
}

static void run_main(void) {
    al_run_main(0, NULL, &c_main);
}
