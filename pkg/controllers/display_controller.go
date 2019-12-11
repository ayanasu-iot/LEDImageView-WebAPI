package controllers

import (
	"fmt"
	"os/exec"
)

func ShowImage(filepath string) {
	cmd := exec.Command("killall", "led-image-viewe")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	cmd = exec.Command(fmt.Sprintf("/usr/local/bin/led-image-viewer %s",filepath), "--led-slowdown-gpio=2")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
