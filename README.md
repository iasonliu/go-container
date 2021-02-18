# Namespaces
- Created with `syscalls`
    - Unix Timesharing System
    - Process IDs
    - Mounts
    - Network
    - User IDs
    - interProcess comms
isolating hostname inside the container with the `NEWUTS` namespace.

# Chroot

chroot can be used to give the container its own root filesystem.

- Limit access to subset of direcotory tree
    - `/path/to/chroot/some-oter-path` on host
    - `/some-other-path` in container

# Container Process IDs and Mounts
NEWPID and NEWNS namespaces are used to give the container a restricted view of processes and mount points. We'll also explore how this corresponds to the processes and mount points you can see on the host machine.