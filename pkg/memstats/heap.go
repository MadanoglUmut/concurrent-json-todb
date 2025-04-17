package memstats

import (
	"log"
	"runtime"
)

func HeapStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %.2f MiB, TotalAlloc = %.2f MiB, Sys = %.2f MiB, NumGC = %d\n",
		float64(m.Alloc)/1024/1024,
		float64(m.TotalAlloc)/1024/1024,
		float64(m.Sys)/1024/1024,
		m.NumGC,
	)
}
