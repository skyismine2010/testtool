package main

import (
	"github.com/skyismine2010/testtool/plat"
	"github.com/skyismine2010/testtool/service/adaptor"
	"github.com/skyismine2010/testtool/service/config"
	"github.com/skyismine2010/testtool/service/gui"
	"github.com/skyismine2010/testtool/service/taskexec"
	"github.com/skyismine2010/testtool/utils"
	"os"
	"path/filepath"
	"strings"
)

var JobEntry = []*plat.JobEntryReg{
	/*平台Job*/
	&plat.JobEntryReg{plat.TIMER_CTRL_JID, "Timer", plat.INIT_STATUS, 1024, plat.TimerCtrl, nil},

	/*业务job*/
	&plat.JobEntryReg{5001, "TaskExec", plat.INIT_STATUS, 1024, taskexec.TaskExecCtrl, nil},
	&plat.JobEntryReg{5002, "HttpAdaptor", plat.INIT_STATUS, 1024, adaptor.HttpAdaptorCtrl, nil},
	&plat.JobEntryReg{5003, "Gui", plat.INIT_STATUS, 1024, gui.GuiCtrl, nil},
}

func GetCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {
	currentPath := GetCurrentDirectory()
	utils.InitLog(currentPath + "/log/trace.log")
	//gui.InitGUI()

	config.LoadXmlCfg(currentPath + "/cfg/db.xml")

	plat.InitPlat(JobEntry)
}
