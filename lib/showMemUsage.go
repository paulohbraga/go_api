package lib

import (
	"fmt"
	"runtime"
)

func PrintMemUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", BToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", BToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", BToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
	return m.Alloc
}

func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
