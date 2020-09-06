

package systemutil

import (
	"net"
	"testing"
)

func Test_IP(t *testing.T) {
	a := net.IPv4(192, 168, 1, 1)
	println(a.IsGlobalUnicast())

	a = net.IPv4(10, 10, 1, 1)
	println(a.IsGlobalUnicast())

	a = net.IPv4(127, 0, 0, 1)
	println(a.IsGlobalUnicast())
}

func Test_GetAgentIp(t *testing.T) {
	println(GetAgentIp())
}
