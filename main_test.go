// SPDX-FileCopyrightText: Copyright 2025 Carabiner Systems, Inc
// SPDX-License-Identifier: Apache-2.0

package main_test

import "testing"

func TestSomething(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		t.Log("I ran!")
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Log("I failed!")
		t.Fail()
	})
}
