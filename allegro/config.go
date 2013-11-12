package allegro

/*
#cgo pkg-config: allegro-5.0
#include <allegro5/allegro.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
)

type Config C.ALLEGRO_CONFIG
type ConfigSectionIterator (*C.ALLEGRO_CONFIG_SECTION)
type ConfigEntryIterator (*C.ALLEGRO_CONFIG_ENTRY)

func CreateConfig() *Config {
	config := (*Config)(C.al_create_config())
	runtime.SetFinalizer(config, config.Destroy)
	return config
}

func LoadConfig(filename string) (*Config, error) {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	cfg := C.al_load_config_file(filename_)
	if cfg == nil {
		return nil, fmt.Errorf("failed to load config file '%s'", filename)
	}
	return (*Config)(cfg), nil
}

func MergeConfig(cfg1, cfg2 *Config) *Config {
	return (*Config)(C.al_merge_config((*C.ALLEGRO_CONFIG)(cfg1), (*C.ALLEGRO_CONFIG)(cfg2)))
}

// Config Instance Methods {{{

func (cfg *Config) AddSection(name string) {
	name_ := C.CString(name)
	defer freeString(name_)
	C.al_add_config_section((*C.ALLEGRO_CONFIG)(cfg), name_)
}

func (cfg *Config) SetValue(section, key, value string) {
	section_ := C.CString(section)
	key_ := C.CString(key)
	value_ := C.CString(value)
	defer freeString(section_)
	defer freeString(key_)
	defer freeString(value_)
	C.al_set_config_value((*C.ALLEGRO_CONFIG)(cfg), section_, key_, value_)
}

func (cfg *Config) Value(section, key string) (string, error) {
	section_ := C.CString(section)
	key_ := C.CString(key)
	defer freeString(section_)
	defer freeString(key_)
	cvalue := C.al_get_config_value((*C.ALLEGRO_CONFIG)(cfg), section_, key_)
	if cvalue == nil {
		return "", fmt.Errorf("config value '%s.%s' not found", section, key)
	}
	return C.GoString(cvalue), nil
}

func (cfg *Config) AddComment(section, comment string) {
	section_ := C.CString(section)
	comment_ := C.CString(comment)
	defer freeString(section_)
	defer freeString(comment_)
	C.al_add_config_comment((*C.ALLEGRO_CONFIG)(cfg), section_, comment_)
}

func (cfg *Config) Save(filename string) error {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	ok := bool(C.al_save_config_file(filename_, (*C.ALLEGRO_CONFIG)(cfg)))
	if !ok {
		return fmt.Errorf("failed to save config file to '%s'", filename)
	}
	return nil
}

func (cfg *Config) Merge(add *Config) {
	C.al_merge_config_into((*C.ALLEGRO_CONFIG)(cfg), (*C.ALLEGRO_CONFIG)(add))
}

func (cfg *Config) Destroy() {
	C.al_destroy_config((*C.ALLEGRO_CONFIG)(cfg))
}

func (cfg *Config) FirstConfigSection() (string, ConfigSectionIterator) {
	var iter ConfigSectionIterator
	section := C.al_get_first_config_section((*C.ALLEGRO_CONFIG)(cfg),
		(**C.ALLEGRO_CONFIG_SECTION)(&iter))
	return C.GoString(section), iter
}

func (cfg *Config) NextConfigSection(iter ConfigSectionIterator) (string, error) {
	section := C.al_get_next_config_section((**C.ALLEGRO_CONFIG_SECTION)(&iter))
	if section == nil {
		return "", errors.New("no more sections in this config")
	}
	return C.GoString(section), nil
}

func (cfg *Config) FirstConfigEntry(section string) (string, ConfigEntryIterator, error) {
	section_ := C.CString(section)
	defer freeString(section_)
	var iter ConfigEntryIterator
	entry := C.al_get_first_config_entry((*C.ALLEGRO_CONFIG)(cfg), section_,
		(**C.ALLEGRO_CONFIG_ENTRY)(&iter))
	if entry == nil {
		return "", nil, fmt.Errorf("section '%s' has no entries", section)
	}
	return C.GoString(entry), iter, nil
}

func (cfg *Config) NextConfigEntry(iter ConfigEntryIterator) (string, error) {
	entry := C.al_get_next_config_entry((**C.ALLEGRO_CONFIG_ENTRY)(&iter))
	if entry == nil {
		return "", fmt.Errorf("no more entries in this section")
	}
	return C.GoString(entry), nil
}
//}}}
