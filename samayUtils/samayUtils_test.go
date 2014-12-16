package samayUtils

import "testing"
import "reflect"

func TestFirst (t *testing.T){
	a := array_equalizer([]string{"aplha","beta","g"})
	for _,j := range(a){
		if len(j) != 5{
			t.Errorf(" Length isn't right ")
		}

	}

}

func TestOfCustom(t *testing.T){

	a := CustomFunction([][]string{[]string{"xxxx"},[]string{"aaa","bbbb","cccc"},[]string{"dd"}})
	b :=  []string{"xxxx    aaa     dd",
	                 		"        bbbb      ",
			 		"        cccc      "} 
	if !reflect.DeepEqual(a,b){
			 t.Error("Actual output")
			 for _,j := range(a){
				t.Error("-"+j+"-")
			 }

			t.Error("Expected output")
			 for _,j := range(b){
				t.Error("-"+j+"-")
			 }

			t.Errorf("CUstom function isn't working the way it is supposed to ");
	 }

}
