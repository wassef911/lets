package pkg

import (
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type ProcServiceInterface interface {
	Processes() (string, error)
	Resources() error
	KillProcessByName(processName string) error
}

var _ ProcServiceInterface = &ProcService{}

type ProcService struct{}

func NewProc() *ProcService {
	return &ProcService{}
}

// lists all running processes (equivalent to `ps aux`).
func (p *ProcService) Processes() (string, error) {
	if !CommandExists("ps") {
		return "", errors.New("'ps' command not found. Please ensure it is installed")
	}

	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		return "", errors.New("failed to list processes: " + err.Error())
	}

	return string(output), nil
}

// provides a live system resource view (equivalent to `htop`).
func (p *ProcService) Resources() error {
	if !CommandExists("htop") {
		return errors.New("'htop' command not found. Please install it to use this feature")
	}

	cmd := exec.Command("htop")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Handle interrupt signal (Ctrl+C) to exit gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		cmd.Process.Signal(os.Interrupt)
	}()

	// Run htop
	err := cmd.Run()
	if err != nil {
		return errors.New("failed to monitor resources: " + err.Error())
	}

	return nil
}

// terminates a process by name (equivalent to `pkill <name>`).
func (p *ProcService) KillProcessByName(processName string) error {
	if !CommandExists("pgrep") {
		return errors.New("'pgrep' command not found. Please ensure it is installed")
	}

	// Find PIDs of processes with the given name
	cmd := exec.Command("pgrep", processName)
	output, err := cmd.Output()
	if err != nil {
		return errors.New("no processes found with name " + processName + " : " + err.Error())
	}

	// Parse PIDs
	pids := strings.Fields(string(output))
	if len(pids) == 0 {
		return errors.New("no processes found with name " + processName)
	}

	// Kill each process
	for _, pidStr := range pids {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			return errors.New("invalid PID " + pidStr + " : " + err.Error())
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			return errors.New("failed to find process with PID " + strconv.Itoa(pid) + " : " + err.Error())
		}

		err = process.Kill()
		if err != nil {
			return errors.New("failed to find process with PID " + strconv.Itoa(pid) + " : " + err.Error())
		}

	}

	return nil
}
