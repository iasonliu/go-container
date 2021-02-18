package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// docker run <image> <cmd> <args>
// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// "go.toolsEnvVars": {"GOOS" : "linux"} only on Linux
	// create Namespace in Go
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// setting hostname in namespace
	must(syscall.Sethostname([]byte("container")))
	// chroot
	must(syscall.Chroot("/home/vagrant/ubuntufs"))
	must(os.Chdir("/"))

	// mounts proc to using ps list container process IDs
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	// mounts tmpfs
	must(syscall.Mount("something", "mytemp", "tmpfs", 0, ""))

	must(cmd.Run())
	// unmount
	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("mytemp", 0))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
