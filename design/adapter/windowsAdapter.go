package adapter

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

// InsertIntoLightningPort Windows 本身没实现这个方法，但是适配器可以将信号适配为windows需要的内容，并返回结果
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}
