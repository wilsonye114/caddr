package gps

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"io/ioutil"
)

/***************************************************************************
* Addr Interface
***************************************************************************/
type Addr interface {
	Lng() float64
	Lat() float64
	Adcode() int64
	Country() string
	Province() string
	City() string
	District() string
	Street() string
	StreetNumber() string
	Direction() string
	Distance() string
	FormattedAddr() string
	Description() string
}

/***************************************************************************
* GpsDmd - GPS Degrees/Minutes/Seconds
*
* Format: DDD MM SS
***************************************************************************/
type GpsDms struct {
	Lat struct {
		Degrees int
		Minutes int
		Seconds int
	}
	Lng struct {
		Degrees int
		Minutes int
		Seconds int
	}
}

/***************************************************************************
* GpsGC - GPS Decimal Degrees
*
* Format: DDD MM.MMMMM
***************************************************************************/
type GpsGC struct {
	Lat struct {
		Degrees int
		Minutes float64
	}
	Lng struct {
		Degrees int
		Minutes float64
	}
}

/***************************************************************************
* GpsDD - GPS Coordinate
*
* Format: DDD.DDDD
***************************************************************************/
type GpsDD struct {
	Lat float64
	Lng float64
}


/***************************************************************************
* GeocodeService
* 
* Geocode service supports geocoding and reverse geocoding.
***************************************************************************/

type GeocodeService interface {
	Geocoding(string) (*GpsDD, error)
	Regeocoding(*GpsDD) (Addr, error)
}


/***************************************************************************
* Baidu Geocide Service
***************************************************************************/
type BaiduGeocodeRes struct {
	Status int `json:"status"`
	Result struct {
		Location struct {
			Lng float64
			Lat float64
		}
		Precise int64
		Confidence int64
		Level string
	}
}

type BaiduRegeocodeRes struct {
	Status int `json:"status"`
	Result struct {
		BaiduAddr
	} `json:"result"`
}

type BaiduRegeocodeResErr struct {
	Status int64 `json:"status"`
	Message string `json:"message"`
}

type BaiduAddr struct {
	Location struct {
		Lng float64 `json:"lng"`
		Lat float64 `json:"lat"`
	} `json:"location"`
	FormattedAddress string `json:"formatted_address"`
	Business string `json:"business"`
	AddressComponent struct {
		Country string `json:"country"`
		CountryCode int `json:"country_code"`
		Province string `json:"province"`
		City string `json:"city"`
		District string `json:"district"`
		Adcode string `json:"adcode"`
		Street string `json:"street"`
		StreetNumber string `json:"street_number"`
		Direction string `json:"direction"`
		Distance string `json:"distance"`
	} `json:"addressComponent"`
	Pois []struct {
		Addr string `json:"addr"`
		Cp string `json:"cp"`
		Direction string `json:"direction"`
		Distance string `json:"distance"`
		Name string `json:"name"`
		PoiType string `json:"poiType"`
		Point struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"point"`
		Tag string `json:"tag"`
		Tel string `json:"tel"`
		Uid string `json:"uid"`
		Zip string `json:"zip"`
	} `json:"pois"`
	PoiRegions []struct {
		DirectionDesc string `json:"direction_desc"`
		Name string `json:"name"`
		Tag string `json:"tag"`
	} `json:"poiRegions"`
	SematicDescription string `json:"sematic_description"`
	Citycode int `json:"cityCode"`
}

func (a *BaiduAddr) Lng() float64 {
	return a.Location.Lng
}

func (a *BaiduAddr) Lat() float64 {
	return a.Location.Lat
}

func (a *BaiduAddr)	Country() string {
	return a.AddressComponent.Country
}

func (a *BaiduAddr)	CountryCode() int {
	return a.AddressComponent.CountryCode
}

func (a *BaiduAddr)	Province() string {
	return a.AddressComponent.Province
}

func (a *BaiduAddr)	City() string {
	return a.AddressComponent.City
}

func (a *BaiduAddr)	CityCode() int {
	return a.Citycode
}

func (a *BaiduAddr)	District() string {
	return a.AddressComponent.District
}

func (a *BaiduAddr)	Adcode() int64 {
	code, err := strconv.ParseInt(a.AddressComponent.Adcode, 10, 64)
	if err != nil {
		return -1
	}
	return code
}

func (a *BaiduAddr)	AdcodeString() string {
	return a.AddressComponent.Adcode
}

func (a *BaiduAddr)	Street() string {
	return a.AddressComponent.Street
}

func (a *BaiduAddr)	StreetNumber() string {
	return a.AddressComponent.StreetNumber
}

func (a *BaiduAddr)	Direction() string {
	return a.AddressComponent.Direction
}

func (a *BaiduAddr)	Distance() string {
	return a.AddressComponent.Distance
}

func (a *BaiduAddr)	FormattedAddr() string {
	return a.FormattedAddress
}

func (a *BaiduAddr)	Description() string {
	return a.SematicDescription
}

// Baidu Coordinate Type
const (
	BC_BD09II = "bd09ll"
	BC_BD09MC = "bd09mc"
	BC_GCJ02II = "gcj02ll"
	BC_WGS84II = "wgs84ll"
)

type BaiduGeocode struct {
	url string
	ak string
	coordtype string
}

func (g *BaiduGeocode) AK() string {
	return g.ak
}

func (g *BaiduGeocode) SetAK(ak string) {
	g.ak = ak
}

func (g *BaiduGeocode) Url() string {
	return g.url
}

func (g *BaiduGeocode) SetUrl(url string) {
	g.url = url
}

func (g *BaiduGeocode) Coordtype() string {
	return g.coordtype
}

func (g *BaiduGeocode) SetCoordtype(coordtype string) {
	g.coordtype = coordtype
}

func (g *BaiduGeocode) Geocoding(addr string) (*GpsDD, error) {
	var res BaiduGeocodeRes
	var reserr BaiduRegeocodeResErr

	url := fmt.Sprintf("%s/?address=%s&output=json&ak=%s", g.url, addr, g.ak)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %v", url, err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &reserr)
	if reserr.Status != 0 {
		return nil, fmt.Errorf("BaiduGeocode Error %d: %s", reserr.Status, reserr.Message)
	}
	json.Unmarshal(b, &res)
	gps := GpsDD {
		Lat:res.Result.Location.Lat,
		Lng:res.Result.Location.Lng}
	return &gps, nil
}

func (g *BaiduGeocode) Regeocoding(gps *GpsDD) (Addr, error) {
	var res BaiduRegeocodeRes
	var reserr BaiduRegeocodeResErr

	url := fmt.Sprintf("%s/?coordtype=%s&location=%f,%f&output=json&pois=1&ak=%s", g.url, g.coordtype, gps.Lat, gps.Lng, g.ak)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return &BaiduAddr{}, fmt.Errorf("fetch %s: %v", url, err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &reserr)
	if reserr.Status != 0 {
		return &BaiduAddr{}, fmt.Errorf("BaiduGeocode Error %d: %s", reserr.Status, reserr.Message)
	}
	json.Unmarshal(b, &res)

	return &res.Result.BaiduAddr, nil
}

func NewBaiduGeocode(ak string) *BaiduGeocode {
	return &BaiduGeocode{
		url: "http://api.map.baidu.com/geocoder/v2",
		ak: ak,
		coordtype: BC_WGS84II}
}

var (
	gGeocodeServiceCache map[string]GeocodeService
	gGeocodeServiceCacheInited bool = false
)
func RegisterGeocode(name string, geo GeocodeService) {
	if gGeocodeServiceCacheInited == false {
		gGeocodeServiceCache = make(map[string]GeocodeService)
		gGeocodeServiceCacheInited = true
	}
	gGeocodeServiceCache[name] = geo
}

func GetGeocodeService(name string) GeocodeService {
	if gGeocodeServiceCacheInited == false {
		panic("No Geocode Service available!")
	}	
	geo, ok := gGeocodeServiceCache[name]
	if !ok {
		err := fmt.Errorf("%s is not registered!", name)
		panic(err)
	}
	return geo
}

