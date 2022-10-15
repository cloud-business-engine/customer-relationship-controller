package internal_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cloud-business-engine/customer-relationship-controller/internal"
)

func TestLogger(t *testing.T) {
	service := "service"

	output := bytes.NewBuffer(nil)
	logger := internal.NewLogger(service, output)
	logger.Log().Send()

	logMessage := make(map[string]string)
	_ = json.Unmarshal(output.Bytes(), &logMessage)

	assert.Equal(t, service, logMessage["service"])
}
