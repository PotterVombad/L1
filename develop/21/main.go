package main

type Import interface {
	LoveCats() []byte
}

type Your struct {
	a string
}

func (y *Your) ConvertDogsToCatsLove() string {
	return y.a
}

type AdapterToCatLover struct {
	yours *Your
}

func (adapter AdapterToCatLover)LoveCats() string {
	return  adapter.yours.ConvertDogsToCatsLove()
}

func main() {
}
