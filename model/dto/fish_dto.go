package dto

import "lilith/model/entity"

type FishAllDto struct {
	V2ex []entity.FishEntity `json:"v2ex"`
	Kr   []entity.FishEntity `json:"kr"`
	Bili []entity.FishEntity `json:"bili"`
}
