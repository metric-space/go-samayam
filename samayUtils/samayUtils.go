package samayUtils

import (
	"strings"
	"time"
	s "strconv"
)


func array_equalizer (x []string) []string {

	max_ := 0
	// find maximu length
	for _, j := range(x) {
		if len(j) > max_ {
			max_ = len(j)		
		}
	}
	temp := make([]string,0)
	for _,j  := range(x) {
		temp = append(temp,j+strings.Repeat(" ",max_-len(j)))	
	}
	return temp
}

func Mod_split(x string, mod int ) []string {

	if len(x) < mod {
		return []string{x}
	}else {
		to_return := make([]string,0)	
		counter :=0
		temp := ""
		for _,j := range(x){
			if counter < mod{
				temp+=string(j)
				counter++
			}else{
				to_return = append(to_return,temp)
				temp = string(j)
				counter = 1
			}
		}
		to_return = append(to_return,temp)
		return to_return	
	}
}

func Format_string( x string, cut int) []string {
	// this function break up user input task from a long horizontal line
	// to multiple vertical lines
	// assume for now the input strings ia long string separated by " " 
	aka := strings.Split(x," ")
	processed_aka := make([]string,0)
	
	// preprocessing step
	for _,j := range aka {
		   processed_aka = append(processed_aka,Mod_split(j,cut)...)
		}

	return (processed_aka)
}

func Formatez(t time.Time ) []string {

	hour,min,sec := t.Clock()
	year,month,day := t.Date()
	a1 := "Day  : "+ t.Weekday().String()
	a2 := "Time : "+s.Itoa(hour)+":"+s.Itoa(min)+":"+s.Itoa(sec)
	a3 := "Date : "+s.Itoa(day)+"/"+month.String()+"/"+s.Itoa(year)

	return ([]string{a1,a2,a3})
}

func CustomFunction(s [][]string) []string {

	// what this function does is that given an input of 
	// [["hello"],["ape","ogre","buffoon"],["odis"]]
	// output will be ["hello    ape        odis",
	//                 "         ogre           ",
	//                 "         buffoon        "]

	end_string_array := make([]string,0) // the final array to be returned
	max_length_index := 0
	max_length := 0

	// find array that has most number of elements
	for i,j := range(s){
		k := len(j)
		if k > max_length {
			max_length = k
			max_length_index = i
		}
	}
	// iterate through the arrays in the array to make them 
	//of equal length, and then equalize the inner array
	for i,_ := range(s){
		if i != max_length_index{
			for len(s[i]) < max_length{ //weird change
				s[i] = append(s[i]," ")
			}	
		}
		s[i] = array_equalizer(s[i])


	}

	for i:=0;i<len(s[0]);i++{

		a := []string{}
		for  j :=0;j<len(s);j++{
			a = append(a,s[j][i])	

		}
		 end_string_array = append(end_string_array,strings.Join(a, "    " ))
	}
	return end_string_array

}
