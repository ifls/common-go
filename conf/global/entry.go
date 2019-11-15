package global

import (
	"fmt"
	"strings"
)

const (
	EnvLocalhost = "localhost"
	EnvInner     = "inner"
	EnvAliChina  = "devTest"
	EnvAliHk     = "aliHk"
	EnvUCHk      = "uc"
	EnvProduct   = "product"
)

//同一个内网的为一个环境
type envInfo struct {
	hosts      []string
	subNetMark string
}

type ServerInfo struct {
	Ip   string
	Port int16
}

func (s ServerInfo) Address() string {
	return fmt.Sprintf("%s:%d", s.Ip, s.Port)
}

//key为环境, value为主机或者域名
var envs map[string]envInfo
var current string

func init() {
	envs = map[string]envInfo{}
	envs[EnvLocalhost] = envInfo{
		hosts: []string{"127.0.0.1"},
	}
	envs[EnvInner] = envInfo{
		hosts: []string{
			"192.168.0.113",
		},
	}
	envs[EnvAliHk] = envInfo{
		hosts: []string{
			"",
		},
	}
	envs[EnvAliChina] = envInfo{
		hosts: []string{
			"47.107.151.251",
		},
	}
	envs[EnvUCHk] = envInfo{
		hosts: []string{
			"23.91.101.147",
		},
	}
	envs[EnvProduct] = envInfo{
		hosts: []string{
			"",
		},
	}
	current = EnvUCHk
}

func GetHostIp(env string, serverType string) string {
	host := ""

	return host
}

func GetServer(env string, serverType ServerType) ServerInfo {
	host := getTcpHost(env, serverType)
	port := getPort(serverType)

	return ServerInfo{
		Ip:   host,
		Port: int16(port),
	}
}

func setEnv(env string) bool {
	if _, ok := envs[env]; ok {
		current = env
		return true
	}

	return false
}

func getTcpHost(env string, serverType ServerType) string {
	host := ""
	return host
}

func getPort(serverType ServerType) int {
	name := ServerType_name[int32(serverType)]
	port := ServerPort_value[strings.Split(name, "_")[1]]
	return int(port)
}
