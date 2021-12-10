package caddyheader

import (
	// "fmt"
	"fmt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp/reverseproxy"
)

func init() {
	caddy.RegisterModule(CaddyHeader{})
}

// caddyheader is an example; put your own type here.
type CaddyHeader struct {
	*reverseproxy.Handler
}

// CaddyModule returns the Caddy module information.
func (CaddyHeader) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "caddyheader",
		New: func() caddy.Module { return new(CaddyHeader) },
	}
}

func (c *CaddyHeader) PrintHead() {

	fmt.Println("This is header", c.Headers)
}
