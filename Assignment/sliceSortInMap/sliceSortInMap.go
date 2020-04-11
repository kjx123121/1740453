package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"module=logTransfer",
		"log_type=console",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
		"# kafka的相关配置",
		"[kafka]",
		"# kafka的地址",
		"addr=127.0.0.1:9092",
		"# logTransfer要消费的topic",
		"topic=nginx_log",
		"[",
	}


	finalPrintOut(ParseIni(iniData))


}
func ParseIni(receivedData []string) map[string]map[string]string {

	var errorTime string = time.Now().Format("2006/01/02 15:04:05")
	var outerMap = make(map[string]map[string]string)
	var interMap = make(map[string]string)

	for i := len(receivedData) - 1; i >= 0; i-- {
		dataLine := receivedData[i]

		if strings.HasPrefix(dataLine, "[") && strings.HasSuffix(dataLine, "]") {
			insertToOuterMap(dataLine, outerMap, interMap)
			interMap = make(map[string]string)

		} else if dataLine == "" || dataLine[0] == '#' || dataLine[0] == ';' {
			continue

		} else if strings.HasPrefix(dataLine, "[") || strings.HasSuffix(dataLine, "]") {
			fmt.Printf("%v parse ini failed, exceptional line content is: %v\n", errorTime, dataLine)

		} else {
			insertToInterMap(dataLine, interMap)

		}

	}
	outerMap["[Global]"] = interMap

	return outerMap
}

func insertToInterMap(dataLine string, interMap map[string]string) map[string]string {
	sliceOfString := strings.Split(dataLine, "=")
	interMap[sliceOfString[0]] = sliceOfString[1]
	return interMap

}

func insertToOuterMap(dataLine string, outerMap map[string]map[string]string, interMap map[string]string) map[string]map[string]string {
	outerMap[dataLine] = interMap
	return outerMap

}

func finalPrintOut (finalMap map[string]map[string]string){
	for outerKey, outerElem := range finalMap {
		fmt.Printf("%v\n", outerKey)

		for interKey, interElem := range outerElem {
			fmt.Printf("%v = %v\n", interKey, interElem)
		}
	}
}
