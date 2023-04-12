## FMTCOLOR LIBRARY FOR ECLA

[![Go Report Card](https://goreportcard.com/badge/github.com/Eclalang/fmtColor)](https://goreportcard.com/report/github.com/Eclalang/fmtColor)

fmtColor is a simple package that allows you to print colored text in the terminal.

# Candidate functions :

|  Func Name  |                         Prototype                         |                              Description                              | Comments |
|:-----------:|:---------------------------------------------------------:|:---------------------------------------------------------------------:|:--------:|
|  PrintRGB   |           PrintRGB(r, g, b int, text string) {}           |                    Prints text in the given color                     |   N/A    |
| PrintlnRGB  |          PrintlnRGB(r, g, b int, text string) {}          |           Prints text in the given color and adds a newline           |   N/A    |
|  PrintfRGB  |  PrintfRGB(r, g, b int, format string, values ...any) {}  |         Prints the values in the given color and formats them         |   N/A    |
| PrintflnRGB | PrintflnRGB(r, g, b int, format string, values ...any) {} | Prints the values in the given color, formats them and adds a newline |   N/A    |
|    Print    |          Print(c color.Color, values ...any) {}           |                 Prints the values in the given color                  |   N/A    |         
|   Println   |         Println(c color.Color, values ...any) {}          |        Prints the values in the given color and adds a newline        |   N/A    |         
|   Printf    |  Printf(c color.Color, format string, values ...any) {}   |         Prints the values in the given color and formats them         |   N/A    |         
|  Printfln   | Printfln(c color.Color, format string, values ...any) {}  | Prints the values in the given color, formats them and adds a newline |   N/A    |         

# Available colors :
| Color Name | RGB Color Code |
|:----------:|:--------------:|
|    Red     |   255, 0, 0    |
|   Green    |   0, 255, 0    |
|    Blue    |   0, 0, 255    |
|   Yellow   |  255, 255, 0   |
|    Cyan    |  0, 255, 255   |
|  Magenta   |  255, 0, 255   |
|   White    | 255, 255, 255  |
|   Black    |    0, 0, 0     |