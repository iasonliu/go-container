package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		Credential: &syscall.Credential{Uid: 0, Gid: 0},
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
	}
	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v\n", os.Args[2:])
	//  cgroup
	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// setting hostname in namespace
	must(syscall.Sethostname([]byte("container")))
	// chroot
	must(syscall.Chroot("/home/user1/ubuntufs"))
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

// setting up cgroup
func cg() {
	cgroups := "/sys/fs/cgroup/"

	mem := filepath.Join(cgroups, "memory")
	os.Mkdir(filepath.Join(mem, "user1"), 0755)
	must(ioutil.WriteFile(filepath.Join(mem, "user1/memory.limit_in_bytes"), []byte("999424"), 0700))
	must(ioutil.WriteFile(filepath.Join(mem, "user1/memory.memsw.limit_in_bytes"), []byte("999424"), 0700))
	must(ioutil.WriteFile(filepath.Join(mem, "user1/memory.ontify_on_release"), []byte("1"), 0700))

	pid := strconv.Itoa(os.Getpid())
	must(ioutil.WriteFile(filepath.Join(mem, "user1/cgroup.procs"), []byte(pid), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
