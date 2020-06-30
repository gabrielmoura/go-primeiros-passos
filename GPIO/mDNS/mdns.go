package mDNS

import (
	"github.com/hashicorp/mdns"
	"net"
	"os"
)
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func SetDNS() {
	// Setup our service export
	host, _ := os.Hostname()
	info := []string{"Raspberry PI"}
	service, _ := mdns.NewMDNSService(host, "_rpi._tcp", "", getLocalIP(), 8000, nil, info)

	// Create the mDNS server, defer shutdown
	server, _ := mdns.NewServer(&mdns.Config{Zone: service})
	defer server.Shutdown()
}
