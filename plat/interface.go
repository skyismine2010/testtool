package plat

import (
	"log"
	"time"
)

const (
	JOB_INIT = iota
	JOB_WORK
)

const (
	POWER_ON_MSG = iota
	POWER_OFF_MSG
	SERVICE_MSG
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
		log.Printf(" %v create chan %v", entry.Name, entry.MsgChan)
		g_JobMap[entry.Jid] = entry
	}

	for _, entry := range entrys {
		log.Printf("go here job = %v chan %v", entry.Name, entry.MsgChan)
		go jobEntry(entry)
	}

	log.Printf("ALL Job is running now!")

	msg := JobMsg{
		POWER_ON_MSG,
		0,
		nil,
	}
	log.Printf("Sending power on msg to all job begin!")
	SendAsyncBroadcastMsg(&msg)
	log.Printf("Sending power on msg to all job finish!")
	time.Sleep(1 * time.Second)
	log.Printf("Finish....")

	return nil
}

func jobEntry(job *JobEntryReg) {
	for {
		log.Printf("job %v begin to loop, chan= %v ", job.Name, job.MsgChan)
		msg := <-job.MsgChan
		log.Printf("service %v  get msg = %v", job.Name, msg.msgType)

		err := job.ServiceEntry(&msg, job.Jid)
		if err != nil {
			log.Printf("service %v get erorr, error: %v", job.Name, err)
		}
	}
}

func SendAsyncMsgByName(jobName string, msg *JobMsg) error {
	return nil
}

func SendAsyncBroadcastMsg(msg *JobMsg) error {
	for _, entry := range g_JobEntry {
		log.Printf("send msg to %v, chan =%v", entry.Name, entry.MsgChan)
		entry.MsgChan <- *msg
	}
	return nil
}

func GetStatus(jid int) int {
	_, ok := g_JobMap[jid]

}
