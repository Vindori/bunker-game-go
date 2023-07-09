package game_engine

type BunkerGameEngine interface {
	GetPlayers() []*Player
}

var _ BunkerGameEngine = (*BunkerGame)(nil)

type BunkerGame struct {
	Players []*Player
}

func (g *BunkerGame) GetPlayers() []*Player {
	return g.Players
}
