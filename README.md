# Namespaces
isolating hostname inside the container with the `NEWUTS` namespace.
- Created with `syscalls`
    - Unix Timesharing System
    - Process IDs
    - Mounts
    - Network
    - User IDs
    - interProcess comms

# Chroot

chroot can be used to give the container its own root filesystem.

- Limit access to subset of direcotory tree
    - `/path/to/chroot/some-oter-path` on host
    - `/some-other-path` in container

# Container Process IDs and Mounts
NEWPID and NEWNS namespaces are used to give the container a restricted view of processes and mount points. We'll also explore how this corresponds to the processes and mount points you can see on the host machine.
- `syscall.Mount("proc", "proc", "proc", 0, "")` mount the `proc` dir to container can list process IDs into container.
    - but if only want look container process requier `syscall.CLONE_NEWPID` namespaces.
    - onnly want look container mounts requier `syscall.CLONE_NEWNS` namespaces.


# other namespaces
`NEWNET`, `NEWIPC`, and `NEWUSER`. We'll briefly discuss how these are used to isolate networking, interprocess communication, and user/group IDs in a container.

# Control Groups(cgroup)
Control groups limit the resources that a container can use, such as memory, network I/O, or CPU usage.
- Filesystem interface
    - Memory
    - CPU
    - I/O
    - Process numbers
    - ...

```
#list cgroup
ls /sys/fs/cgroup
ls /sys/fs/cgroup/memory/docker
```
## Cgroup process assignment
- `/sys/fs/cgroups/memory/cgroup.procs`
    - Default

- `/sys/fs/cgroups/memory/myCgroup/cgroup.procs`
    - Write settings, write <pid> to assign <pid> to myCgroup
    - Children of <pid> assigned here

- `/sys/fs/cgroups/memory/myCgroup/childCgroup`
    - Inherits settings from parent(myCgroup)
