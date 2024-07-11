# README

```
docker build -t interview .
docker run --rm -it -v $PWD:/src/interview interview rake test
docker run --rm -it -v $PWD:/src/interview interview rspec
```
