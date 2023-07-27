package main

// The second-closest thing you'll probably find to this is https://github.com/paxtonhare/demo-magic/ .
//  But it's a different beast.  It only invokes shell stuff, and pretends to be saying what it's doing.
//  What we're going for with ttydrive is that you LITERALLY send "keystrokes" to the terminal.  No fakery possible.
//   Practically speaking, that also means we can keey feeding increments of content to subprocesses, even if they grab the tty themselves.

import (
	"os"
	"strconv"
	"syscall"
	"time"
	"unsafe"
)

// Usage is something like:
// `TTY=/dev/pts/2 td "echo me\n"`
// ... where the pts number is something you have to eyeball.
// You also need sudo to do this.
func main() {
	send, err := strconv.Unquote(`"` + os.Args[1] + `"`)
	if err != nil {
		panic(err)
	}
	handle := uintptr(syscall.Stdin)
	if target := os.Getenv("TTY"); target != "" {
		f, err := os.OpenFile(target, os.O_WRONLY, 0)
		if err != nil {
			panic(err)
		}
		handle = f.Fd()
	}
	for _, c := range send {
		_, _, err := syscall.Syscall(syscall.SYS_IOCTL,
			handle,
			uintptr(syscall.TIOCSTI),
			uintptr(unsafe.Pointer(&c)),
		)
		if err != 0 {
			panic(err)
		}
		time.Sleep(time.Millisecond * 20)
		if c == '\n' {
			time.Sleep(time.Millisecond * 100)
		}
	}

	// can't really figure out how to read back out.
	//  people use 'script' and 'screen' commands for this; maybe strace them.
	//   but i think they're parents of the terminal, so they're maybe just using that position to dup things.
	//  not actually sure if this is possible at all.

	// Polling for info about subprocesses existing (or ceasing to) could be a useful fallback fairly often.
}
