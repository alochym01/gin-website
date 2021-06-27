# gin-website

Content of Makefile
```bash

run:
    go run cmd/web/*


mod:
    # real tab space or get error Makefile:2: *** missing separator.  Stop.
    go mod tidy # remove unused go packages
    go mod vendor # make local copy of third party packages

run cli:

make run
    run main.go in idrac folder
make mod
    run go mod tidy cli
    run go mod vendor cli
```

## Live reload in Development

1. Install air `go get -u github.com/cosmtrek/air`
2. Set shell enviroment by add `GOPATH="/home/hadn/go/bin"` tp `~/.bash_profile`
    ```bash
    export PATH="/home/hadn/.pyenv/bin:$PATH:/opt/hadn/flutter/flutter/bin:/home/hadn/go/bin/"
    ```
3. run `air init` to create air config file should be in same folder with `main.go`
4. run `air` for auto reload
