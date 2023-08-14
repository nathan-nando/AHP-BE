package usecase

import (
	"ahp-be/internal/model"
	"ahp-be/internal/subCriteria"
	"ahp-be/internal/subCriteria/dto"
	"ahp-be/pkg/ahp"
	"ahp-be/pkg/constants"
	"context"
	"encoding/json"
	"gorm.io/gorm"
	"os"
)

type Service struct {
	Repository subCriteria.Repository
	Db         *gorm.DB
}

func New(repo subCriteria.Repository, db *gorm.DB) *Service {
	return &Service{Repository: repo, Db: db}
}

func (s *Service) FindSubCriteria(ctx context.Context) (*model.SubCriteria, error) {
	var result model.SubCriteria

	before, err := s.Repository.FindSubCriteria(ctx)
	if err != nil {
		return nil, err
	}

	after, err := s.Repository.FindSubCriteria(ctx)
	if err != nil {
		return nil, err
	}

	timbulanSampah := after.TimbulanSampah

	row := len(timbulanSampah.Pairwise.PairwiseFromJson)
	col := len(timbulanSampah.Pairwise.PairwiseFromJson[0])

	colSum := make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += timbulanSampah.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			timbulanSampah.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum := make([]float64, row)
	normalColSum := make([]float64, col)
	weights := make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += timbulanSampah.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += timbulanSampah.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += timbulanSampah.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}
	result.TimbulanSampah = model.TimbulanSampah{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.TimbulanSampah.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: timbulanSampah.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		PusatKota:                   weights[0],
		DaerahKomersil:              weights[1],
		DaerahPerumahanTeratur:      weights[2],
		Industri:                    weights[3],
		Jalan:                       weights[4],
		DaerahPerumahanTidakTeratur: weights[5],
	}

	jarakTPA := after.JarakTPA

	row = len(jarakTPA.Pairwise.PairwiseFromJson)
	col = len(jarakTPA.Pairwise.PairwiseFromJson[0])

	colSum = make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += jarakTPA.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			jarakTPA.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum = make([]float64, row)
	normalColSum = make([]float64, col)
	weights = make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += jarakTPA.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += jarakTPA.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += jarakTPA.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}

	result.JarakTPA = model.JarakTPA{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.JarakTPA.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: jarakTPA.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		PelayananIntensif: weights[0],
		PelayananMenengah: weights[1],
		PelayananRendah:   weights[2],
	}

	jarakPemukiman := after.JarakPemukiman

	row = len(jarakPemukiman.Pairwise.PairwiseFromJson)
	col = len(jarakPemukiman.Pairwise.PairwiseFromJson[0])

	colSum = make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += jarakPemukiman.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			jarakPemukiman.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum = make([]float64, row)
	normalColSum = make([]float64, col)
	weights = make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += jarakPemukiman.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += jarakPemukiman.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += jarakPemukiman.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}

	result.JarakPemukiman = model.JarakPemukiman{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.JarakPemukiman.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: jarakPemukiman.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		Jarak1: weights[0],
		Jarak2: weights[1],
		Jarak3: weights[2],
		Jarak4: weights[3],
		Jarak5: weights[4],
	}
	//
	jarakSungai := after.JarakSungai

	row = len(jarakSungai.Pairwise.PairwiseFromJson)
	col = len(jarakSungai.Pairwise.PairwiseFromJson[0])

	colSum = make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += jarakSungai.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			jarakSungai.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum = make([]float64, row)
	normalColSum = make([]float64, col)
	weights = make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += jarakSungai.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += jarakSungai.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += jarakSungai.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}

	result.JarakSungai = model.JarakSungai{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.JarakSungai.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: jarakSungai.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		SangatLayak: weights[0],
		Layak:       weights[1],
		CukupLayak:  weights[2],
		KurangLayak: weights[3],
		TidakLayak:  weights[4],
	}

	partisipasiMasyarakat := after.PartisipasiMasyarakat

	row = len(partisipasiMasyarakat.Pairwise.PairwiseFromJson)
	col = len(partisipasiMasyarakat.Pairwise.PairwiseFromJson[0])

	colSum = make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += partisipasiMasyarakat.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			partisipasiMasyarakat.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum = make([]float64, row)
	normalColSum = make([]float64, col)
	weights = make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += partisipasiMasyarakat.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += partisipasiMasyarakat.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += partisipasiMasyarakat.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}

	result.PartisipasiMasyarakat = model.PartisipasiMasyarakat{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.PartisipasiMasyarakat.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: partisipasiMasyarakat.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		Partisipasi1: weights[0],
		Partisipasi2: weights[1],
		Partisipasi3: weights[2],
		Partisipasi4: weights[3],
		Partisipasi5: weights[4],
	}

	cakupanRumah := after.CakupanRumah

	row = len(cakupanRumah.Pairwise.PairwiseFromJson)
	col = len(cakupanRumah.Pairwise.PairwiseFromJson[0])

	colSum = make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += cakupanRumah.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			cakupanRumah.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum = make([]float64, row)
	normalColSum = make([]float64, col)
	weights = make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += cakupanRumah.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += cakupanRumah.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += cakupanRumah.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}

	result.CakupanRumah = model.CakupanRumah{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.CakupanRumah.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: cakupanRumah.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		Cakupan1: weights[0],
		Cakupan2: weights[1],
		Cakupan3: weights[2],
		Cakupan4: weights[3],
		Cakupan5: weights[4],
	}

	aksesibilitas := after.Aksesibilitas

	row = len(aksesibilitas.Pairwise.PairwiseFromJson)
	col = len(aksesibilitas.Pairwise.PairwiseFromJson[0])

	colSum = make([]float64, row)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			colSum[i] += aksesibilitas.Pairwise.PairwiseFromJson[j][i]
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			aksesibilitas.Pairwise.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalRowSum = make([]float64, row)
	normalColSum = make([]float64, col)
	weights = make([]float64, row)

	for i := 0; i < row; i++ {
		sum := 0.0
		for j := 0; j < col; j++ {
			sum += aksesibilitas.Pairwise.PairwiseFromJson[i][j]
			normalRowSum[i] += aksesibilitas.Pairwise.PairwiseFromJson[i][j]
			normalColSum[i] += aksesibilitas.Pairwise.PairwiseFromJson[j][i]
			weights[i] = sum / float64(len(weights))
		}
	}

	result.Aksesibilitas = model.Aksesibilitas{
		Pairwise: model.Pairwise{
			PairwiseFromJson:        before.Aksesibilitas.Pairwise.PairwiseFromJson,
			PairwiseAfterCalculated: aksesibilitas.Pairwise.PairwiseFromJson,
			Weights:                 weights,
		},
		SangatLayak: weights[0],
		Layak:       weights[1],
		TidakLayak:  weights[2],
	}

	byteResult, err := json.Marshal(result)

	if err != nil {
		return nil, err
	}

	err = os.WriteFile(constants.FileSubCriteria, byteResult, 0644)

	if err != nil {
		return nil, err
	}

	return &result, err
}

func (s *Service) UpdateSubCriteria(ctx context.Context, payload *model.SubCriteria) (*model.SubCriteria, error) {
	jsonFile, err := os.ReadFile(constants.FileSubCriteria)

	if err != nil {
		return nil, err
	}

	var data model.SubCriteria

	err = json.Unmarshal(jsonFile, &data)

	if err != nil {
		return nil, err
	}

	data = *payload

	byteData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	err = os.WriteFile(constants.FileSubCriteria, byteData, 0644)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *Service) CheckConsistencyRatio(ctx context.Context, mode *dto.CheckCRSubCriteria) (bool, error) {
	data, err := s.FindSubCriteria(ctx)

	if err != nil {
		return false, err
	}

	var matrix model.Matrix
	m := mode.Mode

	CM := make([]float64, 0)

	switch m {
	case "timbulan_sampah":
		matrix = data.TimbulanSampah.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.TimbulanSampah.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	case "jarak_tpa":
		matrix = data.JarakTPA.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.JarakTPA.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	case "jarak_pemukiman":
		matrix = data.JarakPemukiman.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.JarakPemukiman.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	case "jarak_sungai":
		matrix = data.JarakSungai.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.JarakSungai.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	case "partisipasi_masyarakat":
		matrix = data.PartisipasiMasyarakat.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.PartisipasiMasyarakat.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	case "cakupan_rumah":
		matrix = data.CakupanRumah.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.CakupanRumah.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	case "aksesibilitas":
		matrix = data.Aksesibilitas.Pairwise.PairwiseFromJson

		row := len(matrix)
		col := len(matrix)

		for i := 0; i < row; i++ {
			var temp float64
			for j := 0; j < col; j++ {
				temp += (matrix[i][j] * data.Aksesibilitas.Pairwise.Weights[j]) / data.TimbulanSampah.Pairwise.Weights[i]
			}

			var sum float64
			sum += temp
			CM = append(CM, temp)
		}

	default:
		return false, err
	}

	var sumCM float64

	for i := 0; i < len(CM); i++ {
		sumCM += CM[i]
	}

	CI := (sumCM/float64(len(matrix)) - float64(len(matrix))) / float64(len(matrix)-1)

	RI := ahp.GetRatioIndex()[len(matrix)-1]

	CR := CI / RI

	if CR <= 0.1 {
		return true, err
	}

	return false, nil
}
