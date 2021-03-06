package serverbox

import (
	"errors"
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/mux"
	"sync"
)

type ServerType uint8

const (
	invalid_server = iota
	http_server
)

type serverInstance interface {
	InitializeServerInstance(ServerConfigurations) error
	RunServerInstance() error
	ShutDownServerInstance() error
	AbortServerInstance() error
	AttachRouterServerInstance(mux.Router) error
}

type Server struct {
	name           string
	sType          ServerType
	uuid           string
	bindIp         string
	bindPort       uint16
	stats          Statistics
	state          State
	serverInstance serverInstance
	enabled        bool
}

func convertServerType(sType string) (ServerType, error) {
	switch sType {
	case "http":
		return http_server, nil
	}
	return invalid_server, errors.New("invalid server type")
}

func getServerInstance(s *Server) (serverInstance, error) {
	switch s.sType {
	case http_server:
		return &ServerHttp{server: s}, nil
	}
	return nil, errors.New("invalid server type")
}

func generateUuid(name string, ip string, port uint16) (string, error) {
	return fmt.Sprintf("%s@%s:%d", name, ip, port), nil
}

func InitializeServers(sbc *SbContext) (err error) {
	sbc.Servers = make(map[string]*Server)

	for serverName, serverConf := range sbc.Conf.Servers {
		server := new(Server)
		server.name = serverName
		server.bindIp = serverConf.Bind_ip
		server.bindPort = serverConf.Bind_port

		server.sType, err = convertServerType(serverConf.Type)
		if err != nil {
			break
		}

		server.uuid, _ = generateUuid(serverName, server.bindIp,
			server.bindPort)

		statsConf := serverConf.Statistics
		if statsConf.Enabled {
			host := fmt.Sprintf("%s:%d", statsConf.Host,
				statsConf.Port)
			err = InitializeStatistics(server.uuid,
				host, &server.stats)
			if err != nil {
				break
			}
		}
		stateConf := serverConf.State
		if stateConf.Enabled {
			host := fmt.Sprintf("%s:%d", stateConf.Host,
				stateConf.Port)
			err = InitializeState(server.uuid,
				host, &server.state)
			if err != nil {
				break
			}
		}

		server.serverInstance, err = getServerInstance(server)
		if err != nil {
			break
		}

		err = server.serverInstance.InitializeServerInstance(serverConf.Configurations)
		if err != nil {
			break
		}

		server.enabled = true
		sbc.Servers[serverName] = server
	}
	return err
}

func RunServers(sbc *SbContext) error {
	for _, server := range sbc.Servers {
		go server.serverInstance.RunServerInstance()
	}
	return nil
}

func ShutDownServers(sbc *SbContext) error {
	var wg sync.WaitGroup
	for _, server := range sbc.Servers {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go func() {
			server.serverInstance.ShutDownServerInstance()
			ShutDownStatistics(&server.stats)
			ShutDownState(&server.state)
			wg.Done()
		}()
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	return nil
}

func AbortServers(sbc *SbContext) (err error) {
	for _, server := range sbc.Servers {
		server.serverInstance.AbortServerInstance()
		ShutDownStatistics(&server.stats)
		ShutDownState(&server.state)
	}
	return nil
}

func AttachRouterToServer(router mux.Router, serName string, sbc *SbContext) (err error) {
	server := sbc.Servers[serName]
	if server != nil {
		err = server.serverInstance.AttachRouterServerInstance(router)
	}
	return err
}
