package GUI


import (
	"bytes"
	"flag"
	"fmt"
	"gioui.org/app/headless"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/op"
	"gioui.org/op/paint"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"net"
	"os"
	"os/exec"
	"time"
	"unicode/utf8"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

)

type scaledConfig struct {
	Scale float32
}

type iconAndTextButton struct {
	theme *material.Theme
}

func NewWindow() {

	flag.Parse()
	exec.Command("amixer", "-q", "sset", "Master", "77%").Start()
	gofont.Register()


	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(480)))
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func saveScreenshot(f string) error {
	const scale = 1.5
	sz := image.Point{X: 800 * scale, Y: 600 * scale}
	w, err := headless.NewWindow(sz.X, sz.Y)
	if err != nil {
		return err
	}
	gtx := new(layout.Context)
	gtx.Reset(nil, &scaledConfig{scale}, sz)
	th := material.NewTheme()
	drawing(gtx, th)
	w.Frame(gtx.Ops)
	img, err := w.Screenshot()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}
	return ioutil.WriteFile(f, buf.Bytes(), 0666)
}

func loop(w *app.Window) error {
	th := material.NewTheme()

	gtx := new(layout.Context)
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx.Reset(e.Queue, e.Config, e.Size)
				drawing(gtx, th)
				e.Frame(gtx.Ops)
			}
			w.Invalidate()
		}
	}
}

var (

	flagev            = true
	relased           = true
	buttonEsc1            = new(widget.Clickable)
	buttonEsc2            = new(widget.Clickable)
	buttonTest            = new(widget.Clickable)
	buttonMowa          = new(widget.Clickable)
	button1            = new(widget.Clickable)
	button2            = new(widget.Clickable)
	button3            = new(widget.Clickable)
	button4            = new(widget.Clickable)
	button5            = new(widget.Clickable)
	button6            = new(widget.Clickable)
	button7            = new(widget.Clickable)
	button8            = new(widget.Clickable)
	button9            = new(widget.Clickable)
	button0            = new(widget.Clickable)
	buttonEnter       = new(widget.Clickable)
	buttonDel          = new(widget.Clickable)
	button            = new(widget.Clickable)
	buttonglosniej    = new(widget.Clickable)
	buttonciszej      = new(widget.Clickable)
	greenButton       = new(widget.Clickable)
	// iconTextButton    = new(widget.Clickable)
	iconButton        = new(widget.Clickable)
	iconButton2       = new(widget.Clickable)
	flatBtn           = new(widget.Clickable)
	radioButtonsGroup = new(widget.Enum)
	list              = &layout.List{
		Axis: layout.Vertical,
	}
	clickedTest         = false
	clickedMowa         = false
	volume float32   = 0.77 //0-100%
	icon                *widget.Icon
	IP                 string// = "000.000.000.000"
	version             = "0.00.00"
	temp                = "25x"
	humidity            = "44x"
	delaytime           =   3*time.Second
	timepress           time.Time
	menuLevel           = 3
	fontsize float32    = 50.0
	CodeLabel           =""
	Code                ="1234"
	wrongpassword       =true
	textBTest           =  "TEST"
	textBMowa           = "Mowa"

)

func (b iconAndTextButton) Layout(gtx *layout.Context, button *widget.Clickable, icon *widget.Icon, word string) {
	material.ButtonLayout(b.theme).Layout(gtx, button, func() {

	})
}

func drawing(gtx *layout.Context, th *material.Theme) {
	if menuLevel==1{
		imgop := getImage("/GUI/pic/satlong.png")
		imgconnok := getImage("/GUI/pic/connok.png")
		imgconnbad := getImage("/GUI/pic/connbad.png")

		widgets := []func(){

			func() {
				widget.Image{Src: imgop, Scale: 1}.Layout(gtx)
			},
			func() {

				for _, e := range gtx.Events(gtx) {
					if e, ok := e.(pointer.Event); ok {
						// println("karol1")

						if e.Type == pointer.Press && relased == true {
							// fmt.Println("get time:",time.Now())
							timepress = time.Now()
							relased = false
						}
						if e.Type == pointer.Release {
							timerelease := time.Now()

							relased = true

							if timerelease.Sub(timepress) > delaytime {
								fmt.Println("3 sekundy")
								menuLevel=2

							}
						}

					}
				}
				var st op.StackOp
				st.Push(gtx.Ops) // Save operation state.
				pointer.Rect(image.Rect(0, 120, 120, 240)).Add(gtx.Ops)
				pointer.InputOp{Key: gtx}.Add(gtx.Ops)
				st.Pop()
				server:="10.10.8.254:80"

				timeout := 1 * time.Second
				_, err := net.DialTimeout("tcp",server, timeout)
				if err != nil {
					widget.Image{Src: imgconnbad, Scale: 1}.Layout(gtx)
				}else {
					widget.Image{Src: imgconnok, Scale: 1}.Layout(gtx)
				}


			},

			func() {
				material.Label(th, th.TextSize.Scale(20/10), "Temp: "+temp+"                Hum: "+humidity+"                 volume: "+fmt.Sprintf("%.0f", volume*100)+"%").Layout(gtx)
			},
			func() {
				material.Label(th, th.TextSize.Scale(20/10),  "MAC:"+getMAC("eno1")+" IP:"+getIP()[2]+" v."+version).Layout(gtx)
			},
		}

		list.Layout(gtx, len(widgets), func(i int) {
			layout.UniformInset(unit.Dp(0)).Layout(gtx, widgets[i]) //vertical spaces
		})
		paint.ColorOp{Color: color.RGBA{A: 0x40}}.Add(gtx.Ops)
		triangle(650, 350, 750, 150, gtx, 1, false)
		triangle(650, 350, 750, 150, gtx, volume, true)
	}
	if menuLevel==2 {

		widgetsLevel2 := []func(){
			func() {
				material.Label(th, th.TextSize.Scale(20/10), CodeLabel).Layout(gtx)
			},
			func() {
				in := layout.UniformInset(unit.Dp(5))//landscape spaces
				layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for buttonDel.Clicked(gtx) {
								fmt.Println("DEL")
								fmt.Println(CodeLabel)
								CodeLabel=trimLastChar(CodeLabel)
								fmt.Println(CodeLabel)
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
							}
							bdel:=material.Button(th, "DEL!")
							bdel.TextSize = unit.Value{   V: fontsize, U: 1, }
							bdel.Layout(gtx, buttonDel)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for buttonEnter.Clicked(gtx) {
								fmt.Println("Enter")
								if(CodeLabel==Code){
									fmt.Println("Sukccess")
									menuLevel=3
								}else{    CodeLabel="Wrong Password"
									wrongpassword=true

								}
							}
							bEnter:=material.Button(th, "Enter!")
							bEnter.TextSize = unit.Value{   V: fontsize, U: 1, }
							bEnter.Layout(gtx, buttonEnter)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for buttonEsc1.Clicked(gtx) {
								CodeLabel=""
								menuLevel=1
							}
							//material.Button(th, "Esc").Layout(gtx, buttonEsc)
							bEsc:=material.Button(th, "Esc!")
							bEsc.TextSize = unit.Value{   V: fontsize, U: 1, }
							bEsc.Layout(gtx, buttonEsc1)
						})
					}),


				)
			},/**/
			func() {
				in := layout.UniformInset(unit.Dp(5))//landscape spaces
				layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button1.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"1"

							}
							//material.Button(th, "1").Layout(gtx, button1)
							b1:=material.Button(th, "1")
							b1.TextSize = unit.Value{   V: fontsize, U: 1, }
							b1.Layout(gtx, button1)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button2.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"2"
							}

							//material.Button(th, "2").Layout(gtx, button2)
							b2:=material.Button(th, "2")
							b2.TextSize = unit.Value{   V: fontsize, U: 1, }
							b2.Layout(gtx, button2)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button3.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"3"
							}
							//material.Button(th, "3").Layout(gtx, button3)
							b3:=material.Button(th, "3")
							b3.TextSize = unit.Value{   V: fontsize, U: 1, }
							b3.Layout(gtx, button3)
						})
					}),


				)
			},/**/
			func() {
				in := layout.UniformInset(unit.Dp(5))//landscape spaces
				layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button4.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"4"
							}
							//material.Button(th, "4").Layout(gtx, button4)
							b4:=material.Button(th, "4")
							b4.TextSize = unit.Value{   V: fontsize, U: 1, }
							b4.Layout(gtx, button4)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							//material.Button(th, "5").Layout(gtx, button5)
							for button5.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"5"
							}
							b5:=material.Button(th, "5")
							b5.TextSize = unit.Value{   V: fontsize, U: 1, }
							b5.Layout(gtx, button5)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button6.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"6"
							}
							//material.Button(th, "6").Layout(gtx, button6)
							bEnter:=material.Button(th, "6")
							bEnter.TextSize = unit.Value{   V: fontsize, U: 1, }
							bEnter.Layout(gtx, button6)
						})
					}),


				)
			},/**/
			func() {
				in := layout.UniformInset(unit.Dp(5))//landscape spaces
				layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button7.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"7"
							}
							//material.Button(th, "7").Layout(gtx, button7)
							b7:=material.Button(th, "7")
							b7.TextSize = unit.Value{   V: fontsize, U: 1, }
							b7.Layout(gtx, button7)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for button8.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"8"
							}
							//material.Button(th, "8").Layout(gtx, button8)
							b8:=material.Button(th, "8")
							b8.TextSize = unit.Value{   V: fontsize, U: 1, }
							b8.Layout(gtx, button8)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							//material.Button(th, "9").Layout(gtx, button9)
							for button9.Clicked(gtx) {
								if wrongpassword==true {
									CodeLabel=""
									wrongpassword=false
								}
								CodeLabel=CodeLabel+"9"
							}
							b9:=material.Button(th, "9")
							b9.TextSize = unit.Value{   V: fontsize, U: 1, }
							b9.Layout(gtx, button9)
						})
					}),


				)
			},/**/

		}
		list.Layout(gtx, len(widgetsLevel2), func(i int) {
			layout.UniformInset(unit.Dp(5)).Layout(gtx, widgetsLevel2[i]) //vertical spaces

		})



	}
	if menuLevel==3{
		widgetsLevel3 := []func(){
			func() {
				material.Label(th, th.TextSize.Scale(20/10), "Service menu").Layout(gtx)
			},
			func() {
				in := layout.UniformInset(unit.Dp(5))//landscape spaces
				layout.Flex{Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for buttonEsc2.Clicked(gtx) {
								CodeLabel=""
								menuLevel=1
							}
							//material.Button(th, "Esc").Layout(gtx, buttonEsc)
							bEsc2:=material.Button(th, "Esc!")
							bEsc2.TextSize = unit.Value{   V: fontsize, U: 1, }
							bEsc2.Layout(gtx, buttonEsc2)
						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {

							//

							for buttonTest.Clicked(gtx) {


								clickedTest= !clickedTest
								if clickedTest==true {
									//clickedTest=false
									volume=0.5
									fmt.Println("Test")
									exec.Command("amixer", "-q", "sset", "Master", "50%").Start()
									exec.Command("aplay","-i", "audio/NTISTIPA.wav").Start()
									textBTest="STOP"
								}else{
									//clickedTest=true
									//exec.Command("^","z").Start()
									exec.Command("pkill","aplay").Start()
									textBTest="TEST"
								}

							}
							bTest := material.Button(th, textBTest)
							if clickedTest==true {
								bTest.Background = color.RGBA{A: 0xff, R: 0xFF, G: 0x00, B: 0x00}
							}
							bTest.TextSize = unit.Value{V: fontsize, U: 1,}
							bTest.Layout(gtx, buttonTest)

						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {

							//

							for buttonMowa.Clicked(gtx) {


								clickedMowa= !clickedMowa
								if clickedMowa==true {
									//clickedTest=false
									volume=0.5
									fmt.Println("Mowa")
									exec.Command("amixer", "-q", "sset", "Master", "50%").Start()
									exec.Command("aplay","-i", "audio/testmowy.wav").Start()
									textBMowa="STOP"
								}else{
									//clickedTest=true
									//exec.Command("^","z").Start()
									exec.Command("pkill","aplay").Start()
									textBMowa="Mowa"

								}

							}
							bMowa := material.Button(th, textBMowa)
							if clickedMowa==true {
								bMowa.Background = color.RGBA{A: 0xff, R: 0xFF, G: 0x00, B: 0x00}
							}

							bMowa.TextSize = unit.Value{V: fontsize, U: 1,}
							bMowa.Layout(gtx, buttonMowa)

						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for buttonglosniej.Clicked(gtx) {


								fmt.Println("+")
								if volume<=0.95{
									exec.Command("amixer", "-q", "sset", "Master", "5%+").Start()
									volume+=0.05
								}else{volume=1
									exec.Command("amixer", "-q", "sset", "Master", "100%").Start()}
								fmt.Println(volume)
							}
							bglosniej:=material.Button(th, "+")
							bglosniej.TextSize = unit.Value{   V: fontsize, U: 1, }
							bglosniej.Layout(gtx, buttonglosniej)


						})
					}),
					layout.Rigid(func() {
						in.Layout(gtx, func() {
							for buttonciszej.Clicked(gtx) {

								fmt.Println("-")
								if volume>=0.05{
									volume-=0.05
									exec.Command("amixer", "-q", "sset", "Master", "5%-").Start()
									// out, _ := exec.Command("amixer", "sget", "Master").Output()
									//fmt.Printf("%c",out[120])
								}else{volume=0
									exec.Command("amixer", "-q", "sset", "Master", "0%").Start()}
								fmt.Println(volume)}
							bciszej:=material.Button(th, "-")
							bciszej.TextSize = unit.Value{   V: fontsize, U: 1, }
							bciszej.Layout(gtx, buttonciszej)


						})
					}),

				)
			},/**/
			func() {
				material.Label(th, th.TextSize.Scale(20/10), "Temp: "+temp+"                Hum: "+humidity+"                volume: "+fmt.Sprintf("%.0f", volume*100)+"%").Layout(gtx)
			},
		}

		paint.ColorOp{Color: color.RGBA{A: 0x40}}.Add(gtx.Ops)
		triangle(650, 350, 750, 150, gtx, 1, false)
		//if volume > 1 {
		//    volume = 0
		//}
		//volume += 0.005
		triangle(650, 350, 750, 150, gtx, volume, true)
		list.Layout(gtx, len(widgetsLevel3), func(i int) {
			layout.UniformInset(unit.Dp(5)).Layout(gtx, widgetsLevel3[i]) //vertical spaces

		})

	}
}

func (s *scaledConfig) Now() time.Time {
	return time.Now()
}

func (s *scaledConfig) Px(v unit.Value) int {
	scale := s.Scale
	if v.U == unit.UnitPx {
		scale = 1
	}
	return int(math.Round(float64(scale * v.V)))
}

func getImage(filePath string) (paint.ImageOp) {

	im,_:=getImg(filePath)
	img:=paint.NewImageOp(im)
	//paint.NewImageOp(img)
	return img
}
func getImg(filePath string) (image.Image, error){
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(f)

	return image, err
}
func triangle(x0 float32, y0 float32, x1 float32, y1 float32, gtx* layout.Context,fillfactor float32, toColor bool){
	deltaX:= x1-x0// 400-100
	deltaY:= y0-y1// 400-100      //triangle(100,400,400,100,gtx)
	increment:= deltaX/deltaY// 300/300
	currentX:=x0//float32(0.1)//x1 //400
	var R uint8=0x20
	var G uint8=0xff

	//var R int8 =0x96
	// inc:=(int)increment
	for i:=0; int(deltaY*fillfactor) >i ; i++{
		//fmt.Println(y1)
		if toColor==true{
			if R==255{
			} else {
				R++}
			if G==0{
			}else {
				G--
			}
			colour(R,G,gtx)
		}
		paint.PaintOp{Rect: f32.Rectangle{
			Min: f32.Point{
				X: x0,
				Y: y0,
			},
			Max: f32.Point{
				X: currentX,//float32(i),//currentX,
				Y: y0+1,
			},
		}}.Add(gtx.Ops)
		y0--
		currentX+=increment/// ---1

	}
}
func colour (R uint8, G uint8, gtx* layout.Context ){

	paint.ColorOp{Color: color.RGBA{
		R: R,
		G: G,
		B: 50,
		A: 0xff,
	}}.Add(gtx.Ops)
}
func trimLastChar(s string) string {
	r, size := utf8.DecodeLastRuneInString(s)
	if r == utf8.RuneError && (size == 0 || size == 1) {
		size = 0
	}
	return s[:len(s)-size]
}

func getIP() []string {
	ipTable:= make([]string,0)
	//index :=0
	ifaces, err := net.Interfaces()
	// handle err

	for _, i := range ifaces {

		addrs, err := i.Addrs()
		// handle err

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipTable=append(ipTable,ip.String())
			// index++
			//   fmt.Println(ip,"    space    ")    // process IP address
		}
		if err!=nil {
			fmt.Println(err)
		}

	}

	if err!=nil {
		fmt.Println(err)
	}

	return ipTable
}
func getMAC(InterfaceName string) string {
	var MAC string
	//index :=0
	ifaces, err := net.Interfaces()
	// handle err

	for _, i := range ifaces {
		if i.Name==InterfaceName{
			MAC=i.HardwareAddr.String()
		}

	}

	if err!=nil {
		fmt.Println(err)
	}

	return MAC
}

