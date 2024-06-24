package cparser

import (
	"encoding/json"
	"errors"
	"fmt"
	cf "github.com/csunibo/config"
	"os"
)

const (
	groupsFile          = "groups.json"
	degreesFile         = "degrees.json"
	teachingsFile       = "teachings.json"
	timetablesFile      = "timetables.json"
	maintainersFile     = "maintainers.json"
	representativesFile = "representatives.json"
)

type Maintainer struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

// config/teachings.json

type Teaching struct {
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	Chat       string   `json:"chat"`
	Website    string   `json:"website"`
	Professors []string `json:"professors"`
}

// config/degrees.json

type YearStudyDiagram struct {
	Mandatory []string `json:"mandatory"`
	Electives []string `json:"electives"`
}

type Year struct {
	Year      int64            `json:"year"`
	Chat      string           `json:"chat"`
	Teachings YearStudyDiagram `json:"teachings"`
}

type Degree struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
	Years []Year `json:"years"`
	Chat  string `json:"chat"`
}

// timetables.json

type Curriculum struct {
	Name     string `json:"name"`
	Callback string `json:"callback"`
}

// Recognized by a callback string
type Timetable struct {
	Course       string `json:"course"`    // Course title
	Name         string `json:"name"`      // Course name
	Type         string `json:"type"`      // Type (laurea|magistrale|2cycle)
	Curriculum   string `json:"curricula"` // Curriculum
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type RepresentativesData struct {
	Description  string `json:"description"`
	Title        string `json:"title"`
	FallbackText string `json:"fallbackText"`
}

type Representative struct {
	Course          string   `json:"course"`
	Representatives []string `json:"representatives"`
}

func ParseTeachings() (teachings []Teaching, err error) {
	file, err := cf.Open(teachingsFile)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", teachingsFile, err)
	}

	err = json.NewDecoder(file).Decode(&teachings)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", teachingsFile, err)
	}
	return
}

func ParseDegrees() (degrees []Degree, err error) {
	file, err := cf.Open(degreesFile)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", degreesFile, err)
	}
	err = json.NewDecoder(file).Decode(&degrees)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", degreesFile, err)
	}
	return
}

func ParseTimetables(configPath string) (timetables map[string]Timetable, err error) {
	file, err := cf.Open(timetablesFile)
	defer file.Close()
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", timetablesFile, err)
	}

	var mapData map[string]Timetable

	err = json.NewDecoder(file).Decode(&mapData)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", timetablesFile, err)
	}

	timetables = mapData
	return
}

func ParseMaintainers(configPath string) (maintainer []Maintainer, err error) {
	file, err := cf.ReadFile(maintainersFile)
	if errors.Is(err, os.ErrNotExist) {
		return maintainer, fmt.Errorf("%s does not exist", maintainersFile)
	} else if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", maintainersFile, err)
	}

	var projects []struct {
		Name        string       `json:"project"`
		Maintainers []Maintainer `json:"maintainers"`
	}

	err = json.Unmarshal(file, &projects)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", maintainersFile, err)
	}

	for _, p := range projects {
		if p.Name == "informabot" {
			return p.Maintainers, nil
		}
	}

	return nil, fmt.Errorf("couldn't found informabot projects after parsing %s", maintainersFile)
}

func ParseRepresentatives(configPath string) (map[string]Representative, error) {
	representatives := make(map[string]Representative)

	byteValue, err := cf.ReadFile(representativesFile)
	if errors.Is(err, os.ErrNotExist) {
		return representatives, fmt.Errorf("%s does not exist", maintainersFile)
	} else if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", representativesFile, err)
	}

	err = json.Unmarshal(byteValue, &representatives)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s file: %w", representativesFile, err)
	}

	if representatives == nil {
		representatives = make(map[string]Representative)
	}

	return representatives, nil
}
