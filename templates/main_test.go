package templates_test

import (
	"testing"

	"github.com/solrac97gr/go-start/templates"
	"github.com/stretchr/testify/assert"
)

func Test_MainTemplate(t *testing.T) {
	expected := `package main

import (
	"log"
	sserver "github.com/user/test/pkg/server"
	"github.com/user/test/pkg/factory"
)

func main() {
	// Sub apps initialize
	user:=factory.BuildUserServer()
	payment:=factory.BuildPaymentServer()

	// Build Super server
	subApps := []sserver.SubApp{
		user,
		payment,
	}

	srv := sserver.NewServer(subApps)
	if err:=srv.Initialize("8080"); err != nil {
		log.Panic(err)
	}
}
`
	result := templates.NewMainTemplate("user", "test", []string{"user", "payment"})
	if result != expected {
		assert.Equal(t, expected, result)
	}
}
