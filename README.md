# go-code-autoradio-renault

## Description

This is a minimal go module that generates a code to unlock a Renault car radio.

## Usage

```go
package main

import (
    "fmt"
    "github.com/rencedegen/go-code-autoradio-renault"
)

func main() {
    code, ok := autoradio.GenerateCode("A123")
    if ok {
        fmt.Println(code)
    } else {
        fmt.Println("The precode is not valid.")
    }
}
```

## Documentation

The documentation is available on [pkg.go.dev](https://pkg.go.dev/github.com/rencedegen/go-code-autoradio-renault).

## Executable

This module also provides an executable that can be used to generate the code.

### Compilation from source

```bash
> go install github.com/rencedegen/go-code-autoradio-renault/cmd/renaultcode
> renaultcode A123
```

### Download

You can download the executable from the [releases](https://github.com/rencedegen/go-code-autoradio-renault/releases) page.

## License

[MIT](LICENSE)


