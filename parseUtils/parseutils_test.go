package parseUtils

import "testing"
import "reflect"

func TestTokenize ( t *testing.T){

	var stringToTest = "command something     something"
	var stringToTest2 = "command'ilva\\n''''#$%t%^&&*"

	expected := []string{"command","something","something" }
	expected2 := []string{"command","ilva","n","t"}

	if !reflect.DeepEqual(Tokenize(stringToTest),expected){

		t.Error("Tokenize test failed on the first test")
		t.Error(" Actual :")
		t.Error(Tokenize(stringToTest))
	}
	if !reflect.DeepEqual(Tokenize(stringToTest2),expected2){

		t.Error("Tokenize test failed on the second test")
		t.Error(" Actual :")
		t.Error(Tokenize(stringToTest2))
	}

}

func TestValidate(t * testing.T) {

	var stringToTest = "ADD something something "
	var stringToTest2 = "SOI gilanujj "

	if !Validate(Tokenize(stringToTest)){
		t.Error(" Validate function failed first test ")
	}

	if Validate(Tokenize(stringToTest2)){
		t.Error(" Validate function failed second test ")
	}

}
