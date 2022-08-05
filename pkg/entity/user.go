package entity

import (
	"encoding/json"
	"strconv"
)

type CachedUser struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func(cu CachedUser) MarshalBinary() ([]byte, error) {
    return json.Marshal(strconv.FormatUint(uint64(cu.ID), 10) + ":" + cu.Username)
}

func(cu *CachedUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, cu)
}