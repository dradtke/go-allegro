// Package util provides some tools for memory management. In particular,
// it helps prevent data stuffed into a uintptr (for e.g. custom events)
// from being cleared out before it's needed.
package util

import (
    "unsafe"
)

var database = make(map[uintptr]interface{})

func Retrieve(addr uintptr) interface{} {
    if val, ok := database[addr]; ok {
        return val
    }
    return nil
}

func Store(val interface{}) uintptr {
    addr := uintptr(unsafe.Pointer(&val))
    database[addr] = val
    return addr
}

func Unstore(val interface{}) {
    addr := uintptr(unsafe.Pointer(&val))
    delete(database, addr)
}
