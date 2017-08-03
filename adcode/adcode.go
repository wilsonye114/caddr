package adcode

import (
	"fmt"
	"regexp"
	"errors"
	"sort"
)

var (
	ErrNoAdcode = errors.New("No Adcode found!")
)

/****************************************************************************
* Administration Division Code
****************************************************************************/
type Adcode struct {
	Code int64
	Province string
	City string
	District string
	Lng float64
	Lat float64
}

func (a *Adcode) String() string {
	re1 := regexp.MustCompile("市辖区$")
	re2 := regexp.MustCompile("省直辖县级行政区划$")
	re3 := regexp.MustCompile("自治区直辖县级行政区划$")

	s := ""
	switch {
	case re1.Match([]byte(a.City)):
		s = a.Province + a.District
	case re2.Match([]byte(a.City)):
		s = a.Province + a.District
	case re3.Match([]byte(a.City)):
		s = a.Province + a.District
	case re1.Match([]byte(a.District)):
		s = a.Province + a.City
	default:
		s = a.Province + a.City + a.District
	}
	return s
}

func (a *Adcode) ToUpperCode() int64 {
	code := int64(0)
	switch {
	case a.IsL1():
		code = a.Code
	case a.IsL2():
		code = a.Code/10000*10000
	case a.IsL3():
		code = a.Code/100*100
	default:
		panic("Unknown Adcode Level")
	}
	return code
}

func (a *Adcode) IsUpper(ac Adcode) bool {
	if a.Level() >= ac.Level() {
		return false
	}

	res := false
	switch {
	case a.IsL1():
		if (a.Code == (ac.Code/10000*10000)) || (a.Code == (ac.Code/100*100)) {
			res = true
		}
	case a.IsL2():
		if a.Code == (ac.Code/100*100) {
			res = true
		}
	}
	return res
}

func (a *Adcode) Level() int {
	level := 0
	switch {
	case a.IsL1():
		level = 1
	case a.IsL2():
		level = 2
	case a.IsL3():
		level = 3
	default:
		panic("Unknown Adcode Level")
	}
	return level
}

// Province
func (a *Adcode) IsL1() bool {
	if a.Code%10000 == 0 {
		return true
	}
	return false
}

// City
func (a *Adcode) IsL2() bool {
	if a.Code%100 == 0 && a.Code%10000 != 0 {
		return true
	}
	return false
}

// District
func (a *Adcode) IsL3() bool {
	if a.Code%100 != 0 {
		return true
	}
	return false
}

/***************************************************************************
* AdcodeCheckFunc
***************************************************************************/
type AdcodeCheckFunc func(a Adcode)bool

// Province 省
func IsProvince(a Adcode) bool {
	re := regexp.MustCompile("省$")
	if a.IsL1() && re.Match([]byte(a.Province)) {
		return true
	}
	return false
}

// municipality region 自治区
func IsMuniRegion(a Adcode) bool {
	re := regexp.MustCompile("自治区$")
	if a.IsL1() && re.Match([]byte(a.Province)) {
		return true
	}
	return false
}

// municipality directly under the Central Government 直辖市
func IsMuniCity(a Adcode) bool {
	re := regexp.MustCompile("市$")
	if a.IsL1() && re.Match([]byte(a.Province)) {
		return true
	}
	return false
}

// Special Administrative Region 特别行政区
func IsSpecialRegion(a Adcode) bool {
	re := regexp.MustCompile("特别行政区$")
	if a.IsL1() && re.Match([]byte(a.Province)) {
		return true
	}
	return false
}

func IsL2City(a Adcode) bool {
	re := regexp.MustCompile("市$")
	if a.IsL2() && re.Match([]byte(a.City)) {
		return true
	}
	return false
}

// county administrative division directly under province 省直辖县级行政规区划
func IsCountyAdminDivisionUnderProvince(a Adcode) bool {
	re := regexp.MustCompile("省直辖县级行政区划$")
	if a.IsL2() && re.Match([]byte(a.City)) {
		return true
	}
	return false
}

// county administrative division directly under region 自治区直辖县级行政区划
func IsCountyAdminDivisionUnderMuniRegion(a Adcode) bool {
	re := regexp.MustCompile("自治区直辖县级行政区划$")
	if a.IsL2() && re.Match([]byte(a.City)) {
		return true
	}
	return false
}

// Division 地区
func IsDivision(a Adcode) bool {
	re := regexp.MustCompile("地区$")
	if a.IsL2() && re.Match([]byte(a.City)) {
		return true
	}
	return false
}

// Autonomous Prefecture 自治州
func IsAutonomousPrefecture(a Adcode) bool {
	re := regexp.MustCompile("自治州$")
	if a.IsL2() && re.Match([]byte(a.City)) {
		return true
	}
	return false
}

// District 区
func IsDistrict(a Adcode) bool {
	re := regexp.MustCompile("区$")
	if a.IsL3() && re.Match([]byte(a.District)) {
		return true
	}
	return false
}

// County 县
func IsCounty(a Adcode) bool {
	re := regexp.MustCompile("县$")
	if a.IsL3() && re.Match([]byte(a.District)) {
		return true
	}
	return false
}

// prefecture-level city 地级市
func IsPrefectureLevelCity(a Adcode) bool {
	re := regexp.MustCompile("市$")
	if a.IsL3() && re.Match([]byte(a.District)) {
		return true
	}
	return false
}

// 市(L1-L3)
func IsCity(a Adcode) bool {
	re := regexp.MustCompile("市$")
	if re.Match([]byte(a.Province)) || re.Match([]byte(a.City)) || re.Match([]byte(a.District)) {
		return true
	}
	return false
}

// Municipal district 市辖区
func IsMuniDistrict(a Adcode) bool {
	re := regexp.MustCompile("市辖区")
	if (a.IsL2() && re.Match([]byte(a.City))) ||
		(a.IsL3() && re.Match([]byte(a.District))) {
		return true
	}
	return false
}

/***************************************************************************
* AdcodeList
***************************************************************************/
type AdcodeList []Adcode

func (l AdcodeList) Len() int {
	return len(l)
}

func (l AdcodeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l AdcodeList) Less(i, j int) bool {
	return l[i].Code < l[j].Code
}

func (l AdcodeList) ToIntSlice() []int64 {
	list := make([]int64, len(l))
	for i, v := range l {
		list[i] = v.Code
	}
	return list
}

/***************************************************************************
* AdcodeRepo
***************************************************************************/
type AdcodeRepo struct {
	Adcodes map[int64]Adcode
}

func (r *AdcodeRepo) Len() int {
	return len(r.Adcodes)
}

func (r *AdcodeRepo) Add(ac Adcode) {
	r.Adcodes[ac.Code] = ac
}

func (r *AdcodeRepo) AddAdcodes(acs AdcodeList) {
	for _, ac := range acs {
		r.Adcodes[ac.Code] = ac
	}
}

func (r *AdcodeRepo) Remove(code int64) {
	delete(r.Adcodes, code)
}

func (r *AdcodeRepo) RemoveAdcodes(acs AdcodeList) {
	for _, ac := range acs {
		delete(r.Adcodes, ac.Code)
	}
}

func (r *AdcodeRepo) RemoveCodes(codes []int64) {
	for _, code := range codes {
		delete(r.Adcodes, code)
	}
}

func (r *AdcodeRepo) RemoveAll() {
	r.Adcodes = make(map[int64]Adcode)
}

func (r *AdcodeRepo) Update(ac Adcode) {
	r.Adcodes[ac.Code] = ac
}

func (r *AdcodeRepo) UpdateAdcodes(acs AdcodeList) {
	for _, ac := range acs {
		r.Adcodes[ac.Code] = ac
	}
}

func (r *AdcodeRepo) IsInRepo(code int64) bool {
	_, ok := r.Adcodes[code]
	return ok
}

func (r *AdcodeRepo) Get(code int64) (Adcode, error) {
	ac, ok := r.Adcodes[code]
	if ok {
		return ac, nil
	}
	return Adcode{}, ErrNoAdcode
}

func (r *AdcodeRepo) GetCodes(codes []int64) (AdcodeList, error) {
	list := make(AdcodeList, len(codes))
	for i, code := range codes {
		if ac, ok := r.Adcodes[code]; ok {
			list[i] = ac
		} else {
			return list, fmt.Errorf("Unknown adcode %d", code)
		}
	}
	return list, nil
}

func (r *AdcodeRepo) GetAll() AdcodeList {
	count := len(r.Adcodes)
	list := make(AdcodeList, count)
	i := 0
	for _, v := range r.Adcodes {
		list[i] = v
		i++
	}
	sort.Sort(list)
	return list
}

func (r *AdcodeRepo) GetL1() AdcodeList {
	list := make(AdcodeList, 0)
	for _, ac := range r.Adcodes {
		if ac.IsL1() {
			list = append(list, ac)
		}
	}
	sort.Sort(list)
	return list
}

func (r *AdcodeRepo) GetL2() AdcodeList {
	list := make(AdcodeList, 0)
	for _, ac := range r.Adcodes {
		if ac.IsL2() {
			list = append(list, ac)
		}
	}
	sort.Sort(list)
	return list
}

func (r *AdcodeRepo) GetL3() AdcodeList {
	list := make(AdcodeList, 0)
	for _, ac := range r.Adcodes {
		if ac.IsL3() {
			list = append(list, ac)
		}
	}
	sort.Sort(list)
	return list
}

func (r *AdcodeRepo) Search(asf AdcodeCheckFunc) AdcodeList {
	list := make(AdcodeList, 0)
	for _, ac := range r.Adcodes {
		if asf(ac) {
			list = append(list, ac)
		}
	}
	sort.Sort(list)
	return list
}

func (r *AdcodeRepo) ToUpper(code int64, level int) (Adcode, error) {
	var acup Adcode

	accur, err := r.Get(code)
	if err != nil {
		return accur, err
	}
	curlevel := accur.Level()
	// ToUpper: [Upper] L1 <--- L2 <--- L3 [Lower]
	switch {
	case curlevel < level:
		return accur, fmt.Errorf("Current adcode level is %d, higher than %d", curlevel, level)
	case curlevel == level:
		return accur, fmt.Errorf("Current adcode level is already %d", curlevel)
	case curlevel > level:
		for actmp := accur; actmp.Level() != level; actmp = acup {
			var err error
			codeup := actmp.ToUpperCode()
			acup, err = r.Get(codeup)
			if err != nil {
				return acup, err
			}
		}
	}
	return acup, nil
}

func (r *AdcodeRepo) ToLower(code int64, level int) (AdcodeList, error) {
	list := make(AdcodeList, 0)
	accur, err := r.Get(code)
	if err != nil {
		list = append(list, accur)
		return list, err
	}
	curlevel := accur.Level()
	// ToLower: [Upper] L1 ---> L2 ---> L3 [Lower]
	switch {
	case curlevel < level:
		for _, ac := range r.Adcodes {
			if  ac.Level() == level && accur.IsUpper(ac) {
				list = append(list, ac)
			}
		}
	case curlevel == level:
		list = append(list, accur)
		return list, fmt.Errorf("Current adcode level is already %d", curlevel)
	case curlevel > level:
		list = append(list, accur)
		return list, fmt.Errorf("Current adcode level is %d, higher than %d", curlevel, level)
	}
	sort.Sort(list)
	return list, nil
}

func (r *AdcodeRepo) Dup() *AdcodeRepo {
	newrepo := AdcodeRepo{Adcodes: make(map[int64]Adcode)}
	for k, v := range r.Adcodes {
		newrepo.Adcodes[k] = v
	}
	return &newrepo
}

func NewAdcodeRepo() *AdcodeRepo {
	repo := AdcodeRepo{Adcodes: make(map[int64]Adcode)}
	return &repo
}

/***************************************************************************
* Global AdcodeRepos
***************************************************************************/
var(
	gAdcodeRepos map[string]*AdcodeRepo
	gAdcodeRepoCacheInited bool = false
)

func InitAdcodeRepos() {
	if gAdcodeRepoCacheInited == false {
		gAdcodeRepos = make(map[string]*AdcodeRepo)
		gAdcodeRepoCacheInited = true
	}

	repo := InitAdcodeRepo2017()
	RegisterAdcodeRepo("2017", repo)
	RegisterAdcodeRepo("default", repo)
}

func RegisterAdcodeRepo(name string, repo *AdcodeRepo) {
	gAdcodeRepos[name] = repo
}

func GetAdcodeRepo(name string) (*AdcodeRepo, error) {
	if v, ok := gAdcodeRepos[name]; ok {
		return v,  nil
	}
	return nil, fmt.Errorf("Unknown AdcodeRepo %s", name)
}

func DupAdcodeRepo(name string) (*AdcodeRepo, error) {
	if v, ok := gAdcodeRepos[name]; ok {
		return v.Dup(),  nil
	}
	return nil, fmt.Errorf("Unknown AdcodeRepo %s", name)
}

func init() {
	InitAdcodeRepos()
}