package  fileUtils

import (
	"os"
	"os/user"
	"testing"
	"reflect"
)

type test_struct struct {
	Index string 
	Samay []int
}


var file_test string

func contains_method(xs []string, contained string) bool{
	for _,j := range xs {
		if j == contained {
			return true
		}
	}
	return false
}

func TestGetFilePointer( t * testing.T ){

	test_dir := ".testdir"
	test_file := "test.test"

	GetFilePointer(test_dir,test_file)
	x,_ := user.Current()
	destination := x.HomeDir
	os.Chdir(destination)
	if _,err :=os.Stat(test_dir);err!=nil {
		t.Error("GetFilePointer function failed while creating directory Error:"+err.Error())
	}
	if _,err :=os.Stat(test_dir+"/"+test_file);err!=nil {
		t.Error("GetFilePointer function failed while creating file")
	}
	//if err :=os.RemoveAll(test_dir+"/"+test_file);err!=nil {
	//	t.Error("Could not delete file and directory")
	//}
}

func TestPutandGet(t *testing.T){


	test_dir := ".testdir"
	test_file := "test.test"
	
	x,_ := user.Current()
	destination := x.HomeDir
	file_test = destination+"/"+test_dir
	err := os.Chdir(file_test)
	file_test = destination+"/"+test_dir+"/"+test_file
	
	xs := []test_struct{test_struct{Samay:[]int{1,2,3,4,5},Index:"1"},test_struct{Samay:[]int{3,4,5},Index:"2"}}
	
	Put(file_test,xs)
	var egg []test_struct
	Get(file_test,&egg)
	if !reflect.DeepEqual(egg,xs){
        		t.Error("put and get failed , got",egg)
	}
}
