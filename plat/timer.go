package plat

//todo: 这里准备使用一个container/heap 来实现比较合理
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

const (
	precision = 10 * time.Millisecond
)

const (
	RELATIVE_TIMER = iota
	LOOP_TIMER
)

type JobTimer struct {
	Jid         int32
	TimerType   byte
	RestDuring  time.Duration
	TotalDuring time.Duration
}

type JobTimerEvent struct {
	TimerId int32
	Timer   JobTimer
}

var timerMap map[int32]*JobTimer

func timerLoopDemo(during time.Duration) {
	ticker := time.NewTicker(during)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			msg := JobMsg{
				PLAT_LOOP_EVENT,
				nil,
			}
			SendAsyncMsgByJid(TIMER_CTRL_JID, &msg)
		}
	}

}

func TimerInit() {
	timerMap = make(map[int32]*JobTimer) //todo: 这里要换成一个最小堆来实现
	log.Printf("Timer Ctrl receive Init message, prepared to power on.")
	SetJobStatus(TIMER_CTRL_JID, WORK_STATUS)
	go timerLoopDemo(precision)
}

func sendTimerOutEvent(timerid int32, timer *JobTimer) {
	msg := JobMsg{timerid, nil}
	SendAsyncMsgByJid(timer.Jid, &msg)
}

func insert2TimerMap(msg *JobMsg) error {
	var event JobTimerEvent
	reader := bytes.NewReader(msg.MsgBuff)
	binary.Read(reader, binary.BigEndian, &event)
	timerId := event.TimerId

	if _, ok := timerMap[timerId]; ok {
		return fmt.Errorf("timerid=%d already exist.", timerId)
	}
	timerMap[timerId] = &event.Timer
	return nil
}

func removeFromTimeMap(msg *JobMsg) error {
	var event JobTimerEvent
	reader := bytes.NewReader(msg.MsgBuff)
	binary.Read(reader, binary.LittleEndian, &event)
	timerId := event.TimerId
	if _, ok := timerMap[timerId]; !ok {
		return fmt.Errorf("timerid=%d does't exist.", timerId)
	}

	delete(timerMap, timerId)
	return nil
}

func TimerCtrl(msg *JobMsg, status int32) error {
	switch status {
	case INIT_STATUS:
		TimerInit()
	case WORK_STATUS:
		{
			switch msg.EventType {
			case PLAT_LOOP_EVENT:
				for timerId, jobTimer := range timerMap {
					jobTimer.RestDuring -= precision
					//log.Printf("scane timer=%d, rest during=%v", TimerId, Timer.RestDuring)
					if jobTimer.RestDuring <= 0 {
						sendTimerOutEvent(timerId, jobTimer)
						if jobTimer.TimerType == LOOP_TIMER {
							jobTimer.RestDuring = jobTimer.TotalDuring
						} else {
							delete(timerMap, timerId)
						}

					}
				}

			case CREATE_TIMER_EVENT:
				err := insert2TimerMap(msg)
				if err != nil {
					log.Printf("Create Timer failed, err=%v", err)
					return err
				}
			case KILL_TIMER_EVENT:
				err := removeFromTimeMap(msg)
				if err != nil {
					log.Printf("kill Timer failed, err=%v", err)
					return err
				}
			}
		}
	}

	return nil
}
