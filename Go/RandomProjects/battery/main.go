package main

import "fmt"

//var(
//	vanSpecs = map[string]float32{"Height": 110.4, "Length": 263.9, "Width": 81.3}
//)

func main() {
	fmt.Printf("Watthours: %5.2f\n", AhToWh(200, 12))

	var ankerAh float32 = 20100.0 / 1000.0 //20Ah
	var switchAh float32 = 4310.0 / 1000.0 //4.3Ah
	var switchDraw float32 = 18            //Wh

	batteryLife := calculateBatteryLife(map[string]float32{"switch": switchDraw}, switchAh)
	fmt.Printf("%5.2f hours: Switch Battery Life\n", batteryLife)

	batteryLife = calculateBatteryLife(map[string]float32{"switch": switchDraw}, switchAh+ankerAh)
	fmt.Printf("%5.2f hours: Switch + Anker Battery Life\n", batteryLife)

	drawOnSystem := map[string]float32{"Mac": 50, "ARB": 60, "Monitor": 85, "Led": 27}
	var sum float32 = 0
	for _, v := range drawOnSystem {
		sum += v
	}
	fmt.Printf("%4.1f Sum of Watts\n", sum)

	var batteryAh float32 = 110
	var solarPanelWatts float32 = 600

	batteryLife = calculateBatteryLife(drawOnSystem, batteryAh)
	fmt.Printf("%5.2f hours: Van Battery Life\n", batteryLife)

	rechargeTime := batteryRechargeTime(solarPanelWatts, 24, batteryAh)
	fmt.Printf("%5.2f hours: Charge Time\n", rechargeTime)

}
func calculateBatteryLife(usage map[string]float32, batteryAh float32) float32 {
	var wattUsageSum float32 = 0
	for _, v := range usage {
		wattUsageSum += v
	}
	batteryWh := AhToWh(batteryAh, 12)
	return batteryWh / wattUsageSum
}
func batteryRechargeTime(wattCharge float32, volts float32, batteryAmpHours float32) float32 {
	ampHours := wattsToAmp(wattCharge, volts)
	//fmt.Printf("%5.2f Amp hours\n", ampHours)
	return batteryAmpHours / (ampHours * .8)
}

func wattsToAmp(watts float32, volts float32) float32 {
	//amp = watts / volts
	//fmt.Printf("%vw / %vv = %vah \n ", watts, volts, watts / volts)
	return watts / volts
}

func AhToWh(ampHour float32, volts float32) float32 {
	//watts = amp * volts
	return ampHour * volts
}
