// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"caster/ent/migrate"

	"caster/ent/episode"
	"caster/ent/profile"
	"caster/ent/show"
	"caster/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
	// Show is the client for interacting with the Show builders.
	Show *ShowClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Episode = NewEpisodeClient(c.config)
	c.Profile = NewProfileClient(c.config)
	c.Show = NewShowClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Episode: NewEpisodeClient(cfg),
		Profile: NewProfileClient(cfg),
		Show:    NewShowClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Episode: NewEpisodeClient(cfg),
		Profile: NewProfileClient(cfg),
		Show:    NewShowClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Episode.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Episode.Use(hooks...)
	c.Profile.Use(hooks...)
	c.Show.Use(hooks...)
	c.User.Use(hooks...)
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Create returns a create builder for Episode.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id string) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EpisodeClient) DeleteOneID(id string) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id string) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id string) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOwner queries the owner edge of a Episode.
func (c *EpisodeClient) QueryOwner(e *Episode) *ShowQuery {
	query := &ShowQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(show.Table, show.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.OwnerTable, episode.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// ProfileClient is a client for the Profile schema.
type ProfileClient struct {
	config
}

// NewProfileClient returns a client for the Profile from the given config.
func NewProfileClient(c config) *ProfileClient {
	return &ProfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profile.Hooks(f(g(h())))`.
func (c *ProfileClient) Use(hooks ...Hook) {
	c.hooks.Profile = append(c.hooks.Profile, hooks...)
}

// Create returns a create builder for Profile.
func (c *ProfileClient) Create() *ProfileCreate {
	mutation := newProfileMutation(c.config, OpCreate)
	return &ProfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profile entities.
func (c *ProfileClient) CreateBulk(builders ...*ProfileCreate) *ProfileCreateBulk {
	return &ProfileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profile.
func (c *ProfileClient) Update() *ProfileUpdate {
	mutation := newProfileMutation(c.config, OpUpdate)
	return &ProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfileClient) UpdateOne(pr *Profile) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfile(pr))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfileClient) UpdateOneID(id string) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfileID(id))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profile.
func (c *ProfileClient) Delete() *ProfileDelete {
	mutation := newProfileMutation(c.config, OpDelete)
	return &ProfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProfileClient) DeleteOne(pr *Profile) *ProfileDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProfileClient) DeleteOneID(id string) *ProfileDeleteOne {
	builder := c.Delete().Where(profile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfileDeleteOne{builder}
}

// Query returns a query builder for Profile.
func (c *ProfileClient) Query() *ProfileQuery {
	return &ProfileQuery{
		config: c.config,
	}
}

// Get returns a Profile entity by its id.
func (c *ProfileClient) Get(ctx context.Context, id string) (*Profile, error) {
	return c.Query().Where(profile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfileClient) GetX(ctx context.Context, id string) *Profile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProfileClient) Hooks() []Hook {
	return c.hooks.Profile
}

// ShowClient is a client for the Show schema.
type ShowClient struct {
	config
}

// NewShowClient returns a client for the Show from the given config.
func NewShowClient(c config) *ShowClient {
	return &ShowClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `show.Hooks(f(g(h())))`.
func (c *ShowClient) Use(hooks ...Hook) {
	c.hooks.Show = append(c.hooks.Show, hooks...)
}

// Create returns a create builder for Show.
func (c *ShowClient) Create() *ShowCreate {
	mutation := newShowMutation(c.config, OpCreate)
	return &ShowCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Show entities.
func (c *ShowClient) CreateBulk(builders ...*ShowCreate) *ShowCreateBulk {
	return &ShowCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Show.
func (c *ShowClient) Update() *ShowUpdate {
	mutation := newShowMutation(c.config, OpUpdate)
	return &ShowUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ShowClient) UpdateOne(s *Show) *ShowUpdateOne {
	mutation := newShowMutation(c.config, OpUpdateOne, withShow(s))
	return &ShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ShowClient) UpdateOneID(id string) *ShowUpdateOne {
	mutation := newShowMutation(c.config, OpUpdateOne, withShowID(id))
	return &ShowUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Show.
func (c *ShowClient) Delete() *ShowDelete {
	mutation := newShowMutation(c.config, OpDelete)
	return &ShowDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ShowClient) DeleteOne(s *Show) *ShowDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ShowClient) DeleteOneID(id string) *ShowDeleteOne {
	builder := c.Delete().Where(show.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ShowDeleteOne{builder}
}

// Query returns a query builder for Show.
func (c *ShowClient) Query() *ShowQuery {
	return &ShowQuery{
		config: c.config,
	}
}

// Get returns a Show entity by its id.
func (c *ShowClient) Get(ctx context.Context, id string) (*Show, error) {
	return c.Query().Where(show.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ShowClient) GetX(ctx context.Context, id string) *Show {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEpisodes queries the episodes edge of a Show.
func (c *ShowClient) QueryEpisodes(s *Show) *EpisodeQuery {
	query := &EpisodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(show.Table, show.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, show.EpisodesTable, show.EpisodesColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ShowClient) Hooks() []Hook {
	return c.hooks.Show
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
