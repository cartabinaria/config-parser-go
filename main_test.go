package cparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const configPath = "./config"

func TestParseTeachings(t *testing.T) {
	_, err := ParseTeachings(configPath)
	assert.Nil(t, err)
}

func TestParseDegrees(t *testing.T) {
	_, err := ParseDegrees(configPath)
	assert.Nil(t, err)
}

func TestParseTimetables(t *testing.T) {
	_, err := ParseTimetables(configPath)
	assert.Nil(t, err)
}

func TestParseMaintainers(t *testing.T) {
	_, err := ParseMaintainers(configPath)
	assert.Nil(t, err)
}

func TestParseRepresentatives(t *testing.T) {
	_, err := ParseRepresentatives(configPath)
	assert.Nil(t, err)
}
