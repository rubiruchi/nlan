package ptn

import (
	"github.com/araobp/go-nlan/nlan/agent/context"
	"github.com/araobp/go-nlan/nlan/model/nlan"
	"github.com/araobp/go-nlan/nlan/util"
)

func AddLinks(links *nlan.Links, con *context.Context, brTun string, brInt string) {
	cmd, cmdp := con.GetCmd()
	logger := con.Logger
	localIp := links.LocalIp
	remoteIps := links.RemoteIps
	for _, remoteIp := range remoteIps {
		inf := util.ZfillIp(remoteIp)
		logger.Printf("Adding a VXLAN tunnel: %s", inf)
		cmdp("ovs-vsctl", "add-port", brTun, inf, "--", "set", "interface", inf, "type=vxlan", "options:in_key=flow", "options:local_ip="+localIp, "options:out_key=flow", "options:remote_ip="+remoteIp)
		infNum := string(util.GetOfport(inf))
		cmd("ovs-ofctl", "add-flow", brTun, "table=0,priority=1,in_port="+infNum+",actions=resubmit(,2)")
		cmd("ip", "link", "set", "dev", brTun, "up")
		cmd("ip", "link", "set", "dev", brInt, "up")
	}
}

func UpdateLinks(links *nlan.Links, con *context.Context, brTun string, brInt string) {
}

func DeleteLinks(links *nlan.Links, con *context.Context, brTun string, brInt string) {
}