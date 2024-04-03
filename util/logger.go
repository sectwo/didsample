package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Enterlog(s string) { log.Println("[INF]=====================<entering>: ", s) }
func Leavelog(s string) { log.Println("[INF]=====================<leaving>: ", s) }

// func errlog(e error)     { log.Println("[ERR]: ", e.Error()) }
// func errString(s string)   { log.Println("[ERR]: ", s) }
func errlog(s, c, r string, e interface{}) {
	Enterlog(s)
	log.Println("[ERR]: Code:", c, ", Msg:", r)
	log.Println("[ERR]:", e)
	Leavelog(s)
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
func createDir(path string) bool {
	// check
	if _, err := os.Stat(path); err != nil {
		err := os.MkdirAll(path, 0711)
		if err != nil {
			errlog("err", "code", "reason", err)
			return false
		}
	}
	return true
}
func baseName() string {
	return filepath.Base(os.Args[0])
}
func createLogDir() string {
	logPath := fmt.Sprintf("%s/log", getCurrentDirectory())
	if createDir(logPath) {
		return logPath
	}
	return ""
}

type myWriter struct {
	createdDate string
	file        *os.File
}

func (t *myWriter) Write(p []byte) (n int, err error) {
	tt := string(p[5:10])
	if t.createdDate != tt {
		if err := t.rotateFile(time.Now()); err != nil {
			log.Printf("[ERRS] %s\n", err.Error())
		}
	}
	return t.file.Write(p)
}
func (t *myWriter) rotateFile(now time.Time) error {
	t.createdDate = fmt.Sprintf("%02d/%02d", now.Month(), now.Day())
	logDir := createLogDir()
	if len(logDir) != 0 {
		//baseName_YYYYMMDD.log
		///path/to/file/<prefix>YYYYMMDD<suffix>
		filePath := fmt.Sprintf("%s/%s%s%s", logDir, baseName()+"_", now.Format("20060102"), ".log")
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			errlog("err", "code", "reason", err)
			return err
		}
		if t.file != nil {
			t.file.Close()
		}
		t.file = file
	}
	return nil
}

// log.Ldate | log.Lmicroseconds
// ./log/baseName_YYYYMMDD.log
func logSetup() {
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	log.SetOutput(io.MultiWriter(&myWriter{} /*os.Stderr,*/, os.Stdout))
}
