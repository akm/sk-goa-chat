// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"applib/time"
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Channel is an object representing the database table.
type Channel struct {
	ID         uint64             `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  time.Time          `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt  time.Time          `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	Name       string             `boil:"name" json:"name" toml:"name" yaml:"name"`
	Visibility ChannelsVisibility `boil:"visibility" json:"visibility" toml:"visibility" yaml:"visibility"`

	R *channelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L channelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChannelColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	Name       string
	Visibility string
}{
	ID:         "id",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	Name:       "name",
	Visibility: "visibility",
}

var ChannelTableColumns = struct {
	ID         string
	CreatedAt  string
	UpdatedAt  string
	Name       string
	Visibility string
}{
	ID:         "channels.id",
	CreatedAt:  "channels.created_at",
	UpdatedAt:  "channels.updated_at",
	Name:       "channels.name",
	Visibility: "channels.visibility",
}

// Generated where

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperChannelsVisibility struct{ field string }

func (w whereHelperChannelsVisibility) EQ(x ChannelsVisibility) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelperChannelsVisibility) NEQ(x ChannelsVisibility) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelperChannelsVisibility) LT(x ChannelsVisibility) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelperChannelsVisibility) LTE(x ChannelsVisibility) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelperChannelsVisibility) GT(x ChannelsVisibility) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelperChannelsVisibility) GTE(x ChannelsVisibility) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelperChannelsVisibility) IN(slice []ChannelsVisibility) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperChannelsVisibility) NIN(slice []ChannelsVisibility) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var ChannelWhere = struct {
	ID         whereHelperuint64
	CreatedAt  whereHelpertime_Time
	UpdatedAt  whereHelpertime_Time
	Name       whereHelperstring
	Visibility whereHelperChannelsVisibility
}{
	ID:         whereHelperuint64{field: "`channels`.`id`"},
	CreatedAt:  whereHelpertime_Time{field: "`channels`.`created_at`"},
	UpdatedAt:  whereHelpertime_Time{field: "`channels`.`updated_at`"},
	Name:       whereHelperstring{field: "`channels`.`name`"},
	Visibility: whereHelperChannelsVisibility{field: "`channels`.`visibility`"},
}

// ChannelRels is where relationship names are stored.
var ChannelRels = struct {
	ChatMessages string
}{
	ChatMessages: "ChatMessages",
}

// channelR is where relationships are stored.
type channelR struct {
	ChatMessages ChatMessageSlice `boil:"ChatMessages" json:"ChatMessages" toml:"ChatMessages" yaml:"ChatMessages"`
}

// NewStruct creates a new relationship struct
func (*channelR) NewStruct() *channelR {
	return &channelR{}
}

func (r *channelR) GetChatMessages() ChatMessageSlice {
	if r == nil {
		return nil
	}
	return r.ChatMessages
}

// channelL is where Load methods for each relationship are stored.
type channelL struct{}

var (
	channelAllColumns            = []string{"id", "created_at", "updated_at", "name", "visibility"}
	channelColumnsWithoutDefault = []string{"name", "visibility"}
	channelColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	channelPrimaryKeyColumns     = []string{"id"}
	channelGeneratedColumns      = []string{}
)

type (
	// ChannelSlice is an alias for a slice of pointers to Channel.
	// This should almost always be used instead of []Channel.
	ChannelSlice []*Channel
	// ChannelHook is the signature for custom Channel hook methods
	ChannelHook func(context.Context, boil.ContextExecutor, *Channel) error

	channelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	channelType                 = reflect.TypeOf(&Channel{})
	channelMapping              = queries.MakeStructMapping(channelType)
	channelPrimaryKeyMapping, _ = queries.BindMapping(channelType, channelMapping, channelPrimaryKeyColumns)
	channelInsertCacheMut       sync.RWMutex
	channelInsertCache          = make(map[string]insertCache)
	channelUpdateCacheMut       sync.RWMutex
	channelUpdateCache          = make(map[string]updateCache)
	channelUpsertCacheMut       sync.RWMutex
	channelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var channelAfterSelectMu sync.Mutex
var channelAfterSelectHooks []ChannelHook

var channelBeforeInsertMu sync.Mutex
var channelBeforeInsertHooks []ChannelHook
var channelAfterInsertMu sync.Mutex
var channelAfterInsertHooks []ChannelHook

var channelBeforeUpdateMu sync.Mutex
var channelBeforeUpdateHooks []ChannelHook
var channelAfterUpdateMu sync.Mutex
var channelAfterUpdateHooks []ChannelHook

var channelBeforeDeleteMu sync.Mutex
var channelBeforeDeleteHooks []ChannelHook
var channelAfterDeleteMu sync.Mutex
var channelAfterDeleteHooks []ChannelHook

var channelBeforeUpsertMu sync.Mutex
var channelBeforeUpsertHooks []ChannelHook
var channelAfterUpsertMu sync.Mutex
var channelAfterUpsertHooks []ChannelHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Channel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Channel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Channel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Channel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Channel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Channel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Channel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Channel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Channel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range channelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChannelHook registers your hook function for all future operations.
func AddChannelHook(hookPoint boil.HookPoint, channelHook ChannelHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		channelAfterSelectMu.Lock()
		channelAfterSelectHooks = append(channelAfterSelectHooks, channelHook)
		channelAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		channelBeforeInsertMu.Lock()
		channelBeforeInsertHooks = append(channelBeforeInsertHooks, channelHook)
		channelBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		channelAfterInsertMu.Lock()
		channelAfterInsertHooks = append(channelAfterInsertHooks, channelHook)
		channelAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		channelBeforeUpdateMu.Lock()
		channelBeforeUpdateHooks = append(channelBeforeUpdateHooks, channelHook)
		channelBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		channelAfterUpdateMu.Lock()
		channelAfterUpdateHooks = append(channelAfterUpdateHooks, channelHook)
		channelAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		channelBeforeDeleteMu.Lock()
		channelBeforeDeleteHooks = append(channelBeforeDeleteHooks, channelHook)
		channelBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		channelAfterDeleteMu.Lock()
		channelAfterDeleteHooks = append(channelAfterDeleteHooks, channelHook)
		channelAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		channelBeforeUpsertMu.Lock()
		channelBeforeUpsertHooks = append(channelBeforeUpsertHooks, channelHook)
		channelBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		channelAfterUpsertMu.Lock()
		channelAfterUpsertHooks = append(channelAfterUpsertHooks, channelHook)
		channelAfterUpsertMu.Unlock()
	}
}

// One returns a single channel record from the query.
func (q channelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Channel, error) {
	o := &Channel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for channels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Channel records from the query.
func (q channelQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChannelSlice, error) {
	var o []*Channel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Channel slice")
	}

	if len(channelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Channel records in the query.
func (q channelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q channelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if channels exists")
	}

	return count > 0, nil
}

// ChatMessages retrieves all the chat_message's ChatMessages with an executor.
func (o *Channel) ChatMessages(mods ...qm.QueryMod) chatMessageQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`chat_messages`.`channel_id`=?", o.ID),
	)

	return ChatMessages(queryMods...)
}

// LoadChatMessages allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (channelL) LoadChatMessages(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChannel interface{}, mods queries.Applicator) error {
	var slice []*Channel
	var object *Channel

	if singular {
		var ok bool
		object, ok = maybeChannel.(*Channel)
		if !ok {
			object = new(Channel)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeChannel)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeChannel))
			}
		}
	} else {
		s, ok := maybeChannel.(*[]*Channel)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeChannel)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeChannel))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &channelR{}
		}
		args[object.ID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &channelR{}
			}
			args[obj.ID] = struct{}{}
		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`chat_messages`),
		qm.WhereIn(`chat_messages.channel_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load chat_messages")
	}

	var resultSlice []*ChatMessage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice chat_messages")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on chat_messages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chat_messages")
	}

	if len(chatMessageAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChatMessages = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &chatMessageR{}
			}
			foreign.R.Channel = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ChannelID {
				local.R.ChatMessages = append(local.R.ChatMessages, foreign)
				if foreign.R == nil {
					foreign.R = &chatMessageR{}
				}
				foreign.R.Channel = local
				break
			}
		}
	}

	return nil
}

// AddChatMessages adds the given related objects to the existing relationships
// of the channel, optionally inserting them as new records.
// Appends related to o.R.ChatMessages.
// Sets related.R.Channel appropriately.
func (o *Channel) AddChatMessages(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChatMessage) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChannelID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `chat_messages` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"channel_id"}),
				strmangle.WhereClause("`", "`", 0, chatMessagePrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChannelID = o.ID
		}
	}

	if o.R == nil {
		o.R = &channelR{
			ChatMessages: related,
		}
	} else {
		o.R.ChatMessages = append(o.R.ChatMessages, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &chatMessageR{
				Channel: o,
			}
		} else {
			rel.R.Channel = o
		}
	}
	return nil
}

// Channels retrieves all the records using an executor.
func Channels(mods ...qm.QueryMod) channelQuery {
	mods = append(mods, qm.From("`channels`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`channels`.*"})
	}

	return channelQuery{q}
}

// FindChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChannel(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*Channel, error) {
	channelObj := &Channel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `channels` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, channelObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from channels")
	}

	if err = channelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return channelObj, err
	}

	return channelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Channel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no channels provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(channelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	channelInsertCacheMut.RLock()
	cache, cached := channelInsertCache[key]
	channelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			channelAllColumns,
			channelColumnsWithDefault,
			channelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(channelType, channelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(channelType, channelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `channels` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `channels` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `channels` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, channelPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into channels")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == channelMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for channels")
	}

CacheNoHooks:
	if !cached {
		channelInsertCacheMut.Lock()
		channelInsertCache[key] = cache
		channelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Channel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Channel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	channelUpdateCacheMut.RLock()
	cache, cached := channelUpdateCache[key]
	channelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			channelAllColumns,
			channelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `channels` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, channelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(channelType, channelMapping, append(wl, channelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update channels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for channels")
	}

	if !cached {
		channelUpdateCacheMut.Lock()
		channelUpdateCache[key] = cache
		channelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q channelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for channels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChannelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `channels` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in channel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all channel")
	}
	return rowsAff, nil
}

var mySQLChannelUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Channel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no channels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(channelColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLChannelUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	channelUpsertCacheMut.RLock()
	cache, cached := channelUpsertCache[key]
	channelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			channelAllColumns,
			channelColumnsWithDefault,
			channelColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			channelAllColumns,
			channelPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert channels, could not build update column list")
		}

		ret := strmangle.SetComplement(channelAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`channels`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `channels` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(channelType, channelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(channelType, channelMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for channels")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == channelMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(channelType, channelMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for channels")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for channels")
	}

CacheNoHooks:
	if !cached {
		channelUpsertCacheMut.Lock()
		channelUpsertCache[key] = cache
		channelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Channel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Channel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Channel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), channelPrimaryKeyMapping)
	sql := "DELETE FROM `channels` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for channels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q channelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no channelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for channels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChannelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(channelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `channels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from channel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for channels")
	}

	if len(channelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Channel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChannel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChannelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), channelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `channels`.* FROM `channels` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, channelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChannelSlice")
	}

	*o = slice

	return nil
}

// ChannelExists checks if the Channel row exists.
func ChannelExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `channels` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if channels exists")
	}

	return exists, nil
}

// Exists checks if the Channel row exists.
func (o *Channel) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ChannelExists(ctx, exec, o.ID)
}
