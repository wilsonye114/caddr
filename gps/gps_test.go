package gps

import (
	"testing"
	"log"
)

func ShowAddr(addr Addr) {
	log.Printf("Lng: %f\n", addr.Lng())
	log.Printf("Lat: %f\n", addr.Lat())
	log.Printf("Country: %s\n", addr.Country())
	log.Printf("Province: %s\n", addr.Province())
	log.Printf("City: %s\n", addr.City())
	log.Printf("District: %s\n", addr.District())
	log.Printf("Adcode: %d\n", addr.Adcode())
	log.Printf("Street: %s\n", addr.Street())
	log.Printf("StreetNumber: %s\n", addr.StreetNumber())
	log.Printf("Direction: %s\n", addr.Direction())
	log.Printf("Distance: %s\n", addr.Distance())
	log.Printf("FormattedAddr: %s\n", addr.FormattedAddr())
	log.Printf("Description: %s\n", addr.Description())
}

func TestRegisterGeocoding(t *testing.T) {
	log.Println("==================== TestRegisterGeocoding ====================")
	geo := NewBaiduGeocode("FLS3ZKCg9A89d5uaeV788TzG")
	RegisterGeocode("default", geo)
	defaultGeo := GetGeocodeService("default")
	gps, err := defaultGeo.Geocoding("北京市")
	if err != nil {
		panic(err)
	}
	log.Printf("%v\n", gps)
}

func TestBaiduGeocoding(t *testing.T) {
	log.Println("==================== TestBaiduGeocoding ====================")
	geo := NewBaiduGeocode("FLS3ZKCg9A89d5uaeV788TzG")
	gps, err := geo.Geocoding("北京市")
	if err != nil {
		panic(err)
	}
	log.Printf("%v\n", gps)
}

func TestBaiduRegeocoding(t *testing.T) {
	log.Println("==================== TestBaiduGeocoding ====================")
	geo := NewBaiduGeocode("FLS3ZKCg9A89d5uaeV788TzG")
	gps := &GpsDD{Lng: 121.328148, Lat: 31.224945}
	addr, err := geo.Regeocoding(gps)
	if err != nil {
		panic(err)
	}
	ShowAddr(addr)
}