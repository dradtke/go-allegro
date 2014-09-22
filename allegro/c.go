package allegro

// #cgo !windows pkg-config: allegro-5.0 allegro_main-5.0
import "C"

// Windows users should set the ALLEGRO_HOME environment variable
// to the root of their Allegro 5 installation, then run `install.bat`
// to install. *Nix users should make sure they have the appropriate
// pkg-config .pc files available, then use the standard `go install`.
