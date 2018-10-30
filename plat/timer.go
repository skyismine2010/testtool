package plat

//todo: 这里准备使用一个container/heap 来实现比较合理
import (
	"log"
	"time"
)

const (
	selfJId = 1
)

type jobTimer struct {
	jid    int
	during time.Duration
}

var timerMap map[int]jobTimer

func SetRelativeTimer() {

}

func SetLoopTimer() {
}

func KillTimer() {

}

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
			SendAsyncMsgByJid(selfJId, &msg)
		}
	}

}

func TimerInit() {
	timerMap = make(map[int]jobTimer)
	log.Printf("Timer Ctrl receive Init message, prepared to power on.")
	SetJobStatus(selfJId, WORK_STATUS)
	go timerLoopDemo(100 * time.Millisecond)
}

func TimerCtrl(msg *JobMsg, status int) error {
	switch GetJobStatus(selfJId) {
	case INIT_STATUS:
		TimerInit()
	case WORK_STATUS:
		{
			switch msg.msgType {
			case PLAT_LOOP_EVENT:

				log.Printf("go here now!")
			case CREATE_TIMER_EVENT:
				log.Printf("go here now!")
			case KILL_TIMER_EVENT:
				log.Printf("go here now!")
			}
		}
	}

	return nil
}
