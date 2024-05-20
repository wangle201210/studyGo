package abstract_factory

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe   // 鞋子
	makeShirt() IShirt // 衬衫
}

// GetSportsFactory 返回一个具体的工厂
func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
