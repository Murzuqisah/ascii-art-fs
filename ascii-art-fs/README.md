
# ASCII ART
The ASCII Art program written in go language, is designed to take a string as an argument and output the string in a graphic representation using ASCII characters. This graphical representation is created by rendering the input string with ASCII characters to form a visual representation.

File Structure: main folder(asciiart), contains(3 banner files; shadow.txt,standard.txt and thinkertoy.txt, a ReadMe.md file containing the documantaion of the program, 5 function folders with their test files and a go mod file that contains modules that help in runing the program )



## Features



## Demo

Insert gif or link to demo

Usage:
+ To run the program, follow these steps:

    2.Navigate to the program folder by copying the command below to the terminal 
```bash
    cd asciiart
```

    3.Run the program using the following command/ copy the command below to the terminal
```bash
    go run . "desired string" specify bannerfile
```

    4. To display the output with newline characters visible, you can add | cat -e at the end of your last argument like shown below 
```bash
    go run . "\n" bannerfile | cat -e
```
                
Example:
For example, running the program with the commands below

command 1.

```bash
go run . "Hello Alice!" shadow
```

would output an ASCII representation of the string "Hello, ASCII Art!" using the shadow banner style as shown below
```bash
                                                                           
_|    _|          _| _|                  _|_|   _| _|                   _| 
_|    _|   _|_|   _| _|   _|_|         _|    _| _|      _|_|_|   _|_|   _| 
_|_|_|_| _|_|_|_| _| _| _|    _|       _|_|_|_| _| _| _|       _|_|_|_| _| 
_|    _| _|       _| _| _|    _|       _|    _| _| _| _|       _|          
_|    _|   _|_|_| _| _|   _|_|         _|    _| _| _|   _|_|_|   _|_|_| _| 
                                                                           
```

command 2.

```bash
go run . "Hello Brian!" thinkertoy
```

would output an ASCII representation of the string "Hello, ASCII Art!" using the thinkertoy banner style as shown below
```bash
                                               
o  o     o o           o--o                  o 
|  |     | |           |   |     o           | 
O--O o-o | | o-o       O--o  o-o    oo  o-o  o 
|  | |-' | | | |       |   | |   | | |  |  |   
o  o o-o o o o-o       o--o  o   | o-o- o  o O 

``` 
command 3.

```bash
go run . "Hello Joe!" standard
```

would output an ASCII representation of the string "Hello, ASCII Art!" using the standard banner style as shown below
```bash

 _    _          _   _                      _                  _  
| |  | |        | | | |                    | |                | | 
| |__| |   ___  | | | |   ___              | |   ___     ___  | | 
|  __  |  / _ \ | | | |  / _ \         _   | |  / _ \   / _ \ | | 
| |  | | |  __/ | | | | | (_) |       | |__| | | (_) | |  __/ |_| 
|_|  |_|  \___| |_| |_|  \___/         \____/   \___/   \___| (_) 
  
```
## Running Tests

To run tests, run the following command

```bash
  npm run test
```
to run the Test
``` go test

## Authors

- [@Aliceokingo](https://learn.zone01kisumu.ke/git/aokingo)
- [@joelamos](https://learn.zone01kisumu.ke/git/jamos)
- [@brianshisia](https://learn.zone01kisumu.ke/git/bshisia)
## License

https://github.com/go-gitea/gitea/blob/main/LICENSE

