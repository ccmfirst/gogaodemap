package getfun

import (
	"fmt"
	"testing"
)

func TestGetDuration(t *testing.T) {
	// 上海市东方明珠
	origin := Location{
		Lng: "121.499740",
		Lat: "31.239853",
	}
	// 上海市普陀区十八英尺公馆
	destination := Location{
		Lng: "121.421205",
		Lat: "31.257776",
	}
	GetDuration(origin, destination, "bicycling")
}

func TestGetLocation(t *testing.T) {
	l, _ := GetLocation("上海市十八英尺公馆")
	fmt.Println(l)
}

func TestGetAddress(t *testing.T) {
	GetAddress(Location{
		Lng: "121.499740",
		Lat: "31.239853",
	})
}
