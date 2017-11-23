package alog

import (
	//"errors"
	"fmt"
	"log"
	"os"
	"time"
)

type slog struct {
	fHourLog *os.File // LOG输出文件
	fDayLog  *os.File // LOG输出文件
}

func (l *slog) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	if l.fHourLog != nil {
		if n, err = l.fHourLog.Write(p); err != nil {
			return n, err
		}
	}
	if l.fDayLog != nil {
		if n, err = l.fDayLog.Write(p); err != nil {
			return n, err
		}
	}
	return n, err
}

func Startup(dir string) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// Create Dir
	if dir != "" {
		if err := os.MkdirAll(dir, 0777); err != nil {
			log.Fatalln(err)
		}
	}
	sl := new(slog)
	created := make(chan bool)
	go func() {
		var oldHour = -1
		var oldDay = -1
		for {
			t := time.Now()
			curHour := t.Hour()
			curDay := t.Day()
			if oldHour != curHour {
				oldHour = curHour
				// 创建小时日志
				{
					fileName := dir + "/" + t.Format("2006-01-02_15.log")
					log.Println("Create log file:", fileName)
					if file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666); err != nil {
						log.Println(err)
					} else {
						old := sl.fHourLog
						sl.fHourLog = file
						if old != nil {
							old.Close()
						}
					}
				}
				// 创建当日日志
				if oldDay != curDay {
					oldDay = curDay
					fileName := dir + "/today.log"
					log.Println("Create log file:", fileName)
					if file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666); err != nil {
						log.Println(err)
					} else {
						old := sl.fDayLog
						sl.fDayLog = file
						if old != nil {
							old.Close()
						}
					}
				}
				if created != nil {
					close(created)
				}
				// 等到下一小时前一秒
				time.Sleep(time.Duration(59-t.Minute())*time.Minute +
					time.Duration(59-t.Second())*time.Second)
			} else {
				time.Sleep(1 * time.Millisecond)
			}
		}
	}()
	<-created
	created = nil
	log.SetOutput(sl)
}
