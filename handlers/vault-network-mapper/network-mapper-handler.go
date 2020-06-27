package vault_network_mapper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

var (
	signals = make(chan os.Signal, 100)
)

func GetConnectedDevices(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	_, result,_ := runBash()
	out := parseBashOutput(result)
	clientResponse(out, w)
}

func parseBashOutput(result string) []byte {
	var clients []*client
	for _, line := range strings.Split(result, "\n") {
		if line != "" {
			s := strings.Split(line, " ")
			clients = append(clients, &client{Hostname: s[0], IP: s[1], MAC: s[2]})
		}
	}
	connected_clients, _ := json.Marshal(connection{Clients: clients})
	fmt.Printf("%s\n", connected_clients)
	return connected_clients
}

func runBash()(bool, string, string){
	cmd := exec.Command("/bin/sh", "-s")
	cmd.Stdin = strings.NewReader(connectedDevicesBashWrapper())
	return finishRunning(cmd)
}

func connectedDevicesBashWrapper() string{
	return `
#!/bin/sh
for interface in ` + "`iwinfo | grep ESSID | cut -f 1 -s -d\" \"`;" +
		`do
  maclist=` + "`iwinfo $interface assoclist | grep dBm | cut -f 1 -s -d\" \"`" +
		`
  for mac in $maclist;
  do
    # If a DHCP lease has been given out by dnsmasq,
    # save it.
    ip="UNKN"
    host=""
    ip=` + "`cat /tmp/dhcp.leases | cut -f 2,3,4 -s -d\" \" | grep -i $mac | cut -f 2 -s -d\" \"`\n" +
		`host=` + "`cat /tmp/dhcp.leases | cut -f 2,3,4 -s -d\" \" | grep -i $mac | cut -f 3 -s -d\" \"`\n" +
		`echo $host $ip $mac
done
done
`
}

func finishRunning(cmd *exec.Cmd) (bool, string, string) {
	//log.Printf("Running Connected Clients Script")
	stdout, stderr := bytes.NewBuffer(nil), bytes.NewBuffer(nil)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	done := make(chan struct{})
	defer close(done)
	go func() {
		for {
			select {
			case <-done:
				return
			case s := <-signals:
				cmd.Process.Signal(s)
			}
		}
	}()
	if err := cmd.Run(); err != nil {
		log.Printf("Error running %v", err)
		return false, string(stdout.Bytes()), string(stderr.Bytes())
	}
	return true, string(stdout.Bytes()), string(stderr.Bytes())
}

// When you want to run this locally
func replicateRouter() string{
	return `echo Osamas-MBP 192.168.8.127 F4:0F:24:24:B4:F0
			echo Mo-MBP 192.168.8.2 F0:F0:F0:F0:F0`
}