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

make run - to run application

make mod - to run go module
```

## Live reload in Development

1. Install air `go get -u github.com/cosmtrek/air`
2. Set shell enviroment by add `GOPATH="/home/hadn/go/bin"` to `~/.bash_profile`
    ```bash
    export PATH="$PATH:/home/hadn/go/bin/"
    ```
3. run `air init` to create air config file should be in same folder with `main.go`
4. air config file - `cmd/web/.air.toml`
   ```bash
   root = "."
   tmp_dir = "tmp"

   [build]
     bin = "./tmp/main"
     cmd = "go build -o ./tmp/main ."
     delay = 1000
     exclude_dir = ["assets", "tmp", "vendor"]
     exclude_file = []
     exclude_regex = []
     exclude_unchanged = false
     follow_symlink = false
     full_bin = ""
     include_dir = []
     include_ext = ["go", "tpl", "tmpl", "html"]
     kill_delay = "0s"
     log = "build-errors.log"
     send_interrupt = false
     stop_on_error = true

   [color]
     app = ""
     build = "yellow"
     main = "magenta"
     runner = "green"
     watcher = "cyan"

   [log]
     time = false

   [misc]
     clean_on_exit = false
   ```
5. run `air` for auto reload

## Using curl for GET/POST/PUT Method

1. Syntax **curl -X POST -H "Content-Type: application/json" -d @FILENAME DESTINATION**
2. POST Method with a record data in json format
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"id": "4", "title": "Blue Train 04", "artist": "Do Nguyen Ha", "price": 6.99}' http://localhost:8080/albums

    curl -X POST -H "Content-Type: application/json" -d @album.json localhost:8080/albums

    // album.json file content
    {"id": "4", "title": "Blue Train 04", "artist": "Do Nguyen Ha", "price": 6.99}
    ```
3. GET Method
    ```bash
    curl -H "Content-Type: application/json" localhost:8080/albums
    ```
4. PUT Method
    ```bash
    curl -X PUT -H "Content-Type: application/json" -d @album.json localhost:8080/albums

    // album.json file content
    {"id": "4", "title": "Blue Train 04", "artist": "Do Nguyen Ha", "price": 66.99}
    ```
