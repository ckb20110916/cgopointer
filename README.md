# cgopointer

A Go library that provides a safe way to pass Go values through CGo pointers.

## Problem

When working with CGo, Go's pointer passing rules do not allow passing Go pointers directly to C code. This makes it difficult to pass arbitrary Go values across the CGo boundary.

## Solution

`cgopointer` provides a workaround by storing Go values in a `sync.Map` and passing a C-allocated pointer as a handle. The original Go value can later be retrieved, updated, and released using that handle.

## Installation

```bash
go get github.com/ckb20110916/cgopointer
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/ckb20110916/cgopointer"
)

func main() {
	cp := &cgopointer.CgoPointer{}

	// Store a Go value and get a C-compatible pointer
	ptr := cp.Save("hello, cgo!")

	// Restore the value from the pointer
	val := cp.Restore(ptr)
	fmt.Println(val) // hello, cgo!

	// Release the pointer and free memory
	cp.Unref(ptr)
}
```

## API

### `CgoPointer`

The main type that manages the mapping between C-allocated pointers and Go values.

### `func (cgo *CgoPointer) Save(val any) unsafe.Pointer`

Stores a Go value and returns a C-allocated `unsafe.Pointer` that can be safely passed to C code. The pointer serves as a key to retrieve the original value later.

- If `val` is `nil`, returns `nil`.
- Panics if memory allocation fails.

### `func (cgo *CgoPointer) Restore(ptr unsafe.Pointer) any`

Retrieves the Go value associated with the given pointer.

- Returns the stored value, or `nil` if the pointer is `nil` or not found.

### `func (cgo *CgoPointer) Unref(ptr unsafe.Pointer)`

Removes the value associated with the pointer from the internal map and frees the allocated C memory.

- Safe to call with a `nil` pointer (no-op).

## License

See [LICENSE](LICENSE).
