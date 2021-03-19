package validate

import (
	"errors"
	"time"
)

type ValidationService struct {

}

func NewValidationService() *ValidationService {
	return &ValidationService{}
}

func (vs *ValidationService) Validate(data map[string]interface{}) (error, string) {
	date := data["date"]
	views := data["views"]
	clicks := data["clicks"]
	cost := data["cost"]

	if vs.typeOf(date) != "string" {
		return errors.New(""), "Field 'date' must be of type string and contain the date in the format YYYY-MM-DD"
	}
	_, err := time.Parse("2006-01-02", date.(string))
	if err != nil {
		return errors.New(""), "Field 'date' must contain the date in the format YYYY-MM-DD"
	}

	if views != nil {
		err, msg := vs.checkViews(views)
		if err != nil {
			return err, msg
		}
	}

	if clicks != nil {
		err, msg := vs.checkClicks(clicks)
		if err != nil {
			return err, msg
		}
	}

	if cost != nil {
		err, msg := vs.checkCost(cost)
		if err != nil {
			return err, msg
		}
	}

	return nil, "ok"
}

func (vs *ValidationService) checkViews(views interface{}) (error, string) {
	if vs.typeOf(views) != "float64" {
		return errors.New(""), "Field 'views' must be of type int"
	}
	if !vs.isIntegral(views.(float64)) {
		return errors.New(""), "Field 'views' must be of type int"
	}
	if views.(float64) < 0 {
		return errors.New(""), "Field 'views' cannot be less than zero"
	}
	return nil, ""
}

func (vs *ValidationService) checkClicks(clicks interface{}) (error, string) {
	if vs.typeOf(clicks) != "float64" {
		return errors.New(""), "Field 'clicks' must be of type int"
	}
	if clicks.(float64) < 0 {
		return errors.New(""), "Field 'clicks' cannot be less than zero"
	}
	if !vs.isIntegral(clicks.(float64)) {
		return errors.New(""), "Field 'clicks' must be of type int"
	}
	return nil, ""
}

func (vs *ValidationService) checkCost(cost interface{}) (error, string) {
	if vs.typeOf(cost) != "float64" {
		return errors.New(""), "Field 'cost' must be of type float64"
	}
	if cost.(float64) < 0 {
		return errors.New(""), "Field 'cost' cannot be less than zero"
	}
	return nil, ""
}

func (vs *ValidationService) isIntegral(val float64) bool {
	return val == float64(int(val))
}

func (vs *ValidationService) typeOf(value interface{}) string {
	switch value.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
