package factory_method

import "fmt"

// 负责生产的工厂
func getGun(gunType string) (IGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "musket":
		return newMusket(), nil
	default:
		return nil, fmt.Errorf("wrong gun type passed")
	}
}
