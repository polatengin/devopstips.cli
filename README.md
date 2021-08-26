# devopstips.cli

CLI for DevOps Tips and Tricks Web Site

```bash
docker run --rm -it -v $PWD:/src --workdir "/src" golang:1.17beta1-buster

docker build -t devopstips .

docker run --rm -it devopstips

alias devopstips="docker run --rm -it devopstips"
```
