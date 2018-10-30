package main

import (
	"github.com/skyismine2010/testtool/plat"
	"github.com/skyismine2010/testtool/service/adaptor"
	"github.com/skyismine2010/testtool/service/gui"
	"github.com/skyismine2010/testtool/service/taskexec"
	"github.com/skyismine2010/testtool/utils"
)

var JobEntry = []*plat.JobEntryReg{
	&plat.JobEntryReg{1, "TaskExec", plat.JOB_INIT, 1024, taskexec.TaskExecCtrl, nil},
	&plat.JobEntryReg{2, "HttpAdaptor", plat.JOB_INIT, 1024, adaptor.HttpAdaptorCtrl, nil},
	&plat.JobEntryReg{3, "Gui", plat.JOB_INIT, 1024, gui.GuiCtrl, nil},
}

func main() {
	utils.InitLog("../log/trace.log")
	//gui.InitGUI()
	plat.InitPlat(JobEntry)
}
