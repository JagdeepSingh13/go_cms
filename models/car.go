package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

func ValidateRequest(carReq CarRequest) error {
	if err := validateName(carReq.Name); err != nil {
		return err
	}
	if err := validateYear(carReq.Year); err != nil {
		return err
	}
	if err := validatePrice(carReq.Price); err != nil {
		return err
	}
	if err := validateBrand(carReq.Brand); err != nil {
		return err
	}
	if err := validateEngine(carReq.Engine); err != nil {
		return err
	}
	if err := validateFuelType(carReq.FuelType); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("year must be a valid  number")
	}
	currYear := time.Now().Year()
	if yearInt < 1886 || yearInt > currYear {
		return errors.New("year must be between 1886 and current Year")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}
func validateFuelType(fuelType string) error {
	validFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybrid"}
	for _, validType := range validFuelTypes {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("enter valid fuel type")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("engineId is required")
	}
	// same for other engine components
	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("enter valid price")
	}
	return nil
}
