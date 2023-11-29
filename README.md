# Advent of Code 2023

Solutions for the 2023 Advent of Code

## Building

This project makes use of Go 1.19.

```bash
go mod download
go test ./...
```

## Running the Solutions

To run a solution, use the problem name, which consists of the day number followed
by a or b for the problem.

For example, to run problem 2a:

```bash
$ go run ./main.go 2 a
Answer: 9633
Took 999.4µs
```

If `input.txt` is not present in the folder for the specified day, or you want to
use a different file, then you can specify it using the `-i` parameter. For example:

```bash
$ go run ./main.go 2 a -i challenge/day2/alternate_input.txt
Answer: 123
Took 464.3µs
```

## Adding New Solutions

A generator program is included that makes templates for each day, automatically
downloading challenge input and updating the root command to add new subcommands
for each problem. Running `go generate` from the repo root will generate the
following for each day that is currently accessible:

* `challenge/day<N>/import.go`: A "glue" file combining commands for both of the day's problems to simplify wiring up subcommands
* `challenge/day<N>/a.go`: The main problem implementation, containing a cobra command `A` and the implementation `func a(*challenge.Input) int`
* `challenge/day<N>/a_test.go`: A basic test template
* `challenge/day<N>/input.txt`: The challenge input

Additionally, `challenge/cmd/cmd.go` will be regenerated to import and add all
subcommands.

This requires `goimports` be available on your `$PATH`. Additionally, you must be
logged into https://adventofcode.com in your browser so the generator can use your session
cookie to download challenge input.

Existing solutions and challenge inputs will be skipped instead of regenerated.

## License

These solutions are licensed under the MIT License.

See [LICENSE](./LICENSE) for details.

## Attribution

The generator (under the `gen/` folder, and `main.go` are heavily based on [nlowe's aoc2020 GitHub repository](https://github.com/nlowe/aoc2020/tree/4c712a8df8904ecc3b886105041c2d011c58aa4c). For the license for this code, please see ATTRIBUTION for details.)
