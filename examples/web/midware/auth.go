package midware

import (
	"github.com/ardanlabs/kit/log"
	"github.com/ardanlabs/kit/web"
)

// Auth simulates a midware for authentication.
func Auth(h web.Handler) web.Handler {

	// Pretend this get a DB connection
	f := func(c *web.Ctx) error {
		log.Dev(c.SessionID, "Auth", "******> Authorized")
		c.Values["UserID"] = "123"

		return h(c)
	}

	return f
}
