package computer

import (
	"fmt"
	"io"
	"log"
	"os"
)

var LogE func(msg string, args ...any)
var LogS func(msg string, args ...any)

func SetupLogging() func() {
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

	LogE = func(msg string, args ...any) {
		wr := io.MultiWriter(execLog, combinedLog, os.Stdout)
		if len(args) == 0 || args[0] == nil {
			if _, err = fmt.Fprint(wr, msg); err != nil {
				log.Fatal("Failed to write to log file:", err)
			}
		}
		if _, err = fmt.Fprintf(wr, msg, args...); err != nil {

		}
	}

	LogS = func(msg string, args ...any) {
		wr := io.MultiWriter(stateLog, combinedLog /*, os.Stderr*/)
		if len(args) == 0 || args[0] == nil {
			if _, err = fmt.Fprintf(wr, msg, args); err != nil {
				log.Fatal("Failed to write to log file:", err)
			}
		} else {
			if _, err = fmt.Fprint(wr, msg); err != nil {
				log.Fatal("Failed to write to log file:", err)
			}
		}
	}
	return func() {
		_ = execLog.Close()
		_ = stateLog.Close()
		_ = combinedLog.Close()
	}
}
