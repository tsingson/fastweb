// +build !binary_log

package diode

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func ExampleNewWriter() {
	w := NewWriter(os.Stdout, 1000, 0, func(missed int) {
		fmt.Printf("Dropped %d messages\n", missed)
	})
	log := zerolog.New(w)
	log.Print("test")

	w.Close()

	// Output: {"level":"debug","message":"test"}
}
