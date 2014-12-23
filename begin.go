package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"os/user"
	"strconv"
	"fmt"
	utils "github.com/nerorevenge/go-samayam/samayUtils"
	parse "github.com/nerorevenge/go-samayam/parseUtils"
	file "github.com/nerorevenge/go-samayam/fileUtils"
)


// unicode constants
const (
	v = '\u2502'
	h = '\u2500'
	box_ul = '\u256d'
	box_ur = '\u256e'
	box_ll = '\u2570'
	box_lr = '\u256f'
	name = " SAMAYAM "
	FG = termbox.ColorWhite
	BG = termbox.ColorBlack
	H_SPACE = 1
	H_LEN =5
	V_SPACE = 2
	V_LEN =5
	EXPERIMENTAL_MAX = 4
	DIRECTORY_ = ".go-samayam-data"
	FILE_= "data.gob"
)

type BOX struct {

	width,height int	
	x_start,y_start int

}

type INSERT struct {
	x,y int
	buffer string

}

type TASK struct {
	Index int
	Task string
	Start time.Time 
	End time.Time
}
// soon to be replaced by linked list
type TASK_TREE struct {
	Tree []TASK
}

func (t *TASK_TREE) add_to(x string){
	e := len(t.Tree)	
	t.Tree = append(t.Tree,TASK{Index:e,Task:x,Start:time.Now()})
}

func draw_string(fill string,x,y int){
	for i,j := range fill {
			termbox.SetCell(x+i,y,j,FG,BG)
		}

}


func draw_string_box(fill []string,x,y,padding int){


	actual_width := len(fill[0]) + 2*padding
	actual_height :=  len(fill) + 2*padding
	actual_start_x := x + padding
	actual_start_y := y + padding

	for i,j := range(fill){
		draw_string(j,actual_start_x,actual_start_y+i)	
	}
	rect := BOX{ x_start:x,y_start:y,width:actual_width,height:actual_height}
	rect.draw_mainbox()
}

func draw_vertical(x,y,h int){

	for i := y ; i <= y+h-1 ; i++ {
		termbox.SetCell(x,i,v,FG,BG)
	}
}

func draw_horizontal(x,y,w int){

	for i := x ; i <= x+w-1 ; i++ {
		termbox.SetCell(i,y,h,FG,BG)
	}
}

func box_filler(x,y,w int){
	// debug function
	a := [5] termbox.Attribute{ termbox.ColorRed,termbox.ColorGreen,termbox.ColorYellow } 
	for i := x ; i <= x+w-1 ; i++ {
		termbox.SetCell(i,y,h,termbox.ColorWhite,a[i%3])
	}
}

func (r *BOX) draw_mainbox(){

	termbox.SetCell(r.x_start,r.y_start,box_ul,FG,BG)
	draw_horizontal(r.x_start+1,r.y_start,r.width-2)
	termbox.SetCell(r.x_start+r.width-1,r.y_start,box_ur,FG,BG)

	termbox.SetCell(r.x_start,r.y_start+r.height-1,box_ll,FG,BG)
	draw_horizontal(r.x_start+1,r.y_start+r.height-1,r.width-2)
	termbox.SetCell(r.x_start+r.width-1,r.y_start+r.height-1,box_lr,FG,BG)

	draw_vertical(r.x_start,r.y_start+1,r.height-2)
	draw_vertical(r.x_start+r.width-1,r.y_start+1,r.height-2)

}

func ( t *TASK) draw_task (x,y,padding int){

	time_start := []string{}
	time_start = append(time_start,[]string{" START "," ----- "}...)
	time_start = append(time_start,utils.Formatez(t.Start)...)

	s := [][]string{[]string{strconv.Itoa(t.Index)},[]string{t.Task},time_start}

	end_string_array := utils.CustomFunction(s)
	draw_string_box(end_string_array,x,y,padding )

}

func (t* TASK_TREE) draw_tree(start int){

	//length := len(t.tree)
	counter := 0
	for i:=start;i<len(t.Tree);i++{

		if counter > EXPERIMENTAL_MAX-1 {
			break
		}

		box_x := H_SPACE + H_LEN
		box_y := V_SPACE + counter*V_LEN
		t.Tree[i].draw_task(box_x,box_y,1)
		draw_horizontal(H_SPACE,box_y+1,H_LEN)
			
		if i>start {
			draw_vertical(H_SPACE,box_y+1-V_LEN,V_LEN)
		}
		counter++
	}

}

var main_tree TASK_TREE

func act(xs []string, destination string) {
	switch xs[0] {
		case "ADD": 
			main_tree.add_to(xs[1])
			file.Put(destination,main_tree)

	}
}

// log file functions and associated variables



func main(){

	errit := termbox.Init()
	screen_w,screen_h := termbox.Size()

	
	if errit != nil {
		panic(" Trouble somewhere")
	} else {
		//x := BOX{width:20,x_start:3,y_start:3,height:20}
		//x.draw_mainbox()
		
		file.Check(DIRECTORY_,FILE_)

		start_index := 0

		x,_ := user.Current()
		destination := x.HomeDir+"/"+DIRECTORY_+"/"+FILE_
		file.Get(destination,&main_tree)
		fmt.Println(main_tree.Tree)
		for  {
					 
			main_tree.draw_tree(start_index)

			commandbox := BOX{x_start:H_SPACE,y_start:screen_h-V_SPACE-3,width:screen_w-H_SPACE-1,height:3}	
			commandbox.draw_mainbox()
			draw_string("COMMAND : ",H_SPACE+2,screen_h-V_SPACE-2)
			draw_string(" [ Press i to enter commands ][ Press h for help ][ Press z to quit ] ",H_SPACE+1,screen_h-1)
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
						draw_string(ogg.buffer,ogg.x,ogg.y)			
						termbox.Flush()
					}
				    }
			}else if g.Key == termbox.KeyArrowDown {
				start_index= (start_index+1)%len(main_tree.Tree)

			}
			termbox.Clear(FG,BG)
		}
	}
	defer termbox.Close()

}


