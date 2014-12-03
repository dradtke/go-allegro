package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Config C.ALLEGRO_CONFIG
type ConfigSectionIterator (*C.ALLEGRO_CONFIG_SECTION)
type ConfigEntryIterator (*C.ALLEGRO_CONFIG_ENTRY)

// Create an empty configuration structure.
func CreateConfig() *Config {
	config := (*Config)(C.al_create_config())
	//runtime.SetFinalizer(config, config.Destroy)
	return config
}

// Read a configuration file from disk. Returns NULL on error. The
// configuration structure should be destroyed with al_destroy_config.
func LoadConfig(filename string) (*Config, error) {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	cfg := C.al_load_config_file(filename_)
	if cfg == nil {
		return nil, fmt.Errorf("failed to load config file '%s'", filename)
	}
	return (*Config)(cfg), nil
}

// Merge two configuration structures, and return the result as a new
// configuration. Values in configuration 'cfg2' override those in 'cfg1'.
// Neither of the input configuration structures are modified. Comments from
// 'cfg2' are not retained.
func MergeConfig(cfg1, cfg2 *Config) *Config {
	return (*Config)(C.al_merge_config((*C.ALLEGRO_CONFIG)(cfg1), (*C.ALLEGRO_CONFIG)(cfg2)))
}

// Read a configuration file from an already open file.
func (f *File) LoadConfig() (*Config, error) {
	cfg := C.al_load_config_file_f((*C.ALLEGRO_FILE)(f))
	if cfg == nil {
		return nil, errors.New("failed to load config from file")
	}
	return (*Config)(cfg), nil
}

// Write out a configuration file to an already open file.
func (f *File) SaveConfig(cfg *Config) error {
	ok := bool(C.al_save_config_file_f((*C.ALLEGRO_FILE)(f), (*C.ALLEGRO_CONFIG)(cfg)))
	if !ok {
		return errors.New("failed to save config from file")
	}
	return nil
}

// Config Instance Methods {{{

// Add a section to a configuration structure with the given name. If the
// section already exists then nothing happens.
func (cfg *Config) AddSection(name string) {
	name_ := C.CString(name)
	defer freeString(name_)
	C.al_add_config_section((*C.ALLEGRO_CONFIG)(cfg), name_)
}

// Set a value in a section of a configuration. If the section doesn't yet
// exist, it will be created. If a value already existed for the given key, it
// will be overwritten. The section can be NULL or "" for the global section.
func (cfg *Config) SetValue(section, key, value string) {
	section_ := C.CString(section)
	key_ := C.CString(key)
	value_ := C.CString(value)
	defer freeString(section_)
	defer freeString(key_)
	defer freeString(value_)
	C.al_set_config_value((*C.ALLEGRO_CONFIG)(cfg), section_, key_, value_)
}

// Gets a pointer to an internal character buffer that will only remain valid
// as long as the ALLEGRO_CONFIG structure is not destroyed. Copy the value if
// you need a copy. The section can be NULL or "" for the global section.
// Returns NULL if the section or key do not exist.
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

// Add a comment in a section of a configuration. If the section doesn't yet
// exist, it will be created. The section can be NULL or "" for the global
// section.
func (cfg *Config) AddComment(section, comment string) {
	section_ := C.CString(section)
	comment_ := C.CString(comment)
	defer freeString(section_)
	defer freeString(comment_)
	C.al_add_config_comment((*C.ALLEGRO_CONFIG)(cfg), section_, comment_)
}

// Write out a configuration file to disk. Returns true on success, false on
// error.
func (cfg *Config) Save(filename string) error {
	filename_ := C.CString(filename)
	defer freeString(filename_)
	ok := bool(C.al_save_config_file(filename_, (*C.ALLEGRO_CONFIG)(cfg)))
	if !ok {
		return fmt.Errorf("failed to save config file to '%s'", filename)
	}
	return nil
}

// Merge one configuration structure into another. Values in configuration
// 'add' override those in 'master'. 'master' is modified. Comments from 'add'
// are not retained.
func (cfg *Config) Merge(add *Config) {
	C.al_merge_config_into((*C.ALLEGRO_CONFIG)(cfg), (*C.ALLEGRO_CONFIG)(add))
}

// Free the resources used by a configuration structure. Does nothing if passed
// NULL.
func (cfg *Config) Destroy() {
	C.al_destroy_config((*C.ALLEGRO_CONFIG)(cfg))
}

// Returns the name of the first section in the given config file. Usually this
// will return an empty string for the global section. The iterator parameter
// will receive an opaque iterator which is used by al_get_next_config_section
// to iterate over the remaining sections.
func (cfg *Config) FirstConfigSection() (string, *ConfigSectionIterator) {
	var iter ConfigSectionIterator
	section := C.al_get_first_config_section((*C.ALLEGRO_CONFIG)(cfg),
		(**C.ALLEGRO_CONFIG_SECTION)(&iter))
	return C.GoString(section), &iter
}

// Returns the name of the next section in the given config file or NULL if
// there are no more sections. The iterator must have been obtained with
// al_get_first_config_section first.
func (cfg *Config) NextConfigSection(iter *ConfigSectionIterator) (string, error) {
	section := C.al_get_next_config_section((**C.ALLEGRO_CONFIG_SECTION)(iter))
	if section == nil {
		return "", errors.New("no more sections in this config")
	}
	return C.GoString(section), nil
}

// Sections() returns a read-only channel of sections in the config file.
// This makes it easy to iterate over them like so:
//
//    for section := range cfg.Sections() {
//        // do stuff
//    }
func (cfg *Config) Sections() <-chan string {
	sections := make(chan string)
	go func() {
		defer close(sections)
		var (
			s string
			iter *ConfigSectionIterator
			err error
		)
		s, iter = cfg.FirstConfigSection()
		sections <- s
		for {
			s, err = cfg.NextConfigSection(iter)
			if err != nil {
				break
			}
			sections <- s
		}
	}()
	return sections
}

// Returns the name of the first key in the given section in the given config
// or NULL if the section is empty. The iterator works like the one for
// al_get_first_config_section.
func (cfg *Config) FirstConfigEntry(section string) (string, *ConfigEntryIterator, error) {
	section_ := C.CString(section)
	defer freeString(section_)
	var iter ConfigEntryIterator
	entry := C.al_get_first_config_entry((*C.ALLEGRO_CONFIG)(cfg), section_,
		(**C.ALLEGRO_CONFIG_ENTRY)(&iter))
	if entry == nil {
		return "", nil, fmt.Errorf("section '%s' has no entries", section)
	}
	return C.GoString(entry), &iter, nil
}

// Returns the next key for the iterator obtained by al_get_first_config_entry.
// The iterator works like the one for al_get_next_config_section.
func (cfg *Config) NextConfigEntry(iter *ConfigEntryIterator) (string, error) {
	entry := C.al_get_next_config_entry((**C.ALLEGRO_CONFIG_ENTRY)(iter))
	if entry == nil {
		return "", fmt.Errorf("no more entries in this section")
	}
	return C.GoString(entry), nil
}

// Entries() returns a read-only channel of entries in the config file.
// This makes it easy to iterate over them like so:
//
//    for entry := range cfg.Entries("Section Title") {
//        // do stuff
//    }
func (cfg *Config) Entries(section string) <-chan string {
	entries := make(chan string)
	go func() {
		defer close(entries)
		var (
			e string
			iter *ConfigEntryIterator
			err error
		)
		e, iter, err = cfg.FirstConfigEntry(section)
		if err == nil {
			return
		}
		entries <- e
		for {
			e, err = cfg.NextConfigEntry(iter)
			if err != nil {
				break
			}
			entries <- e
		}
	}()
	return entries
}

func (cfg *Config) Float32Value(section, key string) (float32, error) {
	str, err := cfg.Value(section, key)
	if err != nil {
		return 0, err
	}
	str = strings.TrimSpace(stripComment(str))
	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}

func (cfg *Config) Float64Value(section, key string) (float64, error) {
	str, err := cfg.Value(section, key)
	if err != nil {
		return 0, err
	}
	str = strings.TrimSpace(stripComment(str))
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return float64(val), nil
}

func (cfg *Config) IntValue(section, key string) (int64, error) {
	str, err := cfg.Value(section, key)
	if err != nil {
		return 0, err
	}
	str = strings.TrimSpace(stripComment(str))
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (cfg *Config) UintValue(section, key string) (uint64, error) {
	str, err := cfg.Value(section, key)
	if err != nil {
		return 0, err
	}
	str = strings.TrimSpace(stripComment(str))
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func stripComment(str string) string {
	index := strings.Index(str, "#")
	if index > -1 {
		return str[:index]
	}
	return str
}

//}}}
