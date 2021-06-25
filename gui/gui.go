package gui

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"

	//"image"
	//"image/draw"

	"gioui.org/unit"
)

//GUI START OMIT
func GUI() {
	//m := Image{}
	//pic.ShowImage(m)

	go func() {
		w := app.NewWindow(app.Size(unit.Dp(800), unit.Dp(480)))
		loop(w)
	}()
	app.Main()
}

func loop(w *app.Window) {
	th := material.NewTheme(gofont.Collection())
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
			}
			/*	layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return
			}*/
			if count < 3 {
				layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return material.H5(th, fmt.Sprintf("Number of clicks: %d", count)).Layout(gtx)
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
			if count >2{
			layout.W.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.H5(th, fmt.Sprintf("Number: %d", count)).Layout(gtx)
					}),


					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Button(th, &btn, "Click me!").Layout(gtx)
					}),


					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Alignment: layout.Middle}.Layout(gtx,
							layout.Flexed(1, material.Slider(th, float, 0, 1000).Layout),
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

// END OMIT