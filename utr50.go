package utr50

//go:generate go run tool/generate.go -o table.go
//go:generate stringer -type Prop

type Prop int

const (
	N Prop = iota
	U
	R
	Tu
	Tr
)

func RuneProp(r rune) Prop {
	for _, t := range table {
		if t.from <= r && r <= t.to {
			return t.prop
		}
	}
	return N
}
