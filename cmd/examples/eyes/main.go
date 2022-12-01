package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/sdk-wrapper"
	"github.com/digital-dream-labs/vector-go-sdk/pkg/sdk-wrapper/settings"
	"time"
)

func main() {
	var serial = flag.String("serial", "", "Vector's Serial Number")
	var hue = flag.String("hue", "", "Hue (0.0 .. 1.0)")
	var saturation = flag.String("saturation", "", "Saturation (0.0 .. 1.0)")
	flag.Parse()

	sdk_wrapper.InitSDK(*serial)

	ctx := context.Background()
	start := make(chan bool)
	stop := make(chan bool)

	go func() {
		_ = sdk_wrapper.Robot.BehaviorControl(ctx, start, stop)
	}()

	for {
		select {
		case <-start:
			if *hue == "" && *saturation == "" {
				for h := 0; h <= 10; h++ {
					strHue := fmt.Sprintf("%f", float64(h)/10)
					println("Hue: " + strHue)
					for s := 0; s <= 10; s++ {
						strSat := fmt.Sprintf("%f", float64(s)/10)
						settings.SetCustomEyeColor(strHue, strSat)
						time.Sleep(time.Duration(50) * time.Millisecond)
					}
				}
			} else {
				settings.SetCustomEyeColor(*hue, *saturation)
			}

			stop <- true
			return
		}
	}
}
