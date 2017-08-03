package main

import (
	"caddr/adcode"
	"caddr/gps"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"encoding/json"
	"bytes"
	"fmt"
)

type AdcodeGenerator struct {
	Adfile string
	Repo *adcode.AdcodeRepo
}

func (g *AdcodeGenerator) Load(filepath string) {
	// adcodes := make([]adcode.Adcode, 0, 3500)
	adcodes := adcode.NewAdcodeRepo()

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// log.Println(line)
		tuple := strings.Split(line, " ")
		code, err := strconv.ParseInt(tuple[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		name := tuple[1]
		ac := adcode.Adcode{Code:code}
		switch {
		case code%100 != 0:
			ac.District = name
		case code%100 == 0 && code%10000 != 0:
			ac.City = name
		case code%10000 == 0:
			ac.Province = name
		default:
			log.Fatal("Invalid Adcode %d\n", code)
		}
		// adcodes = append(adcodes, ac)
		adcodes.Add(ac)
		log.Println(ac)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	g.Repo = adcodes
}

func (g *AdcodeGenerator) RefreshLocation() {
	adcodes := g.Repo.GetAll()
	for _, ac := range adcodes {
		switch {
		case ac.IsL3():
			acl2, err := g.Repo.ToUpper(ac.Code, 2)
			if err != nil {
				panic(err)
			}
			acl1, err := g.Repo.ToUpper(ac.Code, 1)
			if err != nil {
				panic(err)
			}
			ac.City = acl2.City
			ac.Province = acl1.Province
			g.Repo.Update(ac)
		case ac.IsL2():
			acl1, err := g.Repo.ToUpper(ac.Code, 1)
			if err != nil {
				panic(err)
			}
			ac.Province = acl1.Province
			g.Repo.Update(ac)
		}		
	}
}

func (g *AdcodeGenerator) RefreshGps() {
	geo := gps.NewBaiduGeocode("FLS3ZKCg9A89d5uaeV788TzG")
	adcodes := g.Repo.GetAll()
	for _, ac := range adcodes {
		gps, err := geo.Geocoding(ac.String())
		if err != nil {
			log.Printf("[%d] %s\n", ac.Code, err.Error())
		}
		ac.Lng = gps.Lng
		ac.Lat = gps.Lat
		g.Repo.Update(ac)
	}
}

func (g *AdcodeGenerator) DumpJson() {
	js, err := json.MarshalIndent(g.Repo.GetAll(), "    ", "")
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Create(strings.Replace(g.Adfile, ".txt", ".json", 1))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(js)
}

func (g *AdcodeGenerator) DumpGo() {
	var gocode bytes.Buffer
	gocode.WriteString("package adcode\n\n")
	gocode.WriteString("func InitAdcodeRepo() *AdcodeRepo {\n")
	gocode.WriteString("\tvar code Adcode\n")
	gocode.WriteString("\trepo := NewAdcodeRepo()\n")
	for _, code := range g.Repo.GetAll() {
		line := fmt.Sprintf("\tcode = Adcode{Code: %d, Province: \"%s\", City: \"%s\", District: \"%s\", Lng: %f, Lat: %f}\n",
			code.Code, code.Province, code.City, code.District, code.Lng, code.Lat)
		gocode.WriteString(line)
		gocode.WriteString("\trepo.Add(code)\n")

	}
	gocode.WriteString("\treturn repo\n")
	gocode.WriteString("}\n")

	f, err := os.Create("adcodeinit.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(gocode.Bytes())
}

func (g *AdcodeGenerator) Generate() {
	g.Load(g.Adfile)
	g.RefreshLocation()
	g.RefreshGps()
	g.DumpJson()
	g.DumpGo()
}

func main() {
	gen := AdcodeGenerator{Adfile: "adcode_2017.txt"}
	gen.Generate()
}