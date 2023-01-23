package models

type Layouts int

func TimeLayouts() map[string]string {
	tl := make(map[string]string)

	tl["ANSIC"] = "Mon Jan _2 15:04:05 2006"
	tl["UnixDate"] = "Mon Jan _2 15:04:05 MST 2006"
	tl["RubyDate"] = "Mon Jan 02 15:04:05 -0700 2006"
	tl["RFC822"] = "02 Jan 06 15:04 MST"
	tl["RFC822Z"] = "02 Jan 06 15:04 -0700"
	tl["RFC850"] = "Monday, 02-Jan-06 15:04:05 MST"
	tl["RFC1123"] = "Mon, 02 Jan 2006 15:04:05 MST"
	tl["RFC1123Z"] = "Mon, 02 Jan 2006 15:04:05 -0700"
	tl["RFC3339"] = "2006-01-02T15:04:05Z07:00"
	tl["RFC3339Nano"] = "2006-01-02T15:04:05.999999999Z07:00"
	tl["Kitchen"] = "3:04PM"

	return tl
}
