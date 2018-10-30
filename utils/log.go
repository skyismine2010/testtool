package utils

import (
	"log"
	"os"
)

func InitLog(filePath string) error{
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	logFile, err:= os.Open(filePath)
	if err != nil{
		return err
	}

	log.SetOutput(logFile)
	return nil
}


