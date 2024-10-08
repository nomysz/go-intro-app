# astro-weather

Intro app to start with Go.

## How to use

Grab a binary and run:
```bash
$ ./astro-weather -lat 76.9503 -lon 14.6594
CLOUDCOVER: 1, SEEING: 2, TRANSPARENCY: 2
CLOUDCOVER: 1, SEEING: 2, TRANSPARENCY: 2
CLOUDCOVER: 1, SEEING: 2, TRANSPARENCY: 2
...
```

## Development

### Requirements

- Go (download from [official website](https://go.dev/dl/))
- Go plugin for your editor of choice; Go LSP is powerfull tool

### Run

```bash
# check Spitzbergen astro weather
$ go run main.go -lon 76.9503 -lat 14.6594
CLOUDCOVER: 1, SEEING: 2, TRANSPARENCY: 2
CLOUDCOVER: 1, SEEING: 2, TRANSPARENCY: 2
CLOUDCOVER: 1, SEEING: 2, TRANSPARENCY: 2
...

# run without required flags to trigger error
$ go run main.go
Usage of ./astro-weather:
  -lat float
    	latitude for weather
  -lon float
    	longitude for weather
exit status 1
```

### Build

```bash
# Simple build for current environment
$ go build -o astro-weather

# MacOS M1 optimized build
$ GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o astro-weather

# Linux optimized build
$ GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o astro-weather
```

List of all supported architectures and operating systems can be found [here](https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63).

### Test

```bash
$ go test ./..
ok  	github.com/nomysz/go-intro-app	0.109s
```

### Nice to know

Add missing and remove unused packages:
```bash
$ go mod tidy
```

Install packages:
```bash
$ go install github.com/user/package

# or

$ go install github.com/user/package@version
```

## License

[MIT](LICENSE)
