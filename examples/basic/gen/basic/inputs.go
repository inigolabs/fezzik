package basic

import "time"

type OneEnum int

const (
	OneEnum_ValA OneEnum = iota
	OneEnum_ValB
	OneEnum_ValC
)

type OneInput struct {
	OneInt             *int       `json:"oneInt"`
	OneIntMust         int        `json:"oneIntMust"`
	OneIntList         []*int     `json:"oneIntList"`
	OneIntMustList     []int      `json:"oneIntMustList"`
	OneIntMustListMust []int      `json:"oneIntMustListMust"`
	OneStr             *string    `json:"oneStr"`
	OneBool            *bool      `json:"oneBool"`
	OneEnum            *OneEnum   `json:"oneEnum"`
	OneTime            *time.Time `json:"oneTime"`
}
