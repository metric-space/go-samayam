package main

import ("fmt";"time";"os/exec";"os")

const heading string = " go-samayam "
// experimental version

func gen_iterator(start, end int) func() int{
	start_ := start
	end_ := end
	return func () int {
		if start_<end_ {
			start_ += 1
			return start_ 
		} else {
			return 0
		}

	}
}

func screen_writer( xs [2](func() int) ){


	for {

		command := exec.Command("clear")
		command.Stdout = os.Stdout
		command.Run()
		for y :=0;y<2;y++{			

		   fmt.Println(xs[y]())	   		
		}
		time.Sleep(1000*time.Millisecond)
	}
}


func main(){
	x := [2]( func() int){  gen_iterator(10,100),gen_iterator(300,400)}
	screen_writer(x)

}
