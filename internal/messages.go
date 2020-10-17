package internal

import (
	"fmt"
	"os"

	"github.com/jamillosantos/omg/config"
)

func Fatalf(code int, format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, args...))
	os.Exit(code)
}

func Fatal(code int, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprint(args...))
	os.Exit(code)
}

func Errorf(format string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
}

func Verbosef(format string, args ...interface{}) {
	if !config.Verbose {
		return
	}
	fmt.Fprint(os.Stderr, "VERBOSE: ")
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, args...))
}

func Verbose(args ...interface{}) {
	if !config.Verbose {
		return
	}
	fmt.Fprint(os.Stderr, "VERBOSE: ")
	fmt.Fprintln(os.Stderr, args...)
}
