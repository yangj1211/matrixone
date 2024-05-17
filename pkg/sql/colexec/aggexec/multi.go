// Copyright 2024 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aggexec

import (
	"fmt"
	"github.com/matrixorigin/matrixone/pkg/common/mpool"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
)

// newMultiAggFuncExecRetFixed return the agg executor for multi columns agg which return a fixed length type.
func newMultiAggFuncExecRetFixed(
	mg AggMemoryManager, info multiAggInfo, impl multiColumnAggImplementation) AggFuncExec {

	switch info.retType.Oid {
	case types.T_bool:
		e := &multiAggFuncExec1[bool]{}
		e.init(mg, info, impl)
		return e
	case types.T_int8:
		e := &multiAggFuncExec1[int8]{}
		e.init(mg, info, impl)
		return e
	case types.T_int16:
		e := &multiAggFuncExec1[int16]{}
		e.init(mg, info, impl)
		return e
	case types.T_int32:
		e := &multiAggFuncExec1[int32]{}
		e.init(mg, info, impl)
		return e
	case types.T_int64:
		e := &multiAggFuncExec1[int64]{}
		e.init(mg, info, impl)
		return e
	case types.T_uint8:
		e := &multiAggFuncExec1[uint8]{}
		e.init(mg, info, impl)
		return e
	case types.T_uint16:
		e := &multiAggFuncExec1[uint16]{}
		e.init(mg, info, impl)
		return e
	case types.T_uint32:
		e := &multiAggFuncExec1[uint32]{}
		e.init(mg, info, impl)
		return e
	case types.T_uint64:
		e := &multiAggFuncExec1[uint64]{}
		e.init(mg, info, impl)
		return e
	case types.T_float32:
		e := &multiAggFuncExec1[float32]{}
		e.init(mg, info, impl)
		return e
	case types.T_float64:
		e := &multiAggFuncExec1[float64]{}
		e.init(mg, info, impl)
		return e
	case types.T_decimal64:
		e := &multiAggFuncExec1[types.Decimal64]{}
		e.init(mg, info, impl)
		return e
	case types.T_decimal128:
		e := &multiAggFuncExec1[types.Decimal128]{}
		e.init(mg, info, impl)
		return e
	case types.T_date:
		e := &multiAggFuncExec1[types.Date]{}
		e.init(mg, info, impl)
		return e
	case types.T_datetime:
		e := &multiAggFuncExec1[types.Datetime]{}
		e.init(mg, info, impl)
		return e
	case types.T_time:
		e := &multiAggFuncExec1[types.Time]{}
		e.init(mg, info, impl)
		return e
	case types.T_timestamp:
		e := &multiAggFuncExec1[types.Timestamp]{}
		e.init(mg, info, impl)
		return e
	}

	panic(fmt.Sprintf("unexpected parameter to Init a multiAggFuncExec, aggInfo: %s", info))
}

// newMultiAggFuncExecRetVar return the agg executor for multi columns agg which return a variable length type.
func newMultiAggFuncExecRetVar(
	mg AggMemoryManager, info multiAggInfo, impl multiColumnAggImplementation) AggFuncExec {
	e := &multiAggFuncExec2{}
	e.init(mg, info, impl)
	return e
}

// multiAggFuncExec1 and multiAggFuncExec2 are the executors of multi columns agg.
// 1's return type is a fixed length type.
// 2's return type is bytes.
type multiAggFuncExec1[T types.FixedSizeTExceptStrType] struct {
	multiAggInfo

	args   []mArg1[T]
	ret    aggFuncResult[T]
	groups []MultiAggRetFixed[T]

	initGroup MultiAggInit1[T]
	// todo: it's an optimization to move rowValid into eval.
	rowValid rowValidForMultiAgg1[T]
	merge    MultiAggMerge1[T]
	eval     MultiAggEval1[T]
	flush    MultiAggFlush1[T]

	// method to new the private structure for group growing.
	gGroup func() MultiAggRetFixed[T]
}
type multiAggFuncExec2 struct {
	multiAggInfo

	args   []mArg2
	ret    aggFuncBytesResult
	groups []MultiAggRetVar

	initGroup MultiAggInit2
	rowValid  rowValidForMultiAgg2
	merge     MultiAggMerge2
	eval      MultiAggEval2
	flush     MultiAggFlush2

	// method to new the private structure for group growing.
	gGroup func() MultiAggRetVar
}

func (exec *multiAggFuncExec1[to]) marshal() ([]byte, error) {
	d := exec.multiAggInfo.getEncoded()
	r, err := exec.ret.marshal()
	if err != nil {
		return nil, err
	}
	encoded := &EncodedAgg{
		Info:   d,
		Result: r,
	}
	if len(exec.groups) > 0 {
		encoded.Groups = make([][]byte, len(exec.groups))
		for i := range encoded.Groups {
			encoded.Groups[i] = exec.groups[i].Marshal()
		}
	}
	return encoded.Marshal()
}

func (exec *multiAggFuncExec1[T]) unmarshal(mp *mpool.MPool, result []byte, groups [][]byte) error {
	exec.groups = make([]MultiAggRetFixed[T], len(groups))
	for i := range exec.groups {
		exec.groups[i] = exec.gGroup()
		exec.groups[i].Unmarshal(groups[i])
	}
	return exec.ret.unmarshal(result)
}

func (exec *multiAggFuncExec2) marshal() ([]byte, error) {
	d := exec.multiAggInfo.getEncoded()
	r, err := exec.ret.marshal()
	if err != nil {
		return nil, err
	}
	encoded := &EncodedAgg{
		Info:   d,
		Result: r,
	}
	if len(exec.groups) > 0 {
		encoded.Groups = make([][]byte, len(exec.groups))
		for i := range encoded.Groups {
			encoded.Groups[i] = exec.groups[i].Marshal()
		}
	}
	return encoded.Marshal()
}

func (exec *multiAggFuncExec2) unmarshal(mp *mpool.MPool, result []byte, groups [][]byte) error {
	exec.groups = make([]MultiAggRetVar, len(groups))
	for i := range exec.groups {
		exec.groups[i] = exec.gGroup()
		exec.groups[i].Unmarshal(groups[i])
	}
	return exec.ret.unmarshal(result)
}

func (exec *multiAggFuncExec1[T]) init(
	mg AggMemoryManager,
	info multiAggInfo,
	impl multiColumnAggImplementation) {

	exec.multiAggInfo = info
	exec.args = make([]mArg1[T], len(info.argTypes))
	exec.ret = initFixedAggFuncResult[T](mg, info.retType, info.emptyNull)
	exec.groups = make([]MultiAggRetFixed[T], 0, 1)
	exec.gGroup = impl.generator.(func() MultiAggRetFixed[T])
	exec.args = make([]mArg1[T], len(info.argTypes))

	fillNullWhich := impl.fillNullWhich.([]MultiAggFillNull1[T])
	for i := range exec.args {
		exec.args[i] = newArgumentOfMultiAgg1[T](info.argTypes[i])

		exec.args[i].cacheFill(impl.fillWhich[i], fillNullWhich[i])
	}
	exec.rowValid = impl.rowValid.(rowValidForMultiAgg1[T])
	exec.merge = impl.merge.(MultiAggMerge1[T])
	exec.eval = impl.eval.(MultiAggEval1[T])
	if impl.flush != nil {
		exec.flush = impl.flush.(MultiAggFlush1[T])
	}
	if impl.init != nil {
		exec.initGroup = impl.init.(MultiAggInit1[T])
	}
}

func (exec *multiAggFuncExec1[T]) GroupGrow(more int) error {
	if err := exec.ret.grows(more); err != nil {
		return err
	}
	setter := exec.ret.aggSet
	moreGroup := make([]MultiAggRetFixed[T], more)
	for i := 0; i < more; i++ {
		moreGroup[i] = exec.gGroup()
	}

	if exec.initGroup != nil {
		for i := 0; i < more; i++ {
			exec.ret.groupToSet = i + len(exec.groups)
			exec.initGroup(moreGroup[i], setter, exec.argTypes, exec.retType)
		}
	}

	exec.groups = append(exec.groups, moreGroup...)
	return nil
}

func (exec *multiAggFuncExec1[T]) PreAllocateGroups(more int) error {
	return exec.ret.preAllocate(more)
}

func (exec *multiAggFuncExec1[T]) Fill(groupIndex int, row int, vectors []*vector.Vector) error {
	var err error
	for i, arg := range exec.args {
		arg.prepare(vectors[i])
		if err = arg.doRowFill(exec.groups[groupIndex], uint64(row)); err != nil {
			return err
		}
	}
	exec.ret.groupToSet = groupIndex
	if exec.rowValid(exec.groups[groupIndex]) {
		exec.ret.setGroupNotEmpty(groupIndex)
		if err = exec.eval(exec.groups[groupIndex], exec.ret.aggGet, exec.ret.aggSet); err != nil {
			return err
		}
	}

	return nil
}

func (exec *multiAggFuncExec1[T]) BulkFill(groupIndex int, vectors []*vector.Vector) error {
	var err error
	for i, arg := range exec.args {
		arg.prepare(vectors[i])
	}

	setter := exec.ret.aggSet
	getter := exec.ret.aggGet
	exec.ret.groupToSet = groupIndex
	for i, j := uint64(0), uint64(vectors[0].Length()); i < j; i++ {
		for _, arg := range exec.args {
			if err = arg.doRowFill(exec.groups[groupIndex], i); err != nil {
				return err
			}
		}
		if exec.rowValid(exec.groups[groupIndex]) {
			exec.ret.setGroupNotEmpty(groupIndex)
			if err = exec.eval(exec.groups[groupIndex], getter, setter); err != nil {
				return err
			}
		}
	}

	return nil
}

func (exec *multiAggFuncExec1[T]) BatchFill(offset int, groups []uint64, vectors []*vector.Vector) error {
	var err error
	setter := exec.ret.aggSet
	getter := exec.ret.aggGet
	for i, arg := range exec.args {
		arg.prepare(vectors[i])
	}

	for idx, i, j := 0, uint64(offset), uint64(offset+len(groups)); i < j; i++ {
		if groups[idx] != GroupNotMatched {
			groupIdx := int(groups[idx] - 1)
			for _, arg := range exec.args {
				if err = arg.doRowFill(exec.groups[groupIdx], i); err != nil {
					return err
				}
			}
			exec.ret.groupToSet = groupIdx
			if exec.rowValid(exec.groups[groupIdx]) {
				exec.ret.setGroupNotEmpty(groupIdx)
				if err = exec.eval(exec.groups[groupIdx], getter, setter); err != nil {
					return err
				}
			}

		}
		idx++
	}

	return nil
}

func (exec *multiAggFuncExec1[T]) SetExtraInformation(partialResult any, groupIndex int) error {
	panic("unimplemented SetPreparedResult for multiAggFuncExec1")
}

func (exec *multiAggFuncExec1[T]) Merge(next AggFuncExec, groupIdx1, groupIdx2 int) error {
	other := next.(*multiAggFuncExec1[T])
	exec.ret.groupToSet = groupIdx1
	other.ret.groupToSet = groupIdx2

	exec.ret.mergeEmpty(other.ret.basicResult, groupIdx1, groupIdx2)
	return exec.merge(
		exec.groups[groupIdx1],
		other.groups[groupIdx2],
		exec.ret.aggGet, other.ret.aggGet,
		exec.ret.aggSet)
}

func (exec *multiAggFuncExec1[T]) BatchMerge(next AggFuncExec, offset int, groups []uint64) error {
	other := next.(*multiAggFuncExec1[T])
	setter := exec.ret.aggSet
	getter1, getter2 := exec.ret.aggGet, other.ret.aggGet

	for i := range groups {
		if groups[i] == GroupNotMatched {
			continue
		}
		groupIdx1, groupIdx2 := int(groups[i]-1), i+offset
		exec.ret.groupToSet = groupIdx1
		other.ret.groupToSet = groupIdx2

		exec.ret.mergeEmpty(other.ret.basicResult, groupIdx1, groupIdx2)
		if err := exec.merge(
			exec.groups[groupIdx1],
			other.groups[groupIdx2],
			getter1, getter2,
			setter); err != nil {
			return err
		}
	}
	return nil
}

func (exec *multiAggFuncExec1[T]) Flush() (*vector.Vector, error) {
	setter := exec.ret.aggSet
	getter := exec.ret.aggGet

	if exec.flush == nil {
		return exec.ret.flush(), nil
	}

	if exec.ret.emptyBeNull {
		for i, group := range exec.groups {
			if exec.ret.groupIsEmpty(i) {
				continue
			}
			exec.ret.groupToSet = i
			if err := exec.flush(group, getter, setter); err != nil {
				return nil, err
			}
		}
	} else {
		for i, group := range exec.groups {
			exec.ret.groupToSet = i
			if err := exec.flush(group, getter, setter); err != nil {
				return nil, err
			}
		}
	}
	return exec.ret.flush(), nil
}

func (exec *multiAggFuncExec1[T]) Free() {
	exec.ret.free()
}

func (exec *multiAggFuncExec2) init(
	mg AggMemoryManager,
	info multiAggInfo,
	impl multiColumnAggImplementation) {

	exec.multiAggInfo = info
	exec.args = make([]mArg2, len(info.argTypes))
	exec.ret = initBytesAggFuncResult(mg, info.retType, info.emptyNull)
	exec.groups = make([]MultiAggRetVar, 0, 1)
	exec.gGroup = impl.generator.(func() MultiAggRetVar)
	exec.args = make([]mArg2, len(info.argTypes))

	fillNullWhich := impl.fillNullWhich.([]MultiAggFillNull2)
	for i := range exec.args {
		exec.args[i] = newArgumentOfMultiAgg2(info.argTypes[i])

		exec.args[i].cacheFill(impl.fillWhich[i], fillNullWhich[i])
	}
	exec.rowValid = impl.rowValid.(rowValidForMultiAgg2)
	exec.merge = impl.merge.(MultiAggMerge2)
	exec.eval = impl.eval.(MultiAggEval2)
	if impl.flush != nil {
		exec.flush = impl.flush.(MultiAggFlush2)
	}
	if impl.init != nil {
		exec.initGroup = impl.init.(MultiAggInit2)
	}
}

func (exec *multiAggFuncExec2) GroupGrow(more int) error {
	if err := exec.ret.grows(more); err != nil {
		return err
	}
	setter := exec.ret.aggSet
	moreGroup := make([]MultiAggRetVar, more)
	for i := 0; i < more; i++ {
		moreGroup[i] = exec.gGroup()
	}

	if exec.initGroup != nil {
		for i := 0; i < more; i++ {
			exec.ret.groupToSet = i + len(exec.groups)
			exec.initGroup(moreGroup[i], setter, exec.argTypes, exec.retType)
		}
	}

	exec.groups = append(exec.groups, moreGroup...)
	return nil
}

func (exec *multiAggFuncExec2) PreAllocateGroups(more int) error {
	return exec.ret.preAllocate(more)
}

func (exec *multiAggFuncExec2) Fill(groupIndex int, row int, vectors []*vector.Vector) error {
	var err error
	for i, arg := range exec.args {
		arg.prepare(vectors[i])
		if err = arg.doRowFill(exec.groups[groupIndex], uint64(row)); err != nil {
			return err
		}
	}
	exec.ret.groupToSet = groupIndex
	if exec.rowValid(exec.groups[groupIndex]) {
		exec.ret.setGroupNotEmpty(groupIndex)
		return exec.eval(exec.groups[groupIndex], exec.ret.aggGet, exec.ret.aggSet)
	}

	return nil
}

func (exec *multiAggFuncExec2) BulkFill(groupIndex int, vectors []*vector.Vector) error {
	var err error
	for i, arg := range exec.args {
		arg.prepare(vectors[i])
	}

	setter := exec.ret.aggSet
	getter := exec.ret.aggGet
	exec.ret.groupToSet = groupIndex

	// todo: can do optimization here once all the vectors were constant.

	for i, j := uint64(0), uint64(vectors[0].Length()); i < j; i++ {
		for _, arg := range exec.args {
			if err = arg.doRowFill(exec.groups[groupIndex], i); err != nil {
				return err
			}
		}
		if exec.rowValid(exec.groups[groupIndex]) {
			exec.ret.setGroupNotEmpty(groupIndex)
			if err = exec.eval(exec.groups[groupIndex], getter, setter); err != nil {
				return err
			}
		}
	}

	return nil
}

func (exec *multiAggFuncExec2) BatchFill(offset int, groups []uint64, vectors []*vector.Vector) error {
	var err error
	setter := exec.ret.aggSet
	getter := exec.ret.aggGet
	for i, arg := range exec.args {
		arg.prepare(vectors[i])
	}

	for idx, i, j := 0, uint64(offset), uint64(offset+len(groups)); i < j; i++ {
		if groups[idx] != GroupNotMatched {
			groupIdx := int(groups[idx] - 1)
			for _, arg := range exec.args {
				if err = arg.doRowFill(exec.groups[groupIdx], i); err != nil {
					return err
				}
			}
			exec.ret.groupToSet = groupIdx
			if exec.rowValid(exec.groups[groupIdx]) {
				exec.ret.setGroupNotEmpty(groupIdx)
				if err = exec.eval(exec.groups[groupIdx], getter, setter); err != nil {
					return err
				}
			}

		}
		idx++
	}

	return nil
}

func (exec *multiAggFuncExec2) SetExtraInformation(partialResult any, groupIndex int) error {
	panic("unimplemented SetPreparedResult for multiAggFuncExec2")
}

func (exec *multiAggFuncExec2) Merge(next AggFuncExec, groupIdx1, groupIdx2 int) error {
	other := next.(*multiAggFuncExec2)
	exec.ret.groupToSet = groupIdx1
	other.ret.groupToSet = groupIdx2

	exec.ret.mergeEmpty(other.ret.basicResult, groupIdx1, groupIdx2)
	return exec.merge(
		exec.groups[groupIdx1],
		other.groups[groupIdx2],
		exec.ret.aggGet, other.ret.aggGet,
		exec.ret.aggSet)
}

func (exec *multiAggFuncExec2) BatchMerge(next AggFuncExec, offset int, groups []uint64) error {
	other := next.(*multiAggFuncExec2)
	setter := exec.ret.aggSet
	getter1, getter2 := exec.ret.aggGet, other.ret.aggGet

	for i := range groups {
		if groups[i] == GroupNotMatched {
			continue
		}
		groupIdx1, groupIdx2 := int(groups[i]-1), i+offset
		exec.ret.groupToSet = groupIdx1
		other.ret.groupToSet = groupIdx2

		exec.ret.mergeEmpty(other.ret.basicResult, groupIdx1, groupIdx2)
		if err := exec.merge(
			exec.groups[groupIdx1],
			other.groups[groupIdx2],
			getter1, getter2,
			setter); err != nil {
			return err
		}
	}
	return nil
}

func (exec *multiAggFuncExec2) Flush() (*vector.Vector, error) {
	var err error
	setter := exec.ret.aggSet
	getter := exec.ret.aggGet

	if exec.flush == nil {
		return exec.ret.flush(), nil
	}

	if exec.ret.emptyBeNull {
		for i, group := range exec.groups {
			if exec.ret.groupIsEmpty(i) {
				continue
			}
			exec.ret.groupToSet = i
			if err = exec.flush(group, getter, setter); err != nil {
				return nil, err
			}
		}
	} else {
		for i, group := range exec.groups {
			exec.ret.groupToSet = i
			if err = exec.flush(group, getter, setter); err != nil {
				return nil, err
			}
		}
	}
	return exec.ret.flush(), nil
}

func (exec *multiAggFuncExec2) Free() {
	exec.ret.free()
}
