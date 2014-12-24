/*
	wrapper for reading ,writing and updating
	the peristent data structure instance


*/

package fileUtils

import (
	"os"
 	"os/user"
 	"encoding/gob"
	)

func Check(dirName string, fileName string) {

	// goes to user home and searches for the target directory
	// makes one if not, and then changes working directory to 
	// and repeats for file

	x,_ := user.Current()
	destination := x.HomeDir
	os.Chdir(destination)
	if  _,err := os.Stat(dirName);err!=nil {
		os.Mkdir(dirName,os.ModeDir|os.ModePerm)
	}
	os.Chdir(dirName)
	if _,err := os.Stat(fileName);err!=nil {
		final,_:= os.Create(fileName)
		defer final.Close()
	}

}

func Get(file string, x interface{})  {
	file_,err := os.Open(file)
	defer file_.Close()

	if err!= nil {
		panic("Get :something went wrong while trying to open file : "+err.Error())
	}

	dataDecoder := gob.NewDecoder(file_)
	dataDecoder.Decode(x)
}


func Put( file string,x interface {} ) {
	f,err :=  os.OpenFile(file,os.O_WRONLY,os.ModePerm)
	defer f.Close()	

	if err != nil {
		panic("Put:something went wrong while trying to open file : "+err.Error())
	}

	dataDecoder := gob.NewEncoder(f)
	dataDecoder.Encode(x)


}

