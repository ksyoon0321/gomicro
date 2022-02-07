package adapter

import (
	"errors"
	"strings"
	"time"
)

type NotifyData struct {
	Act      string
	RType    string
	Id       string
	Data     string
	Pongtime time.Time
}

type IAdapter interface {
	Listen(ch chan NotifyData)
}

func ParseRegPayload(payload string) (NotifyData, error) {
	// ID_SERVICEIP_PORT -> REG_MON_1_127.0.0.1_9999
	arrdata := strings.Split(payload, "_")

	ndata := NotifyData{}
	if len(arrdata) != 5 {
		return ndata, errors.New("reg payload parse error :" + payload)
	}

	ndata.Act = arrdata[0]
	ndata.RType = arrdata[1]
	ndata.Id = arrdata[2]
	ndata.Data = arrdata[3] + ":" + arrdata[4]
	ndata.Pongtime = time.Now()

	return ndata, nil
}

func (n *NotifyData) GetId() string {
	return n.Id + ":" + n.Data
}
