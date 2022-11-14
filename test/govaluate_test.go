package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	_ "github.com/stretchr/testify/suite"
	lib "testProject/src/lib"
	"testing"
)

type testVariable struct {
	testValue  string
	wantResult float64
}

type GovaluateSuite struct {
	suite.Suite
	VariablesSuite []testVariable
}

func (suite *GovaluateSuite) SetupTest() {
	suite.VariablesSuite = []testVariable{
		{testValue: "5+5", wantResult: 10.0},
	}
}

func (suite *GovaluateSuite) Test() {

	for _, v := range suite.VariablesSuite {
		assert.Equal(suite.T(), v.wantResult, lib.ConvertFromFormula(v.testValue), "Should be equal")
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(GovaluateSuite))
}
