package abstract_factory

type AdidasShoe struct {
	Shoe
}

// 可以重写某些特定方法
func (x *AdidasShoe) getSize() int {
	return x.size + 1 // 这个牌子的大小普遍偏大，所以返回时会额外+1
}
