package vault_config

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var UbusCall = Ubus(func(arg string) ([]byte, error) {
	args := strings.Split(arg, " ")
	args = append([]string{"-S", "call"}, args...)
	return exec.Command("ubus", args...).Output()
})

func HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	m := make(map[string]string)
	m["Status"] = "Up"
	status, _ := json.Marshal(m)
	log.Printf("%s \n", status)
	healthcheckResponse(status, w)
	//response := fmt.Sprintf(`Up`)
	//WriteOKResponse(w, response)
}

func GetSystemInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.Info, w)
}

func GetBoardInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.BoardInfo, w)
}

func GetWanStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.WanStatus, w)
}

func GetLanStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.LanStatus, w)
}

func GetWirelessStatus(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.WirelessStatus, w)
}

func GetServiceList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.ServiceList, w)
}

func GetInterfaceList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.InterfacesList, w)
}

func GetNetworkConfig(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.NetworkConfig, w)
}

func GetWlanClients(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	UbusResponse(UbusCall.WlanClients, w)
}
