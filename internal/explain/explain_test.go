package explain

import (
	"strings"
	"testing"

	"github.com/nauticalabs/modelmariner/internal/policy"
	"github.com/nauticalabs/modelmariner/internal/routing"
)

func TestExplainWinner(t *testing.T) {
	d := routing.Decision{
