package adaptor

import (
	"github.com/skyismine2010/testtool/plat"
	"github.com/skyismine2010/testtool/service/comm"
	"log"
	"time"
)

func HttpAdaptorCtrl(msg *plat.JobMsg, status int32) error {
	jid := int32(comm.HTTP_ADAPTOR_JID)
	switch status {
	case plat.INIT_STATUS:
		time.Sleep(3 * time.Second)
		plat.SetLoopTimer(plat.TIMER_1, 3*time.Second, jid)
		plat.SetJobStatus(jid, plat.WORK_STATUS)

	case plat.WORK_STATUS:
		{
			switch msg.EventType {
			case plat.TIMER_1:
				log.Printf("Http Adatpor Get Timer Event.")

			}
		}

	}
	return nil
}
