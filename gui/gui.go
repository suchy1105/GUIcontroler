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
	"strconv"
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
	var menulevel int = 2
	var start time.Time
	var begin bool = true
	const size = 50
	counter :=0
	var c bool = true
	var ops op.Ops
	var float               = new(widget.Float)
	var floatsub1               = new(widget.Float)
	var floatsub2               = new(widget.Float)
	var floatsub3               = new(widget.Float)
	var floatsub4               = new(widget.Float)
	tmp,_:=strconv.ParseFloat(state.AlsaVolumeM,64)
	float.Value=float32(tmp)
	tmp,_=strconv.ParseFloat(state.AlsaVolume1,64)
	floatsub1.Value=float32(tmp)
	tmp,_=strconv.ParseFloat(state.AlsaVolume2,64)
	floatsub2.Value=float32(tmp)
	tmp,_=strconv.ParseFloat(state.AlsaVolume3,64)
	floatsub3.Value=float32(tmp)
	tmp,_=strconv.ParseFloat(state.AlsaVolume4,64)
	floatsub4.Value=float32(tmp)
	// UI state.
	var btn widget.Clickable
	var btnMuteMaster widget.Clickable
	var btnMute1 widget.Clickable
	var btnMute2 widget.Clickable
	var btnMute3 widget.Clickable
	var btnMute4 widget.Clickable
	var count int

	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)

			for btn.Clicked() {
				count++
				state.ConnState=!state.ConnState
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

			dirbackground, _ := os.Open("./steel-background.jpg")
			var background, _, _ = image.Decode(dirbackground)


				drawImage(&ops, background)

			dirlogo, _ := os.Open("./sat1.png")
			var logo, _, _ = image.Decode(dirlogo)


			drawImage(&ops, logo)

				image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
		//	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)




			//op.Offset(f32.Pt(-653, 0)).Add(&ops)
			layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,

					/*layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H5(th, fmt.Sprintf("MAC: %s", state.Mac)).Layout(gtx)
					}),*/




					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Start}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("Master" )).Layout,
								)
							}),
							layout.Flexed(100, material.Slider(th, float, 0, 100).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("%.0f", float.Value)).Layout,
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Button(th,&btnMuteMaster,"Mute").Layout,
								)
							}),
						)
					}),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Start}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("CH1   " )).Layout,
								)
							}),
							layout.Flexed(2, material.Slider(th, floatsub1, 0, 100).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("%.0f", floatsub1.Value)).Layout,
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									/*material.Body1(th, fmt.Sprintf("%.0f", floatsub4.Value)).Layout,*/ material.Button(th,&btnMute1,"Mute").Layout,
								)
							}),
						)
					}),


					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("CH2   " )).Layout,
								)
							}),
							layout.Flexed(1, material.Slider(th, floatsub2, 0, 100).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("%.0f", floatsub2.Value)).Layout,
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									/*material.Body1(th, fmt.Sprintf("%.0f", floatsub4.Value)).Layout,*/ material.Button(th,&btnMute2,"Mute").Layout,
								)
							}),
						)
					}),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("CH3   " )).Layout,
								)
							}),
							layout.Flexed(2, material.Slider(th, floatsub3, 0, 100).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("%.0f", floatsub3.Value)).Layout,
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									/*material.Body1(th, fmt.Sprintf("%.0f", floatsub4.Value)).Layout,*/ material.Button(th,&btnMute3,"Mute").Layout,
								)
							}),
						)
					}),


					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("CH4   " )).Layout,
								)
							}),

							layout.Flexed(10, material.Slider(th, floatsub4, 0, 100).Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf("%.0f", floatsub4.Value)).Layout,
								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									/*material.Body1(th, fmt.Sprintf("%.0f", floatsub4.Value)).Layout,*/ material.Button(th,&btnMute4,"Mute").Layout,
								)
							}),
						)
					}),

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						op.Offset(f32.Pt(0, 110)).Add(&ops)
						return material.H5(th, fmt.Sprintf("IP: %s                              MAC: %s", state.Ip, state.Mac)).Layout(gtx)
					}),

				)
			})
			op.Offset(f32.Pt(0, 350)).Add(&ops) //DO USTAWIENIA KURSORA
			if state.ConnState!=true {
				dirlinkred, _ := os.Open("./linkRedS.png")
				var obrazlinkRed, _, _ = image.Decode(dirlinkred)
				drawImage(&ops, obrazlinkRed)
			}else{
				dirlinkGreen, _ := os.Open("./linkGreenS.png")
				var obrazlinkGreen, _, _ = image.Decode(dirlinkGreen)
				drawImage(&ops, obrazlinkGreen)
			}
			dirNTI, _ := os.Open("./NTIs.png")
			var obrazNTI, _, _ = image.Decode(dirNTI)

			//drawImage(&ops, obrazNTI)
			//op.Offset(f32.Pt(0, 0)).Add(&ops)
			dirNTIn, _ := os.Open("./NTIns.png")
			var obrazNTIn, _, _ = image.Decode(dirNTIn)
			op.Offset(f32.Pt(100, 0)).Add(&ops) //DO USTAWIENIA KURSORA
			if state.PlayStipa == true {
				drawImage(&ops, obrazNTIn)

			}else {
				drawImage(&ops, obrazNTI)

			}

			dirSpeach, _ := os.Open("./micon.png")
			var obrazSpeach, _, _ = image.Decode(dirSpeach)

			//drawImage(&ops, obrazNTI)
			//op.Offset(f32.Pt(0, 0)).Add(&ops)
			dirSpeachoff, _ := os.Open("./micoff.png")
			var ObrazSpeachoff, _, _ = image.Decode(dirSpeachoff)
			op.Offset(f32.Pt(200, 0)).Add(&ops) //DO USTAWIENIA KURSORA
			if state.PlayVoice == false {
				drawImage(&ops, obrazSpeach)

			}else {
				drawImage(&ops, ObrazSpeachoff)

			}


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

