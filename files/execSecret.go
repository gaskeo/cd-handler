package files

import (
	"fmt"
	"os/exec"
)

func ExecSecret() error {
	_, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s; %s", GetPath(), "./entry.sh")).Output()
	return err
}
