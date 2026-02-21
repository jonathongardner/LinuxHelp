# Docker images for immutable linux
This is a set of docker images for using with immutable linux.

## Yubi
For making changes to yubi keys
```sh
podman build -t yubi -f DockerFile.yubi
distrobox create --image localhost/yubi:latest --name yubi
distrobox stop ytl
distrobox rm ytl
```

## Dev
For dev, contains zed, golang and git lfs
```sh
podman build -t dev:dev -f DockerFile.dev .
distrobox create --image localhost/dev:dev --name dev
```

## ytl
For downloading video/audio using ytl downloader
```sh
podman build -t ytl -f DockerFile.ytl
```
