# Go Corona Go :neckbeard:

Go Corona Go is a command line utility for getting information related to Covid-19 in India.

## Installation
Installation can be done using the `go get` command which will take care of installation of any libraries and dependencies nescessary. This will also install the `go-corona-go` executable which can be used anywhere in the termnial provided `$GOPATH/bin` is in your `PATH`.  
```sh
go get -u github.com/MadhavJivrajani/go-corona-go
```  
Or you can clone this repository and work directly with the executable `go-corona-go`!  
```
git clone https://github.com/MadhavJivrajani/go-corona-go.git
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
### `state` command
This command information about a particular (valid) state.

Usage:
```
go-corona-go state [state]
```

Example:
```
go-corona-go state Karnataka
```

Result:
```
Info was last updated on: 22/07/2020 21:57:09

+--------------------------+-------+
|        KARNATAKA         |       |
+--------------------------+-------+
| Active cases             | 47066 |
| Confirmed cases          | 75833 |
| No. of deaths            |  1519 |
| No. of recoveries (yay!) | 27239 |
+--------------------------+-------+
```
### `district` command
This command provides information for a particular (valid) district of a particular state.

Usage:
```
go-corona-go district [state] [district]
```

Example:
```
go-corona-go district maharashtra pune
```

Result:
```
Info was last updated on: 22/07/2020 21:57:05

+--------------------------+-------+
|           PUNE           |       |
+--------------------------+-------+
| Active cases             | 39353 |
| Confirmed cases          | 63351 |
| No. of deaths            |  1514 |
| No. of recoveries (yay!) | 22484 |
+--------------------------+-------+
```
### `states` command
This command gives a list of valid states, in alphabetical order, whose data can be retreived.

Usage:
```
go-corona-go states
```
### `districts` command
This command gives a list of valid distrcits of a valid state, in particular order, whose data can be retreived.

Usage:
```
go-corona-go districts [state]
```

Example:
```
go-corona-go distrcits karnataka
```
### `plot` command
This command plots daily data about a particular `type` of data about Covid-19 in India.
Plotting functionality is provided using the [asciigraph](https://github.com/guptarohit/asciigraph) library.

Usage:
```
go-corona-go plot -t [type]
[type]: confirmed (default), deaths, recovered

confirmed: plots trend of number of confirmed cases of Covid-19 in India each day
deaths : plots trend of number of deaths in India due to Covid-19 each day
recovered: plots trend of number of recoveries from Covid-19 in India each day
```

Example:
```
go-corona-go plot -t recovered
```

Result:
```
Plot information:
        IMPORTANT: This plot is intended to show ONLY the trend of data and should not be interpreted accurately

        X Axis: Dates starting from 30th January 2020.
        Y Axis: Number of recoveries from Covid-19 in India.
 57759 ┼                                                ╭ 
 54871 ┤                                                │ 
 51983 ┤                                              ╭─╯ 
 49095 ┤                                              │   
 46207 ┤                                              │   
 43319 ┤                                              │   
 40431 ┤                                              │   
 37543 ┤                                             ╭╯   
 34655 ┤                                           ╭─╯    
 31767 ┤                                           │      
 28880 ┤                                           │      
 25992 ┤                                           │      
 23104 ┤                                          ╭╯      
 20216 ┤                                        ╭─╯       
 17328 ┤                                       ╭╯         
 14440 ┤                                    ╭──╯          
 11552 ┤                                    │             
  8664 ┤                                 ╭──╯             
  5776 ┤                             ╭╮╭─╯                
  2888 ┤                         ╭───╯╰╯                  
     0 ┼─────────────────────────╯                        
          A plot showing the trend of the number of deaths due to Covid-19 per day.

```
### `sike` command
I love this command :eyes:  
It will display one of three different messages everytime it is called, feel free to add more!

Usage
```
go-corona-go sike
```

Result:

```
Meet ken the chicken. ken crossed the road 'cause it ken.

    \\
    (o>
 \\_//)
  \_/_)
   _|_
   


This amazing ASCII art was taken from http://www.ascii-art.de/ascii/c/chicken.txt
```
(there are two more messages like this one XD)

## Credits:
- The API used to get information can be found [here](https://rapidapi.com/spamakashrajtech/api/corona-virus-world-and-india-data?endpoint=apiendpoint_e53bab74-70b7-42e9-9d95-4667fdcfa876).
- The [cobra](https://github.com/spf13/cobra) library was used to develop the CLI.
- The [tablewriter](https://github.com/olekukonko/tablewriter) library was used to format and print the data in a tabular form.
- The [asciigraph](https://github.com/guptarohit/asciigraph) library was used to provide in-terminal plotting functionality.
