package utr50

//go:generate go run tool/generate.go -o table.go
//go:generate stringer -type Prop

// Prop is identify type for properties.
type Prop int

const (
	// N is: Unknown Type
	N Prop = iota
	// U is: Characters which are displayed upright, with the same orientation that appears in the code charts.
	U
	// R is: Characters which are displayed sideways, rotated 90 degrees clockwise compared to the code charts.
	R
	// Tu is: Characters which are not just upright or sideways, but generally require a different glyph than in the code charts when used in vertical texts. In addition, as a fallback, the character can be displayed with the code chart glyph upright.
	Tu
	// Tr is: Same as Tu except that, as a fallback, the character can be displayed with the code chart glyph rotated 90 degrees clockwise.
	Tr
)

// RuneProp return Prop for rune r.
func RuneProp(r rune) Prop {
	for _, t := range table {
		if t.from <= r && r <= t.to {
			return t.prop
		}
	}
	return N
}
