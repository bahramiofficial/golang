package main

import (
	"gorm.io/gorm"
)

type PK uint32

type CompanyName struct {
	gorm.Model
	Name      string      `json:"name,omitempty"`
	CarModels []*CarModel `json:"car_models,omitempty"`
}
type CarModel struct {
	gorm.Model
	Name          string         `json:"name,omitempty"`
	CompanyNameID PK             `json:"company_name_id,omitempty"`
	CompanyName   []*CompanyName `json:"company_name,omitempty"`
	CreateYear    []*CreateYear  `json:"create_year,omitempty"`
}

type CreateYear struct {
	gorm.Model
	Name       string       `json:"name,omitempty"`
	CarModelID PK           `json:"car_model_id,omitempty"`
	CarModel   []*CarModel  `json:"car_model,omitempty"`
	Diversity  []*Diversity `json:"diversities,omitempty"`
	Picture    string       `json:"picture,omitempty"`
}

type Diversity struct {
	gorm.Model
	Name         string        `json:"name,omitempty"`
	CreateYearID PK            `json:"create_year_id,omitempty"`
	CreateYear   []*CreateYear `json:"create_year,omitempty"`
	Picture      string        `json:"picture,omitempty"`
	Price        float64       `json:"price,omitempty"`
}
