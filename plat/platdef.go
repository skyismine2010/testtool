package plat

const (
	POWER_ON_EVENT = iota
	PLAT_LOOP_EVENT
	CREATE_TIMER_EVENT
	KILL_TIMER_EVENT
	CFG_CHG_EVENT
)

const (
	TIMER_1 = 1000 + iota
	TIMER_2
	TIMER_3
	TIMER_4
)

const (
	TIMER_CTRL_JID = 1 + iota
)
