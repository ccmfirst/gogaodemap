package main

import "workspace/gogaodemap/getfun"

func main() {
	origin := getfun.Location{
		Lng: "121.499740",
		Lat: "31.239853",
	}
	// 上海市普陀区十八英尺公馆
	destination := getfun.Location{
		Lng: "121.421205",
		Lat: "31.257776",
	}
	getfun.GetDuration(origin, destination, "bicycling")
}
