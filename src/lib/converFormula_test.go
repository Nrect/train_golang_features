package lib

import (
	"github.com/stretchr/testify/assert"
	mockRepo "testProject/src/repo/mock"
	"testing"
)

func TestGenerateParameters(t *testing.T) {
	repo := mockRepo.NewMockDbResponse()
	parameters, err := generateParameters(repo)

	assert.Nil(t, err, "Should be nil")
	assert.IsType(t, map[string]interface{}{}, parameters, "should be map[string]interface{}{} type")
}
