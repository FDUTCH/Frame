package Frame

import (
	"io"

	"github.com/FDUTCH/Frame/storage"
)

type System interface {
	io.Closer
	Init(f *Frame)
}

// AddSystem adds system to the general storage and calls CloseAfterServer & Init.
func AddSystem[T System](f *Frame, system T) {
	storage.Set(f.GeneralStorage(), system)
	system.Init(f)
	f.CloseAfterServer(system)
}

// GetSystem returns added System or nil.
func GetSystem[T System](f *Frame) T {
	var empty T
	return storage.GetOr(f.GeneralStorage(), empty)
}
