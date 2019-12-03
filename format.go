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
	msg := map[string]interface{}{
		"message": entry.Message,
	}
	//累加
	if len(entry.Data) != 0 {
		for k, v := range entry.Data {
			msg[k] = v
		}
	}
	m, _ := json.Marshal(msg)
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
