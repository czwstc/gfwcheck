package exec

import "time"

type ServerConfig struct {
	Name      string        `yml:"name"`
	Host      string        `yml:"ip"`
	Port      int           `yml:"port"`
	User      string        `yml:"user"`
	Password  string        `yml:"password"`
	Method    string        `yml:"method"`
	Key       string        `yml:"key"`
	Timeout   time.Duration `yml:"timeout"`
	Proxy     string        `yml:"proxy"`
	LocalCmd  string        `yml:"localcmd"`
	RemoteCmd string        `yml:"remotecmd"`
	Cron      string        `yml:"cron"`
}

func ExampleConfig() []*ServerConfig {
	return []*ServerConfig{
		{
			Name:     "test1",
			Host:     "test1.com",
			Port:     22,
			User:     "root",
			Password: "test123",
			Key:      "",
			Method:   "password",
		},
		{
			Name:     "test2",
			Host:     "test2.com",
			Port:     22,
			User:     "root",
			Password: "",
			Key:      "~/.ssh/id_rsa",
			Method:   "pem",
		},
	}
}
