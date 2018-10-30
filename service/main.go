package main

import (
	"github.com/skyismine2010/ATT/plat"
	"github.com/skyismine2010/ATT/service/adaptor"
	"github.com/skyismine2010/ATT/service/gui"
	"github.com/skyismine2010/ATT/service/taskexec"
	"github.com/skyismine2010/ATT/utils"
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
