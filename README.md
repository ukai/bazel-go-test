sample repo to reproduce go_test doesn't work with data on Windows

```
> bazelisk test //:all
INFO: Analyzed 2 targets (0 packages loaded, 0 targets configured).
INFO: Found 1 target and 1 test target...
INFO: Elapsed time: 2.582s, Critical Path: 0.69s
INFO: 2 processes: 2 local.
INFO: Build completed, 1 test FAILED, 2 total actions
//:go_default_test
FAILED in 0.4s
  C:/users/ukai/_bazel_google/ukai/kh6vdq6z/execroot/re_client/bazel-out/x64_windows-fastbuild/testlogs/go_default_test/test.log
INFO: Build completed, 1 test FAILED, 2 total actions
```

test.log
```
Executing tests from //:go_default_test
-----------------------------------------------------------------------------
--- FAIL: TestLoad (0.00s)
    sample_test.go:8: Load open testdata/foo: The system cannot find the path specified.; want nil err
FAIL
```

code should work fine with go tooling.

```
> go test bazel-go-test/sample
ok      bazel-go-test/sample    0.167s
```
