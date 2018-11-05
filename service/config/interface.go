package config

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"
)

var DbConfig DbCfg

func LoadXmlCfg(filePath string) error {
	//doc := etree.NewDocument()
	//if err := doc.ReadFromFile(filePath); err != nil {
	//	log.Printf("Can't read path=%s xml file.", filePath)
	//}
	//dbElement := doc.SelectElement("db")
	//for _, tblElement := range dbElement.SelectElements("tbl") {
	//	for _, recordElement := range tblElement.SelectElements("id") {
	//
	//	}
	//}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("read config filled ,fill path=%s, err=%v", filePath, err)
		return err
	}
	decoder := xml.NewDecoder(bytes.NewReader(content))
	err = decoder.Decode(&DbConfig)
	if err != nil {
		log.Printf("Failed to decode config, err=%v", err)
	}
	return nil
}
