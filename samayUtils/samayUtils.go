package samayUtils

import (
	"strings"
	"time"
	s "strconv"
)


func ArrayEqualizer (x []string) []string {

	max := 0
	// find maximum length
	for _, j := range(x) {
		if len(j) > max {
			max = len(j)		
		}
	}
	temp := make([]string,0)
	for _,j  := range(x) {
		temp = append(temp,j+strings.Repeat(" ",max-len(j)))	
	}
	return temp
}

func ModSplit(x string, mod int ) []string {

	if len(x) < mod {
		return []string{x}
	}else {
		toReturn := make([]string,0)	
		counter :=0
		temp := ""
		for _,j := range(x){
			if counter < mod{
				temp+=string(j)
				counter++
			}else{
				toReturn = append(toReturn,temp)
				temp = string(j)
				counter = 1
			}
		}
		toReturn = append(toReturn,temp)
		return toReturn	
	}
}

func FormatString( x string, cut int) []string {
	// UNUSED FUNCTION !!!! WHY ???????
	// this function break up user input task from a long horizontal line
	// to multiple vertical lines
	// assume for now the input strings ia long string separated by " " 
	aka := strings.Split(x," ")
	processedAka := make([]string,0)
	
	// preprocessing step
	for _,j := range aka {
		   processedAka = append(processedAka,ModSplit(j,cut)...)
		}

	return (processedAka)
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

	endStringArray := make([]string,0) // the final array to be returned
	maxLengthIndex := 0
	maxLength := 0

	// find array that has most number of elements
	for i,j := range(s){
		k := len(j)
		if k > maxLength {
			maxLength = k
			maxLengthIndex = i
		}
	}
	// iterate through the arrays in the array to make them 
	//of equal length, and then equalize the inner array
	for i,_ := range(s){
		if i != maxLengthIndex{
			for len(s[i]) < maxLength{ //weird change
				s[i] = append(s[i]," ")
			}	
		}
		s[i] = ArrayEqualizer(s[i])


	}

	for i:=0;i<len(s[0]);i++{

		a := []string{}
		for  j :=0;j<len(s);j++{
			a = append(a,s[j][i])	

		}
		 endStringArray = append(endStringArray,strings.Join(a, "    " ))
	}
	return endStringArray

}
