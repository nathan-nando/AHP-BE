package ahp

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestTimbulanSampahSubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "jaringan jalan",
		request:  "jaringan jalan",
		expected: jaringanJalan,
	}, {
		name:     "perumahan",
		request:  "perumahan",
		expected: perumahan,
	}, {
		name:     "fasilitas komersial",
		request:  "fasilitas komersial",
		expected: fasilitasKomersial,
	}, {
		name:     "fasilitas umum",
		request:  "fasilitas umum",
		expected: fasilitasUmum,
	}, {
		name:     "fasilitas sosial",
		request:  "fasilitas sosial",
		expected: fasilitasSosial,
	}, {
		name:     "ruang terbuka",
		request:  "ruang terbuka",
		expected: ruangTerbuka,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TimbulanSampahSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJarakTPASubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "alternatif berada di jangkauan layanan tpa",
		request:  "alternatif berada di jangkauan layanan tpa",
		expected: alternatifJangkauanTPA,
	}, {
		name:     "alternatif berada di batas terjauh jangkauan layanan tpa",
		request:  "alternatif berada di batas terjauh jangkauan layanan tpa",
		expected: alternatifBatasTerjauhTPA,
	}, {
		name:     "alternatif tidak berada di jangkauan tpa",
		request:  "alternatif tidak berada di jangkauan tpa",
		expected: alternatifBukanJangkauanTPA,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JarakTPASubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJarakPemukimanSubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  float64
		expected float64
	}{{
		name:     "jarak pemukiman >= 0 && <= 100",
		request:  1,
		expected: jarak1,
	}, {
		name:     "jarak pemukiman >= 0 && <= 100",
		request:  100,
		expected: jarak1,
	}, {
		name:     "jarak pemukiman >= 101 && <= 200",
		request:  101,
		expected: jarak2,
	}, {
		name:     "jarak pemukiman >= 101 && <= 200",
		request:  200,
		expected: jarak2,
	}, {
		name:     "jarak pemukiman >= 201 && <= 300",
		request:  201,
		expected: jarak3,
	}, {
		name:     "jarak pemukiman >= 201 && <= 300",
		request:  300,
		expected: jarak3,
	}, {
		name:     "jarak pemukiman >= 301 && <= 401",
		request:  301,
		expected: jarak4,
	}, {
		name:     "jarak pemukiman >= 301 && <= 400",
		request:  400,
		expected: jarak4,
	}, {
		name:     "jarak pemukiman >= 401 && <= 500",
		request:  401,
		expected: jarak5,
	}, {
		name:     "jarak pemukiman >= 401 && <= 500",
		request:  500,
		expected: jarak5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JarakPemukimanSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJarakSungaiSubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "lokasi memenuhi peli banjir",
		request:  "lokasi memenuhi peli banjir",
		expected: memenuhiPeliBanjir,
	}, {
		name:     "lokasi memenuhi sebagian peli banjir",
		request:  "lokasi memenuhi sebagian peli banjir",
		expected: memenuhiSebagianPeliBanjir,
	}, {
		name:     "lokasi tidak memenuhi peli banjir",
		request:  "lokasi tidak memenuhi peli banjir",
		expected: tidakMemenuhiPeliBanjir,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JarakSungaiSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestPartisipasiMasyarakatSubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  float64
		expected float64
	}{{
		name:     "partisipasi masyarakat <= 20",
		request:  1,
		expected: partisipasi1,
	}, {
		name:     "partisipasi masyarakat <= 20",
		request:  20,
		expected: partisipasi1,
	}, {
		name:     "partisipasi masyarakat >= 21 && <= 40",
		request:  21,
		expected: partisipasi2,
	}, {
		name:     "partisipasi masyarakat >= 21 && <= 40",
		request:  40,
		expected: partisipasi2,
	}, {
		name:     "partisipasi masyarakat >= 41 && <= 60",
		request:  41,
		expected: partisipasi3,
	}, {
		name:     "partisipasi masyarakat >= 41 && <= 60",
		request:  60,
		expected: partisipasi3,
	}, {
		name:     "partisipasi masyarakat >= 61 && <= 80",
		request:  61,
		expected: partisipasi4,
	}, {
		name:     "partisipasi masyarakat >= 61 && <= 80",
		request:  80,
		expected: partisipasi4,
	}, {
		name:     "partisipasi masyarakat >= 81",
		request:  81,
		expected: partisipasi5,
	}, {
		name:     "partisipasi masyarakat >= 81",
		request:  100,
		expected: partisipasi5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PartisipasiMasyarakatSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCakupanRumahSubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  float64
		expected float64
	}{{
		name:     "cakupan rumah <= 40",
		request:  1,
		expected: cakupan1,
	}, {
		name:     "cakupan rumah <= 40",
		request:  40,
		expected: cakupan1,
	}, {
		name:     "cakupan rumah >= 41 && <= 80",
		request:  41,
		expected: cakupan2,
	}, {
		name:     "cakupan rumah >= 41 && <= 80",
		request:  80,
		expected: cakupan2,
	}, {
		name:     "cakupan rumah >= 81 && <= 120",
		request:  81,
		expected: cakupan3,
	}, {
		name:     "cakupan rumah >= 81 && <= 120",
		request:  120,
		expected: cakupan3,
	}, {
		name:     "cakupan rumah >= 121 && <= 160",
		request:  121,
		expected: cakupan4,
	}, {
		name:     "cakupan rumah >= 121 && <= 160",
		request:  160,
		expected: cakupan4,
	}, {
		name:     "cakupan rumah >= 161",
		request:  161,
		expected: cakupan5,
	}, {
		name:     "cakupan rumah >= 161",
		request:  200,
		expected: cakupan5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CakupanRumahSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestAksesibilitasSubCriteria(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "kondisi jalan bagus dan bisa dilewati kendaraan pengangkut sampah",
		request:  "kondisi jalan bagus dan bisa dilewati kendaraan pengangkut sampah",
		expected: jalanBagus,
	}, {
		name:     "kondisi jalan bagus, tetapi tidak bisa dilewati kendaraan pengangkut sampah atau jalan tidak bagus, tetapi bisa dilewati kendaraan pengangkut sampah",
		request:  "kondisi jalan bagus, tetapi tidak bisa dilewati kendaraan pengangkut sampah atau jalan tidak bagus, tetapi bisa dilewati kendaraan pengangkut sampah",
		expected: jalanBagusTapi,
	}, {
		name:     "kondisi jalan tidak bagus dan tidak bisa dilewati kendaraan pengangkut sampah",
		request:  "kondisi jalan tidak bagus dan tidak bisa dilewati kendaraan pengangkut sampah",
		expected: jalanTidakBagus,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AksesibilitasSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
