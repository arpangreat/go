package checkers

import (
	"testing"

	"github.com/go-critic/go-critic/framework/linter"
	"github.com/go-critic/go-critic/framework/linttest"
)

func TestCheckers(t *testing.T) {
	allParams := map[string]map[string]interface{}{
		"captLocal": {"paramsOnly": false},
	}

	for _, info := range linter.GetCheckersInfo() {
		params := allParams[info.Name]
		for key, p := range info.Params {
			v, ok := params[key]
			if ok {
				p.Value = v
			}
		}
	}

	cfg := linttest.CheckersTest{
		IgnoreErrors: []string{
			"caseOrder",
		},
	}

	cfg.Run(t)
}

func TestIntegration(t *testing.T) {
	cfg := linttest.IntegrationTest{
		Main: "github.com/go-critic/go-critic/cmd/gocritic",
		Dir:  "./testdata/_integration",
	}
	cfg.Run(t)
}

func TestTags(t *testing.T) {
	// Verify that we're only using strict set of tags.
	// This helps to avoid typos in tag names.
	//
	// Also check that exactly 1 category tag is used.

	for _, info := range linter.GetCheckersInfo() {
		categories := 0
		for _, tag := range info.Tags {
			switch tag {
			case "diagnostic", "style", "performance":
				// Category tags.
				// Can only have one of them.
				categories++
			case "experimental", "opinionated":
				// Optional tags.
			default:
				t.Errorf("%q checker uses unknown tag %q", info.Name, tag)
			}
		}
		if categories != 1 {
			t.Errorf("%q expected to have 1 category, found %d",
				info.Name, categories)
		}
	}
}

func TestDocs(t *testing.T) {
	for _, info := range linter.GetCheckersInfo() {
		if info.Summary == "" {
			t.Errorf("%q checker lacks summary", info.Name)
		}
		for key, p := range info.Params {
			if p.Usage == "" {
				t.Errorf("%q checker %q param lacks usage docs",
					info.Name, key)
			}
		}
	}
}

func TestStableList(t *testing.T) {
	// Verify that new checker is not added without "experimental"
	// tag by accident. When stable checker is about to be added,
	// slice above should be modified to include new checker name.

	// It is a good practice to keep this list sorted.
	stableList := []string{
		"appendAssign",
		"appendCombine",
		"argOrder",
		"assignOp",
		"badCall",
		"badCond",
		"builtinShadow",
		"captLocal",
		"caseOrder",
		"codegenComment",
		"commentFormatting",
		"defaultCaseOrder",
		"deprecatedComment",
		"dupArg",
		"dupBranchBody",
		"dupCase",
		"dupSubExpr",
		"elseif",
		"exitAfterDefer",
		"flagDeref",
		"flagName",
		"hugeParam",
		"ifElseChain",
		"importShadow",
		"indexAlloc",
		"mapKey",
		"newDeref",
		"offBy1",
		"paramTypeCombine",
		"rangeExprCopy",
		"rangeValCopy",
		"regexpMust",
		"singleCaseSwitch",
		"sloppyLen",
		"stringXbytes",
		"switchTrue",
		"typeSwitchVar",
		"typeUnparen",
		"underef",
		"unlambda",
		"unslice",
		"valSwap",
		"wrapperFunc",
	}

	m := make(map[string]bool)
	for _, name := range stableList {
		m[name] = true
	}

	for _, info := range linter.GetCheckersInfo() {
		if info.HasTag("experimental") {
			continue
		}
		if !m[info.Name] {
			t.Errorf("%q checker misses `experimental` tag", info.Name)
		}
	}
}
