package utr50

//go:generate go run tool/generate.go -o table.go
//go:generate stringer -type prop

type prop int

const (
	N prop = iota
	U
	R
	Tu
	Tr
)

func Prop(r rune) prop {
	for _, t := range table {
		if t.from <= r && r <= t.to {
			return t.prop
		}
	}
	return N
}
