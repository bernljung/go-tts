package tts

import (
	"log"
	"net/url"
	"os/exec"
)

const tts_url = "http://translate.google.com/translate_tts"

type TTS struct {
	Tl, Q string
}

func (t TTS) Play() {
	v := url.Values{}
	v.Set("tl", t.Tl)
	v.Add("q", t.Q)
	v.Add("ie", "UTF-8")
	query := "?" + v.Encode()
	log.Println("Query:", t.Q)
	log.Println("Command:", "mpg123", tts_url+query)
	out, err := exec.Command("mpg123", tts_url+query).Output()
	log.Println("command:", out, err)
}
