package eago

import (
	"github.com/BurntSushi/toml"
	"testing"
)

const TestConfig = `
CrawlerName = "crawler"
Urls = []
Depth = 5
InSite = true
TimeOut = 5
TTL = 2
Retry=3
HttpPort = 12002

Auth=false
UserName=""
Token=""

ClusterName = "eagles"

[Local]
NodeName = "eagle"
IP = "127.0.0.1"
Port = 12001

[[NodeList]]
NodeName = "eagle"
IP = "127.0.0.1"
Port = 12001

[[NodeList]]
NodeName = "eago_shard1"
IP = "127.0.0.1"
Port = 12003

[Redis]
    [Redis.redisShard1]
    Host="127.0.0.1:6374"
    DB=0
    Pool=5

    [Redis.redisShard2]
    Host="127.0.0.1:6374"
    DB=1
    Pool=5`

func LoadTestConfig() {
	LogInit()
	_, err := toml.Decode(TestConfig, Configs)
	AssertNil(err)
	ArbitrateConfigs(Configs)
}

func TestLoadTestConfig(t *testing.T) {
	LoadTestConfig()
	AssertEqual(Configs.CrawlerName == "crawler")
	AssertEqual(len(Configs.Urls) == 0)
	AssertEqual(Configs.Depth == 5)
	AssertEqual(Configs.InSite)
	AssertEqual(Configs.TimeOut == 5)
	AssertEqual(Configs.TTL == 2)
	AssertEqual(Configs.Retry == 3)
	AssertEqual(Configs.HttpPort == 12002)
	AssertEqual(Configs.Auth == false)
	AssertEqual(Configs.UserName == "")
	AssertEqual(Configs.Token == "")
	AssertEqual(Configs.ClusterName == "eagles")
	AssertEqual(*Configs.Local == NodeInfo{NodeName: "eagle", IP: "127.0.0.1", Port: 12001})
	AssertEqual(*Configs.NodeList[0] == NodeInfo{NodeName: "eagle", IP: "127.0.0.1", Port: 12001})
	AssertEqual(*Configs.NodeList[1] == NodeInfo{NodeName: "eago_shard1", IP: "127.0.0.1", Port: 12003})
	AssertEqual(*Configs.Redis["redisShard1"] == RedisInstance{Host: "127.0.0.1:6374", DB: 0, Pool: 5})
	AssertEqual(*Configs.Redis["redisShard2"] == RedisInstance{Host: "127.0.0.1:6374", DB: 1, Pool: 5})
}