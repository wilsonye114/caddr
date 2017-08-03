package adcode

import (
	"testing"
	"log"
)

func TestRegisterGeocoding(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	for _, code := range repo.Adcodes {
		log.Println(code)
	}
}

func TestAdcodeL1(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	count := 0
	for _, code := range repo.Adcodes {
		if code.IsL1() {
			log.Println(code)
			count++
		}
	}
	log.Printf("Get %d L1 code\n", count)
}

func TestAdcodeL2(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	count := 0
	for _, code := range repo.Adcodes {
		if code.IsL2() {
			log.Println(code)
			count++
		}
	}
	log.Printf("Get %d L2 code\n", count)
}

func TestAdcodeL3(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	count := 0
	for _, code := range repo.Adcodes {
		if code.IsL3() {
			log.Println(code)
			count++
		}
	}
	log.Printf("Get %d L3 code\n", count)
}

func TestAdcodeRepoToUpper(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	log.Printf("L3 To L1\n")
	ac, err := repo.ToUpper(110111, 1)
	if err != nil {
		panic(err)
	}
	log.Println(ac)

	log.Printf("L3 To L2\n")
	ac, err = repo.ToUpper(110111, 2)
	if err != nil {
		panic(err)
	}
	log.Println(ac)

	// log.Printf("L3 To L3\n")
	// ac, err = repo.ToUpper(110111, 3)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(ac)

	log.Printf("L2 To L1\n")
	ac, err = repo.ToUpper(130200, 1)
	if err != nil {
		panic(err)
	}
	log.Println(ac)

	// log.Printf("L2 To L2\n")
	// ac, err = repo.ToUpper(130200, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(ac)

	// log.Printf("L1 To L1\n")
	// ac, err = repo.ToUpper(410000, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(ac)

	// log.Printf("L1 To L2\n")
	// ac, err = repo.ToUpper(410000, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(ac)

	// log.Printf("L1 To L3\n")
	// ac, err = repo.ToUpper(410000, 3)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(ac)

	// log.Printf("L2 To L1\n")
	// ac, err = repo.ToUpper(411000, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println(ac)
}

func TestAdcodeRepoToLower(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}

	log.Printf("L1 to L3\n")
	list, err := repo.ToLower(110000, 3)
	if err != nil {
		panic(err)
	}
	for _, ac := range list {
		log.Printf("%v\n", ac)
	}

	log.Printf("L1 to L2\n")
	list, err = repo.ToLower(150000, 2)
	if err != nil {
		panic(err)
	}
	for _, ac := range list {
		log.Printf("%v\n", ac)
	}

	// log.Printf("L1 to L1\n")
	// list, err = repo.ToLower(150000, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, ac := range list {
	// 	log.Printf("%v\n", ac)
	// }

	log.Printf("L2 to L3\n")
	list, err = repo.ToLower(411000, 3)
	if err != nil {
		panic(err)
	}
	for _, ac := range list {
		log.Printf("%v\n", ac)
	}

	// log.Printf("L2 to L2\n")
	// list, err = repo.ToLower(411000, 2)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, ac := range list {
	// 	log.Printf("%v\n", ac)
	// }

	// log.Printf("L2 to L1\n")
	// list, err = repo.ToLower(411000, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, ac := range list {
	// 	log.Printf("%v\n", ac)
	// }

	// log.Printf("L3 to L3\n")
	// list, err = repo.ToLower(411002, 3)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, ac := range list {
	// 	log.Printf("%v\n", ac)
	// }
}

func TestAdcodeRepoGetL1(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.GetL1()
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d L1 Adcodes\n", i)
}

func TestAdcodeRepoGetL2(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.GetL2()
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d L2 Adcodes\n", i)
}

func TestAdcodeRepoGetL3(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.GetL3()
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d L3 Adcodes\n", i)
}

func TestAdcodeRepoGetAll(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.GetAll()
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d Adcodes\n", i)
}

func TestAdcodeCheckIsProvince(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsProvince)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 省\n", i)
}

func TestAdcodeCheckIsMuniRegion(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsMuniRegion)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 自治区\n", i)
}

func TestAdcodeCheckIsMuniCity(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsMuniCity)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 直辖市\n", i)
}

func TestAdcodeCheckIsSpecialRegion(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsSpecialRegion)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 特别行政区\n", i)
}

func TestAdcodeCheckIsL2City(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsL2City)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 市(L2)\n", i)
}

func TestAdcodeCheckIsCountyAdminDivisionUnderProvince(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsCountyAdminDivisionUnderProvince)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 省直辖县级行政规区划\n", i)
}

func TestAdcodeCheckIsCountyAdminDivisionUnderMuniRegion(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsCountyAdminDivisionUnderMuniRegion)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 自治区直辖县级行政区划\n", i)
}

func TestAdcodeCheckIsDivision(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsDivision)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 地区\n", i)
}

func TestAdcodeCheckIsAutonomousPrefecture(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsAutonomousPrefecture)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 自治州\n", i)
}

func TestAdcodeCheckIsDistrict(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsDistrict)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 区\n", i)
}

func TestAdcodeCheckIsCounty(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsCounty)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 县\n", i)
}

func TestAdcodeCheckIsPrefectureLevelCity(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsPrefectureLevelCity)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 地级市\n", i)
}

func TestAdcodeCheckIsCity(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsCity)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 市\n", i)
}

func TestAdcodeCheckIsMuniDistrict(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	acs := repo.Search(IsMuniDistrict)
	i := 0
	for _, ac := range acs {
		log.Println(ac)
		i++
	}
	log.Printf("%d 市辖区\n", i)
}

func TestAdcodeRepoCRUD(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	cds := repo.GetAll()
	len1 := repo.Len()
	repo.Remove(cds[0].Code)
	len2 := repo.Len()
	log.Printf("Remove %d adcode\n", len1-len2)
	if len1 != len2+1 {
		panic("AdcodesRepo Remove Error")
	}
	repo.RemoveCodes(cds.ToIntSlice())
	len3 := repo.Len()
	log.Printf("Remove %d adcode\n", len2-len3)
	if len3 != 0 {
		panic("AdcodesRepo Remove Error")
	}
	repo.AddAdcodes(cds)
	len4 := repo.Len()
	log.Printf("Add %d adcode\n", len4-len3)
	if len4 != len(cds) {
		panic("AdcodesRepo Add Error")
	}
	repo.RemoveAdcodes(cds[3:13])
	len5 := repo.Len()
	log.Printf("Remove %d adcode\n", len4-len5)
	if len4 != len(cds) {
		panic("AdcodesRepo remove Error")
	}
	log.Printf("%d is in repo, %v\n", cds[2].Code, repo.IsInRepo(cds[2].Code))
	log.Printf("%d is in repo, %v\n", cds[3].Code, repo.IsInRepo(cds[3].Code))
}

func TestAdcodeRepoDup(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	newrepo := repo.Dup()
	log.Printf("New repo length = %d\n", newrepo.Len())

	repodup, err := DupAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	repodup.RemoveAll()
	log.Printf("repo %d adcodes, repodup %d adcodes\n", repo.Len(), repodup.Len())
}

func TestAdcodeString(t *testing.T) {
	repo, err := GetAdcodeRepo("default")
	if err != nil {
		panic(err)
	}
	list := repo.GetAll()
	for _, ac := range list {
		log.Printf("%s\n", ac.String())
	}
}

