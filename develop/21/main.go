package main

type Import interface {
	LoveCats() []byte
}

type Your struct {
}

func (y *Your) ConvertDogsToCatsLove() string {
	var a string
	return a
}

type AdapterToCatLover struct {
	yours *Your
}

func (adapter AdapterToCatLover)LoveCats() string {
	return  adapter.yours.ConvertDogsToCatsLove()
}

func main() {
}
