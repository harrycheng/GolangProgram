package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	fmt.Println(source, mtype)
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Print("Unsupported music type")
		return
	}

	p.Play(source)
}