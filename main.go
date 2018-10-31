package main

import (
	"github.com/skyismine2010/testtool/plat"
	"github.com/skyismine2010/testtool/service/adaptor"
	"github.com/skyismine2010/testtool/service/gui"
	"github.com/skyismine2010/testtool/service/taskexec"
	"github.com/skyismine2010/testtool/utils"
)

var JobEntry = []*plat.JobEntryReg{
	/*平台Job*/
	&plat.JobEntryReg{plat.TIMER_CTRL_JID, "Timer", plat.INIT_STATUS, 1024, plat.TimerCtrl, nil},

	/*业务job*/
	&plat.JobEntryReg{5001, "TaskExec", plat.INIT_STATUS, 1024, taskexec.TaskExecCtrl, nil},
	&plat.JobEntryReg{5002, "HttpAdaptor", plat.INIT_STATUS, 1024, adaptor.HttpAdaptorCtrl, nil},
	&plat.JobEntryReg{5003, "Gui", plat.INIT_STATUS, 1024, gui.GuiCtrl, nil},
}

func main() {
	utils.InitLog("../log/trace.log")
	//gui.InitGUI()
	plat.InitPlat(JobEntry)
}

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//
//	input := make(chan interface{})
//
//	//producer - produce the messages
//	go func() {
//		for i := 0; i < 5; i++ {
//			input <- i
//		}
//		input <- "hello, world"
//	}()
//
//	t1 := time.NewTimer(time.Second * 5)
//	t2 := time.NewTimer(time.Second * 10)
//
//	for {
//		select {
//		//consumer - consume the messages
//		case msg := <-input:
//			fmt.Println(msg)
//
//		case <-t1.C:
//			println("5s timer")
//			t1.Reset(time.Second * 5)
//
//		case <-t2.C:
//			println("10s timer")
//			t2.Reset(time.Second * 10)
//		}
//	}
//}
