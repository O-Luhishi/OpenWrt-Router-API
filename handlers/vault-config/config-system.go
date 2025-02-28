package vault_config

type Ubus func(arg string) ([]byte, error)

type NUbus func(arg string) ([]byte, error)

func (u Ubus) Info() ([]byte, error) {
	return u("system info")
}

func (u Ubus) BoardInfo() ([]byte, error) {
	return u("system board")
}

func (u Ubus) WanStatus() ([]byte, error) {
	return u("network.interface.wan status")
}

func (u Ubus) LanStatus() ([]byte, error) {
	return u("network.interface.lan status")
}

func (u Ubus) WirelessStatus() ([]byte, error) {
	return u("network.wireless status")
}

func (u Ubus) ServiceList() ([]byte, error) {
	return u("service list")
}

func (u Ubus) InterfacesList() ([]byte, error) {
	return u("network.device status")
}

func (u Ubus) NetworkConfig() ([]byte, error) {
	return u("uci get {\"config\":\"network\"}")
}

func (u Ubus) WlanClients() ([]byte, error) {
	return u("hostapd.wlan0 get_clients")
}

func (u NUbus) BanClient(x string) ([]byte, error) {
	var macAddress, banTime string
	return u("hostapd.wlan0 del_client `{'addr':`" + macAddress + "`, 'reason':1, 'deauth':true, 'ban_time': " + banTime)
}
