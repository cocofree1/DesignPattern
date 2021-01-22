package BridgePattern

// 不同种类的message
type DifferentMessage interface {
	SendKindMessage(name,to string)
}

// 不同速度的message
type DiffSpeedMessage interface {
	SendSpeedMessage(name,to string)
}