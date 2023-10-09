# Memory game

This will be a simple memory game server written in Go.

### Compilation instructions
1. Download Go 1.21.2 from the [Official Page](https://go.dev/doc/install).
2. Compile by running in the root folder `go build -o build/server main/main.go`
3. Execute by running the `./build/server`. This will deploy the game and will be running in the port 8080.

### Executing test
1. Go to the `test` directory.
2. Run the command `go mod tidy` to download the [testify](https://github.com/stretchr/testify) library. This is a thirdparty library that enhances the native Go's testing system to be similar to the one use at [JUnit 5](https://junit.org/junit5/).
3. Run the command `go test`.


The project's structure will be the one used at [Go's project layout](https://github.com/golang-standards/project-layout), but with some modifications to use the MVC pattern.