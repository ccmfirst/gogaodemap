package getfun

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 高德地图Key
const DefaultKey = "55b99a5145f6cd5fe33a4a92e93add0b"

// 步行路径规划URL前缀
const ReqUrlForWalk = "http://restapi.amap.com/v3/direction/walking?origin="

// 骑行路径规划URL前缀
const ReqUrlForBicycle = "http://restapi.amap.com/v4/direction/bicycling?origin="

// 地理编码：通过地址获取坐标URL前缀
const ReqUrlForGeo = "https://restapi.amap.com/v3/geocode/geo?address="

// 逆地理编码： 通过坐标获取地址URL前缀
const ReqUrlForReGeo = "https://restapi.amap.com/v3/geocode/regeo?output=JSON&radius=1000&extensions=all&location="

type Location struct {
	Lng string
	Lat string
}

type Step struct {
	Instruction string
	//Orientation string
	//Road            string
	Distance string
	Duration string
	Polyline string
	//Action          string
	AssistantAction string
	WalkType        string
}

type WalkPath struct {
	Distance string
	Duration string
	//Steps    []Step
}

type BicyclePath struct {
	Distance int64
	Duration int64
	//Steps    []Step
}

type WalkResp struct {
	Status string
	Info   string
	Count  string
	Route  struct {
		Paths []WalkPath
	}
}

type BicycleResp struct {
	Data struct {
		Origin      string
		Destination string
		Paths       []BicyclePath
	}
}

type GeoCode struct {
	Formatted_Address string
	Country           string
	Province          string
	CityCode          string
	City              string
	District          string
	Township          []string
	Location          string
}

type GetLocationResp struct {
	Status   string
	Info     string
	InfoCode string
	Count    string
	GeoCodes []GeoCode
}

type GetAddressResp struct {
	Status    string
	ReGeoCode struct {
		Formatted_Address string
	}
}

func GetURL(origin, destination Location, repType string) string {
	var url string
	if repType == "walking" {
		url = ReqUrlForWalk + origin.Lng + "," + origin.Lat + "&" + "destination=" +
			destination.Lng + "," + destination.Lat + "&output=json&key=" + DefaultKey
	}

	if repType == "bicycling" {
		url = ReqUrlForBicycle + origin.Lng + "," + origin.Lat + "&" + "destination=" +
			destination.Lng + "," + destination.Lat + "&output=json&key=" + DefaultKey
	}

	return url
}

func GetDuration(origin, destination Location, repType string) (int64, error) {
	var res interface{}
	if repType == "walking" {
		res = new(WalkResp)
	}

	if repType == "bicycling" {
		res = new(BicycleResp)
	}
	url := GetURL(origin, destination, repType)
	fmt.Println(url)
	httpClient := http.Client{}
	resp, err := httpClient.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(bytes))
	if resp.StatusCode == 200 {
		err := json.Unmarshal(bytes, &res)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
	} else {
		return 0, errors.New("请求高德API失败，状态码不等于200")
	}
	fmt.Println(res)

	return 0, nil
}

func GetLocation(address string) (Location, error) {
	url := ReqUrlForGeo + address + "&output=JSO&key=" + DefaultKey
	var res GetLocationResp
	httpClient := http.Client{}
	resp, err := httpClient.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		fmt.Println(err)
		return Location{}, err
	}
	fmt.Println(res.GeoCodes[0].Location, res.GeoCodes[0].Formatted_Address)
	location := strings.Split(res.GeoCodes[0].Location, ",")
	fmt.Println(location)

	return Location{Lng: location[0], Lat: location[1]}, nil
}

func GetAddress(location Location) (string, error) {
	var res GetAddressResp
	url := ReqUrlForReGeo + location.Lng + "," + location.Lat + "&key=" + DefaultKey
	httpClient := http.Client{}
	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(res.ReGeoCode.Formatted_Address)
	return res.ReGeoCode.Formatted_Address, nil
}
