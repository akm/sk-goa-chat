// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Channels", testChannels)
	t.Run("GooseDBVersions", testGooseDBVersions)
}

func TestDelete(t *testing.T) {
	t.Run("Channels", testChannelsDelete)
	t.Run("GooseDBVersions", testGooseDBVersionsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Channels", testChannelsQueryDeleteAll)
	t.Run("GooseDBVersions", testGooseDBVersionsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Channels", testChannelsSliceDeleteAll)
	t.Run("GooseDBVersions", testGooseDBVersionsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Channels", testChannelsExists)
	t.Run("GooseDBVersions", testGooseDBVersionsExists)
}

func TestFind(t *testing.T) {
	t.Run("Channels", testChannelsFind)
	t.Run("GooseDBVersions", testGooseDBVersionsFind)
}

func TestBind(t *testing.T) {
	t.Run("Channels", testChannelsBind)
	t.Run("GooseDBVersions", testGooseDBVersionsBind)
}

func TestOne(t *testing.T) {
	t.Run("Channels", testChannelsOne)
	t.Run("GooseDBVersions", testGooseDBVersionsOne)
}

func TestAll(t *testing.T) {
	t.Run("Channels", testChannelsAll)
	t.Run("GooseDBVersions", testGooseDBVersionsAll)
}

func TestCount(t *testing.T) {
	t.Run("Channels", testChannelsCount)
	t.Run("GooseDBVersions", testGooseDBVersionsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Channels", testChannelsHooks)
	t.Run("GooseDBVersions", testGooseDBVersionsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Channels", testChannelsInsert)
	t.Run("Channels", testChannelsInsertWhitelist)
	t.Run("GooseDBVersions", testGooseDBVersionsInsert)
	t.Run("GooseDBVersions", testGooseDBVersionsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Channels", testChannelsReload)
	t.Run("GooseDBVersions", testGooseDBVersionsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Channels", testChannelsReloadAll)
	t.Run("GooseDBVersions", testGooseDBVersionsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Channels", testChannelsSelect)
	t.Run("GooseDBVersions", testGooseDBVersionsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Channels", testChannelsUpdate)
	t.Run("GooseDBVersions", testGooseDBVersionsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Channels", testChannelsSliceUpdateAll)
	t.Run("GooseDBVersions", testGooseDBVersionsSliceUpdateAll)
}