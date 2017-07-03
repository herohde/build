# build

[![Build Status](https://travis-ci.org/seekerror/build.svg?branch=master)](https://travis-ci.org/seekerror/build)

Build is a small utility for compile-time major.minor.micro versioning
with additional linker hooks for git commit tree size and hash, such as
"1.4.2.120.4f54bc3d". The benefit of the commit tree size is that it is
unique within a branch and monotonically increasing. The hash serves to
both identify the exact commit and prevent collisions between brances.

## Usage

Define the desired version and log it on startup, say:
```
var version = build.NewVersion(1, 4, 2)

func main() {
   log.Printf("Foo %v", version)
   ...
}
```

Add the following flags to `go build` and `go install`:
```
-ldflags "-X github.com/seekerror/build.GitTree=`git rev-list HEAD | wc -l | tr -d '[[:space:]]'` -X github.com/seekerror/build.GitHash=`git rev-parse --short HEAD`"
```

## License

Build is released under the [MIT License](http://opensource.org/licenses/MIT).
