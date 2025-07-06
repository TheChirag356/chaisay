# ChaiSay

A command-line utility and library that generates ASCII art of a chai (or other beverages) with a message.

## Installation

```bash
go get github.com/TheChirag356/chaisay
```

## Usage

Import the package in your file
```bash
import "github.com/TheChirag356/chaisay"
```

Create a ChaiSay struct
```go
c := ChaiSay{
    Text: "Text goes here",
    ArtStyle: "Art Style goes here, default is CHAI.",
}
```

And print the message
```go
c.Print()
```

## Features

- Display custom messages
- Choose your art from a list of items
- Simple and lightweight

## Contributing

Pull requests are welcome. For major changes, please open an issue first.

## License

MIT