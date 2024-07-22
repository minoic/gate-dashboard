package main

import (
	"github.com/minoic/gate-dashboard/plugins/players"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// It's a normal Go program, we only need
// to register our plugins and execute Gate.
func main() {
	// Here we register our plugins with the proxy.
	proxy.Plugins = append(proxy.Plugins,
		players.PlayersPlugin,
	)

	gate.Execute()
}