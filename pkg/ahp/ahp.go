package ahp

import (
	"ahp-be/internal/model"
	"encoding/json"
	"os"
)

type SubCriteria map[string]float64

func GetRatioIndex() [15]float64 {
	return [15]float64{0, 0, 0.58, 0.9, 1.12, 1.24, 1.32, 1.41, 1.46, 1.49, 1.51, 1.48, 1.56, 1.57, 1.59}
}

func TimbulanSampahSubCriteria(str string) float64 {
	if str == "daerah di jalan protokol pusat kota" {
		return pusatKota
	} else if str == "daerah komersil" {
		return komersil
	} else if str == "daerah perumahan teratur" {
		return perumahanTeratur
	} else if str == "daerah industri" {
		return industri
	} else if str == "jalan taman dan hutan kota" {
		return jalan
	} else if str == "daerah perumahan tidak teratur" {
		return perumahanTidakTeratur
	} else {
		return 0
	}
}

func JarakTPASubCriteria(str string) float64 {
	if str == "pelayanan intensif" {
		return pelayananIntensif
	} else if str == "pelayanan menengah" {
		return pelayananMenengah
	} else if str == "pelayanan rendah" {
		return PelayananRendah
	} else {
		return 0
	}
}

func JarakPemukimanSubCriteria(num float64) float64 {
	if num >= 0 && num <= 100 {
		return jarak1
	} else if num >= 101 && num <= 200 {
		return jarak2
	} else if num >= 201 && num <= 300 {
		return jarak3
	} else if num >= 301 && num <= 400 {
		return jarak4
	} else if num >= 401 && num <= 500 {
		return jarak5
	} else {
		return 0
	}
}

func JarakSungaiSubCriteria(str string) float64 {
	if str == "sangat layak" {
		return sungaiSangatLayak
	} else if str == "layak" {
		return sungaiLayak
	} else if str == "cukup layak" {
		return sungaiCukupLayak
	} else if str == "kurang layak" {
		return sungaiKurangLayak
	} else if str == "tidak layak" {
		return sungaiTidakLayak
	} else {
		return 0
	}
}

func PartisipasiMasyarakatSubCriteria(num float64) float64 {
	if num <= 20 {
		return partisipasi1
	} else if num >= 21 && num <= 40 {
		return partisipasi2
	} else if num >= 41 && num <= 60 {
		return partisipasi3
	} else if num >= 61 && num <= 80 {
		return partisipasi4
	} else if num >= 81 {
		return partisipasi5
	} else {
		return 0
	}
}

func CakupanRumahSubCriteria(num float64) float64 {
	if num <= 40 {
		return cakupan1
	} else if num >= 41 && num <= 80 {
		return cakupan2
	} else if num >= 81 && num <= 120 {
		return cakupan3
	} else if num >= 121 && num <= 160 {
		return cakupan4
	} else if num >= 161 {
		return cakupan5
	} else {
		return 0
	}
}

func AksesibilitasSubCriteria(str string) float64 {
	if str == "sangat layak" {
		return aksesibilitasSangatLayak
	} else if str == "layak" {
		return aksesibilitasLayak
	} else if str == "tidak layak" {
		return aksesibilitasTidakLayak
	} else {
		return 0
	}
}

func CreateSubCriteriaFile() {
	data := model.SubCriteria{}

	file, _ := json.MarshalIndent(data, "", "")

	_ = os.WriteFile("assets/subCriteria.json", file, 0644)
}
