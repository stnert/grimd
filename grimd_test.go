package main

import (
	"testing"

	"github.com/miekg/dns"
)

func BenchmarkResolver(b *testing.B) {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(testDomain), dns.TypeA)

	c := new(dns.Client)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, err := c.Exchange(m, testNameserver)
		if err != nil {
			logger.Error(err)
		}
	}
}
