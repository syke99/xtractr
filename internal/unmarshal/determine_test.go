package unmarshal

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetermine_QuerySQL_False(t *testing.T) {
	tags := "query"

	b, err := DetermineSQL(tags)

	assert.NoError(t, err)
	assert.Equal(t, false, b)
}

func TestDetermine_PathSQL_False(t *testing.T) {
	tags := "path"

	b, err := DetermineSQL(tags)

	assert.NoError(t, err)
	assert.Equal(t, false, b)
}

func TestDetermine_QuerySQL_True(t *testing.T) {
	tags := "query,sql"

	b, err := DetermineSQL(tags)

	assert.NoError(t, err)
	assert.Equal(t, true, b)
}

func TestDetermine_PathSQL_True(t *testing.T) {
	tags := "path,sql"

	b, err := DetermineSQL(tags)

	assert.NoError(t, err)
	assert.Equal(t, true, b)
}

func TestDetermine_TooMany(t *testing.T) {
	tags := "query,path,sql"

	_, err := DetermineSQL(tags)

	assert.Equal(t, errors.New("incorrect xtractr value(s) provided"), err)
}

func TestDetermineSQL_InvalidValue(t *testing.T) {
	tags := "test"

	_, err := DetermineSQL(tags)

	assert.Equal(t, errors.New("incorrect xtractr value(s) provided"), err)
}

func TestDetermineSQL_NoLocation(t *testing.T) {
	tags := "sql"

	_, err := DetermineSQL(tags)

	assert.Equal(t, errors.New("sql type specified without location of parameter"), err)
}
