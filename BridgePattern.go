package main

import "DesignPattern/BridgePattern"

func main(){
	b1 := BridgePattern.NewCommMessage(BridgePattern.NewEmailMessage())
	b1.SendSpeedMessage("lqh","yaya")
	b2 := BridgePattern.NewQuickMessage(BridgePattern.NewPhoneMessage())
	b2.SendSpeedMessage("lqh","yaya")
	b3 := BridgePattern.NewQuickMessage(BridgePattern.NewPhoneMessage())
	b3.SendSpeedMessage("lqh","yaya")
}
