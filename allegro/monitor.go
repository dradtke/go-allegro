package allegro

// #include <allegro5/allegro.h>
import "C"
import (
	"fmt"
)

type MonitorInfo C.struct_ALLEGRO_MONITOR_INFO

// Get information about a monitor's position on the desktop. adapter is a number from
// 0 to al_get_num_video_adapters()-1.
func GetMonitorInfo(adapter int) (*MonitorInfo, error) {
	var info C.struct_ALLEGRO_MONITOR_INFO
	success := bool(C.al_get_monitor_info(C.int(adapter), &info))
	if !success {
		return nil, fmt.Errorf("error getting monitor info for adapter '%d'", adapter)
	}
	return (*MonitorInfo)(&info), nil
}

func (m *MonitorInfo) X1() int {
	return int(m.x1)
}

func (m *MonitorInfo) X2() int {
	return int(m.x2)
}

func (m *MonitorInfo) Y1() int {
	return int(m.y1)
}

func (m *MonitorInfo) Y2() int {
	return int(m.y2)
}

func (m *MonitorInfo) Width() int {
	return m.X2() - m.X1()
}

func (m *MonitorInfo) Height() int {
	return m.Y2() - m.Y1()
}

// Convenience method for testing whether or not this is the primary monitor.
// Returns true iff x1 and y1 are both 0.
func (m *MonitorInfo) IsPrimary() bool {
	return m.X1() == 0 && m.Y1() == 0
}

// Gets the video adapter index where new displays will be created by the
// calling thread, if previously set with al_set_new_display_adapter. Otherwise
// returns ALLEGRO_DEFAULT_DISPLAY_ADAPTER.
func NewDisplayAdapter() int {
	return int(C.al_get_new_display_adapter())
}

// Sets the adapter to use for new displays created by the calling thread. The
// adapter has a monitor attached to it. Information about the monitor can be
// gotten using al_get_num_video_adapters and al_get_monitor_info.
func SetNewDisplayAdapter(adapter int) {
	C.al_set_new_display_adapter(C.int(adapter))
}


// Get the number of video "adapters" attached to the computer. Each video card
// attached to the computer counts as one or more adapters. An adapter is thus
// really a video port that can have a monitor connected to it.
func NumVideoAdapters() int {
	return int(C.al_get_num_video_adapters())
}
