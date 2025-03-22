package log

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

type Logger struct {
	name string
	hash []byte
}

func colorize(message string, color string) string {
	return fmt.Sprintf("%v%v%v", color, message, Reset)
}

func GetLogger(name string) Logger {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(time.Now().UnixNano()))
	hash := sha256.Sum256(b)
	return Logger{
		name: name,
		hash: hash[:4],
	}
}

func (l *Logger) Debug(message string, a ...any) {
	level := "DEBUG   "
	color := Cyan
	l.log(fmt.Sprintf(message, a...), level, color)
}

func (l *Logger) Info(message string, a ...any) {
	level := "INFO    "
	color := Green
	l.log(fmt.Sprintf(message, a...), level, color)
}

func (l *Logger) Warning(message string, a ...any) {
	level := "WARNING "
	color := Yellow
	l.log(fmt.Sprintf(message, a...), level, color)
}

func (l *Logger) Error(message string, a ...any) {
	level := "ERROR   "
	color := Red
	l.log(fmt.Sprintf(message, a...), level, color)
}

func (l *Logger) Critical(message string, a ...any) {
	level := "CRITICAL"
	color := Purple
	l.log(fmt.Sprintf(message, a...), level, color)
}

func (l *Logger) log(message, level, color string) {
	name := colorize(l.name, White)
	levelOffset := 20 - len(l.name)
	level = strings.Repeat(" ", levelOffset) + level
	level = colorize(level, color)
	hash := colorize(fmt.Sprintf("%x", l.hash), White)
	now := time.Now()
	timeFormat := fmt.Sprintf(
		"%02d:%02d:%02d.%03d",
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond()/1000000,
	)
	fmt.Printf("%v %v %v %v: %v\n", timeFormat, hash, name, level, message)
}
