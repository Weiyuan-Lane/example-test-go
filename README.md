

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
