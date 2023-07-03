// Copyright (c) 2023, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package econfig

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"golang.org/x/exp/maps"
)

// TestSubConfig is a sub-struct with special params
type TestSubConfig struct {
	NPats      int     `def:"10" desc:"number of patterns to create"`
	Sparseness float32 `def:"0.15" desc:"proportion activity of created params"`
}

// TestConfig is a testing config
type TestConfig struct {
	Includes     []string       `desc:"specify include files here, and after configuration, it contains list of include files added"`
	GUI          bool           `def:"true" desc:"open the GUI -- does not automatically run -- if false, then runs automatically and quits"`
	GPU          bool           `def:"true" desc:"use the GPU for computation"`
	Debug        bool           `desc:"log debugging information"`
	PatParams    TestSubConfig  `desc:"important for testing . notation etc"`
	Network      map[string]any `desc:"network parameters applied after built-in params -- use toml map format: '{key = val, key2 = val2}' where key is 'selector:path' (e.g., '.PFCLayer:Layer.Inhib.Layer.Gi' where '.PFCLayer' is a class) and values should be strings to be consistent with standard params format"`
	ParamSet     string         `desc:"ParamSet name to use -- must be valid name as listed in compiled-in params or loaded params"`
	ParamFile    string         `desc:"Name of the JSON file to input saved parameters from."`
	ParamDocFile string         `desc:"Name of the file to output all parameter data. If not empty string, program should write file(s) and then exit"`
	Tag          string         `desc:"extra tag to add to file names and logs saved from this run"`
	Note         string         `def:"testing is fun" desc:"user note -- describe the run params etc -- like a git commit message for the run"`
	Run          int            `def:"0" desc:"starting run number -- determines the random seed -- runs counts from there -- can do all runs in parallel by launching separate jobs with each run, runs = 1"`
	Runs         int            `def:"10" desc:"total number of runs to do when running Train"`
	Epochs       int            `def:"100" desc:"total number of epochs per run"`
	NTrials      int            `def:"128" desc:"total number of trials per epoch.  Should be an even multiple of NData."`
	NData        int            `def:"16" desc:"number of data-parallel items to process in parallel per trial -- works (and is significantly faster) for both CPU and GPU.  Results in an effective mini-batch of learning."`
	SaveWts      bool           `desc:"if true, save final weights after each run"`
	EpochLog     bool           `def:"true" desc:"if true, save train epoch log to file, as .epc.tsv typically"`
	RunLog       bool           `def:"true" desc:"if true, save run log to file, as .run.tsv typically"`
	TrialLog     bool           `def:"true" desc:"if true, save train trial log to file, as .trl.tsv typically. May be large."`
	TestEpochLog bool           `def:"false" desc:"if true, save testing epoch log to file, as .tst_epc.tsv typically.  In general it is better to copy testing items over to the training epoch log and record there."`
	TestTrialLog bool           `def:"false" desc:"if true, save testing trial log to file, as .tst_trl.tsv typically. May be large."`
	NetData      bool           `desc:"if true, save network activation etc data from testing trials, for later viewing in netview"`
	Enum         TestEnum       `desc:"can set these values by string representation if stringer and registered as an enum with kit"`
}

func (cfg *TestConfig) IncludesPtr() *[]string { return &cfg.Includes }

func TestDefaults(t *testing.T) {
	cfg := &TestConfig{}
	SetFromDefaults(cfg)
	if cfg.Epochs != 100 || cfg.EpochLog != true || cfg.Note != "testing is fun" {
		t.Errorf("Main defaults failed to set")
	}
	if cfg.PatParams.NPats != 10 || cfg.PatParams.Sparseness != 0.15 {
		t.Errorf("PatParams defaults failed to set")
	}
}

func TestArgsPrint(t *testing.T) {
	t.Skip("prints all possible args")

	cfg := &TestConfig{}
	allArgs := FieldArgNames(cfg)

	keys := maps.Keys(allArgs)
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	fmt.Println("Args:")
	fmt.Println(strings.Join(keys, "\n"))
}

func TestArgs(t *testing.T) {
	cfg := &TestConfig{}
	SetFromDefaults(cfg)
	// note: cannot use "-Includes=testcfg.toml",
	args := []string{"-save-wts", "-nogui", "-no-epoch-log", "--NoRunLog", "--runs=5", "--run", "1", "--TAG", "nice", "--PatParams.Sparseness=0.1", "--Network", "{'.PFCLayer:Layer.Inhib.Gi' = '2.4', '#VSPatchPrjn:Prjn.Learn.LRate' = '0.01'}", "-Enum=TestValue2", "leftover1", "leftover2"}
	leftovers, err := parseArgs(cfg, args)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(leftovers)
	if cfg.Runs != 5 || cfg.Run != 1 || cfg.Tag != "nice" || cfg.PatParams.Sparseness != 0.1 || cfg.SaveWts != true || cfg.GUI != false || cfg.EpochLog != false || cfg.RunLog != false {
		t.Errorf("args not set properly: %#v", cfg)
	}
	if cfg.Enum != TestValue2 {
		t.Errorf("args enum from string not set properly: %#v", cfg)
	}

	// if cfg.Network != nil {
	// 	mv := cfg.Network
	// 	for k, v := range mv {
	// 		fmt.Println(k, " = ", v)
	// 	}
	// }
}

func TestOpen(t *testing.T) {
	IncludePaths = []string{".", "testdata"}
	cfg := &TestConfig{}
	err := OpenWithIncludes(cfg, "testcfg.toml")
	if err != nil {
		t.Errorf(err.Error())
	}

	// fmt.Println("includes:", cfg.Includes)

	// if cfg.Network != nil {
	// 	mv := cfg.Network
	// 	for k, v := range mv {
	// 		fmt.Println(k, " = ", v)
	// 	}
	// }

	if cfg.GUI != true || cfg.Tag != "testing" {
		t.Errorf("testinc.toml not parsed\n")
	}
	if cfg.Epochs != 500 || cfg.GPU != true {
		t.Errorf("testinc2.toml not parsed\n")
	}
	if cfg.Note != "something else" {
		t.Errorf("testinc3.toml not parsed\n")
	}
	if cfg.Runs != 8 {
		t.Errorf("testinc3.toml didn't overwrite testinc2\n")
	}
	if cfg.NTrials != 32 {
		t.Errorf("testinc.toml didn't overwrite testinc2\n")
	}
	if cfg.NData != 12 {
		t.Errorf("testcfg.toml didn't overwrite testinc3\n")
	}
	if cfg.Enum != TestValue2 {
		t.Errorf("testinc.toml Enum value not parsed\n")
	}
}

func TestUsage(t *testing.T) {
	t.Skip("prints usage string")
	cfg := &TestConfig{}
	us := Usage(cfg)
	fmt.Println(us)
}

func TestSave(t *testing.T) {
	// t.Skip("prints usage string")
	IncludePaths = []string{".", "testdata"}
	cfg := &TestConfig{}
	OpenWithIncludes(cfg, "testcfg.toml")
	Save(cfg, "testdata/testwrite.toml")
}
