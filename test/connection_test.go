package test

import (
	"blog-mongo/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnection(t *testing.T) {
	conn,err := helper.Connection()
	assert.Nil(t, err)
	assert.NotNil(t, conn)
}
