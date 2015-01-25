package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"os/user"
	"strconv"
	"strings"
	utils "github.com/nerorevenge/go-samayam/samayUtils"
	parse "github.com/nerorevenge/go-samayam/parseUtils"
	file "github.com/nerorevenge/go-samayam/fileUtils"
	draw "github.com/nerorevenge/go-samayam/drawUtils"
)


// constants
const (
	HSPACE = 1
	HLEN =5
	VSPACE = 2
	VLEN =5
	EXPMAX = 4
	DIRECTORY = ".go-samayam-data"
	FILE = "data.gob"
)

type Insert struct {
	x,y int
	buffer string

}

type Task struct {
	Index int
	TaskString string
	Start time.Time 
	End time.Time
	Len int
	Hours int
	Minutes int
	Seconds int
}
// soon to be replaced by linked list
type TaskTree struct {
	Tree []Task
}

func (t *TaskTree) AddTo(x string){
	e := len(t.Tree)	
	t.Tree = append(t.Tree,Task{Index:e,TaskString:x,Start:time.Now()})
}

func (t *TaskTree) DeleteFrom(x int){
	e := len(t.Tree)
	if x <= e-1{
		t.Tree = append(t.Tree[:x],t.Tree[x+1:]...) // this part is responsible for the deletion

		for i,_ := range(t.Tree){
			t.Tree[i].Index = i
		}
	}
}

func (t *TaskTree) Edit(index int, newTask string){
	e := len(t.Tree)
	if index <= e-1{
		 t.Tree[index].TaskString = newTask
		}
}

func ( t *Task) DrawTask (x,y,padding int){

	timeStart := []string{}
	timeStart = append(timeStart,[]string{" START "," ----- "}...)
	timeStart = append(timeStart,utils.Formatez(t.Start)...)

	s := [][]string{[]string{strconv.Itoa(t.Index)},[]string{t.TaskString},timeStart}

	

	if t.End.IsZero(){

		// calculate the time difference compared to time.Now()
		temp := time.Now()
		b:= temp.Sub(t.Start)
		hours := int(b.Hours())
		minutes :=  int(b.Minutes())-60*int(b.Hours())
		seconds := int(b.Seconds())-(3600*hours+60*minutes)	

		//------------------------------------------------------



		interval := []string{}
		interval = append(interval,[]string{" INTERVAL "," ----- "}...)
		temp1 := []string{"Seconds :"+strconv.Itoa(seconds), "Minutes :"+strconv.Itoa(minutes), "Hours   :"+strconv.Itoa(hours)}
		interval = append(interval,temp1...)
		s = append(s,interval)	
	}else{
		interval := []string{}
		interval = append(interval,[]string{" INTERVAL "," ----- "}...)
		temp1 := []string{"Seconds :"+strconv.Itoa(t.Seconds), "Minutes :"+strconv.Itoa(t.Minutes), "Hours   :"+strconv.Itoa(t.Hours)}
		interval = append(interval,temp1...)
		s = append(s,interval)

		timeStop := []string{}
		timeStop = append(timeStop,[]string{" STOP "," ----- "}...)
		timeStop = append(timeStop,utils.Formatez(t.End)...)
		s = append(s,timeStop)	

	}

	endStringArray := utils.CustomFunction(s)
	draw.StringBox(endStringArray,x,y,padding )
	t.Len = len(endStringArray)+2*padding
}

func (t* TaskTree) DrawTree(start int){

	//length := len(t.tree)
	yCounter := 0

	counter := 0
	for i:=start;;i=(i+1)%len(t.Tree){

		if counter > EXPMAX-1 || counter > len(t.Tree)-1 {
			break
		}
		box_x := HSPACE + HLEN
		box_y := VSPACE + yCounter
		t.Tree[i].DrawTask(box_x,box_y,1)
		yCounter += t.Tree[i].Len
		draw.Horizontal(HSPACE,box_y+1,HLEN)
		if counter>0 {
			draw.Vertical(HSPACE,box_y+1-VLEN,t.Tree[i].Len)
		}
		counter++
	}

}

var mainTree TaskTree

func WithinRange (x int) bool {
		if x < len(mainTree.Tree){
			return true	
		}
		return false
	}


func act(xs []string, destination string) {

	switch xs[0] {
		case "ADD": 
			mainTree.AddTo(strings.Join(xs[1:]," "))
		case "DELETE":
			a ,_:= strconv.Atoi(xs[1])
			mainTree.DeleteFrom(a)
		case "EDIT":
			a, _:=strconv.Atoi(xs[1])
			if !WithinRange(a){
				break
			}
			mainTree.Edit(a,xs[2])
		case "STOP":
			a, _:=strconv.Atoi(xs[1])
			if !WithinRange(a){
				break
			}
			mainTree.Tree[a].End = time.Now()
			b:= mainTree.Tree[a].End.Sub(mainTree.Tree[a].Start)
			mainTree.Tree[a].Hours = int(b.Hours())
			mainTree.Tree[a].Minutes =  int(b.Minutes())-60*int(b.Hours())
			d := mainTree.Tree[a]
			mainTree.Tree[a].Seconds = int(b.Seconds())-(3600*d.Hours+60*d.Minutes)

		case "RESTART":
			a, _:=strconv.Atoi(xs[1])
			if !WithinRange(a){
				break
			}
			var b time.Time
			mainTree.Tree[a].Start = time.Now()
			mainTree.Tree[a].End = b

	}
		file.Put(destination,mainTree)
}

// log file functions and associated variables



func main(){

	errit := termbox.Init()
	defer termbox.Close()
	screenW,screenH := termbox.Size()

	if errit != nil {
		panic(" Trouble somewhere")
	} else {
		
		file.Check(DIRECTORY,FILE)

		startIndex := 0

		x,_ := user.Current()
		destination := x.HomeDir+"/"+DIRECTORY+"/"+FILE
		file.Get(destination,&mainTree)

		commandbox := draw.BOX{Xstart:HSPACE,Ystart:screenH-VSPACE-3,Width:screenW-HSPACE-1,Height:3}
		ogg := Insert{x:HSPACE+3+len("COMMAND :"),y:screenH-VSPACE-2,buffer:""}
	
		// various channels
		
		hideUnhide := make(chan int,1)


		//----------------- dance of functions starts here--------------------------------------------


		go func(){
			for {
				//termbox.Clear(draw.FG,draw.BG)
				atb := time.Now()
				toDisplay := atb.Format("Jan 2 15:04:05pm ")
				draw.String(toDisplay,screenW-len(toDisplay),VSPACE-2)
				termbox.Flush()
				time.Sleep(time.Second)
			}
		}()

		go func(){

			for {
				select {

				case <- hideUnhide:
					<-hideUnhide // blocking action
				default:
					draw.String("COMMAND : ",HSPACE+2,screenH-VSPACE-2)
					draw.String(" [ Press i to enter commands ][ Press h for help ][ Press z to quit ] ",HSPACE+1,screenH-1)
					mainTree.DrawTree(startIndex)
					commandbox.Box()
					termbox.Flush()
					time.Sleep(100*time.Millisecond)
			
				}
			}
		}()

		for  {
			termbox.Flush()
			g := termbox.PollEvent()
			
			if g.Ch == 'z'{
				break
			} else if g.Ch == 'h' {
				termbox.Clear(draw.FG,draw.BG) // clears the screen to display options
				hideUnhide <- 3 // the value here doesn't matter

				draw.String(" [ Press any button to exit ] ",HSPACE+1,screenH-1)
				draw.String("  ( ADD taskString) || ( DELETE INDEX) || ( EDIT INDEX taskString  )||(STOP INDEX) || ( RESTART INDEX )",HSPACE,int(screenH/2))
				termbox.Flush()
				termbox.PollEvent()

				hideUnhide <-3

			}else if g.Ch == 'i'{
				for {

					g :=  termbox.PollEvent()
					if  g.Key == termbox.KeyEnter {
						a:= parse.Tokenize(ogg.buffer)
						ogg.buffer = ""
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
			}else if g.Key == termbox.KeyArrowUp {
				if  startIndex > 0 {
					startIndex= (startIndex-1)%len(mainTree.Tree)
				} else {
					startIndex = len(mainTree.Tree)-1
				}
			}else if g.Key == termbox.KeyArrowDown {
				startIndex= (startIndex+1)%len(mainTree.Tree)
			}

			termbox.Clear(draw.FG,draw.BG)
		}
	}

}


