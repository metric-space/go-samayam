package drawUtils

import "github.com/nsf/termbox-go"

// unicode constants
const (
	V = '\u2502'
	H = '\u2500'
	BOXUL = '\u256d'
	BOXUR = '\u256e'
	BOXLL = '\u2570'
	BOXLR = '\u256f'
	FG = termbox.ColorWhite
	BG = termbox.ColorBlack
)

type BOX struct {

	Width,Height int	
	Xstart,Ystart int

}


func String(fill string,x,y int){
	for i,j := range fill {
			termbox.SetCell(x+i,y,j,FG,BG)
		}

}


func StringBox(fill []string,x,y,padding int){


	actualWidth := len(fill[0]) + 2*padding
	actualHeight :=  len(fill) + 2*padding
	actualStartX := x + padding
	actualStartY := y + padding

	for i,j := range(fill){
		String(j,actualStartX,actualStartY+i)	
	}
	rect := BOX{ Xstart:x,Ystart:y,Width:actualWidth,Height:actualHeight}
	rect.Box()
}

func Vertical(x,y,h int){

	for i := y ; i <= y+h-1 ; i++ {
		termbox.SetCell(x,i,V,FG,BG)
	}
}

func Horizontal(x,y,w int){

	for i := x ; i <= x+w-1 ; i++ {
		termbox.SetCell(i,y,H,FG,BG)
	}
}


func (r *BOX) Box(){

	termbox.SetCell(r.Xstart,r.Ystart,BOXUL,FG,BG)
	Horizontal(r.Xstart+1,r.Ystart,r.Width-2)
	termbox.SetCell(r.Xstart+r.Width-1,r.Ystart,BOXUR,FG,BG)

	termbox.SetCell(r.Xstart,r.Ystart+r.Height-1,BOXLL,FG,BG)
	Horizontal(r.Xstart+1,r.Ystart+r.Height-1,r.Width-2)
	termbox.SetCell(r.Xstart+r.Width-1,r.Ystart+r.Height-1,BOXLR,FG,BG)

	Vertical(r.Xstart,r.Ystart+1,r.Height-2)
	Vertical(r.Xstart+r.Width-1,r.Ystart+1,r.Height-2)

}

