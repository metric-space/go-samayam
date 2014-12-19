package parseUtils

import "testing"
import "reflect"

func TestTokenize ( t *testing.T){

	var string_to_test = "command something     something"
	var string_to_test2 = "command'ilva\\n''''#$%t%^&&*"

	expected := []string{"command","something","something" }
	expected2 := []string{"command","ilva","n","t"}

	if !reflect.DeepEqual(Tokenize(string_to_test),expected){

		t.Error("Tokenize test failed on the first test")
		t.Error(" Actual :")
		t.Error(Tokenize(string_to_test))
	}
	if !reflect.DeepEqual(Tokenize(string_to_test2),expected2){

		t.Error("Tokenize test failed on the second test")
		t.Error(" Actual :")
		t.Error(Tokenize(string_to_test2))
	}

}

func TestValidate(t * testing.T) {

	var string_to_test = "ADD something something"
	var string_to_test2 = "SOI gilanujj "

	if !Validate(Tokenize(string_to_test)){
		t.Error(" Validate function failed first test ")
	}

	if Validate(Tokenize(string_to_test2)){
		t.Error(" Validate function failed second test ")
	}

}
