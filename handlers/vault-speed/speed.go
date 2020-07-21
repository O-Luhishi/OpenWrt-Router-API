package vault_speed

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"github.com/ddo/go-fast"
)

func Get_Download_Speed(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Printf("Getting Download Speed")
	status := ""
	// output
	fastCom := fast.New()

	// init
	err := fastCom.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get urls
	urls, err := fastCom.GetUrls()
	if err != nil {
		log.Println(err)
		return
	}

	// measure
	KbpsChan := make(chan float64)

	go func() {
		for Kbps := range KbpsChan {
			status = format(Kbps)
		}
		m := make(map[string]string)
		m["Download-Speed"] = status
		speed, _ := json.Marshal(m)
		log.Printf("%s \n", speed)
		speedResponse(speed, w)
	}()

	err = fastCom.Measure(urls, KbpsChan)
	if err != nil {
		log.Println(err)
	}
	return
}

func format(Kbps float64) string {
	unit := "Kbps"
	f := "%.f %s"

	if Kbps > 1000000 { // Gbps
		f = "%.2f %s"
		unit = "Gbps"
		Kbps /= 1000000

	} else if Kbps > 1000 { // Mbps
		f = "%.2f %s"
		unit = "Mbps"
		Kbps /= 1000
	}
	return fmt.Sprintf(f, Kbps, unit)
}
