package plat

//todo: 这里准备使用一个container/heap 来实现比较合理
import (
	"fmt"
	"log"
	"time"
	"unsafe"
)

const (
	precision = 10 * time.Millisecond
)

const (
	RELATIVE_TIMER = iota
	LOOP_TIMER
)

type JobTimer struct {
	jid         int
	timerType   byte
	restDuring  time.Duration
	totalDuring time.Duration
}

type JobTimerEvent struct {
	timerId  int
	jobTimer JobTimer
}

var timerMap map[int]*JobTimer

func timerLoopDemo(during time.Duration) {
	ticker := time.NewTicker(during)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			msg := JobMsg{
				PLAT_LOOP_EVENT,
				0,
				nil,
			}
			SendAsyncMsgByJid(TIMER_CTRL_JID, &msg)
		}
	}

}

func TimerInit() {
	timerMap = make(map[int]*JobTimer) //todo: 这里要换成一个最小堆来实现
	log.Printf("Timer Ctrl receive Init message, prepared to power on.")
	SetJobStatus(TIMER_CTRL_JID, WORK_STATUS)
	go timerLoopDemo(precision)
}

func sendTimerOutEvent(timerid int, timer *JobTimer) {
	msg := JobMsg{timerid, 0, nil}
	SendAsyncMsgByJid(timer.jid, &msg)
}

func createTimer(msg *JobMsg) error {

	pJobTimerEvent := *(*JobTimerEvent)(unsafe.Pointer(msg))
	timerId := pJobTimerEvent.timerId
	if _, ok := timerMap[timerId]; ok {
		return fmt.Errorf("timerid=%d already exist.", timerId)
	}
	timerMap[timerId] = &pJobTimerEvent.jobTimer
	return nil
}

func killTimer(msg *JobMsg) error {
	pJobTimerEvent := *(**JobTimerEvent)(unsafe.Pointer(msg))
	timerId := pJobTimerEvent.timerId
	if _, ok := timerMap[timerId]; !ok {
		return fmt.Errorf("timerid=%d does't exist.", timerId)
	}

	delete(timerMap, timerId)
	return nil
}

func TimerCtrl(msg *JobMsg, status int) error {
	switch status {
	case INIT_STATUS:
		TimerInit()
	case WORK_STATUS:
		{
			switch msg.EventType {
			case PLAT_LOOP_EVENT:
				for timerId, jobTimer := range timerMap {
					jobTimer.restDuring -= precision
					log.Printf("scane timer=%d, rest during=%v", timerId, jobTimer.restDuring)
					if jobTimer.restDuring <= 0 {
						sendTimerOutEvent(timerId, jobTimer)
						if jobTimer.timerType == LOOP_TIMER {
							jobTimer.restDuring = jobTimer.totalDuring
						} else {
							delete(timerMap, timerId)
						}

					}
				}

			case CREATE_TIMER_EVENT:
				err := createTimer(msg)
				if err != nil {
					log.Printf("Create Timer failed, err=%v", err)
					return err
				}
			case KILL_TIMER_EVENT:
				err := killTimer(msg)
				if err != nil {
					log.Printf("kill Timer failed, err=%v", err)
					return err
				}
			}
		}
	}

	return nil
}
