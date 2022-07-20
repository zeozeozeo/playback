package effect

import (
	"fmt"

	"github.com/gotracker/playback/format/s3m/layout/channel"
	"github.com/gotracker/playback/player/intf"
)

// Tremor defines a tremor effect
type Tremor ChannelCommand // 'I'

// Start triggers on the first tick, but before the Tick() function is called
func (e Tremor) Start(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback) error {
	cs.ResetRetriggerCount()
	return nil
}

// Tick is called on every tick
func (e Tremor) Tick(cs intf.Channel[channel.Memory, channel.Data], p intf.Playback, currentTick int) error {
	mem := cs.GetMemory()
	x, y := mem.LastNonZeroXY(channel.DataEffect(e))
	return doTremor(cs, currentTick, int(x)+1, int(y)+1)
}

func (e Tremor) String() string {
	return fmt.Sprintf("I%0.2x", channel.DataEffect(e))
}
