// Code generated by "stringer -type Prop"; DO NOT EDIT.

package utr50

import "fmt"

const _Prop_name = "NURTuTr"

var _Prop_index = [...]uint8{0, 1, 2, 3, 5, 7}

func (i Prop) String() string {
	if i < 0 || i >= Prop(len(_Prop_index)-1) {
		return fmt.Sprintf("Prop(%d)", i)
	}
	return _Prop_name[_Prop_index[i]:_Prop_index[i+1]]
}
