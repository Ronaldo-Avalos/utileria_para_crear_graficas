#### Comenzamos el proyecto con la creacion del archivo main en Go

```GO
package main

import (
	
)
func main() {
	
}
 ```
 #### Se añadieron las opciones de casos que se presentaran
 
  ```GO
  func main() {
  op1 := os.Args[1]
	op2 := os.Args[2]
	if op2 == "--generate" {
		switch op1 {
		case "--bar":
		case "--line":	
		case "--pie":
		}
	} else {

		switch op1 {
		case "--bar":
		case "--line":
		case "--pie":

		}
	}
}
   ```
