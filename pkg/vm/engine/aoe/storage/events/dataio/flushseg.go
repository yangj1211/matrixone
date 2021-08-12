package dataio

import (
	// log "github.com/sirupsen/logrus"
	"matrixone/pkg/container/batch"
	"matrixone/pkg/vm/engine/aoe/storage/common"
	"matrixone/pkg/vm/engine/aoe/storage/layout/dataio"
	"matrixone/pkg/vm/engine/aoe/storage/layout/table/v2/iface"
	"matrixone/pkg/vm/engine/aoe/storage/sched"
)

type flushSegEvent struct {
	baseEvent
	Segment iface.ISegment
}

func NewFlushSegEvent(ctx *Context, seg iface.ISegment) *flushSegEvent {
	e := &flushSegEvent{Segment: seg}
	e.baseEvent = baseEvent{
		Ctx:       ctx,
		BaseEvent: *sched.NewBaseEvent(e, sched.MergeSortEvent, ctx.DoneCB, ctx.Waitable),
	}
	return e
}

func (e *flushSegEvent) Execute() error {
	defer e.Segment.Unref()
	ids := e.Segment.BlockIds()
	meta := e.Segment.GetMeta()
	batches := make([]*batch.Batch, len(ids))
	colCnt := len(meta.Table.Schema.ColDefs)
	nodes := make([]*common.MemNode, 0)
	for i := 0; i < len(ids); i++ {
		blk := e.Segment.WeakRefBlock(ids[i])
		bat := &batch.Batch{}
		for colIdx := 0; colIdx < colCnt; colIdx++ {
			vec, err := blk.GetVectorWrapper(colIdx)
			if err != nil {
				panic(err)
			}
			bat.Vecs = append(bat.Vecs, &vec.Vector)
			if vec.Vector.Length() == 0 {
				panic("")
			}
			nodes = append(nodes, vec.MNode)
		}
		batches[i] = bat
	}

	release := func() {
		for _, node := range nodes {
			common.GPool.Free(node)
		}
	}
	defer release()

	w := dataio.NewSegmentWriter(batches, meta, e.Ctx.Opts.Meta.Conf.Dir)
	return w.Execute()
}
