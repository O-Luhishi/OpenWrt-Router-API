package vault_port_scanner

import (
	"fmt"
	portscanner "github.com/anvie/port-scanner"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := fmt.Sprintf(`{Status: "UP"}`)
	writeOKResponse(w, response)
}

func ScanLocalHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "Scanning Beginning!\n")
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
	writeOKResponse(w, m)
}

func ScanNetworkDevice(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	fmt.Fprint(w, "Scanning Beginning!\n")
	ip_add := params.ByName("ip_add")
	ps := portscanner.NewPortScanner(ip_add, 2*time.Second, 5)
	m := make(map[int]string)
	// get opened port
	fmt.Printf("scanning port %d-%d...\n", 20, 30000)
	openedPorts := ps.GetOpenedPort(20, 30000)
	for i := 0; i < len(openedPorts); i++ {
		port := openedPorts[i]
		fmt.Print(" ", port, " [open]")
		fmt.Println("  -->  ", ps.DescribePort(port))
		m[port] = ps.DescribePort(port)
	}
	writeOKResponse(w, m)
}

func TestJson(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	x := params.ByName("try")
	if len(x) < 0{
		writeErrorResponse(w, http.StatusNotFound, "Record Not Found")
	}
	writeOKResponse(w, x)
}

