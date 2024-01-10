package salpha

import "time"

type attributes struct {
	PublishOn time.Time // if this  doesn't work use a strings
	Title     string
}
type article struct {
	Id         string
	Attributes attributes
}
type SeekingAlphaResponse struct {
	Data []article
}
