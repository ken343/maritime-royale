// Package game will keep track of the games state auxilory functions
// for accessing and changing that state.
// State will be necessary for keeping track of the game's phases
// and map data. The games phases should be able to lock what actions
// can be taken by players.
package game

// Init will set all initial values for this packge when it is imported.
func Init() {
	// TODO: Add shotclock initialization here when adding that function.

	// always have the admin option available every game for now.
	Players["admin"] = Player{Name: "admin", IsAdmin: true, IsFinished: false}

}

// These values represent the phase of the game, and will determine what update function will be in play.
const (
	PhaseCommand = "command"
	PhaseAnimate = "animate"
	PhaseCalculate = "calculate"
	PixelsPerTile = 64 // I don't know if this is the proper place to put this value.
)

var (
	// IsAllPlayersMoved will update from false to true once all players have completed turns OR shotclock runs out.
	IsAllPlayersMoved bool = false

	// Game will eventually have a Civ 5 like clock to give players limited time to complete all actions.
	// Shotclock time.Duration

	// CurrentPhase represents what actions are allowed and help the client/server know what computations to perform.
	CurrentPhase string = PhaseCommand

	// Players is a Go map of different player objects, where keys are represented by user names.
	Players map[string]Player = make(map[string]Player, 0)
)

// Player object represents the player as having a name and privilege that represents 
// if they can perform admin functions. They also have an indicator records if they have complete all available moves.
type Player struct {
	Name string
	IsAdmin bool
	IsFinished bool
}

// Dimensions represent the map size (measured in tiles, not pixels!)
type Dimensions struct {
	MapX int
	MapY int
}

