package usecase

import (
	"ahp-be/internal/alternative"
	"ahp-be/internal/criteria"
	"ahp-be/internal/criteria/dto"
	"ahp-be/internal/model"
	ahpFunc "ahp-be/pkg/ahp"
	"ahp-be/pkg/constants"
	"ahp-be/pkg/response"
	"ahp-be/pkg/utils"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
)

type Service struct {
	Repository            criteria.Repository
	RepositoryAlternative alternative.Repository
	Db                    *gorm.DB
}

func New(repo criteria.Repository, repoAlternative alternative.Repository, db *gorm.DB) *Service {
	return &Service{
		Repository:            repo,
		RepositoryAlternative: repoAlternative,
		Db:                    db,
	}
}

func (s *Service) FindCriteria(ctx context.Context) (*model.Pairwise, error) {
	var result *model.Pairwise

	jsonFile, err := os.ReadFile(constants.FileCriteria)

	if err != nil {
		return nil, err
	}

	var criteriaData model.Pairwise
	err = json.Unmarshal(jsonFile, &criteriaData)
	if err != nil {
		return nil, err
	}

	var preCriteriaData model.Pairwise
	err = json.Unmarshal(jsonFile, &preCriteriaData)
	if err != nil {
		return nil, err
	}

	rowsPC := len(criteriaData.PairwiseFromJson)
	colsPC := len(criteriaData.PairwiseFromJson[0])

	//MENCARI SUM DARI MASING MASING COL
	colSum := make([]float64, len(criteriaData.PairwiseFromJson))
	for i := 0; i < rowsPC; i++ {
		for j := 0; j < colsPC; j++ {
			colSum[i] += criteriaData.PairwiseFromJson[j][i]
		}
	}

	//NORMALISASI
	for i := 0; i < rowsPC; i++ {
		for j := 0; j < colsPC; j++ {
			criteriaData.PairwiseFromJson[i][j] /= colSum[j]
		}
	}

	normalColSum := make([]float64, len(criteriaData.PairwiseFromJson))
	normalRowSum := make([]float64, len(criteriaData.PairwiseFromJson[0]))
	weights := make([]float64, len(criteriaData.PairwiseFromJson))

	for i := 0; i < rowsPC; i++ {
		sum := 0.0
		for j := 0; j < colsPC; j++ {
			sum += criteriaData.PairwiseFromJson[i][j]
			normalColSum[i] += criteriaData.PairwiseFromJson[j][i]
			normalRowSum[i] += criteriaData.PairwiseFromJson[i][j]
			weights[i] = sum / float64(len(weights))
		}
	}

	result = &model.Pairwise{
		PairwiseFromJson:        preCriteriaData.PairwiseFromJson,
		PairwiseAfterCalculated: criteriaData.PairwiseFromJson,
		Weights:                 weights,
	}

	return result, err
}

func (s *Service) UpdateCriteria(ctx context.Context, payload *dto.CriteriaUpdateRequest) (*model.Pairwise, error) {
	jsonFile, err := os.ReadFile(constants.FileCriteria)

	var criteriaData model.Pairwise

	err = json.Unmarshal(jsonFile, &criteriaData)

	if err != nil {
		return nil, err
	}
	criteriaData.PairwiseFromJson = payload.Pairwise

	byteCriteria, err := json.Marshal(criteriaData)

	if err != nil {
		return nil, err
	}

	err = os.WriteFile(constants.FileCriteria, byteCriteria, 0644)

	if err != nil {
		return nil, err
	}

	result := &model.Pairwise{
		PairwiseFromJson: criteriaData.PairwiseFromJson,
	}

	return result, nil
}

func (s *Service) CheckConsistencyRatio(ctx context.Context) (bool, error) {
	criteriaData, err := s.FindCriteria(ctx)

	if err != nil {
		return false, err
	}

	matrix := criteriaData.PairwiseFromJson

	rows := len(matrix)
	cols := len(matrix[0])

	consistencyMatrix := make([]float64, 0)
	for i := 0; i < rows; i++ {
		var temp float64
		for j := 0; j < cols; j++ {
			temp += (matrix[i][j] * criteriaData.Weights[j]) / criteriaData.Weights[i]
		}
		var sum float64
		sum += temp
		consistencyMatrix = append(consistencyMatrix, temp)
	}

	var sumConsistencyMatrix float64
	for i := 0; i < len(consistencyMatrix); i++ {
		sumConsistencyMatrix += consistencyMatrix[i]
	}

	consistencyIndex := (sumConsistencyMatrix/float64(len(matrix)) - float64(len(matrix))) / float64(len(matrix)-1)

	ratioIndex := ahpFunc.GetRatioIndex()[len(matrix)-1]

	consistencyRatio := consistencyIndex / ratioIndex

	if consistencyRatio <= 0.1 {
		return true, nil
	}

	return false, nil
}

func (s *Service) FindScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	datas, err := s.Repository.FindScoreByCollectionID(ctx, collectionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return datas, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	if len(datas) == 0 {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}
	return datas, nil
}

func (s *Service) FindFinalScoresByCollectionID(ctx context.Context, collectionID *string) ([]model.AlternativeModel, error) {
	datas, err := s.Repository.FindFinalScoreByCollectionID(ctx, collectionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return datas, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return datas, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}
	if len(datas) == 0 {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}
	return datas, nil
}

func (s *Service) CalculateAlternativeToPoint(ctx context.Context, collectionID *string) (model.Matrix, error) {
	alternatives, err := s.RepositoryAlternative.FindsByCollectionID(ctx, collectionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if len(alternatives) == 0 {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	matrix := make(model.Matrix, 0)

	for i := 0; i < len(alternatives); i++ {
		row := model.Matrix{
			{
				ahpFunc.TimbulanSampahSubCriteria(alternatives[i].TimbulanSampah),
				ahpFunc.JarakTPASubCriteria(alternatives[i].JarakTpa),
				ahpFunc.JarakPemukimanSubCriteria(alternatives[i].JarakPemukiman),
				ahpFunc.JarakSungaiSubCriteria(alternatives[i].JarakSungai),
				ahpFunc.PartisipasiMasyarakatSubCriteria(alternatives[i].PartisipasiMasyarakat),
				ahpFunc.CakupanRumahSubCriteria(alternatives[i].CakupanRumah),
				ahpFunc.AksesibilitasSubCriteria(alternatives[i].Aksesibilitas),
			},
		}

		matrix = append(matrix, row...)
	}

	return matrix, nil
}

func (s *Service) CalculateScore(ctx context.Context, collectionID *string) ([]model.ScoreModel, error) {
	alternatives, err := s.RepositoryAlternative.FindsByCollectionID(ctx, collectionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if len(alternatives) == 0 {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	checkScores, err := s.Repository.FindScoreByCollectionID(ctx, collectionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	if len(checkScores) > 0 {
		_, err := s.Repository.DeleteScoresByCollectionID(ctx, collectionID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
			}
			return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
	}

	matrix := make(model.Matrix, 0)
	matrix, err = s.CalculateAlternativeToPoint(ctx, collectionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
		}
		return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
	}

	criteriaData, _ := s.FindCriteria(ctx)
	weights := criteriaData.Weights

	rowsACS := len(matrix)
	colsACS := len(matrix[0])

	//PERKALIAN MATRIKS ALTERNATIF DENGAN MATRIKS BOBOT
	for i := 0; i < rowsACS; i++ {
		for j := 0; j < colsACS; j++ {
			matrix[i][j] *= weights[j]
		}
	}

	scores := make([]model.ScoreModel, 0)

	for i := 0; i < len(matrix); i++ {
		scores = append(scores, model.ScoreModel{
			BaseModel: model.BaseModel{ID: uuid.NewString()},
			Score: model.Score{
				TimbulanSampah:        utils.RoundFloat(matrix[i][0], 3),
				JarakTpa:              utils.RoundFloat(matrix[i][1], 3),
				JarakPemukiman:        utils.RoundFloat(matrix[i][2], 3),
				JarakSungai:           utils.RoundFloat(matrix[i][3], 3),
				PartisipasiMasyarakat: utils.RoundFloat(matrix[i][4], 3),
				CakupanRumah:          utils.RoundFloat(matrix[i][5], 3),
				Aksesibilitas:         utils.RoundFloat(matrix[i][6], 3),
			},
			AlternativeID: alternatives[i].ID,
			CollectionID:  alternatives[i].CollectionID,
		})
	}

	_, err = s.Repository.CreateScore(ctx, scores)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	var collection *model.CollectionModel

	collection = &model.CollectionModel{
		Collection: model.Collection{
			ScoreIsCalculated: true,
		},
	}

	_, err = s.Repository.UpdateCollection(ctx, collectionID, collection)

	if err != nil {
		return nil, response.ErrorBuilder(&response.ErrorConstant.UnprocessableEntity, err)
	}

	return scores, nil
}

func (s *Service) CalculateFinalScore(ctx context.Context, collectionID *string) ([]model.FinalScoreModel, error) {
	alternativeScores, err := s.FindScoresByCollectionID(ctx, collectionID)
	var collection *model.CollectionModel

	if err != nil {
		return nil, err
	}

	if len(alternativeScores) == 0 {
		return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
	}

	checkFinalScores, err := s.Repository.FindFinalScoreByCollectionID(ctx, collectionID)

	if err != nil {
		return nil, err
	}

	if len(checkFinalScores) > 0 {
		_, err := s.Repository.DeleteFinalScoresByCollectionID(ctx, collectionID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, response.ErrorBuilder(&response.ErrorConstant.NotFound, err)
			}
			return nil, response.ErrorBuilder(&response.ErrorConstant.InternalServerError, err)
		}
	}

	finalScores := make([]model.FinalScoreModel, 0)

	for i := 0; i < len(alternativeScores); i++ {
		finalScores = append(finalScores, model.FinalScoreModel{
			BaseModel: model.BaseModel{ID: uuid.NewString()},
			FinalScore: model.FinalScore{
				FinalScore: (alternativeScores[i].Score.TimbulanSampah + alternativeScores[i].Score.JarakTpa + alternativeScores[i].Score.JarakPemukiman + alternativeScores[i].Score.JarakSungai + alternativeScores[i].Score.PartisipasiMasyarakat + alternativeScores[i].Score.CakupanRumah + alternativeScores[i].Score.Aksesibilitas) * 100,
				Rank:       0,
			},
			AlternativeID: alternativeScores[i].ID,
			CollectionID:  alternativeScores[i].CollectionID,
		})
	}

	_, err = s.Repository.CreateFinalScore(ctx, finalScores)

	collection = &model.CollectionModel{
		Collection: model.Collection{
			FinalScoreIsCalculated: true,
		},
	}

	_, err = s.Repository.UpdateCollection(ctx, collectionID, collection)

	if err != nil {
		return nil, err
	}

	return finalScores, nil
}
