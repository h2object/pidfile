package pidfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

type PidFile struct {
	path string
}

func getPidProcess(path string) (*os.Process, error) {
	pidString, err := ioutil.ReadFile(path);
	if err != nil {
		return nil, err
	}

	pid, err := strconv.Atoi(string(pidString))
	if err != nil {
		return nil, fmt.Errorf("%s fake", path)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return nil, err
	}
	return proc, nil
}

func checkPidFileAlreadyExists(path string) error {
	if pidString, err := ioutil.ReadFile(path); err == nil {
		if pid, err := strconv.Atoi(string(pidString)); err == nil {
			if _, err := os.Stat(filepath.Join("/proc", string(pid))); err == nil {
				return fmt.Errorf("pid process is running")
			}
		}
	}
	return nil
}

func (file *PidFile) remove() error {
	if err := os.Remove(file.path); err != nil {
		return err
	}
	return nil
}

func New(path string) (*PidFile, error) {
	if err := checkPidFileAlreadyExists(path); err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(path, []byte(fmt.Sprintf("%d", os.Getpid())), 0644); err != nil {
		return nil, err
	}

	return &PidFile{path: path}, nil
}

func Load(path string) (*PidFile, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}
	return &PidFile{path: path}, nil
}

func (file *PidFile) Kill() error {
	defer file.remove()
	proc, err := getPidProcess(file.path)
	if err != nil {
		return err
	}

	return proc.Kill()
}

