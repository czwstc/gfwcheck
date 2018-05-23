package exec

import (
	"os/exec"

	"log"

	"os"

	"strings"

	"github.com/mritd/gfwcheck/proxy"
	"github.com/robfig/cron"
	"github.com/spf13/viper"
)

func (server *ServerConfig) RemoteExec() bool {
	client, err := server.Connection()
	defer client.Close()
	if err != nil {
		log.Printf("Connect to server [%s] failed!\n", server.Host)
		return false
	}
	session, err := client.NewSession()
	defer session.Close()
	if err != nil {
		log.Println("Session create failed:", err)
		return false
	}
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Run(server.RemoteCmd)
	if err != nil {
		log.Printf("Server %s remote command [%s] exec failed!\n", server.Name, server.RemoteCmd)
		log.Println(err.Error())
		return false
	} else {
		log.Printf("Server %s remote command [%s] exec success!\n", server.Name, server.RemoteCmd)
		return true
	}
}

func (server *ServerConfig) LocalExec() bool {
	var cmd *exec.Cmd
	localCmd := strings.Fields(server.LocalCmd)
	if len(localCmd) < 1 {
		log.Printf("Local command missing,Server %s\n", server.Name)
		return false
	} else if len(localCmd) == 1 {
		cmd = exec.Command(localCmd[0])
	} else {
		cmd = exec.Command(localCmd[0], localCmd[1:]...)
	}

	err := cmd.Run()
	if err != nil {
		log.Printf("Server %s local command [%s] exec failed!\n", server.Name, server.LocalCmd)
		log.Println(err.Error())
		return false
	} else {
		log.Printf("Server %s local command [%s] exec success!\n", server.Name, server.LocalCmd)
		return true
	}
}

func (server *ServerConfig) CheckGFWAndExec() {
	log.Printf("%s checking...\n", server.Name)
	if !proxy.Check(server.Proxy) {
		server.RemoteExec()
		server.LocalExec()
	}
}

func Start() {
	var servers []ServerConfig
	err := viper.UnmarshalKey("servers", &servers)
	if err != nil {
		log.Println("Can't parse server config!")
		return
	}
	c := cron.New()
	for i, _ := range servers {
		x := i
		c.AddFunc(servers[x].Cron, func() {
			servers[x].CheckGFWAndExec()
		})
	}
	c.Start()
	select {}
}
