package ahp

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadSubCriteriaFile(t *testing.T) {
	data, err := ReadSubCriteriaFile()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(data.TimbulanSampah.DaerahPerumahanTeratur)
}

func TestTimbulanSampahSubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	TS := data.TimbulanSampah

	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "daerah di jalan protokol pusat kota",
		request:  "daerah di jalan protokol pusat kota",
		expected: TS.PusatKota,
	}, {
		name:     "daerah komersil",
		request:  "daerah komersil",
		expected: TS.DaerahKomersil,
	}, {
		name:     "daerah perumahan teratur",
		request:  "daerah perumahan teratur",
		expected: TS.DaerahPerumahanTeratur,
	}, {
		name:     "daerah industri",
		request:  "daerah industri",
		expected: TS.Industri,
	}, {
		name:     "jalan taman dan hutan kota",
		request:  "jalan taman dan hutan kota",
		expected: TS.Jalan,
	}, {
		name:     "daerah perumahan tidak teratur",
		request:  "daerah perumahan tidak teratur",
		expected: TS.DaerahPerumahanTidakTeratur,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := TimbulanSampahSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJarakTPASubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	test := data.JarakTPA

	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "pelayanan intensif",
		request:  "pelayanan intensif",
		expected: test.PelayananIntensif,
	}, {
		name:     "pelayanan menengah",
		request:  "pelayanan menengah",
		expected: test.PelayananMenengah,
	}, {
		name:     "pelayanan rendah",
		request:  "pelayanan rendah",
		expected: test.PelayananRendah,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JarakTPASubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJarakPemukimanSubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	test := data.JarakPemukiman

	tests := []struct {
		name     string
		request  float64
		expected float64
	}{{
		name:     "jarak pemukiman >= 0 && <= 100",
		request:  1,
		expected: test.Jarak1,
	}, {
		name:     "jarak pemukiman >= 0 && <= 100",
		request:  100,
		expected: test.Jarak1,
	}, {
		name:     "jarak pemukiman >= 101 && <= 200",
		request:  101,
		expected: test.Jarak2,
	}, {
		name:     "jarak pemukiman >= 101 && <= 200",
		request:  200,
		expected: test.Jarak2,
	}, {
		name:     "jarak pemukiman >= 201 && <= 300",
		request:  201,
		expected: test.Jarak3,
	}, {
		name:     "jarak pemukiman >= 201 && <= 300",
		request:  300,
		expected: test.Jarak3,
	}, {
		name:     "jarak pemukiman >= 301 && <= 401",
		request:  301,
		expected: test.Jarak4,
	}, {
		name:     "jarak pemukiman >= 301 && <= 400",
		request:  400,
		expected: test.Jarak4,
	}, {
		name:     "jarak pemukiman >= 401 && <= 500",
		request:  401,
		expected: test.Jarak5,
	}, {
		name:     "jarak pemukiman >= 401 && <= 500",
		request:  500,
		expected: test.Jarak5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JarakPemukimanSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJarakSungaiSubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	test := data.JarakSungai

	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "sangat layak",
		request:  "sangat layak",
		expected: test.SangatLayak,
	}, {
		name:     "layak",
		request:  "layak",
		expected: test.Layak,
	}, {
		name:     "cukup layak",
		request:  "cukup layak",
		expected: test.CukupLayak,
	}, {
		name:     "kurang layak",
		request:  "kurang layak",
		expected: test.KurangLayak,
	}, {
		name:     "tidak layak",
		request:  "tidak layak",
		expected: test.TidakLayak,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := JarakSungaiSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestPartisipasiMasyarakatSubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	test := data.PartisipasiMasyarakat

	tests := []struct {
		name     string
		request  float64
		expected float64
	}{{
		name:     "partisipasi masyarakat <= 20",
		request:  1,
		expected: test.Partisipasi1,
	}, {
		name:     "partisipasi masyarakat <= 20",
		request:  20,
		expected: test.Partisipasi1,
	}, {
		name:     "partisipasi masyarakat >= 21 && <= 40",
		request:  21,
		expected: test.Partisipasi2,
	}, {
		name:     "partisipasi masyarakat >= 21 && <= 40",
		request:  40,
		expected: test.Partisipasi2,
	}, {
		name:     "partisipasi masyarakat >= 41 && <= 60",
		request:  41,
		expected: test.Partisipasi3,
	}, {
		name:     "partisipasi masyarakat >= 41 && <= 60",
		request:  60,
		expected: test.Partisipasi3,
	}, {
		name:     "partisipasi masyarakat >= 61 && <= 80",
		request:  61,
		expected: test.Partisipasi4,
	}, {
		name:     "partisipasi masyarakat >= 61 && <= 80",
		request:  80,
		expected: test.Partisipasi4,
	}, {
		name:     "partisipasi masyarakat >= 81",
		request:  81,
		expected: test.Partisipasi5,
	}, {
		name:     "partisipasi masyarakat >= 81",
		request:  100,
		expected: test.Partisipasi5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PartisipasiMasyarakatSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestCakupanRumahSubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	test := data.CakupanRumah

	tests := []struct {
		name     string
		request  float64
		expected float64
	}{{
		name:     "cakupan rumah <= 40",
		request:  1,
		expected: test.Cakupan1,
	}, {
		name:     "cakupan rumah <= 40",
		request:  40,
		expected: test.Cakupan1,
	}, {
		name:     "cakupan rumah >= 41 && <= 80",
		request:  41,
		expected: test.Cakupan2,
	}, {
		name:     "cakupan rumah >= 41 && <= 80",
		request:  80,
		expected: test.Cakupan2,
	}, {
		name:     "cakupan rumah >= 81 && <= 120",
		request:  81,
		expected: test.Cakupan3,
	}, {
		name:     "cakupan rumah >= 81 && <= 120",
		request:  120,
		expected: test.Cakupan3,
	}, {
		name:     "cakupan rumah >= 121 && <= 160",
		request:  121,
		expected: test.Cakupan4,
	}, {
		name:     "cakupan rumah >= 121 && <= 160",
		request:  160,
		expected: test.Cakupan4,
	}, {
		name:     "cakupan rumah >= 161",
		request:  161,
		expected: test.Cakupan5,
	}, {
		name:     "cakupan rumah >= 161",
		request:  200,
		expected: test.Cakupan5,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := CakupanRumahSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestAksesibilitasSubCriteria(t *testing.T) {
	data, _ := ReadSubCriteriaFile()
	test := data.Aksesibilitas

	tests := []struct {
		name     string
		request  string
		expected float64
	}{{
		name:     "sangat layak",
		request:  "sangat layak",
		expected: test.SangatLayak,
	}, {
		name:     "layak",
		request:  "layak",
		expected: test.Layak,
	}, {
		name:     "tidak layak",
		request:  "tidak layak",
		expected: test.TidakLayak,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := AksesibilitasSubCriteria(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
