/*

Package to validate user input for the application go-samayam

	forms:
	 ( ADD task_string) || ( DELETE INDEX) || ( EDIT INDEX task_string  )||(STOP INDEX) || ( RESTART INDEX )

*/

package parseUtils

import "regexp"

func Tokenize ( command string) []string {

	a:= regexp.MustCompile(`\w+`)	
	return a.FindAllString(command,-1)

}

func Validate ( command_string []string) bool {

	switch command_string[0] {

	case "ADD","DELETE","EDIT","RESTART","STOP":
		return true
	}

	return false
}
