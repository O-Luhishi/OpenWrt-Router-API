package vault_network_mapper

import (
	"encoding/json"
	"fmt"
	"github.com/Vioft/Vault-API/common"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func GetConnectedDevices(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, result, _ := common.RunBash(connectedDevicesBashWrapper())
	out := parseBashOutput(result)
	clientResponse(out, w)
}

func parseBashOutput(result string) []byte {
	var clients []*client
	for _, line := range strings.Split(result, "\n") {
		if line != "" {
			s := strings.Split(line, " ")
			clients = append(clients, &client{Hostname: s[0], IP: s[1], MAC: s[2], STATUS: true})
		}
	}
	connectedClients, _ := json.Marshal(connection{Clients: clients})
	fmt.Printf("%s\n", connectedClients)
	return connectedClients
}

func connectedDevicesBashWrapper() string {
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

// When you want to run this locally
func replicateRouter() string {
	return `echo Osamas-MBP 192.168.8.127 F4:0F:24:24:B4:F0
			echo Mo-MBP 192.168.8.2 F0:F0:F0:F0:F0`
}
