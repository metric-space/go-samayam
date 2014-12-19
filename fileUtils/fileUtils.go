/*
	wrapper for reading ,writing and updating
	the peristent data structure instance


*/

package fileUtils

import "os"
import "os/user"
import "encoding/gob"
//import "fmt"

func GetFilePointer(dirName string, fileName string) {
	x,_ := user.Current()
	destination := x.HomeDir
	os.Chdir(destination)
	if  _,err := os.Stat(dirName);err!=nil {
		os.Mkdir(dirName,os.ModeDir|os.ModePerm)
	}
	os.Chdir(dirName)
	final,_:= os.Create(fileName)
	//if err!=nil {
	//		panic("error whhile opening file")
	//}
	final.Close()

}

func Get(file string, x interface{})  {
	file_,_ := os.Open(file)
	dataDecoder := gob.NewDecoder(file_)
	dataDecoder.Decode(x)
	//if err != nil {
	//		fmt.Println("something went wrong while decoding ,error:"+err.Error())
	//}
	file_.Close()
}


func Put( file string,x interface {} ) {
	f,_ :=  os.OpenFile(file,os.O_WRONLY,os.ModePerm)

	dataDecoder := gob.NewEncoder(f)
	dataDecoder.Encode(x)

	//if err != nil {
        //		fmt.Println("something went wrong while encoding ,error:"+err.Error())
	//}
	f.Close()

}

