package gui

import (
	"GUIcontroler/api"
	"fmt"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/widget/material"
	"image/jpeg"
	"image/png"
	"os"
	"time"

	//"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"image"
	//"image"
	//"image/draw"

	"gioui.org/unit"
)

//GUI  dasa START OMIT
func GUI(state *api.GuiState) {
	//m := Image{}
	//pic.ShowImage(m)

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(480)))
		loop(w, state)
	}()
	app.Main()
}

func loop(w *app.Window, state *api.GuiState) {
	th := material.NewTheme(gofont.Collection())
	timeSpace:=0*time.Second
	menulevel:= 1
	var start time.Time
	var begin bool = true
	const size = 50
	counter :=0
	var c bool = true
	var ops op.Ops
	var float               = new(widget.Float)
	// UI state.
	var btn widget.Clickable
	var count int

	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)

			for btn.Clicked() {
				count++
				fmt.Println(count)
			}
			// Process events using the key, c.
			for _, e := range gtx.Events(c) {
				if e, ok := e.(pointer.Event); ok {
					if e.Type == pointer.Press {

						if begin == false {
							start = time.Now()
							begin = true
						}
						timeSpace = (time.Now().Sub(start))
						fmt.Println(timeSpace)

						if timeSpace > 3*time.Second {
							counter = 0
							begin = false
							//timeSpace = 0
						}
						if counter > 5 && timeSpace < 3*time.Second {
							fmt.Println("level up")
							menulevel++
						}
						c = !c
						counter++
					}
				}
			}
		if menulevel==1{	// Confine input to the area covered by the checkbox.
			image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
			dir, _ := os.Open("./GUIsat.png")

			var obrazownia, _, _ = image.Decode(dir)

			drawImage(&ops, obrazownia)
			pointer.Rect(image.Rectangle{Max: image.Point{
				X: size,
				Y: size,
			}}).Add(gtx.Ops)
			// Declare the filter with the key, c.
			pointer.InputOp{Tag: c, Types: pointer.Press}.Add(gtx.Ops)
			fmt.Println(counter)
		}
			/*	layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return
			}*/
		if menulevel==2{

		//	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
			if state.ConnState!=true {
				dirlinkred, _ := os.Open("./linkRed.png")
				var obrazlinkRed, _, _ = image.Decode(dirlinkred)
				drawImage(&ops, obrazlinkRed)
			}else{
				dirlinkGreen, _ := os.Open("./linkGreen.png")
				var obrazlinkGreen, _, _ = image.Decode(dirlinkGreen)
				drawImage(&ops, obrazlinkGreen)
			}
			image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)

			dirNTI, _ := os.Open("./NTI.jpg")
			var obrazNTI, _, _ = image.Decode(dirNTI)

			op.Offset(f32.Pt(660, 0)).Add(&ops) //DO USTAWIENIA KURSORA


			drawImage(&ops, obrazNTI)
			//op.Offset(f32.Pt(0, 0)).Add(&ops)
			dirNTIn, _ := os.Open("./NTIn.jpg")
			var obrazNTIn, _, _ = image.Decode(dirNTIn)

			drawImage(&ops, obrazNTIn)
			op.Offset(f32.Pt(-650, 0)).Add(&ops)

			layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H5(th, fmt.Sprintf("IP: %s", state.Ip)).Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H5(th, fmt.Sprintf("MAC: %s", state.Mac)).Layout(gtx)
					}),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(th, &btn, "Click me!").Layout(gtx)
					}),


					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, material.Slider(th, float, 0, 100).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(8)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("%.0f", float.Value)).Layout,
								)
							}),
						)
					}),
				)
			})
		}

			e.Frame(gtx.Ops)
		}
	}
}

func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(100, 100), f32.Pt(400, 400)))
	paint.PaintOp{}.Add(ops)
}
