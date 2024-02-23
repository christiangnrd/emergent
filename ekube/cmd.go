// Copyright (c) 2024, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os/exec"

	"cogentcore.org/core/xe"
)

// Build builds a Docker image for the emergent model in the current directory.
func Build(c *Config) error { //gti:add
	err := InstallDeps()
	if err != nil {
		return err
	}
	return nil
}

func InstallDeps() error {
	if _, err := exec.LookPath("fyne-cross"); err != nil {
		err := xe.Run("go", "install", "-tags", "k8s", "github.com/fyne-io/fyne-cross@latest")
		if err != nil {
			return err
		}
	}
	return nil
}
