package plat

import (
	"fmt"
	"log"
	"time"
)

const (
	INIT_STATUS = iota
	WORK_STATUS
	UNKNOWN_STATUS
)

type JobMsg struct {
	msgType int
	msgLen  int
	msg     []byte
}

type JobEntryReg struct {
	Jid          int
	Name         string
	Status       int
	BuffSize     int
	ServiceEntry func(msg *JobMsg, jid int) error
	//Lock         sync.Mutex
	MsgChan chan JobMsg
}

var g_JobMap map[int](*JobEntryReg)
var g_JobEntry []*JobEntryReg

func InitPlat(entrys []*JobEntryReg) error {
	g_JobEntry = entrys
	g_JobMap = make(map[int](*JobEntryReg), len(entrys))

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
		0,
		nil,
	}
	log.Printf("Sending power on msg to all job begin!")
	SendAsyncBroadcastMsg(&msg)
	log.Printf("Sending power on msg to all job finish!")
	time.Sleep(100 * time.Second)
	log.Printf("Finish....")

	return nil
}

func jobEntry(job *JobEntryReg) {
	for {
		msg := <-job.MsgChan
		err := job.ServiceEntry(&msg, job.Jid)
		if err != nil {
			log.Printf("service %v get erorr, error: %v", job.Name, err)
		}
	}
}

func SendAsyncMsgByJid(jId int, msg *JobMsg) error {
	entry, ok := g_JobMap[jId]
	if !ok {
		return fmt.Errorf("Can't find error")
	}
	entry.MsgChan <- *msg

	return nil
}

func SendAsyncBroadcastMsg(msg *JobMsg) error {
	for _, entry := range g_JobEntry {
		log.Printf("send msg to %v, chan =%v", entry.Name, entry.MsgChan)
		entry.MsgChan <- *msg
	}
	return nil
}

func GetJobStatus(jid int) int {
	entry, ok := g_JobMap[jid]
	if !ok {
		return UNKNOWN_STATUS
	}
	return entry.Status
}

func SetJobStatus(jid int, status int) {
	entry, ok := g_JobMap[jid]
	if !ok {
		log.Printf("Can't find jid = %d", jid)
		return
	}
	entry.Status = status
}
