package abstract_factory

import "fmt"

// 让我们假设一下， 如果你想要购买一组运动装备， 比如一双鞋与一件衬衫这样由两种不同产品组合而成的套装。 相信你会想去购买同一品牌的商品， 这样商品之间能够互相搭配起来。
// 如果我们把这样的行为转换成代码的话， 帮助我们创建此类产品组的工具就是抽象工厂， 便于产品之间能够相互匹配。
func client() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s IShoe) {
	fmt.Printf("shoe logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("shoe size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s IShirt) {
	fmt.Printf("shirt logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("shirt size: %d", s.getSize())
	fmt.Println()
}
