package hello

import (
	"rsc.io/quote"
	"rsc.io/sampler"

func Hello() string {
    return quote.Hello() + sampler.Glass()
}
