package psmetrics

import "net"

func getDevWithIP() map[string]string {
	res := make(map[string]string)
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, dev := range interfaces {
			devaddr, _ := dev.Addrs()
			var devip string
			for _, v := range devaddr {
				if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						devip = ipnet.IP.String()
						break
					}
				}
			}
			if devip != "" {
				res[dev.Name] = devip
			}
		}
	}
	return res
}
