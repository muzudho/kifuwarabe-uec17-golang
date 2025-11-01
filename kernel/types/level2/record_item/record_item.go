package kernel

// Level 1
import point "github.com/muzudho/kifuwarabe-uec17/kernel/level_2_conceptual/sublevel_1/point"

// RecordItem - 棋譜の一手分
type RecordItem struct {
	// 着手点
	PlacePlay point.Point

	// [O22o7o1o0] コウの位置
	Ko point.Point
}

// NewRecordItem - 棋譜の一手分
func NewRecordItem() *RecordItem {
	var ri = new(RecordItem)
	return ri
}

// Clear - 空っぽにします
func (ri *RecordItem) Clear() {
	ri.PlacePlay = point.Point(0)
	ri.Ko = point.Point(0)
}
