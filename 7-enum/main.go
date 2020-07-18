package main

import "log"

// Action 行為
type Action int

const (
	// // Play 播放
	// Play Action = iota // 0
	// _
	// // Pause 暫停
	// Pause // 2
	// // Stop 停止
	// Stop // 3

	// // Play 播放
	// Play Action = iota // 0
	// // Pause 暫停
	// Pause // 1
	// // Stop 停止
	// Stop //2

	// Play 播放
	Play Action = iota + 1 //  1
	// Pause 暫停
	Pause // 2
	// Stop 停止
	Stop // 3
)

// Text 訊息
func (a Action) Text() string {
	switch a {
	case Play:
		return "play"
	case Pause:
		return "pause"
	case Stop:
		return "stop"
	default:
		return "undefind"
	}
}

func main() {
	print(Play)
}

func print(a Action) {
	log.Println(a.Text(), a)
}
