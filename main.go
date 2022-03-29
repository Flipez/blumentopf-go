package main

import (
	"github.com/Tinkerforge/go-api-bindings/industrial_quad_relay_v2_bricklet"
	"github.com/Tinkerforge/go-api-bindings/ipconnection"
	"github.com/go-co-op/gocron"

	"time"
)

const ADDR string = "192.168.178.46:4223"
const UID string = "NW1"

func main() {
	ipcon := ipconnection.New()
	defer ipcon.Close()
	iqr, _ := industrial_quad_relay_v2_bricklet.New(UID, &ipcon)

	ipcon.Connect(ADDR)
	defer ipcon.Disconnect()

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("19:00").Do(func() {
		iqr.SetMonoflop(0, true, 10000)
		time.Sleep(10 * time.Second)
		iqr.SetMonoflop(1, true, 10000)
		time.Sleep(10 * time.Second)
		iqr.SetMonoflop(2, true, 10000)
		time.Sleep(10 * time.Second)
		iqr.SetMonoflop(3, true, 10000)
	})

	s.StartBlocking()
}
