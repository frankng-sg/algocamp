package main

import (
	"fmt"
	"math"
	"os"
)

type Vector struct {
	X, Y int
}

func (v *Vector) FromPolar(angle, radius int) {
	v.X = int(Cos(angle) * float64(radius))
	v.Y = int(Sin(angle) * float64(radius))
}

func (v *Vector) Add(u Vector) {
	v.X += u.X
	v.Y += u.Y
}

type Control struct {
	Dir Vector
	Pwr int
}

type Pod struct {
	Pos       Vector
	Speed     Vector
	Angle     int
	NextChkpt int
	Ctrl      Control
}

func (p *Pod) Output() string {
	return fmt.Sprintf("%d %d %d", p.Ctrl.Dir.X, p.Ctrl.Dir.Y, p.Ctrl.Pwr)
}

func (p *Pod) CalcFit(next Vector, angle int) int {
	return Dist(next, g.Chkpts[p.NextChkpt]) + angle
}

func (p *Pod) Plan() {
	var dir Vector

	bestFit := math.MaxInt
	for _, angle := range []int{-18, 0, 18} {
		for _, pwr := range []int{0, 50, 100, 150, 200} {
			dir.FromPolar(p.Angle+angle, pwr)
			fitScore := p.CalcFit(Add(p.Pos, Add(p.Speed, dir)), angle)
			if fitScore < bestFit {
				dir.Add(p.Pos)
				bestFit, p.Ctrl.Dir.X, p.Ctrl.Dir.Y, p.Ctrl.Pwr = fitScore, dir.X, dir.Y, pwr
			}
		}
	}
}

type GameMaster struct {
	NLaps   int
	NChkpts int
	Chkpts  []Vector
	You     []Pod
	Enemy   []Pod
}

func (g *GameMaster) Init() {
	fmt.Scan(&g.NLaps)
	fmt.Scan(&g.NChkpts)
	g.Chkpts = make([]Vector, g.NChkpts)
	for i := 0; i < g.NChkpts; i++ {
		fmt.Scan(&g.Chkpts[i].X, &g.Chkpts[i].Y)
	}
	g.You = make([]Pod, 2)
	g.Enemy = make([]Pod, 2)
}

func (g *GameMaster) Read(p *Pod) {
	fmt.Scan(&p.Pos.X, &p.Pos.Y, &p.Speed.X, &p.Speed.Y, &p.Angle, &p.NextChkpt)
}

func (g *GameMaster) Write(out string) {
	fmt.Println(out)
}

var g GameMaster
var cos [361]float64
var sin [361]float64

func main() {
	g.Init()
	InitCache()
	for {
		for i := 0; i < 2; i++ {
			g.Read(&g.You[i])
		}
		for i := 0; i < 2; i++ {
			g.Read(&g.Enemy[i])
		}
		for i := 0; i < 2; i++ {
			g.You[i].Plan()
			g.Write(g.You[i].Output())
		}
		fmt.Fprintln(os.Stderr, "Pod 1 Pos: ", g.You[0].Pos, " Speed ", g.You[0].Speed, " Angle ", g.You[0].Angle)
		fmt.Fprintln(os.Stderr, "Pod 1 Dir: ", g.You[0].Ctrl)
	}
}

func InitCache() {
	for i := 0; i < 361; i++ {
		cos[i] = math.Cos(float64(i))
		sin[i] = math.Sin(float64(i))
	}
}

func Cos(x int) float64 {
	if x < 0 {
		x += 360
	}
	return cos[x]
}

func Sin(x int) float64 {
	if x < 0 {
		x += 360
	}
	return sin[x]
}

func Add(u, v Vector) Vector {
	var r Vector
	r.X = v.X + u.X
	r.Y = v.Y + u.Y
	return r
}

func Dist(u, v Vector) int {
	x := float64(u.X - v.X)
	y := float64(u.Y - v.Y)
	return int(math.Sqrt(x*x + y*y))
}
