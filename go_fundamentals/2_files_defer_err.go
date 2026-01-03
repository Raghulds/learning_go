package gofundamentals

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"time"
)

func KillServer() {
	time.Sleep(3 * time.Second)
	err := killServer("server.pid")
	if err != nil {
		fmt.Println("Error:", err)
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}
		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("close", "file", pidFile)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %s", pidFile, err)
	}

	slog.Info("killing", "pid", pid)
	if err := os.Remove(pidFile); err != nil {
		slog.Warn("delete", "file", pidFile, "error", err)
	}

	return nil
}

/*
defer
------
happens when the func exists
works at func level
executed in reverse order (LIFO)

Try to aquire a resource, check for error, defer release
*/
