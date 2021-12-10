package caddyheader

import (
	"github.com/FiReBeNdEr07/caddy-header"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/dustin/go-humanize"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("caddyheader", parseCaddyfile)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	rb := new(caddyheader)

	for h.Next() {
		// configuration should be in a block
		for h.NextBlock(0) {
			switch h.Val() {
			case "max_size":
				var sizeStr string
				if !h.AllArgs(&sizeStr) {
					return nil, h.ArgErr()
				}
				size, err := humanize.ParseBytes(sizeStr)
				if err != nil {
					return nil, h.Errf("parsing max_size: %v", err)
				}
				rb.MaxSize = int64(size)
			default:
				return nil, h.Errf("unrecognized servers option '%s'", h.Val())
			}
		}
	}

	return rb, nil
}
