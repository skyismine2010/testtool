package plat

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

const (
	INIT_STATUS = iota
	WORK_STATUS
	UNKNOWN_STATUS
)

type JobMsg struct {
	EventType int32
	msg       []byte
}

type JobEntryReg struct {
	Jid          int32
	Name         string
	Status       int32
	BuffSize     int32
	ServiceEntry func(msg *JobMsg, status int32) error
	//Lock         sync.Mutex
	MsgChan chan JobMsg
}

var g_JobMap map[int32](*JobEntryReg)
var g_JobEntry []*JobEntryReg

func InitPlat(entrys []*JobEntryReg) error {
	g_JobEntry = entrys
	g_JobMap = make(map[int32](*JobEntryReg), len(entrys))

	for _, entry := range entrys {
		entry.MsgChan = make(chan JobMsg, entry.BuffSize)
		g_JobMap[entry.Jid] = entry
	}

	for _, entry := range entrys {
		go jobEntry(entry)
	}

	log.Printf("ALL Job is running now!")

	msg := JobMsg{
		POWER_ON_EVENT,
		nil,
	}
	log.Printf("Sending power on msg to all job begin!")
	SendAsyncBroadcastMsg(&msg)
	log.Printf("Sending power on msg to all job finish!")
	time.Sleep(10000 * time.Second)
	log.Printf("Finish....")

	return nil
}

func jobEntry(job *JobEntryReg) {
	for {
		msg := <-job.MsgChan
		err := job.ServiceEntry(&msg, job.Status)
		if err != nil {
			log.Printf("service %v get erorr, error: %v", job.Name, err)
		}
	}
}

func SendAsyncMsgByJid(jId int32, msg *JobMsg) {
	entry, ok := g_JobMap[jId]
	if !ok {
		log.Printf("Can't find error Jid=%d", jId)
		return
	}
	entry.MsgChan <- *msg
}

func SendAsyncBroadcastMsg(msg *JobMsg) error {
	for _, entry := range g_JobEntry {
		entry.MsgChan <- *msg
	}
	return nil
}

func GetJobStatus(jid int32) int32 {
	entry, ok := g_JobMap[jid]
	if !ok {
		return UNKNOWN_STATUS
	}
	return entry.Status
}

func SetJobStatus(jid int32, status int32) {
	entry, ok := g_JobMap[jid]
	if !ok {
		log.Printf("Can't find Jid = %d", jid)
		return
	}
	entry.Status = status
}

func sendMsg2TimerCtrl(event *JobTimerEvent) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, *event)
	if err != nil {
		log.Printf("Can't write to Buffer. err=%v", err)
	}
	log.Printf("Write to Buffer, buff=% x", buff.Bytes())
	jobMsg := JobMsg{CREATE_TIMER_EVENT, buff.Bytes()}
	SendAsyncMsgByJid(TIMER_CTRL_JID, &jobMsg)

}

func SetRelativeTimer(timerID int32, during time.Duration, jId int32) {
	event := JobTimerEvent{timerID, JobTimer{
		jId, RELATIVE_TIMER, during, during}}
	sendMsg2TimerCtrl(&event)
}

func SetLoopTimer(timerID int32, during time.Duration, jId int32) {
	event := JobTimerEvent{timerID, JobTimer{
		jId, LOOP_TIMER, during, during}}
	sendMsg2TimerCtrl(&event)

}

func KillTimer(timerID int) {
}
