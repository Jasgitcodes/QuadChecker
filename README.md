# quadchecker

`quadchecker` is a simple Go program that reads ASCII art from standard input and checks whether the input matches one of five predefined "quad" rectangle patterns.

## What it does

- Reads the full input from stdin.
- Trims trailing whitespace and determines the width and height of the shape.
- Compares the input with generated output from each of the pattern functions:
  - `QuadA`
  - `QuadB`
  - `QuadC`
  - `QuadD`
  - `QuadE`
- Prints matched patterns with dimensions in the format:
  - `[quadA] [width] [height]`
- If no patterns match, prints:
  - `Not a quad function`

## Supported patterns

- `quadA`: corners `o`, horizontal edges `-`, vertical edges `|`, spaces inside.
- `quadB`: corner slashes `/` and `\\`, edges `*`, spaces inside.
- `quadC`: top corners `A`, bottom corners `C`, other edges `B`, spaces inside.
- `quadD`: all four corners are `A`/`C` depending on position, other edges `B`, spaces inside.
- `quadE`: top corners `A`, bottom corners `C`, top/bottom edges `B`, left/right edges `B`, spaces inside.

## Usage

Build the program:

```bash
go build -o quadchecker main.go quadA.go quadB.go quadC.go quadD.go quadE.go
```

Then run it by piping a quad pattern into it:

```bash
cat input.txt | ./quadchecker
```

Or use a here-document:

```bash
./quadchecker <<EOF
o--o
o  o
o--o
EOF
```

## Generating quad A-E binaries dynamically with `main-test.go`

The file `main-test.go` is part of the same `main` package as `quadA.go` through `quadE.go`, so it can be used as a reusable build harness for the individual quad generators.

1. Uncomment or add a `main` function in `main-test.go` that parses two numeric arguments and prints one of the generators.
2. Set the generator call to the desired pattern, for example:

```go
fmt.Print(QuadA(x, y))
```

3. Build the binary for each quad:

```bash
go build -o quadA main-test.go quadA.go
go build -o quadB main-test.go quadB.go
go build -o quadC main-test.go quadC.go
go build -o quadD main-test.go quadD.go
go build -o quadE main-test.go quadE.go
```

The `-o quadX` suffix names each executable after the quad generator it produces. This makes the binaries easy to distinguish and run directly, e.g. `./quadA 5 5` for the `QuadA` shape.

4. Run the binary with width and height:

```bash
./quadA 5 5
```

Because `main-test.go` compiles in the same package as the quad generator files, the quad implementations are included automatically during build.

## Project structure

- `main.go` – main program: reads stdin and compares input against quad generators.
- `quadA.go` – implements the `QuadA` generation function.
- `quadB.go` – implements the `QuadB` generation function.
- `quadC.go` – implements the `QuadC` generation function.
- `quadD.go` – implements the `QuadD` generation function.
- `quadE.go` – implements the `QuadE` generation function.
- `main-test.go` – contains the `QuadA` implementation used during testing / development.
- `go.mod` / `go.sum` – Go module files.

## Notes

- The program expects clean ASCII input with consistent line lengths.
- Empty or whitespace-only input results in `Not a quad function`.
