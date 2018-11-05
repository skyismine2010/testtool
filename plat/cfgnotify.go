package plat

import "sync"

type watchers struct {
	Jids []int
}

var mapLock sync.Mutex
var registerCfgMap map[string]watchers

func init() {
	registerCfgMap = make(map[string]watchers)
}

func SendCfgNotify(fileName string) {

}

//这个函数应该支持线程安全调用
func RegisterWatchCfgFile(jid int, cfgTblName string) {
	mapLock.Lock()
	if _, ok := registerCfgMap[cfgTblName]; !ok {
		registerCfgMap[cfgTblName] = watchers{make([]int, 64)}
	}
	lists := registerCfgMap[cfgTblName].Jids
	lists = append(lists, jid)
	mapLock.Unlock()
}
