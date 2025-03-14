package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type ProcServiceInterface interface {
	Processes() error
	Resources() error
	KillProcessByName(processName string) error
}

type ProcService struct{}

func NewProc() *ProcService {
	return &ProcService{}
}

// lists all running processes (equivalent to `ps aux`).
func (p *ProcService) Processes() error {
	if !CommandExists("ps") {
		return fmt.Errorf("'ps' command not found. Please ensure it is installed")
	}

	cmd := exec.Command("ps", "aux")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to list processes: %v", err)
	}

	fmt.Println(string(output))
	return nil
}

// provides a live system resource view (equivalent to `htop`).
func (p *ProcService) Resources() error {
	if !CommandExists("htop") {
		return fmt.Errorf("'htop' command not found. Please install it to use this feature")
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
		return fmt.Errorf("failed to monitor resources: %v", err)
	}

	return nil
}

// terminates a process by name (equivalent to `pkill <name>`).
func (p *ProcService) KillProcessByName(processName string) error {
	if !CommandExists("pgrep") {
		return fmt.Errorf("'pgrep' command not found. Please ensure it is installed")
	}

	// Find PIDs of processes with the given name
	cmd := exec.Command("pgrep", processName)
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("no processes found with name %s: %v", processName, err)
	}

	// Parse PIDs
	pids := strings.Fields(string(output))
	if len(pids) == 0 {
		return fmt.Errorf("no processes found with name %s", processName)
	}

	// Kill each process
	for _, pidStr := range pids {
		pid, err := strconv.Atoi(pidStr)
		if err != nil {
			return fmt.Errorf("invalid PID %s: %v", pidStr, err)
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			return fmt.Errorf("failed to find process with PID %d: %v", pid, err)
		}

		err = process.Kill()
		if err != nil {
			return fmt.Errorf("failed to kill process with PID %d: %v", pid, err)
		}

		fmt.Printf("Successfully killed process %s (PID %d)\n", processName, pid)
	}

	return nil
}
