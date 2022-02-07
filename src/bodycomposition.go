package bodycompositionwebapi

import (
	"fmt"
	"github.com/davidkroell/bodycomposition"
)

type BodyCompositionRequest struct {
	TimeStamp         int64   `json:"timeStamp"`
	Weight            float64 `json:"weight"`
	PercentFat        float64 `json:"percentFat"`
	PercentHydration  float64 `json:"percentHydration"`
	BoneMass          float64 `json:"boneMass"`
	MuscleMass        float64 `json:"muscleMass"`
	VisceralFatRating float64 `json:"visceralFatRating"`
	PhysiqueRating    float64 `json:"physiqueRating"`
	MetabolicAge      float64 `json:"metabolicAge"`
	CaloriesActiveMet float64 `json:"caloriesActiveMet"`
	BodyMassIndex     float64 `json:"bodyMassIndex"`
	Email	          string  `json:"email"`
	Password          string  `json:"password"`
}

func UploadBodyComposition(request BodyCompositionRequest) string {
	bc := bodycomposition.NewBodyComposition(
		request.Weight,
		request.PercentFat,
		request.PercentHydration,
		request.BoneMass,
		request.MuscleMass,
		request.VisceralFatRating,
		request.PhysiqueRating,
		request.MetabolicAge,
		request.CaloriesActiveMet,
		request.BodyMassIndex,
		request.TimeStamp)

	var result string
		if err := bodycomposition.Upload(request.Email, request.Password, bc); err != nil {
			result = fmt.Sprintf("%s%s","Error uploading weight to Garmin Connect: ", err.Error())
		} else {
			result = "Ok"	
		}
	return result
}