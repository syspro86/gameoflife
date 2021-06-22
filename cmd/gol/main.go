package main

import (
	"math/rand"
	"syscall/js"
)

type World struct {
	Width           int
	Height          int
	ActiveTile      int
	Tiles           [][][]bool
	RandomGenFactor float64
}

func (world *World) Update() {
	activeTile := world.Tiles[world.ActiveTile]
	nextTile := world.Tiles[1-world.ActiveTile]

	for x := 0; x < world.Width; x++ {
		for y := 0; y < world.Height; y++ {
			neighbor := 0
			for a := -1; a <= 1; a++ {
				for b := -1; b <= 1; b++ {
					if a == 0 && b == 0 {
						continue
					}
					xx := x + a
					yy := y + b
					if xx < 0 || xx >= world.Width {
						continue
					}
					if yy < 0 || yy >= world.Height {
						continue
					}
					if activeTile[yy][xx] {
						neighbor++
					}
				}
			}
			if activeTile[y][x] {
				if neighbor == 2 || neighbor == 3 {
					nextTile[y][x] = true
				} else {
					nextTile[y][x] = false
				}
			} else {
				if neighbor == 3 {
					nextTile[y][x] = true
				} else if neighbor == 0 {
					nextTile[y][x] = rand.Float64() < world.RandomGenFactor
				} else {
					nextTile[y][x] = false
				}
			}
		}
	}
	world.ActiveTile = 1 - world.ActiveTile
}

func makeRender(world *World) func(js.Value, []js.Value) interface{} {
	return func(this js.Value, p []js.Value) interface{} {
		world.Update()

		var canvas js.Value = p[0] //js.Global().Get("document").Call("getElementById", "canvas")
		var context js.Value = canvas.Call("getContext", "2d")

		canvas.Set("height", world.Height*10)
		canvas.Set("width", world.Width*10)
		context.Call("clearRect", 0, 0, world.Width*10, world.Height*10)

		for y := 0; y < world.Height; y++ {
			for x := 0; x < world.Width; x++ {
				if world.Tiles[world.ActiveTile][y][x] {
					context.Call("fillRect", x*10, y*10, 10, 10)
				}
			}
		}
		return js.Null()
	}
}

func main() {
	world := new(World)
	makeWorld := func(this js.Value, p []js.Value) interface{} {
		world.Width = p[0].Int() / 10
		world.Height = p[1].Int() / 10
		world.ActiveTile = 0
		world.Tiles = make([][][]bool, 2)
		world.Tiles[0] = make([][]bool, world.Height)
		world.Tiles[1] = make([][]bool, world.Height)
		world.RandomGenFactor = 0
		for i := 0; i < world.Height; i++ {
			world.Tiles[0][i] = make([]bool, world.Width)
			world.Tiles[1][i] = make([]bool, world.Width)

			for j := 0; j < world.Width; j++ {
				world.Tiles[0][i][j] = rand.Float64() >= 0.5
			}
		}
		return js.Null()
	}

	c := make(chan struct{}, 0)
	js.Global().Set("initWorld", js.FuncOf(makeWorld))
	js.Global().Set("renderFrame", js.FuncOf(makeRender(world)))
	<-c
}
