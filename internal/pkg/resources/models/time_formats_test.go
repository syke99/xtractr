package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeLayouts(t *testing.T) {
	tl := TimeLayouts()

	assert.Equal(t, "Mon Jan _2 15:04:05 2006", tl[ANSIC])
	assert.Equal(t, "Mon Jan _2 15:04:05 MST 2006", tl[UnixDate])
	assert.Equal(t, "Mon Jan 02 15:04:05 -0700 2006", tl[RubyDate])
	assert.Equal(t, "02 Jan 06 15:04 MST", tl[RFC822])
	assert.Equal(t, "02 Jan 06 15:04 -0700", tl[RFC822Z])
	assert.Equal(t, "Monday, 02-Jan-06 15:04:05 MST", tl[RFC850])
	assert.Equal(t, "Mon, 02 Jan 2006 15:04:05 MST", tl[RFC1123])
	assert.Equal(t, "Mon, 02 Jan 2006 15:04:05 -0700", tl[RFC1123Z])
	assert.Equal(t, "2006-01-02T15:04:05Z07:00", tl[RFC3339])
	assert.Equal(t, "2006-01-02T15:04:05.999999999Z07:00", tl[RFC339NANO])
	assert.Equal(t, "3:04PM", tl[Kitchen])
}
