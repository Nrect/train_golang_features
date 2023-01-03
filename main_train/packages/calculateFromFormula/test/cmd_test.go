package calculateFromFormula

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestAllSuites(t *testing.T) {
	suite.Run(t, new(CalculateFromFormulaSuite))
}
