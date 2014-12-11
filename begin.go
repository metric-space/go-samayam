package main

import ("github.com/nsf/termbox-go";"time")


// unicode constants
const (
	v = '\u2502'
	h = '\u2500'
	box_ul = '\u256d'
	box_ur = '\u256e'
	box_ll = '\u2570'
	box_lr = '\u256f'
	name = " SAMAYAM "
)

type BOX struct {

	width,height int	
	x_start,y_start int

}

func draw_vertical(x,y,h int){

	for i := y ; i <= y+h-1 ; i++ {
		termbox.SetCell(x,i,v,termbox.ColorRed,termbox.ColorBlack)
	}
}

func draw_horizontal(x,y,w int){

	for i := x ; i <= x+w-1 ; i++ {
		termbox.SetCell(i,y,h,termbox.ColorRed,termbox.ColorBlack)
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

	//draw_vertical(r.x_start,r.y_start+1,r.height-1)
	
	termbox.SetCell(r.x_start,r.y_start,box_ul,termbox.ColorRed,termbox.ColorBlack)
	draw_horizontal(r.x_start+1,r.y_start,r.width-2)
	termbox.SetCell(r.x_start+r.width-1,r.y_start,box_ur,termbox.ColorRed,termbox.ColorBlack)

	termbox.SetCell(r.x_start,r.y_start+r.height-1,box_ll,termbox.ColorRed,termbox.ColorBlack)
	draw_horizontal(r.x_start+1,r.y_start+r.height-1,r.width-2)
	termbox.SetCell(r.x_start+r.width-1,r.y_start+r.height-1,box_lr,termbox.ColorRed,termbox.ColorBlack)

	draw_vertical(r.x_start,r.y_start+1,r.height-2)
	draw_vertical(r.x_start+r.width-1,r.y_start+1,r.height-2)

}

func main(){

	errit := termbox.Init()
	if errit != nil {
		panic(" Trouble somewhere")
	} else {
		x := BOX{width:20,x_start:3,y_start:3,height:20}
		x.draw_mainbox()
		termbox.Flush()
		time.Sleep(10*time.Second)
	}
	defer termbox.Close()

}


