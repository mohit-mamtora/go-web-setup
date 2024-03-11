package logger

// refrence chatgpt
import (
	"os"
	"sync"
	"time"
)

type rotatingFile struct {
	filename  string
	maxSizeMB int
	file      *os.File
	size      int64
	mu        sync.Mutex
}

func (rf *rotatingFile) Write(p []byte) (n int, err error) {

	if rf.file == nil {
		if err := rf.rotateFile(); err != nil {
			return 0, err
		}
	}

	// Write to the file
	n, err = rf.file.Write(p)
	rf.size += int64(n)

	// Check if file size exceeds the limit
	if rf.size >= int64(rf.maxSizeMB)*1024*1024 {
		if err := rf.rotateFile(); err != nil {
			return n, err
		}
	}

	return n, err
}

func (rf *rotatingFile) Close() error {
	if rf.file != nil {
		return rf.file.Close()
	}
	return nil
}

func (rf *rotatingFile) rotateFile() error {
	defer rf.mu.Unlock()
	rf.mu.Lock()
	if rf.file != nil {
		if err := rf.file.Close(); err != nil {
			return err
		}
	}

	// Rename the current file
	err := os.Rename(rf.filename, rf.filename+"."+time.Now().Format(time.DateTime))
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// Create a new file
	file, err := os.Create(rf.filename)
	if err != nil {
		return err
	}
	rf.file = file
	rf.size = 0

	return nil
}
