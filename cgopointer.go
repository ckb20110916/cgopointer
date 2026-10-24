package cgopointer

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"
)

type CgoPointer struct {
	mapper sync.Map
}

func (cgo *CgoPointer) Save(val any) unsafe.Pointer {
	if val == nil {
		return nil
	}
	var ptr unsafe.Pointer = C.malloc(C.size_t(1))
	if ptr == nil {
		panic("can't allocate 'cgo-pointer hack index pointer': ptr == nil")
	}
	cgo.mapper.Store(ptr, val)
	return ptr
}

func (cgo *CgoPointer) Restore(ptr unsafe.Pointer) (v any) {
	if ptr != nil {
		vv, ok := cgo.mapper.Load(ptr)
		if ok {
			return vv
		}
	}
	return
}

func (cgo *CgoPointer) Unref(ptr unsafe.Pointer) {
	if ptr != nil {
		_, ok := cgo.mapper.LoadAndDelete(ptr)
		if ok {
			C.free(ptr)
		}
	}
}
