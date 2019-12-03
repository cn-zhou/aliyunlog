package aliyunlog

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strings"
)

//自定义格式
type AliLogFormat struct {
}

func (a *AliLogFormat) Format(entry *log.Entry) ([]byte, error) {
	eData := make(map[string]interface{})
	//累加
	eData["content"] = entry.Message
	if len(entry.Data) != 0 {
		for k, v := range entry.Data {
			eData[k] = v
		}
	}
	fullData := map[string]interface{}{
		"message": eData,
	}
	m, _ := json.Marshal(fullData)
	var bf bytes.Buffer
	bf.WriteString(GetLocalIp())
	bf.WriteString("|")
	bf.WriteString(GetUnixMsToTime(entry.Time))
	bf.WriteString("|")
	bf.WriteString(strings.ToUpper(entry.Level.String()))
	bf.WriteString("|null|null|")
	bf.WriteString(string(m))
	bf.WriteString("\n")
	return bf.Bytes(), nil
}
