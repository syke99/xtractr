package models

type Layouts int

func TimeLayouts() map[string]string {
	tl := make(map[string]string)

	tl[ANSIC] = "Mon Jan _2 15:04:05 2006"
	tl[UnixDate] = "Mon Jan _2 15:04:05 MST 2006"
	tl[RubyDate] = "Mon Jan 02 15:04:05 -0700 2006"
	tl[RFC822] = "02 Jan 06 15:04 MST"
	tl[RFC822Z] = "02 Jan 06 15:04 -0700"
	tl[RFC850] = "Monday, 02-Jan-06 15:04:05 MST"
	tl[RFC1123] = "Mon, 02 Jan 2006 15:04:05 MST"
	tl[RFC1123Z] = "Mon, 02 Jan 2006 15:04:05 -0700"
	tl[RFC3339] = "2006-01-02T15:04:05Z07:00"
	tl[RFC339NANO] = "2006-01-02T15:04:05.999999999Z07:00"
	tl[Kitchen] = "3:04PM"

	return tl
}

const (
	ANSIC      = "ANSIC"
	UnixDate   = "UnixDate"
	RubyDate   = "RubyDate"
	RFC822     = "RFC822"
	RFC822Z    = "RFC822Z"
	RFC850     = "RFC850"
	RFC1123    = "RFC1123"
	RFC1123Z   = "RFC1123Z"
	RFC3339    = "RFC3339"
	RFC339NANO = "RFC3339Nano"
	Kitchen    = "Kitchen"
)
