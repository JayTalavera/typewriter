# typewriter

Typewriter package can be used to have Go print strings, character by character with a set wait time between each printed character. You can also have words be printed out one by one.

# Usage
There are three typewriter functions:
- `typewriter.Printf` - Will only print out a string without a newline at the end.
- `typewriter.Println` - Will only print out a string with a newline appended.
- `typewriter.Words` - Will only print out words one by one with a space added to the end of each word. Newlines aren't currently supported.

Each function takes a string and an integer:
- `typewriter.Printf(string, int)`
- `typewriter.Println(string, int)`
- `typewriter.Words(string, int)`

The `int` is the number of milliseconds set to wait between each character print.

If the `int` is set to `0` it will default to `50`:
- `typewriter.Printf("Hello, World!", 0)` is the same as `typewriter.Printf("Hello, World!", 50)`
```go
package main

import "github.com/JayTalavera/typewriter"

func main() {
	typewriter.Printf("This is using \"typewriter.Printf\" and we must use \\n to add a newline. \n", 75)
  
  	typewriter.Println("This is using \"typewriter.Println\" and a newline is appended.", 75)

	typewriter.Words("This is using \"typewriter.Words\"  and it prints out one word at a time.", 250)
} 
```

![typewriter](https://user-images.githubusercontent.com/95598769/173625796-2e64f535-7506-49e1-8e69-be4f83912c75.gif)

### Demo
![typewriter_demo](https://user-images.githubusercontent.com/95598769/173627027-73b41fc3-250a-4ed5-ac6c-b544acc9e1eb.gif)
