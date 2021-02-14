# Example-test-go

This is an example project to support a [blogpost](https://weiyuan-liu.medium.com/tips-to-testing-and-coverage-in-go-ea5b9091700f) that I was intending to write

In this project, we will look at three different areas - using `ginkgo`, `gomock` and additional logic I have added to make `go test` and `go cover` work more seamlessly, with a UI to boot.

To run this project simply run the following command (make sure to have `docker` installed):
```
docker-compose up
```

Then shell into the container, using the following:
```
docker exec -it service sh
```

## Using Ginkgo and Gomega for Behavior-Driven Development ("BDD")
In this project, tests are written for `internal/server/utils/queryparams`

To replicate, here are some guiding steps on the above:

1) Change directory to the target package that you're are testing. In this case, `internal/server/utils/queryparams`.
2) For each distinct package to be tested, bootstrap should be run. Use the following command in the same folder.
```
ginkgo bootstrap
```
3) If there are multiple files in your package, you can separate the tests in their own files. Replace `%FILENAME%` to the name of your file (e.g. `queryparams` for `queryparams.go`)
```
ginkgo generate %FILENAME%
```

After completing the steps above, you should an empty test file resembling `internal/server/utils/queryparams/queryparams_test.go`, which you can complete with tests like those in that file.


## Using Mockgen to mock interfaces and methods
In this project, there is a mock file at `test/server/utils/queryparams/queryparams.go`

To create a mockfile for your own project, run the following command:

```
mockgen -source=pathTo/file.go -destination=pathTo/mockFile.go
```

## Running tests and coverage

One of my pet peeves with `go test` is having to remember what each option does, along with reviewing coverage.

I've decided to create a Makefile in this project, so that I can reference it from other Go projects.

To run all the tests, simply call:
```
make run_tests
```

To get a view of different coverage review of your code, AND also get a UI in a webpage to depict what code has been covered / not covered, run:
```
make coverage
```
