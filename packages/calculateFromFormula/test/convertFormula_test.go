package calculateFromFormula

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	_ "github.com/stretchr/testify/suite"
	"testProject/packages/calculateFromFormula"
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
	MockDbResponse         *calculateFromFormula.FormulaResponseDb
}

func (suite *CalculateFromFormulaSuite) SetupTest() {
	suite.GoodVariable = []variableTestCalculateFromFormula{
		{testValue: "5 + 5", wantResult: 10},
		{testValue: "5 - 5", wantResult: 0},
		{testValue: "5 * 5", wantResult: 25},
		{testValue: "5 / 5", wantResult: 1},
		{testValue: "5 + (5 * 2)", wantResult: 15},
		{testValue: "5 + 5 * 2", wantResult: 15},
		{testValue: "5.5 + 5 * 2", wantResult: 15.5},
		{testValue: "-10 - 5.5", wantResult: -15.5},
		{testValue: "-10 + 5.5", wantResult: -4.5},
	}

	suite.BadVariableErrorNul = []variableTestCalculateFromFormula{
		{testValue: "0 / 0"},
		{testValue: "5 / 0"},
	}

	suite.BadVariableErrorNotNul = []variableTestCalculateFromFormula{
		{testValue: "5 / "},
		{testValue: "5 * some text"},
	}

	suite.MockDbResponse = calculateFromFormula.NewMockDbResponse()
	dbResponse := suite.MockDbResponse

	suite.MockRequestVariable = []variableTestCalculateFromFormula{
		{testValue: calculateFromFormula.DetailsCard.ToString(), wantResult: dbResponse.DetailsCard},
		{
			testValue: fmt.Sprintf(
				"%s + %s + %s",
				calculateFromFormula.CashIncome,
				calculateFromFormula.DetailsCard,
				calculateFromFormula.Tip,
			),
			wantResult: dbResponse.CashIncome + dbResponse.DetailsCard + dbResponse.Tip,
		},
	}
}

// TestGoodVariable
func (suite *CalculateFromFormulaSuite) TestGoodVariable() {
	for _, v := range suite.GoodVariable {
		result, err := calculateFromFormula.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.Nil(suite.T(), err, fmt.Sprintf("Should be nil.Value: %+v. Result: %f", v, err))
		assert.Equal(suite.T(), v.wantResult, result, "Should be equal")
	}
}

// TestBadVariableErrorNul
func (suite *CalculateFromFormulaSuite) TestBadVariableErrorNul() {
	for _, v := range suite.BadVariableErrorNul {
		_, err := calculateFromFormula.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.Nil(suite.T(), err, "Should be nil")
	}
}

// TestBadVariableErrorNotNul
func (suite *CalculateFromFormulaSuite) TestBadVariableErrorNotNul() {
	for _, v := range suite.BadVariableErrorNotNul {
		_, err := calculateFromFormula.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.NotNil(suite.T(), err, "Should not be nil")
	}
}

// TestMockDbVariables
func (suite *CalculateFromFormulaSuite) TestMockDbVariables() {
	for _, v := range suite.MockRequestVariable {
		result, err := calculateFromFormula.CalculateFromFormula(v.testValue, suite.MockDbResponse)
		assert.Nil(suite.T(), err, "Should be nil")
		assert.Equal(suite.T(), v.wantResult, result, "Should be equal")
	}
}

// TestGenerateParameters
func (suite *CalculateFromFormulaSuite) TestGenerateParameters() {
	repo := calculateFromFormula.NewMockDbResponse()
	parameters, err := calculateFromFormula.GenerateParameters(repo)

	assert.Nil(suite.T(), err, "Should be nil")
	assert.IsType(suite.T(), map[string]interface{}{}, parameters, "should be map[string]interface{}{} type")
}

// TestSuite run tests
func TestCalculateFromFormulaSuite(t *testing.T) {
	suite.Run(t, new(CalculateFromFormulaSuite))
}
