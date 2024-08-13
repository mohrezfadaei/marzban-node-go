package xray

import (
	"bufio"
	"fmt"
	"os/exec"
	"sync"
)

type XRayCore struct {
	executablePath string
	assetsPath     string
	started        bool
	connected      bool
	sessionID      string
	clientIP       string
	mutex          sync.Mutex
}

func NewXRayCore(executablePath, assetsPath string) *XRayCore {
	return &XRayCore{
		executablePath: executablePath,
		assetsPath:     assetsPath,
	}
}

func (x *XRayCore) GetVersion() string {
	cmd := exec.Command(x.executablePath, "version")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(output)
}

func (x *XRayCore) Start(config *XRayConfig) error {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	if x.started {
		return fmt.Errorf("xray is already started")
	}

	cmd := exec.Command(x.executablePath, "run", "-config", "stdin:")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	if _, err := stdin.Write([]byte(config.ToJson())); err != nil {
		return err
	}

	if err := stdin.Close(); err != nil {
		return err
	}

	x.started = true
	return nil
}

func (x *XRayCore) Stop() error {
	x.mutex.Lock()
	defer x.mutex.Unlock()

	if !x.started {
		return nil
	}

	cmd := exec.Command("pkill", "-f", x.executablePath)
	if err := cmd.Run(); err != nil {
		return err
	}

	x.started = false
	return nil
}

func (x *XRayCore) Restart(config *XRayConfig) error {
	if err := x.Stop(); err != nil {
		return err
	}
	return x.Start(config)
}

func (x *XRayCore) GetLogs() chan string {
	logs := make(chan string)
	go func() {
		cmd := exec.Command("tail", "-f", "/var/log/xray.log")
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			close(logs)
			return
		}

		if err := cmd.Start(); err != nil {
			close(logs)
			return
		}

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			logs <- scanner.Text()
		}

		if err := cmd.Wait(); err != nil {
			close(logs)
		}
	}()
	return logs
}

func (x *XRayCore) Connected() bool {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	return x.connected
}

func (x *XRayCore) Started() bool {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	return x.started
}

func (x *XRayCore) Connect(sessionID, clientIP string) {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	x.sessionID = sessionID
	x.clientIP = clientIP
	x.connected = true
}

func (x *XRayCore) Disconnect() {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	x.sessionID = ""
	x.clientIP = ""
	x.connected = false
}

func (x *XRayCore) MatchSessionID(sessionID string) bool {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	return x.sessionID == sessionID
}

func (x *XRayCore) ClientIP() string {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	return x.clientIP
}

func (x *XRayCore) SessionID() string {
	x.mutex.Lock()
	defer x.mutex.Unlock()
	return x.sessionID
}
