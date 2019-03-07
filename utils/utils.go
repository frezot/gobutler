package utils

import "log"
import "net"

// BuildAppioidBaseURL detects ip and concat with port to URL
func BuildAppioidBaseURL(port string) string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println("[DEBUG] ip autodetect failed. localhost will be used")
		return "http://127.0.0.1:" + port
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return "http://" + localAddr.IP.String() + ":" + port
}

// AppiumServerURL build URL for interact with appium
func AppiumServerURL(port string) string {
	return "http://127.0.0.1:" + port + "/wd/hub"
}
