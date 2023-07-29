package models

type RaceCollection interface {
	AddRace(race Race)
	GetRaceById(id int) Race
	GetLength() int
	GetRaces() []Race
}

type raceCollection []Race

func NewRaceCollection() *raceCollection {
	return &raceCollection{}
}
func (rc *raceCollection) AddRace(race Race) {
	*rc = append(*rc, race)
}

func (rc *raceCollection) GetRaceById(id int) Race {
	var race Race
	for _, currentRace := range *rc {
		if currentRace.GetId() == id {
			race = currentRace
			break
		}
	}
	return race
}

func (rc *raceCollection) GetLength() int {
	return len(*rc)
}

func (rc *raceCollection) GetRaces() []Race {
	return *rc
}
