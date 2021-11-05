package main

import (
	"flag"
	"fmt"
	tplink "github.com/Pluslab/tplink-api"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

func init() {
	flag.String("user", "", "Username")
	flag.String("pass", "", "Password")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)

	if err != nil {
		fmt.Println("Error in bindpflags")
	}
}

func main() {
	api, err := tplink.Connect(viper.GetString("user"), viper.GetString("pass"))
	if err != nil {
		log.Fatal(err)
	}

	// devices, err := api.GetDevicesInfo()
	// fmt.Printf("Devices=%v\n", devices)

	// fmt.Printf("================================================================")
	// fmt.Println(devices)

	hs100, err := api.GetHS100("Shed Large")
	if err != nil {
		log.Fatal(err)
	}

	info, err := hs100.GetInfo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(info)

	aps, err := hs100.ScanAPs()

	fmt.Println(aps)
	hs100.TurnOff()
}
