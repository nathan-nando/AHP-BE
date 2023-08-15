package model

type Matrix [][]float64

type Pairwise struct {
	PairwiseFromJson        Matrix    `json:"pairwise"`
	PairwiseAfterCalculated Matrix    `json:"pairwise_after_calculated"`
	Weights                 []float64 `json:"weights"`
}

type SubCriteria struct {
	TimbulanSampah        TimbulanSampah        `json:"timbulan_sampah"`
	JarakTPA              JarakTPA              `json:"jarak_tpa"`
	JarakPemukiman        JarakPemukiman        `json:"jarak_pemukiman"`
	JarakSungai           JarakSungai           `json:"jarak_sungai"`
	PartisipasiMasyarakat PartisipasiMasyarakat `json:"partisipasi_masyarakat"`
	CakupanRumah          CakupanRumah          `json:"cakupan_rumah"`
	Aksesibilitas         Aksesibilitas         `json:"aksesibilitas"`
}

type TimbulanSampah struct {
	Pairwise                    Pairwise `json:"pairw ise"`
	PusatKota                   float64  `json:"pusat_kota"`
	DaerahKomersil              float64  `json:"daerah_komersil"`
	DaerahPerumahanTeratur      float64  `json:"daerah_perumahan_teratur"`
	Industri                    float64  `json:"industri"`
	Jalan                       float64  `json:"jalan"`
	DaerahPerumahanTidakTeratur float64  `json:"daerah_perumahan_tidak_teratur"`
}

type JarakTPA struct {
	Pairwise          Pairwise `json:"pairwise"`
	PelayananIntensif float64  `json:"pelayanan_intensif"`
	PelayananMenengah float64  `json:"pelayanan_menengah"`
	PelayananRendah   float64  `json:"pelayanan_rendah"`
}

type JarakPemukiman struct {
	Pairwise Pairwise `json:"pairwise"`
	Jarak1   float64  `json:"jarak_1"`
	Jarak2   float64  `json:"jarak_2"`
	Jarak3   float64  `json:"jarak_3"`
	Jarak4   float64  `json:"jarak_4"`
	Jarak5   float64  `json:"jarak_5"`
}

type JarakSungai struct {
	Pairwise    Pairwise `json:"pairwise"`
	SangatLayak float64  `json:"sangat_layak"`
	Layak       float64  `json:"layak"`
	CukupLayak  float64  `json:"cukup_layak"`
	KurangLayak float64  `json:"kurang_layak"`
	TidakLayak  float64  `json:"tidak_layak"`
}

type PartisipasiMasyarakat struct {
	Pairwise     Pairwise `json:"pairwise"`
	Partisipasi1 float64  `json:"partisipasi_1"`
	Partisipasi2 float64  `json:"partisipasi_2"`
	Partisipasi3 float64  `json:"partisipasi_3"`
	Partisipasi4 float64  `json:"partisipasi_4"`
	Partisipasi5 float64  `json:"partisipasi_5"`
}

type CakupanRumah struct {
	Pairwise Pairwise `json:"pairwise"`
	Cakupan1 float64  `json:"cakupan_1"`
	Cakupan2 float64  `json:"cakupan_2"`
	Cakupan3 float64  `json:"cakupan_3"`
	Cakupan4 float64  `json:"cakupan_4"`
	Cakupan5 float64  `json:"cakupan_5"`
}

type Aksesibilitas struct {
	Pairwise    Pairwise `json:"pairwise"`
	SangatLayak float64  `json:"sangat_layak"`
	Layak       float64  `json:"layak"`
	TidakLayak  float64  `json:"tidak_layak"`
}
