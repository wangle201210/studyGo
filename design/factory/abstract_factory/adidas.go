package abstract_factory

// Adidas 这是一个具体工厂
type Adidas struct {
}

// 工厂需要具有 生产鞋子 的能力
func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 14,
		},
	}
}

// 工厂需要具有 生产衬衫 的能力
func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 14,
		},
	}
}
