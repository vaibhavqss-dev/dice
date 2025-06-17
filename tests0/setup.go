package ironhawk

// Copyright (c) 2022-present, DiceDB contributors
// All rights reserved. Licensed under the BSD 3-Clause License. See LICENSE file in the project root for full license information.

import (
	"time"

	"github.com/dicedb/dice/config"
	"github.com/dicedb/dicedb-go"
)

type TestCase struct {
	name     string
	commands []string
	expected []interface{}
	delay    []time.Duration
}

//nolint:unused
func getLocalConnection() *dicedb.Client {
	client, err := dicedb.NewClient("localhost", config.Config.Port)
	if err != nil {
		panic(err)
	}
	return client
}
