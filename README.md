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

