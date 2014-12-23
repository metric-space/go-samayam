package drawUtils

import "github.com/nsf/termbox-go"

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
)

type BOX struct {

	Width,Height int	
	X_start,Y_start int

}


func String(fill string,x,y int){
	for i,j := range fill {
			termbox.SetCell(x+i,y,j,FG,BG)
		}

}


func StringBox(fill []string,x,y,padding int){


	actual_width := len(fill[0]) + 2*padding
	actual_height :=  len(fill) + 2*padding
	actual_start_x := x + padding
	actual_start_y := y + padding

	for i,j := range(fill){
		String(j,actual_start_x,actual_start_y+i)	
	}
	rect := BOX{ X_start:x,Y_start:y,Width:actual_width,Height:actual_height}
	rect.Box()
}

func Vertical(x,y,h int){

	for i := y ; i <= y+h-1 ; i++ {
		termbox.SetCell(x,i,v,FG,BG)
	}
}

func Horizontal(x,y,w int){

	for i := x ; i <= x+w-1 ; i++ {
		termbox.SetCell(i,y,h,FG,BG)
	}
}


func (r *BOX) Box(){

	termbox.SetCell(r.X_start,r.Y_start,box_ul,FG,BG)
	Horizontal(r.X_start+1,r.Y_start,r.Width-2)
	termbox.SetCell(r.X_start+r.Width-1,r.Y_start,box_ur,FG,BG)

	termbox.SetCell(r.X_start,r.Y_start+r.Height-1,box_ll,FG,BG)
	Horizontal(r.X_start+1,r.Y_start+r.Height-1,r.Width-2)
	termbox.SetCell(r.X_start+r.Width-1,r.Y_start+r.Height-1,box_lr,FG,BG)

	Vertical(r.X_start,r.Y_start+1,r.Height-2)
	Vertical(r.X_start+r.Width-1,r.Y_start+1,r.Height-2)

}

