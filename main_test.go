package cparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const configPath = "./config"

func TestParseTeachings(t *testing.T) {
	_, err := ParseTeachings()
	assert.Nil(t, err)
}

func TestParseDegrees(t *testing.T) {
	_, err := ParseDegrees()
	assert.Nil(t, err)
}

func TestParseTimetables(t *testing.T) {
	_, err := ParseTimetables()
	assert.Nil(t, err)
}

func TestParseMaintainers(t *testing.T) {
	_, err := ParseMaintainers()
	assert.Nil(t, err)
}

func TestParseRepresentatives(t *testing.T) {
	_, err := ParseRepresentatives()
	assert.Nil(t, err)
}
