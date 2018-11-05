package adaptor

import (
	"github.com/skyismine2010/testtool/plat"
	"github.com/skyismine2010/testtool/service/comm"
	"log"
	"time"
)

func httpCfgChgHandle(msg []byte) {

}

func HttpAdaptorCtrl(msg *plat.JobMsg, status int32) error {
	jid := int32(comm.HTTP_ADAPTOR_JID)
	switch status {
	case plat.INIT_STATUS:
		time.Sleep(1 * time.Second)
		plat.SetLoopTimer(plat.TIMER_1, 1*time.Second, jid)
		plat.SetJobStatus(jid, plat.WORK_STATUS)

	case plat.WORK_STATUS:
		{
			switch msg.EventType {
			case plat.TIMER_1:
				log.Printf("Http Adatpor Get Timer Event.")
			case plat.CFG_CHG_EVENT:
				log.Printf("Http adaptor receive config change event.")
				httpCfgChgHandle(msg.MsgBuff)
			}
		}
	}
	return nil
}
