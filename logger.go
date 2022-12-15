package logger

import "io"

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type Logger struct {
	out       io.Writer
	level     LogLevel
	formatter Formatter
}

// NewLogger creates a new Logger instance
func NewLogger(out io.Writer, level LogLevel) *Logger {
	return &Logger{
		out:   out,
		level: level,
	}
}

// SetFormatter sets the formatter used by the logger
func (l *Logger) SetFormatter(formatter Formatter) {
	l.formatter = formatter
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	if l.level <= DebugLevel {
		l.log(DebugLevel, msg)
	}
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	if l.level <= InfoLevel {
		l.log(InfoLevel, msg)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	if l.level <= WarnLevel {
		l.log(WarnLevel, msg)
	}
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	if l.level <= ErrorLevel {
		l.log(ErrorLevel, msg)
	}
}

// Panic logs a panic message
func (l *Logger) Panic(msg string) {
	if l.level <= PanicLevel {
		l.log(PanicLevel, msg)
	}
}

// Fatal logs a fatal message
func (l *Logger) Fatal(msg string) {
	if l.level <= FatalLevel {
		l.log(FatalLevel, msg)
	}
}

// log formats and writes a log message to the logger's output
func (l *Logger) log(level LogLevel, msg string) {
	if l.formatter != nil {
		msg = l.formatter.Format(level, msg)
	}
	l.out.Write([]byte(msg))
}

// Formatter is the interface for log message formatters
type Formatter interface {
	Format(level LogLevel, msg string) string
}
