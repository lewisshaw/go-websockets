package models

import "encoding/json"

type Competitor interface {
	getId() int
	getName() string
	getFinishTime() int
}

type competitor struct {
	id         int
	name       string
	finishTime int
}

func NewCompetitor(id int, name string) *competitor {
	return &competitor{
		id:         id,
		name:       name,
		finishTime: 0,
	}
}

func (c *competitor) getId() int {
	return c.id
}

func (c *competitor) getName() string {
	return c.name
}

func (c *competitor) getFinishTime() int {
	return c.finishTime
}

func (c *competitor) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(jsonMap{
		"id":         c.id,
		"name":       c.name,
		"finishTime": c.finishTime,
	})
	if err != nil {
		return []byte(""), err
	}
	return json, nil
}
