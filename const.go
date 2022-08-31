package ehs

import "errors"

var (
	topRootDomains = []string{
		"com", "net", "org", "gov", "xyz", "app",
		"cn", "com.cn", "io", "me", "cc", "co",
	}
)

const (
	UA       = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:104.0) Gecko/20100101 Firefox/104.0"
	MAX_SIZE = 2 * 1024 * 1024
)

var (
	ERR_TOO_LARGE = errors.New("Body Too Large.")
)
