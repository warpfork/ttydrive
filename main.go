package main

// The second-closest thing you'll probably find to this is https://github.com/paxtonhare/demo-magic/ .
//  But it's a different beast.  It only invokes shell stuff, and pretends to be saying what it's doing.
//  What we're going for with ttydrive is that you LITERALLY send "keystrokes" to the terminal.  No fakery possible.
//   Practically speaking, that also means we can keey feeding increments of content to subprocesses, even if they grab the tty themselves.


import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	w := window{}
	syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&w)),
	)
	fmt.Printf("%v\n", w)

	d := "send this\n"
	for _, c := range d {
		syscall.Syscall(syscall.SYS_IOCTL,
			uintptr(syscall.Stdin),
			uintptr(syscall.TIOCSTI),
			uintptr(unsafe.Pointer(&c)),
		)
	}

	// can't really figure out how to read back out.
	//  people use 'script' and 'screen' commands for this; maybe strace them.
	//   but i think they're parents of the terminal, so they're maybe just using that position to dup things.
	//  not actually sure if this is possible at all.
	
	// Polling for info about subprocesses existing (or ceasing to) could be a useful fallback fairly often.
}

type window struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
