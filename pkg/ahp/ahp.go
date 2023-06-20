package ahp

type SubCriteria map[string]float64

func GetRatioIndex() [15]float64 {
	return [15]float64{0, 0, 0.58, 0.9, 1.12, 1.24, 1.32, 1.41, 1.46, 1.49, 1.51, 1.48, 1.56, 1.57, 1.59}
}

func TimbulanSampahSubCriteria(str string) float64 {
	if str == "jaringan jalan" {
		return jaringanJalan
	} else if str == "perumahan" {
		return perumahan
	} else if str == "fasilitas komersial" {
		return fasilitasKomersial
	} else if str == "fasilitas umum" {
		return fasilitasUmum
	} else if str == "fasilitas sosial" {
		return fasilitasSosial
	} else if str == "ruang terbuka" {
		return ruangTerbuka
	} else {
		return 0
	}
}

func JarakTPASubCriteria(str string) float64 {
	if str == "alternatif berada di jangkauan layanan tpa" {
		return alternatifJangkauanTPA
	} else if str == "alternatif berada di batas terjauh jangkauan layanan tpa" {
		return alternatifBatasTerjauhTPA
	} else if str == "alternatif tidak berada di jangkauan tpa" {
		return alternatifBukanJangkauanTPA
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
	if str == "lokasi memenuhi peli banjir" {
		return memenuhiPeliBanjir
	} else if str == "lokasi memenuhi sebagian peli banjir" {
		return memenuhiSebagianPeliBanjir
	} else if str == "lokasi tidak memenuhi peli banjir" {
		return tidakMemenuhiPeliBanjir
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
	if str == "kondisi jalan bagus dan bisa dilewati kendaraan pengangkut sampah" {
		return jalanBagus
	} else if str == "kondisi jalan bagus, tetapi tidak bisa dilewati kendaraan pengangkut sampah atau jalan tidak bagus, tetapi bisa dilewati kendaraan pengangkut sampah" {
		return jalanBagusTapi
	} else if str == "kondisi jalan tidak bagus dan tidak bisa dilewati kendaraan pengangkut sampah" {
		return jalanTidakBagus
	} else {
		return 0
	}
}
