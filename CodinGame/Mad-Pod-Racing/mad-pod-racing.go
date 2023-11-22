package main

import "fmt"

type Vector struct {
	X, Y int
}

type Pod struct {
	Pos       Vector
	Speed     Vector
	Angle     int
	NextChkpt int
	NextPos   Vector
	NextPwr   int
}

func (p *Pod) Output() string {
	return fmt.Sprintf("%d %d %d", p.NextPos.X, p.NextPos.Y, 200)
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
	p.NextPos = g.Chkpts[p.NextChkpt]
}

func (g *GameMaster) Write(out string) {
	fmt.Println(out)
}

var g GameMaster

func main() {
	g.Init()
	for {
		for i := 0; i < 2; i++ {
			g.Read(&g.You[i])
		}
		for i := 0; i < 2; i++ {
			g.Read(&g.Enemy[i])
		}
		for i := 0; i < 2; i++ {
			g.Write(g.You[i].Output())
		}
	}
}
