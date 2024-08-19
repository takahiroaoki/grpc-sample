package validator

import "regexp"

var mailRegexp = regexp.MustCompile(`^[a-zA-Z0-9\._%+-]+@[a-zA-Z0-9\.-]+\.[a-zA-Z]{2,}$`)

func MailRegexp() *regexp.Regexp {
	return mailRegexp
}
