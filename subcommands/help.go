package subcommands

import "fmt"

func Help() {
	fmt.Print(`Commands:	
add    - to add data to the list.
delete - to delete a single data entry to the list. 
ls   - to display the list.
clear  - to clear the list.  
`)
}
