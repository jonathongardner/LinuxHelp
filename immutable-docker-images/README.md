# Docker images for immutable linux
This is a set of docker images for using with immutable linux.

## GoogleChrome
```sh
podman stop google-chrome
toolbox rm google-chrome
podman image rm localhost/google-chrome:latest
podman build -t google-chrome -f DockerFile.google-chrome
toolbox create --image localhost/google-chrome:latest google-chrome
```

## Yubi
For making changes to yubi keys
```sh
podman build -t yubi -f DockerFile.yubi
toolbox create --image localhost/yubi:latest yubi
podman stop ytl
toolbox rm ytl
```

## Dev
For dev, contains zed, golang and git lfs
```sh
podman build -t dev:dev -f DockerFile.dev .
toolbox create --image localhost/dev:dev dev
```

## AI Dev
For ai dev, contains antigravity, golang and git lfs
```sh
podman build -t dev:ai -f DockerFile.ai .
toolbox create --image localhost/dev:ai dev-ai
```

## ytl
For downloading video/audio using ytl downloader
```sh
podman build -t ytl -f DockerFile.ytl
```