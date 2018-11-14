package utils

import (
	"log"
)

// LogFatalError it's a snippet for
//     if err != nil {
//         log.Fatalf("Error: %s", err)
//     }
func LogFatalError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

// RequiredStrFatal it's a snippet for
//     if str == "" {
//         log.Fatalf("Error: %s is empty", name)
//     }
func RequiredStrFatal(name, str string) {
	if str == "" {
		log.Fatalf("Error: %s is empty", name)
	}
}
