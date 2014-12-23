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

func Validate ( commandString []string) bool {

	length := len(commandString)

	switch commandString[0] {

	case "ADD":
		if length >= 2 {
			return true
		}
		break
	case "DELETE":
		if length == 2 {
			return true
		}
		break
	case "EDIT":
		if length == 3 {
			return true
		}
		break
	case "RESTART":
		if length == 1 {
			return true
		}
		break
	case "STOP":
		if length == 1 {
			return true
		}
		break	
	
	}

	return false
}
