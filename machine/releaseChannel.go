package machine

// ReleaseChannel enum represents a CoreOS release channel
type ReleaseChannel int

// Stable etc. release channels
const (
	Stable ReleaseChannel = iota
	Beta
	Alpha
)

var channelByID = map[ReleaseChannel]string{
	Stable: "stable",
	Beta:   "beta",
	Alpha:  "alpha",
}

func (c ReleaseChannel) String() string {
	return channelByID[c]
}
