package vault_device_manager

import (
	"encoding/json"
	"github.com/Vioft/Vault-API/common"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func ubusBashWrapper(deviceAddr, time string) string {
	return `ubus call hostapd.wlan0 del_client "{'addr': '` + deviceAddr + `', 'reason':1, 'deauth':true, 'ban_time': ` + time + `}"`
}
func BanClient(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	deviceAddr := params.ByName("mac")
	time := params.ByName("time")
	m := make(map[string]string)
	log.Print("Banning: ", deviceAddr, " For Duration: ", time)
	_, _, err := common.RunBash(ubusBashWrapper(deviceAddr, time))
	if err != "" {
		m["Device-Ban"] = "Error"
		log.Println(err)
	} else {
		m["Device-Ban"] = "Success"
	}
	response, _ := json.Marshal(m)
	log.Printf("%s \n", response)
	banClientResponse(response, w)
}
