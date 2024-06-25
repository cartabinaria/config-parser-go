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

type Year struct {
	Year int64  `json:"year"`
	Chat string `json:"chat"`
}

// This is temporary, i think we can join Teachins and Degrees
type DegreeTeaching struct {
	Name      string `json:"name"`
	Year      int64  `json:"year"`
	Mandatory bool   `json:"mandatory"`
}

type Degree struct {
	Id        string           `json:"id"`
	Name      string           `json:"name"`
	Icon      string           `json:"icon"`
	Teachings []DegreeTeaching `json:"teachings"`
	Years     []Year           `json:"years"`
	Chat      string           `json:"chat"`
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

func ParseTimetables() (timetables map[string]Timetable, err error) {
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

func ParseMaintainers() (maintainer []Maintainer, err error) {
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

func ParseRepresentatives() (map[string]Representative, error) {
	representatives := make(map[string]Representative)

	byteValue, err := cf.ReadFile(representativesFile)
	if errors.Is(err, os.ErrNotExist) {
		return representatives, fmt.Errorf("%s does not exist", representativesFile)
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

func GetAllMandatoryTeachingsFromDegree(d Degree) (dt []DegreeTeaching) {
	for _, i := range d.Teachings {
		if i.Mandatory {
			dt = append(dt, i)
		}
	}
	return
}

func GetAllElectivesTeachingsFromDegree(d Degree) (dt []DegreeTeaching) {
	for _, i := range d.Teachings {
		if !i.Mandatory {
			dt = append(dt, i)
		}
	}
	return
}

func GetYearMandatoryTeachingsFromDegree(d Degree, year int64) (dt []DegreeTeaching) {
	for _, i := range d.Teachings {
		if i.Mandatory && i.Year == year {
			dt = append(dt, i)
		}
	}
	return
}

func GetYearElectivesTeachingsFromDegree(d Degree, year int64) (dt []DegreeTeaching) {
	for _, i := range d.Teachings {
		if !i.Mandatory && i.Year == year {
			dt = append(dt, i)
		}
	}
	return
}
