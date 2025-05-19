package computer

import (
	"fmt"
	"io"
	"log"
	"os"
)

type CpuLogger interface {
	LogE(msg string, args ...any)
	LogS(msg string, args ...any)
	SetCycle(cycle uint)
	Close() error
}

type MultiCpuLogger struct {
	cycle uint

	executeLog  io.Writer
	stateLog    io.Writer
	combinedLog io.Writer

	close func()
}

func (l *MultiCpuLogger) Close() error {
	l.close()
	return nil
}

func (l *MultiCpuLogger) log(wr io.Writer, m string, args ...any) {
	msg := "  " + m
	if len(args) != 0 && args[0] != nil {
		if _, err := fmt.Fprintf(wr, msg, args...); err != nil {
			log.Fatal("Failed to write to log file:", err)
		}
	} else {
		if _, err := fmt.Fprint(wr, msg); err != nil {
			log.Fatal("Failed to write to log file:", err)
		}
	}
}
func (l *MultiCpuLogger) LogE(msg string, args ...any) {
	l.log(l.executeLog, msg, args...)
	l.log(os.Stdout, msg, args...)

}

func (l *MultiCpuLogger) LogS(msg string, args ...any) {
	l.log(l.stateLog, msg, args...)
	l.log(os.Stdout, msg, args...)
}

func (l *MultiCpuLogger) SetCycle(cycle uint) {
	//if l.cycle != 0 {
	//	_, _ = fmt.Fprintln(l.stateLog, "}")
	//	_, _ = fmt.Fprintln(l.executeLog, "}")
	//	_, _ = fmt.Fprintln(os.Stdout, "}")
	//
	//}
	l.cycle = cycle
	_, _ = fmt.Fprintf(l.stateLog, "%d\n", l.cycle)
	_, _ = fmt.Fprintf(l.executeLog, "%d\n", l.cycle)
	_, _ = fmt.Fprintf(os.Stdout, "%d\n", l.cycle)
}

func SetupLogging() MultiCpuLogger {
	folder, ok := os.LookupEnv("LOG_FOLDER")
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "LOG_FOLDER environment variable not set")
	}

	execLog, err := os.OpenFile(folder+"exec.log" /*os.O_APPEND|*/, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	stateLog, err := os.OpenFile(folder+"state.log" /*os.O_APPEND|*/, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	combinedLog, err := os.OpenFile(folder+"combined.log" /* os.O_APPEND|*/, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	logger := MultiCpuLogger{
		cycle:       0,
		executeLog:  execLog,
		stateLog:    stateLog,
		combinedLog: combinedLog,
		close: func() {
			_ = execLog.Close()
			_ = stateLog.Close()
			_ = combinedLog.Close()
		},
	}

	return logger
}
