package global

type ServerConf struct {
	svrType    string
	sid        uint32
	ip         string
	port       uint16
	debugLevel uint8
	user       string
	password   string
}

var ports map[string]int

type ServerInstance struct {
	wanIp    string
	lanIP    string
	user     string
	password string
}

var instances map[string]ServerInstance

var deploy map[string][]ServerConf

func init() {
	deploy = map[string][]ServerConf{}
	instances = map[string]ServerInstance{}

	initLocalhost()
	initDevTest()
	initTest()
	initProduct()
}

func initLocalhost() {
	lhip := "127.0.0.1"
	deploy["localhost"] = make([]ServerConf, 0)
	deploy["localhost"] = append(deploy["localhost"], ServerConf{
		svrType:    "log",
		sid:        39001,
		ip:         lhip,
		port:       49001,
		debugLevel: 0,
	})
}

func initDevTest() {
	instances["aliyun_sz_1"] = ServerInstance{
		wanIp:    "47.107.151.251",
		lanIP:    "172.18.36.104",
		user:     "root",
		password: "Wfs123456",
	}

	insIp := instances["aliyun_sz_1"].wanIp

	deploy["devtest"] = make([]ServerConf, 0)
	deploy["devtest"] = append(deploy["devtest"], ServerConf{
		svrType:    "mysql",
		sid:        3306,
		ip:         insIp,
		port:       3306,
		debugLevel: 0,
		user:       "root",
		password:   "02#F20ebac",
	})

	deploy["devtest"] = append(deploy["devtest"], ServerConf{
		svrType:    "redis",
		sid:        6379,
		ip:         insIp,
		port:       6379,
		debugLevel: 0,
		user:       "",
		password:   "",
	})
}

func initTest() {

}

func initProduct() {

}
