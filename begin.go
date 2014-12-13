package main

import ("github.com/nsf/termbox-go";"time";"strings")


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

	task string
	start time.Time 
	end time.Time
}

func draw_string(fill string,x,y int){
	
	for i,j := range fill {
			termbox.SetCell(x+i,y,j,FG,BG)
		}


}


func draw_string_box(fill string,x,y,padding int){


	actual_width := len(fill) + 2*padding
	actual_height :=  1 + 2*padding
	actual_start_x := x + padding
	actual_start_y := y + padding

	draw_string(fill,actual_start_x,actual_start_y)

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
	s := []string{t.task,t.start.String()}
	string_to_draw :=  strings.Join(s, "    " )
	draw_string_box(string_to_draw,x,y,padding )

}

func draw_tree(tree []TASK){

	//length := len(t.tree)
	for i,j :=range(tree){
		box_x := H_SPACE + H_LEN
		box_y := V_SPACE + i*V_LEN
		j.draw_task(box_x,box_y,1)
		draw_horizontal(H_SPACE,box_y+1,H_LEN)
		if i>=1 {
			draw_vertical(H_SPACE,box_y+1-V_LEN,V_LEN)
		}
	}

}

func main(){

	errit := termbox.Init()
	screen_w,screen_h := termbox.Size()


	if errit != nil {
		panic(" Trouble somewhere")
	} else {
		//x := BOX{width:20,x_start:3,y_start:3,height:20}
		//x.draw_mainbox()
		for  {
			task_1 := []TASK{ TASK{task:" take meeko for a walk",start:time.Now()}, TASK{task:" finish samayam",start:time.Now()}, TASK{task:" finish python",start:time.Now()}}
			draw_tree(task_1)

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
					if g.Key == termbox.KeyEnter {
						break	
					}else if g.Key == termbox.KeySpace{
						break
					}else{
						ogg.buffer+=string(g.Ch)
						draw_string(ogg.buffer,ogg.x,ogg.y)			
						termbox.Flush()
					}
				}
			}
			termbox.Clear(FG,BG)
		}
	}
	defer termbox.Close()

}


