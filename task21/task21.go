package main

import "fmt"

// клиентский интерфейс / Target
type Logger interface {
	Log(message string)
}

// существующий (несовместимый) интерфейс "адаптируемого" кода / Adaptee
type LegacyLogWriter struct{}

func (w *LegacyLogWriter) WriteLog(msg string) {
	fmt.Printf("[Legacy Log] %s\n", msg)
}

// Adapter
type LoggerAdapter struct {
	writer *LegacyLogWriter
}

func NewLoggerAdapter(w *LegacyLogWriter) Logger {
	return &LoggerAdapter{writer: w}
}

func (l *LoggerAdapter) Log(message string) {
	l.writer.WriteLog(message)
}

func ProcessSomething(l Logger) {
	l.Log("Starting process")
	// какая‑то работа...
	l.Log("Finished process")
}

func main() {
	// У нас есть только LegacyLogWriter, но клиент хочет Logger
	legacy := &LegacyLogWriter{}

	// Создаём адаптер
	logger := NewLoggerAdapter(legacy)
	ProcessSomething(logger)
}
