package config

type RecordT struct {
	recordID  int32
	oneRecord map[string]string
}

type TableT struct {
	tblName string
	records []RecordT
}

type HttpCfgRecord struct {
	Id   int    `xml:id,attr`
	Host string `xml:host,attr`
}

type HttpCfgTbl struct {
	Records []HttpCfgRecord `xml:Record`
}

type TcpCfgRecord struct {
	Id       int    `xml:"id,attr"`
	SrcIP    string `xml:"src_ip,attr"`
	DestIP   string `xml:"dest_ip,attr"`
	SrcPort  string `xml:"src_port,attr"`
	DestPort string `xml:"dest_port,attr"`
}

type TcpCfgTbl struct {
	Records []TcpCfgRecord `xml:Record`
}

type DbCfg struct {
	HttpCfg HttpCfgTbl `xml:"http"`
	TcpCfg  TcpCfgTbl  `xml:"tcp"`
}
