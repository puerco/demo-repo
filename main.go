// SPDX-FileCopyrightText: Copyright 2025 Carabiner Systems, Inc
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/carabiner-dev/lab-vexable-repo/pkg/insult"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
)

func main() {
	logrus.Info("Sending world greetings")
	fmt.Println("Hello World!")
	logrus.Info("Trying to connect:")
	causeCVE()
}

// causeCVE is a function that uses a demo go module known to
// contain these (non-exploitable) vulnerabilities:
//
//	CVE-2025-22870 (not exploitable by insult dialer)
//	CVE-2025-22872 (not reachable)
//	CVE-2025-22871 (not reachable)
func causeCVE() {
	address := "[::1%25.example.com]:80"

	// This calls the Dial function of the proxy.PerHost struct which is affected
	// by CVE-2025-22870.
	// //
	// But note that we are not really affected as we using a fake dialer that
	// whose purpose is to insult you and always fails to connect:
	ph := proxy.NewPerHost(insult.NewDialer(), insult.NewDialer())

	// Open a connection with the InsultDialer:
	_, err := ph.Dial("tcp", address)

	// Say hi.
	fmt.Printf("\nError: %v\n\n", err)
}
