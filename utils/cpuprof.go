package utils

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpupprof = flag.String("cpupprof", "", "CPU profile output file path")

// CPUProf runs the profiler and returns a function you need to defer
func CPUProf() func() {
	flag.Parse()
	out := *cpupprof
	if out != "" {
		file, err := os.Create(out)
		if err != nil {
			log.Fatalf("can't open CPU profile file %q", out)
		}
		pprof.StartCPUProfile(file)
		return func() {
			pprof.StopCPUProfile()
			file.Close()
		}
	}
	return func() {}
}
