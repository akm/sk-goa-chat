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
	t.Run("ChatMessages", testChatMessages)
	t.Run("GooseDBVersions", testGooseDBVersions)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Channels", testChannelsDelete)
	t.Run("ChatMessages", testChatMessagesDelete)
	t.Run("GooseDBVersions", testGooseDBVersionsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Channels", testChannelsQueryDeleteAll)
	t.Run("ChatMessages", testChatMessagesQueryDeleteAll)
	t.Run("GooseDBVersions", testGooseDBVersionsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Channels", testChannelsSliceDeleteAll)
	t.Run("ChatMessages", testChatMessagesSliceDeleteAll)
	t.Run("GooseDBVersions", testGooseDBVersionsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Channels", testChannelsExists)
	t.Run("ChatMessages", testChatMessagesExists)
	t.Run("GooseDBVersions", testGooseDBVersionsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Channels", testChannelsFind)
	t.Run("ChatMessages", testChatMessagesFind)
	t.Run("GooseDBVersions", testGooseDBVersionsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Channels", testChannelsBind)
	t.Run("ChatMessages", testChatMessagesBind)
	t.Run("GooseDBVersions", testGooseDBVersionsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Channels", testChannelsOne)
	t.Run("ChatMessages", testChatMessagesOne)
	t.Run("GooseDBVersions", testGooseDBVersionsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Channels", testChannelsAll)
	t.Run("ChatMessages", testChatMessagesAll)
	t.Run("GooseDBVersions", testGooseDBVersionsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Channels", testChannelsCount)
	t.Run("ChatMessages", testChatMessagesCount)
	t.Run("GooseDBVersions", testGooseDBVersionsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Channels", testChannelsHooks)
	t.Run("ChatMessages", testChatMessagesHooks)
	t.Run("GooseDBVersions", testGooseDBVersionsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Channels", testChannelsInsert)
	t.Run("Channels", testChannelsInsertWhitelist)
	t.Run("ChatMessages", testChatMessagesInsert)
	t.Run("ChatMessages", testChatMessagesInsertWhitelist)
	t.Run("GooseDBVersions", testGooseDBVersionsInsert)
	t.Run("GooseDBVersions", testGooseDBVersionsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChatMessageToChannelUsingChannel", testChatMessageToOneChannelUsingChannel)
	t.Run("ChatMessageToUserUsingUser", testChatMessageToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("ChannelToChatMessages", testChannelToManyChatMessages)
	t.Run("UserToChatMessages", testUserToManyChatMessages)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChatMessageToChannelUsingChatMessages", testChatMessageToOneSetOpChannelUsingChannel)
	t.Run("ChatMessageToUserUsingChatMessages", testChatMessageToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ChatMessageToUserUsingChatMessages", testChatMessageToOneRemoveOpUserUsingUser)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("ChannelToChatMessages", testChannelToManyAddOpChatMessages)
	t.Run("UserToChatMessages", testUserToManyAddOpChatMessages)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("UserToChatMessages", testUserToManySetOpChatMessages)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("UserToChatMessages", testUserToManyRemoveOpChatMessages)
}

func TestReload(t *testing.T) {
	t.Run("Channels", testChannelsReload)
	t.Run("ChatMessages", testChatMessagesReload)
	t.Run("GooseDBVersions", testGooseDBVersionsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Channels", testChannelsReloadAll)
	t.Run("ChatMessages", testChatMessagesReloadAll)
	t.Run("GooseDBVersions", testGooseDBVersionsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Channels", testChannelsSelect)
	t.Run("ChatMessages", testChatMessagesSelect)
	t.Run("GooseDBVersions", testGooseDBVersionsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Channels", testChannelsUpdate)
	t.Run("ChatMessages", testChatMessagesUpdate)
	t.Run("GooseDBVersions", testGooseDBVersionsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Channels", testChannelsSliceUpdateAll)
	t.Run("ChatMessages", testChatMessagesSliceUpdateAll)
	t.Run("GooseDBVersions", testGooseDBVersionsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
