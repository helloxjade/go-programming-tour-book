package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

// 1、 日志分级： 预定义应用日志的level和Fields的具体类型
type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	}
	return ""
}

// 2、日志标准化
//对日志的实例初始化和标准化参数进行绑定
type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) Clone() *Logger {
	nl := *l
	return &nl
}

//设置日志公共字段
func (l *Logger) WithFields(f Fields) *Logger {
	ll := l.Clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}

	for k, v := range f {
		ll.fields[k] = v
	}
	return ll
}

//设置日志上下文属性
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.Clone()
	ll.ctx = ctx
	return ll
}

//设置当前某一层调用栈的信息（程序计数器、文件信息和行号）
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.Clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

// 设置当前整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		s := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, s)
		if !more {
			break
		}
	}
	ll := l.Clone()
	ll.callers = callers
	return ll
}

//日志格式化和输出
func (l *Logger) JSONFormat(levle Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = levle.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}
func (l *Logger) Output(level Level, message string) {
	body, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(body)
	switch level {
	case LevelDebug:
		l.newLogger.Println(content)
	case LevelInfo:
		l.newLogger.Println(content)
	case LevelError:
		l.newLogger.Println(content)
	case LevelFatal:
		l.newLogger.Println(content)
	case LevelPanic:
		l.newLogger.Println(content)
	}
}

//日志分级输出
func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}
func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelFatal, fmt.Sprintf(format, v...))
}
func (l *Logger) Error(v ...interface{}) {
	l.Output(LevelError, fmt.Sprint(v...))
}
func (l *Logger) Errorf(format string,v ...interface{}) {
	l.Output(LevelError, fmt.Sprintf(format,v...))
}
func (l *Logger) Warn(v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprint(v...))
}
func (l *Logger) Warnf(format string,v ...interface{}) {
	l.Output(LevelWarn, fmt.Sprintf(format,v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprint(v...))
}
func (l *Logger) Debugf(format string,v ...interface{}) {
	l.Output(LevelDebug, fmt.Sprintf(format,v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprint(v...))
}
func (l *Logger) Panicf(format string,v ...interface{}) {
	l.Output(LevelPanic, fmt.Sprintf(format,v...))
}