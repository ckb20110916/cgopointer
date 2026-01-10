package cgopointer

// #include <stdlib.h>
import "C"
import (
	"sync"
	"unsafe"
)

func New() *CgoPointer {
	return &CgoPointer{}
}

type CgoPointer struct {
	mapper sync.Map
}

func (p *CgoPointer) Save(val any) unsafe.Pointer {
	if val == nil {
		return nil
	}
	var ptr unsafe.Pointer = C.malloc(C.size_t(1))
	if ptr == nil {
		panic("can't allocate 'cgo-pointer hack index pointer': ptr == nil")
	}
	p.mapper.Store(ptr, val)
	return ptr
}

func (p *CgoPointer) Restore(ptr unsafe.Pointer) (v any) {
	if ptr != nil {
		vv, ok := p.mapper.Load(ptr)
		if ok {
			return vv
		}
	}
	return
}

func (p *CgoPointer) Unref(ptr unsafe.Pointer) {
	if ptr != nil {
		_, ok := p.mapper.LoadAndDelete(ptr)
		if ok {
			C.free(ptr)
		}
	}
}
