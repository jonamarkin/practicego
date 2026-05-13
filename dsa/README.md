# DSA Exercises

Each problem lives in its own folder as a separate runnable Go package.
That keeps every exercise free to have its own `main()` function.

Run one exercise:

```sh
go run ./dsa/simple-array-sum
go run ./dsa/very-big-sum
go run ./dsa/diagonal-difference
```

Run all tests:

```sh
go test ./...
```

When adding a new problem, create a folder under `dsa/`:

```text
dsa/problem-name/main.go
dsa/problem-name/main_test.go
```

Use `package main` inside each problem folder.
