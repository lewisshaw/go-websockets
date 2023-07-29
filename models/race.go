package models

import "encoding/json"

type Race interface {
	GetId() int
	GetName() string
	GetCompetitors() []Competitor
	GetElapsedTime() int
	IncrementElapsedTime()
	AddCompetitors(competitors ...Competitor)
	GetCompetitorCount() int
	GetStartSequence() StartSequence
}

type race struct {
	id            int
	name          string
	competitors   []Competitor
	elapsedTime   int
	startSequence StartSequence
}

func NewRace(id int, name string, startSequence StartSequence) *race {
	return &race{
		id:            id,
		name:          name,
		competitors:   []Competitor{},
		elapsedTime:   0,
		startSequence: startSequence,
	}
}

func (r *race) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(jsonMap{
		"id":          r.id,
		"name":        r.name,
		"competitors": r.competitors,
		"elapsedTime": r.elapsedTime,
	})
	if err != nil {
		return []byte(""), err
	}

	return json, nil
}

func (r *race) AddCompetitors(competitors ...Competitor) {
	r.competitors = append(r.competitors, competitors...)
}

func (r *race) GetId() int {
	return r.id
}

func (r *race) GetName() string {
	return r.name
}

func (r *race) GetCompetitors() []Competitor {
	return r.competitors
}

func (r *race) GetElapsedTime() int {
	return r.elapsedTime
}

func (r *race) IncrementElapsedTime() {
	r.elapsedTime++
}

func (r *race) GetCompetitorCount() int {
	return len(r.competitors)
}

func (r *race) GetStartSequence() StartSequence {
	return r.startSequence
}
