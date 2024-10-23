package main

import (
	"os/exec"
	"runtime"
)

func main() {
	url := "http://localhost:30020"

	var err error
	switch runtime.GOOS {
	case "windows":
		// Windows에서는 start 명령어를 사용하여 Edge 실행
		err = exec.Command("cmd", "/c", "start", "msedge", url).Start()
	default:
		// 기타 OS에서는 xdg-open (Linux) 또는 open (macOS)를 사용
		err = exec.Command("xdg-open", url).Start()
		if err != nil {
			err = exec.Command("open", url).Start()
		}
	}

	if err != nil {
		panic(err)
	}
}

