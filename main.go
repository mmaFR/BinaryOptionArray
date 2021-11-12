package BinaryOptionArray

import "fmt"

type pair struct {
	name string
	enabled bool
}

type array8 struct {
	array uint8
	msbMode bool
	option [8]*pair
}
func (a8 *array8) RegisterOption(name string, index uint8) *array8 {
	if index > 7 {
		return a8
	}
	a8.option[index] = &pair{name: name}
	return a8
}
func (a8 *array8) UnregisterOption(index uint8) *array8 {
	if index > 7 {
		return a8
	}
	a8.option[index] = nil
	return a8
}
func (a8 *array8) SetOptionValueByName(name string, value bool) *array8 {
	for _, o := range a8.option {
		if o != nil {
			if o.name == name {
				o.enabled = value
			}
		}
	}
	return a8
}
func (a8 *array8) SetOptionValueByIndex(index uint8, value bool) *array8 {
	if index > 7 {return a8}
	if a8.option[index] != nil {
		a8.option[index].enabled = value
	}
	return a8
}
func (a8 *array8) GetOptionValueByName(name string) *bool {
	var r bool
	for _, o := range a8.option {
		if o != nil {
			if o.name == name {
				if o.enabled {
					r = true
				}
				return &r
			}
		}
	}
	return nil
}
func (a8 *array8) GetOptionValueByIndex(index uint8) *bool {
	var r bool
	if index > 7 {return nil}
	if a8.option[index] != nil {
		if a8.option[index].enabled {
			r = true
		}
		return &r
	}
	return nil
}
func (a8 *array8) SetModeLSB() *array8 {
	a8.msbMode = false
	return a8
}
func (a8 *array8) SetModeMSB() *array8 {
	a8.msbMode = true
	return a8
}
func (a8 *array8) Encode() uint8 {
	a8.array = 0
	if a8.msbMode {
		for i:=0;i<8;i++ {
			if a8.option[i] != nil {
				if a8.option[i].enabled {
					a8.array += 0b1 << i
				}
			}
		}
	} else {
		for i:=7;i>-1;i-- {
			if a8.option[i] != nil {
				if a8.option[i].enabled {
					a8.array += 0b1 << i
				}
			}
		}
	}
	return a8.array
}
func (a8 *array8) Decode(v uint8) {
	var val string = fmt.Sprintf("%b", v)
	for i, b := range []byte(val) {
		if a8.option[i] != nil {
			if b == 49 {
				a8.option[i].enabled = true
			} else {
				a8.option[i].enabled = false
			}
		}
	}
}

type array16 struct {
	array uint16
	msbMode bool
	option [16]*pair
}
func (a16 *array16) RegisterOption(name string, index uint8) *array16 {
	if index > 15 {
		return a16
	}
	a16.option[index] = &pair{name: name}
	return a16
}
func (a16 *array16) UnregisterOption(index uint8) *array16 {
	if index > 15 {
		return a16
	}
	a16.option[index] = nil
	return a16
}
func (a16 *array16) SetOptionValueByName(name string, value bool) *array16 {
	for _, o := range a16.option {
		if o != nil {
			if o.name == name {
				o.enabled = value
			}
		}
	}
	return a16
}
func (a16 *array16) SetOptionValueByIndex(index uint8, value bool) *array16 {
	if index > 15 {return a16}
	if a16.option[index] != nil {
		a16.option[index].enabled = value
	}
	return a16
}
func (a16 *array16) GetOptionValueByName(name string) *bool {
	var r bool
	for _, o := range a16.option {
		if o != nil {
			if o.name == name {
				if o.enabled {
					r = true
				}
				return &r
			}
		}
	}
	return nil
}
func (a16 *array16) GetOptionValueByIndex(index uint8) *bool {
	var r bool
	if index > 15 {return nil}
	if a16.option[index] != nil {
		if a16.option[index].enabled {
			r = true
		}
		return &r
	}
	return nil
}
func (a16 *array16) SetModeLSB() *array16 {
	a16.msbMode = false
	return a16
}
func (a16 *array16) SetModeMSB() *array16 {
	a16.msbMode = true
	return a16
}
func (a16 *array16) Encode() uint16 {
	a16.array = 0
	if a16.msbMode {
		for i:=0;i<16;i++ {
			if a16.option[i] != nil {
				if a16.option[i].enabled {
					a16.array += 0b1 << i
				}
			}
		}
	} else {
		for i:=15;i>-1;i-- {
			if a16.option[i] != nil {
				if a16.option[i].enabled {
					a16.array += 0b1 << i
				}
			}
		}
	}
	return a16.array
}
func (a16 *array16) Decode(v uint16) {
	var val string = fmt.Sprintf("%b", v)
	for i, b := range []byte(val) {
		if a16.option[i] != nil {
			if b == 49 {
				a16.option[i].enabled = true
			} else {
				a16.option[i].enabled = false
			}
		}
	}
}

type array32 struct {
	array uint32
	msbMode bool
	option [32]*pair
}
func (a32 *array32) RegisterOption(name string, index uint8) *array32 {
	if index > 31 {
		return a32
	}
	a32.option[index] = &pair{name: name}
	return a32
}
func (a32 *array32) UnregisterOption(index uint8) *array32 {
	if index > 31 {
		return a32
	}
	a32.option[index] = nil
	return a32
}
func (a32 *array32) SetOptionValueByName(name string, value bool) *array32 {
	for _, o := range a32.option {
		if o != nil {
			if o.name == name {
				o.enabled = value
			}
		}
	}
	return a32
}
func (a32 *array32) SetOptionValueByIndex(index uint8, value bool) *array32 {
	if index > 31 {return a32
	}
	if a32.option[index] != nil {
		a32.option[index].enabled = value
	}
	return a32
}
func (a32 *array32) GetOptionValueByName(name string) *bool {
	var r bool
	for _, o := range a32.option {
		if o != nil {
			if o.name == name {
				if o.enabled {
					r = true
				}
				return &r
			}
		}
	}
	return nil
}
func (a32 *array32) GetOptionValueByIndex(index uint8) *bool {
	var r bool
	if index > 31 {return nil}
	if a32.option[index] != nil {
		if a32.option[index].enabled {
			r = true
		}
		return &r
	}
	return nil
}
func (a32 *array32) SetModeLSB() *array32 {
	a32.msbMode = false
	return a32
}
func (a32 *array32) SetModeMSB() *array32 {
	a32.msbMode = true
	return a32
}
func (a32 *array32) Encode() uint32 {
	a32.array = 0
	if a32.msbMode {
		for i:=0;i<32;i++ {
			if a32.option[i] != nil {
				if a32.option[i].enabled {
					a32.array += 0b1 << i
				}
			}
		}
	} else {
		for i:=31;i>-1;i-- {
			if a32.option[i] != nil {
				if a32.option[i].enabled {
					a32.array += 0b1 << i
				}
			}
		}
	}
	return a32.array
}
func (a32 *array32) Decode(v uint32) {
	var val string = fmt.Sprintf("%b", v)
	for i, b := range []byte(val) {
		if a32.option[i] != nil {
			if b == 49 {
				a32.option[i].enabled = true
			} else {
				a32.option[i].enabled = false
			}
		}
	}
}

type array64 struct {
	array uint64
	msbMode bool
	option [64]*pair
}
func (a64 *array64) RegisterOption(name string, index uint8) *array64 {
	if index > 63 {
		return a64
	}
	a64.option[index] = &pair{name: name}
	return a64
}
func (a64 *array64) UnregisterOption(index uint8) *array64 {
	if index > 63 {
		return a64
	}
	a64.option[index] = nil
	return a64
}
func (a64 *array64) SetOptionValueByName(name string, value bool) *array64 {
	for _, o := range a64.option {
		if o != nil {
			if o.name == name {
				o.enabled = value
			}
		}
	}
	return a64
}
func (a64 *array64) SetOptionValueByIndex(index uint8, value bool) *array64 {
	if index > 63 {return a64
	}
	if a64.option[index] != nil {
		a64.option[index].enabled = value
	}
	return a64
}
func (a64 *array64) GetOptionValueByName(name string) *bool {
	var r bool
	for _, o := range a64.option {
		if o != nil {
			if o.name == name {
				if o.enabled {
					r = true
				}
				return &r
			}
		}
	}
	return nil
}
func (a64 *array64) GetOptionValueByIndex(index uint8) *bool {
	var r bool
	if index > 63 {return nil}
	if a64.option[index] != nil {
		if a64.option[index].enabled {
			r = true
		}
		return &r
	}
	return nil
}
func (a64 *array64) SetModeLSB() *array64 {
	a64.msbMode = false
	return a64
}
func (a64 *array64) SetModeMSB() *array64 {
	a64.msbMode = true
	return a64
}
func (a64 *array64) Encode() uint64 {
	a64.array = 0
	if a64.msbMode {
		for i:=0;i<64;i++ {
			if a64.option[i] != nil {
				if a64.option[i].enabled {
					a64.array += 0b1 << i
				}
			}
		}
	} else {
		for i:=63;i>-1;i-- {
			if a64.option[i] != nil {
				if a64.option[i].enabled {
					a64.array += 0b1 << i
				}
			}
		}
	}
	return a64.array
}
func (a64 *array64) Decode(v uint64) {
	var val string = fmt.Sprintf("%b", v)
	for i, b := range []byte(val) {
		if a64.option[i] != nil {
			if b == 49 {
				a64.option[i].enabled = true
			} else {
				a64.option[i].enabled = false
			}
		}
	}
}

func NewBinaryOptionArray8() *array8 {return new(array8)}
func NewBinaryOptionArray16() *array16 {return new(array16)}
func NewBinaryOptionArray32() *array32 {return new(array32)}
func NewBinaryOptionArray64() *array64 {return new(array64)}