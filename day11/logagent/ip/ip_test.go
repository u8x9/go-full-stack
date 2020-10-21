package ip

import "testing"

func TestGetOutboundIP(t *testing.T) {
	ip, err := GetOutboundIP()
	if err != nil {
		t.Fatal("can not get outbound ip:", err)
	}
	t.Log(ip)
}
