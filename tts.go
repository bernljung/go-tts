package tts

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const tts_url = "http://translate.google.com/translate_tts"

type TTS struct {
	Filename, Tl, Q string
}

func (t TTS) Play() bool {
	out, err := exec.Command("mpg123", t.Filename).Output()
	log.Println("command:", out, err)
	return true
}

func (t TTS) Download() bool {
	out, err := os.Create(t.Filename)
	defer out.Close()
	log.Println(tts_url + "?tl=" + t.Tl + "&q=" + t.Q)
	resp, err := http.Get(tts_url + "?tl=" + t.Tl + "&q=" + t.Q)
	log.Println("resp:", resp, err)
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
	return true
}
