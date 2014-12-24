package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"os/user"
	"strconv"
	"fmt"
	"strings"
	utils "github.com/nerorevenge/go-samayam/samayUtils"
	parse "github.com/nerorevenge/go-samayam/parseUtils"
	file "github.com/nerorevenge/go-samayam/fileUtils"
	draw "github.com/nerorevenge/go-samayam/drawUtils"
)


// constants
const (
	H_SPACE = 1
	H_LEN =5
	V_SPACE = 2
	V_LEN =5
	EXPERIMENTAL_MAX = 4
	DIRECTORY = ".go-samayam-data"
	FILE = "data.gob"
)

type INSERT struct {
	x,y int
	buffer string

}

type TASK struct {
	Index int
	Task string
	Start time.Time 
	End time.Time
	Len int
	Hours int
	Minutes int
	Seconds int
}
// soon to be replaced by linked list
type TASK_TREE struct {
	Tree []TASK
}

func (t *TASK_TREE) add_to(x string){
	e := len(t.Tree)	
	t.Tree = append(t.Tree,TASK{Index:e,Task:x,Start:time.Now()})
}

func (t *TASK_TREE) deleteFrom(x int){
	e := len(t.Tree)
	if x <= e-1{
		t.Tree = append(t.Tree[:x],t.Tree[x+1:]...)

		for i,_ := range(t.Tree){
			t.Tree[i].Index = i
		}
	}
}

func (t *TASK_TREE) edit(index int, newTask string){
	e := len(t.Tree)
	if index <= e-1{
		 t.Tree[index].Task = newTask
		}
}

func ( t *TASK) stop_task (x,y,padding int){

	time_start := []string{}
	time_start = append(time_start,[]string{" START "," ----- "}...)
	time_start = append(time_start,utils.Formatez(t.Start)...)

	s := [][]string{[]string{strconv.Itoa(t.Index)},[]string{t.Task},time_start}

	end_string_array := utils.CustomFunction(s)
	draw.StringBox(end_string_array,x,y,padding )
	t.Len = len(end_string_array)+2*padding
}


func ( t *TASK) draw_task (x,y,padding int){

	time_start := []string{}
	time_start = append(time_start,[]string{" START "," ----- "}...)
	time_start = append(time_start,utils.Formatez(t.Start)...)

	s := [][]string{[]string{strconv.Itoa(t.Index)},[]string{t.Task},time_start}
	
	if !t.End.IsZero(){
		time_stop := []string{}
		time_stop = append(time_stop,[]string{" STOP "," ----- "}...)
		time_stop = append(time_stop,utils.Formatez(t.End)...)
		s = append(s,time_stop)	

		interval := []string{}
		interval = append(interval,[]string{" INTERVAL "," ----- "}...)
		temp := []string{"Seconds :"+strconv.Itoa(t.Seconds), "Minutes :"+strconv.Itoa(t.Minutes), "Hours   :"+strconv.Itoa(t.Hours)}
		interval = append(interval,temp...)
		s = append(s,interval)	
	}

	end_string_array := utils.CustomFunction(s)
	draw.StringBox(end_string_array,x,y,padding )
	t.Len = len(end_string_array)+2*padding
}

func (t* TASK_TREE) DrawTree(start int){

	//length := len(t.tree)
	yCounter := 0

	counter := 0
	for i:=start;i<len(t.Tree);i++{

		if counter > EXPERIMENTAL_MAX-1 {
			break
		}
		box_x := H_SPACE + H_LEN
		box_y := V_SPACE + yCounter
		t.Tree[i].draw_task(box_x,box_y,1)
		yCounter += t.Tree[i].Len
		draw.Horizontal(H_SPACE,box_y+1,H_LEN)
		if i>start {
			draw.Vertical(H_SPACE,box_y+1-V_LEN,t.Tree[i].Len)
		}
		counter++
	}

}

var main_tree TASK_TREE

func act(xs []string, destination string) {
	switch xs[0] {
		case "ADD": 
			main_tree.add_to(strings.Join(xs[1:]," "))
		case "DELETE":
			a ,_:= strconv.Atoi(xs[1])
			main_tree.deleteFrom(a)
		case "EDIT":
			a, _:=strconv.Atoi(xs[1])
			main_tree.edit(a,xs[2])
		case "STOP":
			a, _:=strconv.Atoi(xs[1])
			main_tree.Tree[a].End = time.Now()
			b:= main_tree.Tree[a].End.Sub(main_tree.Tree[a].Start)
			main_tree.Tree[a].Hours = int(b.Hours())
			main_tree.Tree[a].Minutes =  int(b.Minutes())-60*int(b.Hours())
			d := main_tree.Tree[a]
			main_tree.Tree[a].Seconds = int(b.Seconds())-(3600*d.Hours+60*d.Minutes)

		case "RESTART":
			a, _:=strconv.Atoi(xs[1])
			var b time.Time
			main_tree.Tree[a].Start = time.Now()
			main_tree.Tree[a].End = b

	}
		file.Put(destination,main_tree)
}

// log file functions and associated variables



func main(){

	errit := termbox.Init()
	defer termbox.Close()
	screen_w,screen_h := termbox.Size()

	
	if errit != nil {
		panic(" Trouble somewhere")
	} else {
		//x := BOX{width:20,x_start:3,y_start:3,height:20}
		//x.draw_mainbox()
		file.Check(DIRECTORY,FILE)

		start_index := 0

		x,_ := user.Current()
		destination := x.HomeDir+"/"+DIRECTORY+"/"+FILE
		file.Get(destination,&main_tree)
		fmt.Println(main_tree.Tree)
		for  {
			main_tree.DrawTree(start_index)

			commandbox := draw.BOX{X_start:H_SPACE,Y_start:screen_h-V_SPACE-3,Width:screen_w-H_SPACE-1,Height:3}	
			commandbox.Box()
			draw.String("COMMAND : ",H_SPACE+2,screen_h-V_SPACE-2)
			draw.String(" [ Press i to enter commands ][ Press h for help ][ Press z to quit ] ",H_SPACE+1,screen_h-1)
			termbox.Flush()
			g := termbox.PollEvent()
			ogg := INSERT{x:H_SPACE+2+len("COMMAND :"),y:screen_h-V_SPACE-2,buffer:""}
			if g.Ch == 'z'{
				break
			} else if g.Ch == 'i'{
				for {

					g :=  termbox.PollEvent()
					if  g.Key == termbox.KeyEnter {
						a:= parse.Tokenize(ogg.buffer)
						if parse.Validate(a){
							act(a,destination)
						}
						break
					}else if g.Key == termbox.KeySpace{
							ogg.buffer+=" "
					}else {	
						ogg.buffer+=string(g.Ch)
						draw.String(ogg.buffer,ogg.x,ogg.y)			
						termbox.Flush()
					}
				    }
			}else if g.Key == termbox.KeyArrowDown {
				start_index= (start_index+1)%len(main_tree.Tree)

			}
			termbox.Clear(draw.FG,draw.BG)
		}
	}

}


