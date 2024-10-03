package gadget

import "fmt"

type TapePlayer struct {
	batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
