package gui

import (
	"github.com/skyismine2010/testtool/plat"
)

func initGUI(msg *plat.JobMsg, jid int) error {
	return nil
}

func GuiCtrl(msg *plat.JobMsg, jid int) error {
	status := plat.GetJobStatus(jid)
	switch status {
	case plat.INIT_STATUS:
		initGUI(msg, jid)
	}

	return nil
}
