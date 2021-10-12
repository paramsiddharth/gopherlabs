package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// जहाँ तक मैं समझता हूँ, यह विंडोज़ पर काम नहीं करेगा
	prc := exec.Command("echo", "Hi!")
	output := bytes.NewBuffer([]byte{})
	prc.Stdout = output

	if e := prc.Run(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}

	if prc.ProcessState.Success() {
		fmt.Printf("Output: " + output.String())
	}
}
