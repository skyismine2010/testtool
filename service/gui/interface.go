package gui

import "log"
import "github.com/skyismine2010/ATT/plat"

func initGUI(msg *plat.JobMsg, jid int) error {
	log.Printf("receive ")

	return nil
}

func GuiCtrl(msg *plat.JobMsg, jid int) error {
	status := plat.GetStatus(jid)
	switch status {
	case plat.JOB_INIT:
		initGUI(msg, jid)
	}

	return nil
}
