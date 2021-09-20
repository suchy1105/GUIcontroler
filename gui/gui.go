package gui

import (
	"GUIcontroler/api"
	"fmt"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	//"gioui.org/op/clip"
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
	var menulevel  		= 	2
	var start time.Time
	var begin  			= 	true
	const size 			=	50
	//const sizeXNTIMin   =   0
	//const sizeYNTIMin   =   0
	const sizeXNTI 		=	100
	const sizeYNTI 		=	100
	//const sizeXVoiceMin	=	0
	//const sizeYVoiceMin	=	0
	//const sizeXVoiceMax	=	60
	//const sizeYVoiceMax	=	380
	counter 			:=	0
	var c  				=	true

	var ops op.Ops
	//var ops1 op.Ops
	var float          		    = new(widget.Float)
	var floatsub1               = new(widget.Float)
	var floatsub2               = new(widget.Float)
	var floatsub3               = new(widget.Float)
	var floatsub4               = new(widget.Float)
	volM,_:=strconv.ParseFloat(state.AlsaVolumeM,32)
	float.Value=float32(volM)
	vol1,_:=strconv.ParseFloat(state.AlsaVolume1,32)
	floatsub1.Value=float32(vol1)
	vol2,_:=strconv.ParseFloat(state.AlsaVolume2,32)
	floatsub2.Value=float32(vol2)
	vol3,_:=strconv.ParseFloat(state.AlsaVolume3,32)
	floatsub3.Value=float32(vol3)
	vol4,_:=strconv.ParseFloat(state.AlsaVolume4,32)
	floatsub4.Value=float32(vol4)
	// UI state.
	var btn widget.Clickable
	var btnMuteMaster widget.Clickable
	var btnMute1 widget.Clickable
	var btnMute2 widget.Clickable
	var btnMute3 widget.Clickable
	var btnMute4 widget.Clickable
	var btnNTI widget.Clickable
	var btnVoice widget.Clickable
	var count int
	var changes = false
	for e := range w.Events() {
		if e, ok := e.(system.FrameEvent); ok {
			gtx := layout.NewContext(&ops, e)
			//gtx1 := layout.NewContext(&ops1,e)
			for btn.Clicked() {
				count++
				state.ConnState=!state.ConnState
				fmt.Println(count)
			}
			for btnMuteMaster.Clicked(){
				fmt.Println("klikniety Master",float.Value)
				float.Value=0
			}
			for btnMute1.Clicked(){
				fmt.Println("klikniety Master1",floatsub1.Value)
				floatsub1.Value=0
			}
			for btnMute2.Clicked(){
				fmt.Println("klikniety Master2",floatsub2.Value)
				floatsub2.Value=0
			}
			for btnMute3.Clicked(){
				fmt.Println("klikniety Master3",floatsub3.Value)
				floatsub3.Value=0
			}
			for btnMute4.Clicked(){
				fmt.Println("klikniety Maste4r", floatsub4.Value)
				floatsub4.Value=0
			}
			for btnNTI.Clicked(){
				fmt.Println("klikniety NTI", floatsub4.Value)
				state.PlayStipa=!state.PlayStipa
				changes=true
			}
			for btnVoice.Clicked(){
				fmt.Println("klikniety Voice", floatsub4.Value)
				state.PlayVoice=!state.PlayVoice
				changes=true
			}
			// Process events using the key, c.
			 for _, e := range gtx.Events(c)   {
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
							//menulevel++
						}

						c = !c
						counter++
					}
					//if e, ok = e.(pointer.Event); ok {
					//	if e.Type == pointer.ButtonPrimary {

					//		fmt.Println("levebutttonl up")

					//	}
					//}
					fmt.Println(floatsub4.Value, "float syb 4 event CE")
				}
				fmt.Println(floatsub4.Value, "float syb 4")


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
						return layout.Flex{Alignment: layout.End}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Body1(th, fmt.Sprintf(" 																																																				" )).Layout,
								)
							}),

							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Button(th,&btnNTI,"NTIStipa").Layout,

								)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.UniformInset(unit.Dp(4)).Layout(gtx,
									material.Button(th,&btnVoice,"Voice").Layout,
								)
							}),


						)

					}),

					//layout.Flexed(100, material.Button(th, &btnMuteMaster, "play").Layout),
					//layout.Flexed(10, material.Slider(th, floatsub4, 0, 100).Layout),
					/*layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(40)).Layout(gtx,
							material.Button(th,&btnMuteMaster,"Play STIPA").Layout,
						)
					}),*/

					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						op.Offset(f32.Pt(0, 90)).Add(&ops)
						return material.H5(th, fmt.Sprintf("IP: %s                              MAC: %s", state.Ip, state.Mac)).Layout(gtx)
					}),

				)
			})

/*			pointer.Rect(image.Rectangle{//Min: image.Point{
			//	X: sizeXNTIMin,
			//	Y: sizeYNTIMin,
			//},
			Max: image.Point{

				X: sizeXNTI,
				Y: sizeYNTI,
			}}).Add(gtx.Ops)
						// Declare the filter with the key, c.
			pointer.InputOp{Tag: cd, Types: pointer.Press}.Add(gtx.Ops)
*/
			//col := color.NRGBA{A: 0xff, G: 0xff} // Green.
			// Draw checkbox. Red for unchecked, green for checked.

			//paint.PaintOp{}.Add(gtx.Ops)

			/*pointer.Rect(image.Rectangle{
				Min: image.Point{
					X: sizeXVoiceMin,
					Y: sizeYVoiceMin,
				},
				Max: image.Point{
					X: sizeXVoiceMax,
					Y: sizeYVoiceMax,
			}}).Add(gtx.Ops)*/

			//paint.ColorOp{Color: col}.Add(gtx.Ops)
/*			clip.Rect{
				/*Min: image.Point{
					X: sizeXNTIMin,
					Y: sizeYNTIMin,
				},
				Max: image.Point{

					X: sizeXNTI,
					Y: sizeYNTI,
				},
			}.Add(gtx.Ops)
			paint.PaintOp{}.Add(gtx.Ops)*/
			/*clip.Rect{
				Min: image.Point{
					X: sizeXVoiceMin,
					Y: sizeYVoiceMin,
				},
				Max: image.Point{
					X: sizeXVoiceMax,
					Y: sizeYVoiceMax,
				},
			}.Add(gtx.Ops)*/
		//paint.PaintOp{}.Add(gtx.Ops)
			// Declare the filter with the key, c.
		//	pointer.InputOp{Tag: ce, Types: pointer.Press}.Add(gtx.Ops)

			switch  {
			case int(volM)!=int(float.Value):
				volM=float64(float.Value)
				fmt.Println("uneq1",volM)
				changes=true
			case int(vol1)!=int(floatsub1.Value):
				vol1=float64(floatsub1.Value)
				fmt.Println("uneq2",vol1)
				changes=true
			case int(vol2)!=int(floatsub2.Value):
				vol2=float64(floatsub2.Value)
				fmt.Println("uneq3", vol2)
				changes=true
			case int(vol3)!=int(floatsub3.Value):
				vol3=float64(floatsub3.Value)
				fmt.Println("uneq4", vol3)
				changes=true
			case int(vol4)!=int(floatsub4.Value):
				vol4=float64(floatsub4.Value)
				fmt.Println("uneq5",vol4)
				changes=true
			}
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
		//	pointer.InputOp{Tag: c, Types: pointer.Press}.Add(gtx.Ops)
		}
		if changes==true{


			state.AlsaVolumeM=strconv.Itoa(int(volM))
			state.AlsaVolume1=strconv.Itoa(int(vol1))
			state.AlsaVolume2=strconv.Itoa(int(vol2))
			state.AlsaVolume3=strconv.Itoa(int(vol3))
			state.AlsaVolume4=strconv.Itoa(int(vol4))
			push(state)
			changes=false
		}
	}
}

func drawImage(ops *op.Ops, img image.Image) {
	imageOp := paint.NewImageOp(img)
	imageOp.Add(ops)
	op.Affine(f32.Affine2D{}.Scale(f32.Pt(100, 100), f32.Pt(400, 400)))
	paint.PaintOp{}.Add(ops)
}
func push(s *api.GuiState){
  fmt.Println("Pushed data: ",s)
}

