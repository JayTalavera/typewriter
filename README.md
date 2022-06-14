# typewriter

Typewriter package can be used to have Go print strings, character by character with a set wait time between each printed character. You can also have words be printed out one by one.

### Usage

```go
package main

import "github.com/JayTalavera/typewriter"

func main() {
	typewriter.Printf("This is the first test. \n", 75)
  
  	typewriter.Println("This is the second test.", 90)

	typewriter.Words("This is typing out one word at a time.", 250)
} 
```
