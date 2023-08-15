package ahp

import (
	"ahp-be/internal/model"
	"ahp-be/pkg/constants"
	"encoding/json"
	"os"
)

type SubCriteria map[string]float64

func GetRatioIndex() [15]float64 {
	return [15]float64{0, 0, 0.58, 0.9, 1.12, 1.24, 1.32, 1.41, 1.46, 1.49, 1.51, 1.48, 1.56, 1.57, 1.59}
}

func TimbulanSampahSubCriteria(str string) float64 {
	data, _ := ReadSubCriteriaFile()
	timbulanSampah := data.TimbulanSampah

	if str == "daerah di jalan protokol pusat kota" {
		return timbulanSampah.PusatKota
	} else if str == "daerah komersil" {
		return timbulanSampah.DaerahKomersil
	} else if str == "daerah perumahan teratur" {
		return timbulanSampah.DaerahPerumahanTeratur
	} else if str == "daerah industri" {
		return timbulanSampah.Industri
	} else if str == "jalan taman dan hutan kota" {
		return timbulanSampah.Jalan
	} else if str == "daerah perumahan tidak teratur" {
		return timbulanSampah.DaerahPerumahanTidakTeratur
	} else {
		return 0
	}
}

func JarakTPASubCriteria(str string) float64 {
	data, _ := ReadSubCriteriaFile()
	jarakTPA := data.JarakTPA

	if str == "pelayanan intensif" {
		return jarakTPA.PelayananIntensif
	} else if str == "pelayanan menengah" {
		return jarakTPA.PelayananMenengah
	} else if str == "pelayanan rendah" {
		return jarakTPA.PelayananRendah
	} else {
		return 0
	}
}

func JarakPemukimanSubCriteria(num float64) float64 {
	data, _ := ReadSubCriteriaFile()
	jarakPemukiman := data.JarakPemukiman

	if num >= 0 && num <= 100 {
		return jarakPemukiman.Jarak1
	} else if num >= 101 && num <= 200 {
		return jarakPemukiman.Jarak2
	} else if num >= 201 && num <= 300 {
		return jarakPemukiman.Jarak3
	} else if num >= 301 && num <= 400 {
		return jarakPemukiman.Jarak4
	} else if num >= 401 && num <= 500 {
		return jarakPemukiman.Jarak5
	} else {
		return 0
	}
}

func JarakSungaiSubCriteria(str string) float64 {
	data, _ := ReadSubCriteriaFile()
	jarakSungai := data.JarakSungai

	if str == "sangat layak" {
		return jarakSungai.SangatLayak
	} else if str == "layak" {
		return jarakSungai.Layak
	} else if str == "cukup layak" {
		return jarakSungai.CukupLayak
	} else if str == "kurang layak" {
		return jarakSungai.KurangLayak
	} else if str == "tidak layak" {
		return jarakSungai.TidakLayak
	} else {
		return 0
	}
}

func PartisipasiMasyarakatSubCriteria(num float64) float64 {
	data, _ := ReadSubCriteriaFile()
	PM := data.PartisipasiMasyarakat

	if num <= 20 {
		return PM.Partisipasi1
	} else if num >= 21 && num <= 40 {
		return PM.Partisipasi2
	} else if num >= 41 && num <= 60 {
		return PM.Partisipasi3
	} else if num >= 61 && num <= 80 {
		return PM.Partisipasi4
	} else if num >= 81 {
		return PM.Partisipasi5
	} else {
		return 0
	}
}

func CakupanRumahSubCriteria(num float64) float64 {
	data, _ := ReadSubCriteriaFile()
	CR := data.CakupanRumah

	if num <= 40 {
		return CR.Cakupan1
	} else if num >= 41 && num <= 80 {
		return CR.Cakupan2
	} else if num >= 81 && num <= 120 {
		return CR.Cakupan3
	} else if num >= 121 && num <= 160 {
		return CR.Cakupan4
	} else if num >= 161 {
		return CR.Cakupan5
	} else {
		return 0
	}
}

func AksesibilitasSubCriteria(str string) float64 {
	data, _ := ReadSubCriteriaFile()
	aksesibilitas := data.Aksesibilitas

	if str == "sangat layak" {
		return aksesibilitas.SangatLayak
	} else if str == "layak" {
		return aksesibilitas.Layak
	} else if str == "tidak layak" {
		return aksesibilitas.TidakLayak
	} else {
		return 0
	}
}

func CreateSubCriteriaFile() {
	data := model.SubCriteria{}

	file, _ := json.MarshalIndent(data, "", "")

	_ = os.WriteFile(constants.FileSubCriteria, file, 0644)
}

func ReadSubCriteriaFile() (*model.SubCriteria, error) {
	jsonFile, err := os.ReadFile(constants.FileSubCriteriaPKG)

	if err != nil {
		return nil, err
	}

	var data model.SubCriteria

	err = json.Unmarshal(jsonFile, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
