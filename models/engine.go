package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylenders int64     `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_ange"`
}

type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylenders int64 `json:"noOfCylinders"`
	CarRange      int64 `json:"carRange"`
}

func ValidateEngineRequest(engineRequest EngineRequest) error {
	if err := validateDisplacement(engineRequest.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCylinders(engineRequest.NoOfCylenders); err != nil {
		return err
	}
	if err := validatecarRange(engineRequest.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(disp int64) error {
	if disp <= 0 {
		return errors.New("displacement is required")
	}
	return nil
}

func validateNoOfCylinders(cyl int64) error {
	if cyl <= 0 {
		return errors.New("cylinder is required")
	}
	return nil
}

func validatecarRange(rangeCar int64) error {
	if rangeCar <= 0 {
		return errors.New("car range is required")
	}
	return nil
}
