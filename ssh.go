package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unsafe"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
)

func sshServer() {
	log.Println("Running ssh server")
	ssh.Handle(func(s ssh.Session) {
		cmd := exec.Command(os.Args[0])
		f, err := pty.Start(cmd)
		if err != nil {
			_, _ = s.Stderr().Write([]byte("Failed to start bbs"))
			_ = s.Exit(1)
		}

		_, winCh, isPty := s.Pty()
		if !isPty {
			return
		}
		go func() {
			for win := range winCh {
				_, _, _ = syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
					uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{
						uint16(win.Height),
						uint16(win.Width),
						0, 0,
					}),
					))
			}
		}()

		go func() {
			if _, err := io.Copy(f, s); err != nil {
				_, _ = s.Stderr().Write([]byte("Failed to connect input"))
				_ = s.Exit(1)
			}
		}()
		if _, err = io.Copy(s, f); err != nil {
			return
		}
	})
	log.Fatal(ssh.ListenAndServe(":22", nil))
}
