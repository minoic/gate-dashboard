package players

import (
	"context"
	"sync"

	"github.com/robinbraemer/event"
	"go.minekube.com/gate/pkg/edition/java/proxy"
	"go.minekube.com/gate/pkg/util/uuid"
)

var players = make(map[uuid.UUID]proxy.Player)
var playersLock sync.RWMutex

var PlayersPlugin = proxy.Plugin{
	Name: "Ping",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		event.Subscribe(p.Event(), 0, onLogin())
		event.Subscribe(p.Event(), 0, onDisconnect())
		return nil
	},
}

func onLogin() func(*proxy.LoginEvent) {
	return func(e *proxy.LoginEvent) {
		playersLock.Lock()
		defer playersLock.Unlock()
		if e.Allowed() {
			players[e.Player().ID()] = e.Player()
		}
	}
}

func onDisconnect() func(*proxy.DisconnectEvent) {
	return func(e *proxy.DisconnectEvent) {
		playersLock.Lock()
		defer playersLock.Unlock()
		delete(players, e.Player().ID())
	}
}

func PlayerByUUID(id uuid.UUID) proxy.Player{
	playersLock.RLock()
	defer playersLock.RUnlock()
	return players[id]
}

func PlayerByName(username string) proxy.Player{
	playersLock.RLock()
	defer playersLock.RUnlock()
	for _, v := range players {
		if v.Username() == username{
			return v
		}
	}
	return nil
}

func PlayerList() []proxy.Player {
	playersLock.RLock()
	defer playersLock.RUnlock()
	var ret []proxy.Player
	for _, v := range players {
		ret = append(ret, v)
	}
	return ret
}
