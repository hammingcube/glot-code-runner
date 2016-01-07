# A template to use glot-code-runner

```sh
docker pull rsmmr/clang
docker run --rm -it -v $PWD:/go/src/glot -w /go/src/glot  golang go build runner.go
cat test_with_stdin.txt | docker run --rm -i -v "$(pwd)":/app -w /app rsmmr/clang ./runner
```
