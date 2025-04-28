package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// Declare a global variable for the logger instance
var logInstance *Logger
var once sync.Once

// Logger struct to hold log configuration
type Logger struct {
	*log.Logger
	file *os.File
	mu   sync.Mutex
}

// NewLogger initializes a new logger instance with the desired settings and returns it
func NewLogger(debugEnabled bool) (*Logger, error) {
	// Use sync.Once to ensure the logger is only initialized once
	once.Do(func() {
		// Open the log file (or create it if it doesn't exist)
		file, err := os.OpenFile("provider-debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		// Create the global logger instance
		logInstance = &Logger{
			Logger: log.New(file, "Info : ", log.Ldate|log.Ltime|log.Lshortfile),
			file:   file,
		}

		// Write initial message to the log if debugging is enabled
		if debugEnabled {
			logInstance.Println("Logger initialized")
		}
	})

	// Return the global logger instance
	if logInstance == nil {
		return nil, fmt.Errorf("logger not initialized")
	}

	return logInstance, nil
}

// Debug logs a debug message, only if DebugEnabled is true
func Debug(message string) {
	if logInstance != nil {
		logInstance.mu.Lock()
		defer logInstance.mu.Unlock()
		logInstance.Println("DEBUG: ", message)
	}
}

// Info logs an info message
func Info(message string) {
	if logInstance != nil {
		logInstance.mu.Lock()
		defer logInstance.mu.Unlock()
		logInstance.Println("INFO: ", message)
	}
}

// Error logs an error message
func Error(message string) {
	if logInstance != nil {
		logInstance.mu.Lock()
		defer logInstance.mu.Unlock()
		logInstance.Println("ERROR: ", message)
	}
}

// Close closes the log file
func Close() {
	if logInstance != nil && logInstance.file != nil {
		logInstance.file.Close()
	}
}
