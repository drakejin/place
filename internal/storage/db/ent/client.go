// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/drakejin/place/internal/storage/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/drakejin/place/internal/storage/db/ent/page"
	"github.com/drakejin/place/internal/storage/db/ent/pagereferred"
	"github.com/drakejin/place/internal/storage/db/ent/pagesource"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Page is the client for interacting with the Page builders.
	Page *PageClient
	// PageReferred is the client for interacting with the PageReferred builders.
	PageReferred *PageReferredClient
	// PageSource is the client for interacting with the PageSource builders.
	PageSource *PageSourceClient
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
	c.Page = NewPageClient(c.config)
	c.PageReferred = NewPageReferredClient(c.config)
	c.PageSource = NewPageSourceClient(c.config)
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Page:         NewPageClient(cfg),
		PageReferred: NewPageReferredClient(cfg),
		PageSource:   NewPageSourceClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:          ctx,
		config:       cfg,
		Page:         NewPageClient(cfg),
		PageReferred: NewPageReferredClient(cfg),
		PageSource:   NewPageSourceClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Page.
//		Query().
//		Count(ctx)
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
	c.Page.Use(hooks...)
	c.PageReferred.Use(hooks...)
	c.PageSource.Use(hooks...)
}

// PageClient is a client for the Page schema.
type PageClient struct {
	config
}

// NewPageClient returns a client for the Page from the given config.
func NewPageClient(c config) *PageClient {
	return &PageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `page.Hooks(f(g(h())))`.
func (c *PageClient) Use(hooks ...Hook) {
	c.hooks.Page = append(c.hooks.Page, hooks...)
}

// Create returns a builder for creating a Page entity.
func (c *PageClient) Create() *PageCreate {
	mutation := newPageMutation(c.config, OpCreate)
	return &PageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Page entities.
func (c *PageClient) CreateBulk(builders ...*PageCreate) *PageCreateBulk {
	return &PageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Page.
func (c *PageClient) Update() *PageUpdate {
	mutation := newPageMutation(c.config, OpUpdate)
	return &PageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PageClient) UpdateOne(pa *Page) *PageUpdateOne {
	mutation := newPageMutation(c.config, OpUpdateOne, withPage(pa))
	return &PageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PageClient) UpdateOneID(id uuid.UUID) *PageUpdateOne {
	mutation := newPageMutation(c.config, OpUpdateOne, withPageID(id))
	return &PageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Page.
func (c *PageClient) Delete() *PageDelete {
	mutation := newPageMutation(c.config, OpDelete)
	return &PageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PageClient) DeleteOne(pa *Page) *PageDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PageClient) DeleteOneID(id uuid.UUID) *PageDeleteOne {
	builder := c.Delete().Where(page.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PageDeleteOne{builder}
}

// Query returns a query builder for Page.
func (c *PageClient) Query() *PageQuery {
	return &PageQuery{
		config: c.config,
	}
}

// Get returns a Page entity by its id.
func (c *PageClient) Get(ctx context.Context, id uuid.UUID) (*Page, error) {
	return c.Query().Where(page.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PageClient) GetX(ctx context.Context, id uuid.UUID) *Page {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPageSource queries the page_source edge of a Page.
func (c *PageClient) QueryPageSource(pa *Page) *PageSourceQuery {
	query := &PageSourceQuery{config: c.config}
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(page.Table, page.FieldID, id),
			sqlgraph.To(pagesource.Table, pagesource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, page.PageSourceTable, page.PageSourceColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PageClient) Hooks() []Hook {
	return c.hooks.Page
}

// PageReferredClient is a client for the PageReferred schema.
type PageReferredClient struct {
	config
}

// NewPageReferredClient returns a client for the PageReferred from the given config.
func NewPageReferredClient(c config) *PageReferredClient {
	return &PageReferredClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pagereferred.Hooks(f(g(h())))`.
func (c *PageReferredClient) Use(hooks ...Hook) {
	c.hooks.PageReferred = append(c.hooks.PageReferred, hooks...)
}

// Create returns a builder for creating a PageReferred entity.
func (c *PageReferredClient) Create() *PageReferredCreate {
	mutation := newPageReferredMutation(c.config, OpCreate)
	return &PageReferredCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PageReferred entities.
func (c *PageReferredClient) CreateBulk(builders ...*PageReferredCreate) *PageReferredCreateBulk {
	return &PageReferredCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PageReferred.
func (c *PageReferredClient) Update() *PageReferredUpdate {
	mutation := newPageReferredMutation(c.config, OpUpdate)
	return &PageReferredUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PageReferredClient) UpdateOne(pr *PageReferred) *PageReferredUpdateOne {
	mutation := newPageReferredMutation(c.config, OpUpdateOne, withPageReferred(pr))
	return &PageReferredUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PageReferredClient) UpdateOneID(id uuid.UUID) *PageReferredUpdateOne {
	mutation := newPageReferredMutation(c.config, OpUpdateOne, withPageReferredID(id))
	return &PageReferredUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PageReferred.
func (c *PageReferredClient) Delete() *PageReferredDelete {
	mutation := newPageReferredMutation(c.config, OpDelete)
	return &PageReferredDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PageReferredClient) DeleteOne(pr *PageReferred) *PageReferredDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PageReferredClient) DeleteOneID(id uuid.UUID) *PageReferredDeleteOne {
	builder := c.Delete().Where(pagereferred.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PageReferredDeleteOne{builder}
}

// Query returns a query builder for PageReferred.
func (c *PageReferredClient) Query() *PageReferredQuery {
	return &PageReferredQuery{
		config: c.config,
	}
}

// Get returns a PageReferred entity by its id.
func (c *PageReferredClient) Get(ctx context.Context, id uuid.UUID) (*PageReferred, error) {
	return c.Query().Where(pagereferred.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PageReferredClient) GetX(ctx context.Context, id uuid.UUID) *PageReferred {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PageReferredClient) Hooks() []Hook {
	return c.hooks.PageReferred
}

// PageSourceClient is a client for the PageSource schema.
type PageSourceClient struct {
	config
}

// NewPageSourceClient returns a client for the PageSource from the given config.
func NewPageSourceClient(c config) *PageSourceClient {
	return &PageSourceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pagesource.Hooks(f(g(h())))`.
func (c *PageSourceClient) Use(hooks ...Hook) {
	c.hooks.PageSource = append(c.hooks.PageSource, hooks...)
}

// Create returns a builder for creating a PageSource entity.
func (c *PageSourceClient) Create() *PageSourceCreate {
	mutation := newPageSourceMutation(c.config, OpCreate)
	return &PageSourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PageSource entities.
func (c *PageSourceClient) CreateBulk(builders ...*PageSourceCreate) *PageSourceCreateBulk {
	return &PageSourceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PageSource.
func (c *PageSourceClient) Update() *PageSourceUpdate {
	mutation := newPageSourceMutation(c.config, OpUpdate)
	return &PageSourceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PageSourceClient) UpdateOne(ps *PageSource) *PageSourceUpdateOne {
	mutation := newPageSourceMutation(c.config, OpUpdateOne, withPageSource(ps))
	return &PageSourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PageSourceClient) UpdateOneID(id uuid.UUID) *PageSourceUpdateOne {
	mutation := newPageSourceMutation(c.config, OpUpdateOne, withPageSourceID(id))
	return &PageSourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PageSource.
func (c *PageSourceClient) Delete() *PageSourceDelete {
	mutation := newPageSourceMutation(c.config, OpDelete)
	return &PageSourceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PageSourceClient) DeleteOne(ps *PageSource) *PageSourceDeleteOne {
	return c.DeleteOneID(ps.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PageSourceClient) DeleteOneID(id uuid.UUID) *PageSourceDeleteOne {
	builder := c.Delete().Where(pagesource.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PageSourceDeleteOne{builder}
}

// Query returns a query builder for PageSource.
func (c *PageSourceClient) Query() *PageSourceQuery {
	return &PageSourceQuery{
		config: c.config,
	}
}

// Get returns a PageSource entity by its id.
func (c *PageSourceClient) Get(ctx context.Context, id uuid.UUID) (*PageSource, error) {
	return c.Query().Where(pagesource.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PageSourceClient) GetX(ctx context.Context, id uuid.UUID) *PageSource {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PageSourceClient) Hooks() []Hook {
	return c.hooks.PageSource
}
