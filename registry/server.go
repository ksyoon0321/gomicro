package registry

import (
	"time"

	"github.com/ksyoon0321/gomicro/registry/adapter"
)

type RegistryChannelId struct {
	Monitor string
	Fetcher string
}

func NewRegistryChannel(mon string, fet string) RegistryChannelId {
	return RegistryChannelId{Monitor: mon, Fetcher: fet}
}

type ServiceRegistry struct {
	regchans RegistryChannelId
	adapters []adapter.IAdapter
	notify   chan adapter.NotifyData
	monlist  map[string]adapter.NotifyData
}

func NewServiceRegistry(regChans RegistryChannelId) *ServiceRegistry {
	srv := &ServiceRegistry{
		regchans: regChans,
		monlist:  make(map[string]adapter.NotifyData),
	}

	srv.notify = make(chan adapter.NotifyData)
	return srv
}

func (s *ServiceRegistry) Listen() {
	go func() {
		for data := range s.notify {
			s.doActOrSave(data)
		}
	}()

	for _, item := range s.adapters {
		item.Listen(s.notify)
	}

	for {
		time.Sleep(time.Second)
	}
}

func (s *ServiceRegistry) doActOrSave(data adapter.NotifyData) {
	if data.Act == "REG" {
		switch data.RType {
		case "MON":
			s.monlist[data.GetId()] = data
		}
	} else if data.Act == "OUT" {
		switch data.RType {
		case "MON":
			delete(s.monlist, data.GetId())
		}
	} else if data.Act == "PONG" {
		if s.isRegisterd(data) {
			switch data.RType {
			case "MON":
				tmp := s.monlist[data.GetId()]
				tmp.Pongtime = time.Now()
				s.monlist[data.GetId()] = tmp
			}
		}
	}
}

func (s *ServiceRegistry) isRegisterd(data adapter.NotifyData) bool {
	switch data.RType {
	case "MON":
		if _, ok := s.monlist[data.GetId()]; ok {
			return true
		}
	}
	return false
}

func (s *ServiceRegistry) RegistAdapter(adap adapter.IAdapter) {
	s.adapters = append(s.adapters, adap)
}
