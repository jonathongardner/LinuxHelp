# Docker images for immutable linux
This is a set of docker images for using with immutable linux.

```sh
podman build -t yubi -f DockerFile.yubi
toolbox create --image localhost/yubi:latest yubi
```