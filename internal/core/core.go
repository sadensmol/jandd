package core

type IComponent interface {
	Init()
	Update()
	Draw(w float32, h float32)
	CleanUp()
}
