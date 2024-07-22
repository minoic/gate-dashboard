package metrics

import (
	"context"

	"github.com/robinbraemer/event"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)


var MetricsPlugin = proxy.Plugin{
	Name: "Metrics",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		event.Subscribe(p.Event(), 0, onLogin())
		event.Subscribe(p.Event(), 0, onDisconnect())
		return nil
	},
}

func onLogin() func(*proxy.LoginEvent) {

}

func onDisconnect() func(*proxy.DisconnectEvent) {

}