package vault_port_scanner

import (
	"encoding/json"
	"fmt"
	"github.com/anvie/port-scanner"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"time"
)

func ScanLocalHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("Scanning NOW Beginning!")
	ps := portscanner.NewPortScanner("localhost", 2*time.Second, 5)
	m := make(map[int]string)
	fmt.Printf("scanning port %d-%d...\n", 20, 30000)
	openedPorts := ps.GetOpenedPort(20, 30000)
	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
		m[port] = ps.DescribePort(port)
	}
	connectedClients, _ := json.Marshal(m)
	clientResponse(connectedClients, w)
}

func ScanNetworkDevice(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Printf("Scanning NOW Beginning!")
	ip_add := params.ByName("ip_add")
	ps := portscanner.NewPortScanner(ip_add, 2*time.Second, 5)
	m := make(map[int]string)
	// get opened port
	fmt.Printf("scanning port %d-%d...\n", 20, 30000)
	openedPorts := ps.GetOpenedPort(20, 30000)
	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		log.Print(" ", port, " [open]")
		log.Println("  -->  ", ps.DescribePort(port))
		m[port] = ps.DescribePort(port)
	}
	connectedClients, _ := json.Marshal(m)
	clientResponse(connectedClients, w)
}

func TestJson(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	x := params.ByName("try")
	if len(x) < 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error"))
	}
}
