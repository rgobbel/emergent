// Code generated by numeric.gen.go.tmpl. DO NOT EDIT.

// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etensor

import (
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/apache/arrow/go/arrow/tensor"
	"github.com/emer/emergent/bitslice"
)

// Int64 is an n-dim array of int64s.
type Int64 struct {
	Shape
	Values []int64
	Nulls  bitslice.Slice
}

// NewInt64 returns a new n-dimensional array of int64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewInt64(shape, strides []int, names []string) *Int64 {
	tsr := &Int64{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]int64, tsr.Len())
	return tsr
}

// NewInt64Shape returns a new n-dimensional array of int64s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewInt64Shape(shape *Shape) *Int64 {
	tsr := &Int64{}
	tsr.CopyShape(shape)
	tsr.Values = make([]int64, tsr.Len())
	return tsr
}

func (tsr *Int64) Value(i []int) int64    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Int64) Set(i []int, val int64) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Int64) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Int64) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Int64) ToArrow() *tensor.Int64 {
	bld := array.NewInt64Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewInt64Array()
	return tensor.NewInt64(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Int64) FromArrow(arw *tensor.Int64, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Int64Values()
		tsr.Values = make([]int64, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Int64Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Uint64 is an n-dim array of uint64s.
type Uint64 struct {
	Shape
	Values []uint64
	Nulls  bitslice.Slice
}

// NewUint64 returns a new n-dimensional array of uint64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewUint64(shape, strides []int, names []string) *Uint64 {
	tsr := &Uint64{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]uint64, tsr.Len())
	return tsr
}

// NewUint64Shape returns a new n-dimensional array of uint64s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewUint64Shape(shape *Shape) *Uint64 {
	tsr := &Uint64{}
	tsr.CopyShape(shape)
	tsr.Values = make([]uint64, tsr.Len())
	return tsr
}

func (tsr *Uint64) Value(i []int) uint64    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Uint64) Set(i []int, val uint64) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Uint64) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Uint64) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Uint64) ToArrow() *tensor.Uint64 {
	bld := array.NewUint64Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewUint64Array()
	return tensor.NewUint64(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Uint64) FromArrow(arw *tensor.Uint64, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Uint64Values()
		tsr.Values = make([]uint64, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Uint64Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Float64 is an n-dim array of float64s.
type Float64 struct {
	Shape
	Values []float64
	Nulls  bitslice.Slice
}

// NewFloat64 returns a new n-dimensional array of float64s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewFloat64(shape, strides []int, names []string) *Float64 {
	tsr := &Float64{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]float64, tsr.Len())
	return tsr
}

// NewFloat64Shape returns a new n-dimensional array of float64s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewFloat64Shape(shape *Shape) *Float64 {
	tsr := &Float64{}
	tsr.CopyShape(shape)
	tsr.Values = make([]float64, tsr.Len())
	return tsr
}

func (tsr *Float64) Value(i []int) float64    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Float64) Set(i []int, val float64) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Float64) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Float64) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Float64) ToArrow() *tensor.Float64 {
	bld := array.NewFloat64Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewFloat64Array()
	return tensor.NewFloat64(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Float64) FromArrow(arw *tensor.Float64, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Float64Values()
		tsr.Values = make([]float64, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Float64Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Int32 is an n-dim array of int32s.
type Int32 struct {
	Shape
	Values []int32
	Nulls  bitslice.Slice
}

// NewInt32 returns a new n-dimensional array of int32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewInt32(shape, strides []int, names []string) *Int32 {
	tsr := &Int32{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]int32, tsr.Len())
	return tsr
}

// NewInt32Shape returns a new n-dimensional array of int32s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewInt32Shape(shape *Shape) *Int32 {
	tsr := &Int32{}
	tsr.CopyShape(shape)
	tsr.Values = make([]int32, tsr.Len())
	return tsr
}

func (tsr *Int32) Value(i []int) int32    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Int32) Set(i []int, val int32) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Int32) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Int32) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Int32) ToArrow() *tensor.Int32 {
	bld := array.NewInt32Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewInt32Array()
	return tensor.NewInt32(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Int32) FromArrow(arw *tensor.Int32, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Int32Values()
		tsr.Values = make([]int32, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Int32Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Uint32 is an n-dim array of uint32s.
type Uint32 struct {
	Shape
	Values []uint32
	Nulls  bitslice.Slice
}

// NewUint32 returns a new n-dimensional array of uint32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewUint32(shape, strides []int, names []string) *Uint32 {
	tsr := &Uint32{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]uint32, tsr.Len())
	return tsr
}

// NewUint32Shape returns a new n-dimensional array of uint32s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewUint32Shape(shape *Shape) *Uint32 {
	tsr := &Uint32{}
	tsr.CopyShape(shape)
	tsr.Values = make([]uint32, tsr.Len())
	return tsr
}

func (tsr *Uint32) Value(i []int) uint32    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Uint32) Set(i []int, val uint32) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Uint32) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Uint32) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Uint32) ToArrow() *tensor.Uint32 {
	bld := array.NewUint32Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewUint32Array()
	return tensor.NewUint32(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Uint32) FromArrow(arw *tensor.Uint32, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Uint32Values()
		tsr.Values = make([]uint32, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Uint32Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Float32 is an n-dim array of float32s.
type Float32 struct {
	Shape
	Values []float32
	Nulls  bitslice.Slice
}

// NewFloat32 returns a new n-dimensional array of float32s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewFloat32(shape, strides []int, names []string) *Float32 {
	tsr := &Float32{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]float32, tsr.Len())
	return tsr
}

// NewFloat32Shape returns a new n-dimensional array of float32s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewFloat32Shape(shape *Shape) *Float32 {
	tsr := &Float32{}
	tsr.CopyShape(shape)
	tsr.Values = make([]float32, tsr.Len())
	return tsr
}

func (tsr *Float32) Value(i []int) float32    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Float32) Set(i []int, val float32) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Float32) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Float32) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Float32) ToArrow() *tensor.Float32 {
	bld := array.NewFloat32Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewFloat32Array()
	return tensor.NewFloat32(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Float32) FromArrow(arw *tensor.Float32, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Float32Values()
		tsr.Values = make([]float32, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Float32Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Int16 is an n-dim array of int16s.
type Int16 struct {
	Shape
	Values []int16
	Nulls  bitslice.Slice
}

// NewInt16 returns a new n-dimensional array of int16s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewInt16(shape, strides []int, names []string) *Int16 {
	tsr := &Int16{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]int16, tsr.Len())
	return tsr
}

// NewInt16Shape returns a new n-dimensional array of int16s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewInt16Shape(shape *Shape) *Int16 {
	tsr := &Int16{}
	tsr.CopyShape(shape)
	tsr.Values = make([]int16, tsr.Len())
	return tsr
}

func (tsr *Int16) Value(i []int) int16    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Int16) Set(i []int, val int16) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Int16) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Int16) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Int16) ToArrow() *tensor.Int16 {
	bld := array.NewInt16Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewInt16Array()
	return tensor.NewInt16(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Int16) FromArrow(arw *tensor.Int16, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Int16Values()
		tsr.Values = make([]int16, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Int16Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Uint16 is an n-dim array of uint16s.
type Uint16 struct {
	Shape
	Values []uint16
	Nulls  bitslice.Slice
}

// NewUint16 returns a new n-dimensional array of uint16s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewUint16(shape, strides []int, names []string) *Uint16 {
	tsr := &Uint16{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]uint16, tsr.Len())
	return tsr
}

// NewUint16Shape returns a new n-dimensional array of uint16s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewUint16Shape(shape *Shape) *Uint16 {
	tsr := &Uint16{}
	tsr.CopyShape(shape)
	tsr.Values = make([]uint16, tsr.Len())
	return tsr
}

func (tsr *Uint16) Value(i []int) uint16    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Uint16) Set(i []int, val uint16) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Uint16) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Uint16) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Uint16) ToArrow() *tensor.Uint16 {
	bld := array.NewUint16Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewUint16Array()
	return tensor.NewUint16(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Uint16) FromArrow(arw *tensor.Uint16, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Uint16Values()
		tsr.Values = make([]uint16, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Uint16Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Int8 is an n-dim array of int8s.
type Int8 struct {
	Shape
	Values []int8
	Nulls  bitslice.Slice
}

// NewInt8 returns a new n-dimensional array of int8s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewInt8(shape, strides []int, names []string) *Int8 {
	tsr := &Int8{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]int8, tsr.Len())
	return tsr
}

// NewInt8Shape returns a new n-dimensional array of int8s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewInt8Shape(shape *Shape) *Int8 {
	tsr := &Int8{}
	tsr.CopyShape(shape)
	tsr.Values = make([]int8, tsr.Len())
	return tsr
}

func (tsr *Int8) Value(i []int) int8    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Int8) Set(i []int, val int8) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Int8) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Int8) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Int8) ToArrow() *tensor.Int8 {
	bld := array.NewInt8Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewInt8Array()
	return tensor.NewInt8(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Int8) FromArrow(arw *tensor.Int8, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Int8Values()
		tsr.Values = make([]int8, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Int8Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}

// Uint8 is an n-dim array of uint8s.
type Uint8 struct {
	Shape
	Values []uint8
	Nulls  bitslice.Slice
}

// NewUint8 returns a new n-dimensional array of uint8s.
// If strides is nil, row-major strides will be inferred.
// If names is nil, a slice of empty strings will be created.
// Nulls are initialized to nil.
func NewUint8(shape, strides []int, names []string) *Uint8 {
	tsr := &Uint8{}
	tsr.SetShape(shape, strides, names)
	tsr.Values = make([]uint8, tsr.Len())
	return tsr
}

// NewUint8Shape returns a new n-dimensional array of uint8s.
// Using shape structure instead of separate data.
// Nulls are initialized to nil.
func NewUint8Shape(shape *Shape) *Uint8 {
	tsr := &Uint8{}
	tsr.CopyShape(shape)
	tsr.Values = make([]uint8, tsr.Len())
	return tsr
}

func (tsr *Uint8) Value(i []int) uint8    { j := tsr.Offset(i); return tsr.Values[j] }
func (tsr *Uint8) Set(i []int, val uint8) { j := tsr.Offset(i); tsr.Values[j] = val }
func (tsr *Uint8) IsNull(i []int) bool {
	if tsr.Nulls == nil {
		return false
	}
	j := tsr.Offset(i)
	return tsr.Nulls.Index(j)
}
func (tsr *Uint8) SetNull(i []int, nul bool) {
	if tsr.Nulls == nil {
		tsr.Nulls = bitslice.Make(tsr.Len(), 0)
	}
	j := tsr.Offset(i)
	tsr.Nulls.Set(j, nul)
}

// ToArrow returns the apache arrow equivalent of the tensor
func (tsr *Uint8) ToArrow() *tensor.Uint8 {
	bld := array.NewUint8Builder(memory.DefaultAllocator)
	if tsr.Nulls != nil {
		bld.AppendValues(tsr.Values, tsr.Nulls.ToBools())
	} else {
		bld.AppendValues(tsr.Values, nil)
	}
	vec := bld.NewUint8Array()
	return tensor.NewUint8(vec.Data(), tsr.Shape64(), tsr.Strides64(), tsr.DimNames())
}

// FromArrow intializes this tensor from an arrow tensor of same type
// cpy = true means make a copy of the arrow data, otherwise it directly
// refers to its values slice -- we do not Retain() on that data so it is up
// to the go GC and / or your own memory management policies to ensure the data
// remains intact!
func (tsr *Uint8) FromArrow(arw *tensor.Uint8, cpy bool) {
	nms := make([]string, arw.NumDims()) // todo: would be nice if it exposed DimNames()
	for i := range nms {
		nms[i] = arw.DimName(i)
	}
	tsr.SetShape64(arw.Shape(), arw.Strides(), nms)
	if cpy {
		vls := arw.Uint8Values()
		tsr.Values = make([]uint8, tsr.Len())
		copy(tsr.Values, vls)
	} else {
		tsr.Values = arw.Uint8Values()
	}
	// todo: doesn't look like the Data() exposes the nulls themselves so it is not
	// clear we can copy the null values -- nor does it seem that the tensor class
	// exposes it either!  https://github.com/apache/arrow/issues/3496
	// nln := arw.Data().NullN()
	// if nln > 0 {
	// }
}
