package  fileUtils

import (
	"os"
	"os/user"
	"testing"
	"reflect"
)

type TestStruct struct {
	Index string 
	Samay []int
}


var fileTest string

func ContainsMethod(xs []string, contained string) bool{
	for _,j := range xs {
		if j == contained {
			return true
		}
	}
	return false
}

func TestCheck( t * testing.T ){

	testDir := ".testdir"
	testFile := "test.test"

	Check(testDir,testFile)
	x,_ := user.Current()
	destination := x.HomeDir
	os.Chdir(destination)
	if _,err := os.Stat(testDir);err!=nil {
		t.Error("GetFilePointer function failed while creating directory Error:"+err.Error())
	}
	if _,err := os.Stat(testDir+"/"+testFile);err!=nil {
		t.Error("GetFilePointer function failed while creating file")
	}
}

func TestPutAndGet(t *testing.T){

	testDir := ".testdir"
	testFile := "test.test"
	
	x,_ := user.Current()
	destination := x.HomeDir
	fileTest = destination+"/"+testDir+"/"+testFile
	
	xs := []TestStruct{TestStruct{Samay:[]int{22,336},Index:"1"},TestStruct{Samay:[]int{3,4,5},Index:"2"}}
	
	Put(fileTest,xs)
	var egg []TestStruct
	Get(fileTest,&egg)
	if !reflect.DeepEqual(egg,xs){
       	 		t.Error("put and get failed , got",egg)
	}
}
