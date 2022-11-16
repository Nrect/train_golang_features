package unit

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	_ "github.com/stretchr/testify/suite"
	"testProject/src/constants"
	lib "testProject/src/lib"
	mockRepo "testProject/src/repo/mock"
	"testing"
)

type variableTestCalculateFromFormula struct {
	testValue  string
	wantResult float64
}

type CalculateFromFormulaSuite struct {
	suite.Suite
	GoodVariable           []variableTestCalculateFromFormula
	BadVariableErrorNul    []variableTestCalculateFromFormula
	BadVariableErrorNotNul []variableTestCalculateFromFormula
	MockRequestVariable    []variableTestCalculateFromFormula
	MockDbResponse         *mockRepo.FormulaResponseDb
}

func (suite *CalculateFromFormulaSuite) SetupTest() {
	suite.GoodVariable = []variableTestCalculateFromFormula{
		{testValue: "5+5", wantResult: 10},
		{testValue: "5-5", wantResult: 0},
		{testValue: "5*5", wantResult: 25},
		{testValue: "5/5", wantResult: 1},
		{testValue: "5 + (5 * 2)", wantResult: 15},
		{testValue: "5 + 5 * 2", wantResult: 15},
		{testValue: "5.5 + 5 * 2", wantResult: 15.5},
	}

	suite.BadVariableErrorNul = []variableTestCalculateFromFormula{
		{testValue: "0 / 0"},
		{testValue: "5 / 0"},
	}

	suite.BadVariableErrorNotNul = []variableTestCalculateFromFormula{
		{testValue: "5 / "},
		{testValue: "5 * some text"},
	}

	suite.MockDbResponse = mockRepo.NewMockDbResponse()
	dbResponse := suite.MockDbResponse

	suite.MockRequestVariable = []variableTestCalculateFromFormula{
		{testValue: constants.DetailsTip.ToString(), wantResult: dbResponse.DetailsTip},
		{
			testValue: fmt.Sprintf(
				"%s + %s + %s",
				constants.DetailsIncome,
				constants.DetailsTip,
				constants.DetailsCard,
			),
			wantResult: dbResponse.DetailsIncome + dbResponse.DetailsTip + dbResponse.DetailsCard,
		},
		{
			testValue: fmt.Sprintf(
				"%s * %s - %s * %s + 3000",
				constants.DetailsIncome,
				constants.DriverPercent,
				constants.DetailsCard,
				constants.DriverPercent,
			),
			wantResult: dbResponse.DetailsIncome*dbResponse.DriverPercent - dbResponse.DetailsCard*dbResponse.DriverPercent + 3000,
		},
	}
}

// TestGoodVariable
func (suite *CalculateFromFormulaSuite) TestGoodVariable() {
	for _, v := range suite.GoodVariable {
		result, err := lib.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.Nil(suite.T(), err, "Should be nil")
		assert.Equal(suite.T(), v.wantResult, result, "Should be equal")
	}
}

// TestBadVariableErrorNul
func (suite *CalculateFromFormulaSuite) TestBadVariableErrorNul() {
	for _, v := range suite.BadVariableErrorNul {
		_, err := lib.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.Nil(suite.T(), err, "Should be nil")
	}
}

// TestBadVariableErrorNotNul
func (suite *CalculateFromFormulaSuite) TestBadVariableErrorNotNul() {
	for _, v := range suite.BadVariableErrorNotNul {
		_, err := lib.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.NotNil(suite.T(), err, "Should not be nil")
	}
}

// TestMockDbVariables
func (suite *CalculateFromFormulaSuite) TestMockDbVariables() {
	for _, v := range suite.MockRequestVariable {
		result, err := lib.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.Nil(suite.T(), err, "Should be nil")
		assert.Equal(suite.T(), v.wantResult, result, "Should be equal")
	}
}

// TestSuite run tests
func TestCalculateFromFormulaSuite(t *testing.T) {
	suite.Run(t, new(CalculateFromFormulaSuite))
}
