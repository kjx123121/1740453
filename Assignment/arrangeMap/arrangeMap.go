package main

import (
	"fmt"
	"sort"
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
		"[Kafka]",
		"# kafka的地址",
		"addr=127.0.0.1:9092",
		"# logTransfer要消费的topic",
		"topic=nginx_log",
		"[",
	}

	PrintIni(ParseIni(iniData))

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
			outerMap["[Global]"] = interMap
		}

	}

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

//以下为第二部分

func PrintIni(receivedMap map[string]map[string]string) {
	outerKeysSlice := make([]string, 0, len(receivedMap))

	for outerKeys := range receivedMap {
		outerKeysSlice = append(outerKeysSlice, outerKeys)
	}

	sort.Strings(outerKeysSlice)

	for numOfKey, outerKeys := range outerKeysSlice {
		if numOfKey != 0 {
			fmt.Printf("\n")
		}

		fmt.Printf("%s\n", outerKeys)
		arrangeInterMap(receivedMap[outerKeys])
	}
}

func arrangeInterMap(interMap map[string]string) {
	interKeysSlice := make([]string, 0, len(interMap))

	for interKeys := range interMap {
		interKeysSlice = append(interKeysSlice, interKeys)
	}

	sort.Strings(interKeysSlice)

	for _, interKeys := range interKeysSlice {
		fmt.Printf("%s=%s\n", interKeys, interMap[interKeys])
	}
}
