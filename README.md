# ApprovalTests.Go and the Go test cache

This repository shows a problem with ApprovalTests.Go. If a Go package runs an
approval test, `go test` does not put the result of that package in the test
cache. The tests run again each time, but no files changed.

This repository has two packages. The root package runs an approval test. The
`control` package has an empty test. The `control` package shows the usual
behavior of the test cache.

## How to show the problem

Run these commands:

```sh
go clean -testcache
go test ./...
go test ./...
go test ./...
```

After the first run, the `control` package shows `(cached)`. The root package
does not show `(cached)`. The root package runs again each time.

```text
ok  github.com/cgt/approval-tests-cache-repro          0.301s
ok  github.com/cgt/approval-tests-cache-repro/control  0.490s
ok  github.com/cgt/approval-tests-cache-repro          0.152s
ok  github.com/cgt/approval-tests-cache-repro/control  (cached)
ok  github.com/cgt/approval-tests-cache-repro          0.144s
ok  github.com/cgt/approval-tests-cache-repro/control  (cached)
```

Go can show the cause. Run this command:

```sh
GODEBUG=gocachetest=1 go test .
```

The output contains this line:

```text
input file /.../.approval_tests_temp/.gitignore: file used as input is too new
```

## The cause

ApprovalTests.Go makes a directory with the name `.approval_tests_temp`. It
makes this directory in the source directory of the package. Then it writes
three files in that directory. It writes these files again at each test run.

`go test` makes a list of each file that a test opens. If a file is in the
module of the package, `go test` examines that file again. If the test opened a
file that is less than two seconds old, `go test` does not put the result in the
cache.

The library writes these three files immediately before the test stops. The
files are always less than two seconds old. Because of this, `go test` does not
make a cache entry. There is no cache entry to use later.

## More information

- [The `go test` documentation](https://pkg.go.dev/cmd/go#hdr-Test_packages)
- [The `modTimeCutoff` code in `cmd/go`](https://go.dev/src/cmd/go/internal/test/test.go)

---

*Generated with Claude Code / Claude Opus 5*
