# Go Corona Go

Go Corona Go is a command line utility to provide information about Covid-19 related cases in India.

## Installation
Installation can be done using the `go get` command which will take care of installation of any libraries and dependencies nescessary. This will also install the `go-corona-go` executable which can be used anywhere in the termnial provided `$GOPATH/bin` is in your `PATH`.

```sh
go get -u github.com/MadhavJivrajani/go-corona-go
```

## Usage
```txt
A command line utility for getting information related to Covid-19 in India.

Usage:
  go-corona-go [command]

Available Commands:
  district    Provides district-wise stats of a particular state for Covid-19 in India
  districts   Gets a list of valid districts of a particular state whose stats can be retreived
  help        Help about any command
  plot        Plots daily data about a particular statistic about Covid-19 in India
  sike        sike is sike
  state       Provides state wise stats of Covid-19 in India
  states      Gets a list of valid states whose stats can be retreived

Flags:
      --config string   config file (default is $HOME/.go-corona-go.yaml)
  -h, --help            help for go-corona-go
  -t, --toggle          Help message for toggle

Use "go-corona-go [command] --help" for more information about a command.
```