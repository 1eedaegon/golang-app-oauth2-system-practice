// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent/image"
	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent/predicate"
	"github.com/1eedaegon/golang-app-oauth2-system-practice/db/ent/tenant"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeImage  = "Image"
	TypeTenant = "Tenant"
)

// ImageMutation represents an operation that mutates the Image nodes in the graph.
type ImageMutation struct {
	config
	op            Op
	typ           string
	id            *int
	image_id      *uuid.UUID
	name          *string
	created_at    *time.Time
	updated_at    *time.Time
	tenant_id     *uuid.UUID
	clearedFields map[string]struct{}
	tenant        *int
	clearedtenant bool
	done          bool
	oldValue      func(context.Context) (*Image, error)
	predicates    []predicate.Image
}

var _ ent.Mutation = (*ImageMutation)(nil)

// imageOption allows management of the mutation configuration using functional options.
type imageOption func(*ImageMutation)

// newImageMutation creates new mutation for the Image entity.
func newImageMutation(c config, op Op, opts ...imageOption) *ImageMutation {
	m := &ImageMutation{
		config:        c,
		op:            op,
		typ:           TypeImage,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withImageID sets the ID field of the mutation.
func withImageID(id int) imageOption {
	return func(m *ImageMutation) {
		var (
			err   error
			once  sync.Once
			value *Image
		)
		m.oldValue = func(ctx context.Context) (*Image, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Image.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withImage sets the old Image of the mutation.
func withImage(node *Image) imageOption {
	return func(m *ImageMutation) {
		m.oldValue = func(context.Context) (*Image, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ImageMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ImageMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ImageMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ImageMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Image.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetImageID sets the "image_id" field.
func (m *ImageMutation) SetImageID(u uuid.UUID) {
	m.image_id = &u
}

// ImageID returns the value of the "image_id" field in the mutation.
func (m *ImageMutation) ImageID() (r uuid.UUID, exists bool) {
	v := m.image_id
	if v == nil {
		return
	}
	return *v, true
}

// OldImageID returns the old "image_id" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldImageID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldImageID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldImageID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldImageID: %w", err)
	}
	return oldValue.ImageID, nil
}

// ResetImageID resets all changes to the "image_id" field.
func (m *ImageMutation) ResetImageID() {
	m.image_id = nil
}

// SetName sets the "name" field.
func (m *ImageMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *ImageMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *ImageMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ImageMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ImageMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ImageMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ImageMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ImageMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ImageMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetTenantID sets the "tenant_id" field.
func (m *ImageMutation) SetTenantID(u uuid.UUID) {
	m.tenant_id = &u
}

// TenantID returns the value of the "tenant_id" field in the mutation.
func (m *ImageMutation) TenantID() (r uuid.UUID, exists bool) {
	v := m.tenant_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTenantID returns the old "tenant_id" field's value of the Image entity.
// If the Image object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ImageMutation) OldTenantID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTenantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTenantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTenantID: %w", err)
	}
	return oldValue.TenantID, nil
}

// ClearTenantID clears the value of the "tenant_id" field.
func (m *ImageMutation) ClearTenantID() {
	m.tenant_id = nil
	m.clearedFields[image.FieldTenantID] = struct{}{}
}

// TenantIDCleared returns if the "tenant_id" field was cleared in this mutation.
func (m *ImageMutation) TenantIDCleared() bool {
	_, ok := m.clearedFields[image.FieldTenantID]
	return ok
}

// ResetTenantID resets all changes to the "tenant_id" field.
func (m *ImageMutation) ResetTenantID() {
	m.tenant_id = nil
	delete(m.clearedFields, image.FieldTenantID)
}

// SetTenantID sets the "tenant" edge to the Tenant entity by id.
func (m *ImageMutation) SetTenantID(id int) {
	m.tenant = &id
}

// ClearTenant clears the "tenant" edge to the Tenant entity.
func (m *ImageMutation) ClearTenant() {
	m.clearedtenant = true
}

// TenantCleared reports if the "tenant" edge to the Tenant entity was cleared.
func (m *ImageMutation) TenantCleared() bool {
	return m.clearedtenant
}

// TenantID returns the "tenant" edge ID in the mutation.
func (m *ImageMutation) TenantID() (id int, exists bool) {
	if m.tenant != nil {
		return *m.tenant, true
	}
	return
}

// TenantIDs returns the "tenant" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// TenantID instead. It exists only for internal usage by the builders.
func (m *ImageMutation) TenantIDs() (ids []int) {
	if id := m.tenant; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetTenant resets all changes to the "tenant" edge.
func (m *ImageMutation) ResetTenant() {
	m.tenant = nil
	m.clearedtenant = false
}

// Where appends a list predicates to the ImageMutation builder.
func (m *ImageMutation) Where(ps ...predicate.Image) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ImageMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ImageMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Image, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ImageMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ImageMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Image).
func (m *ImageMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ImageMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.image_id != nil {
		fields = append(fields, image.FieldImageID)
	}
	if m.name != nil {
		fields = append(fields, image.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, image.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, image.FieldUpdatedAt)
	}
	if m.tenant_id != nil {
		fields = append(fields, image.FieldTenantID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ImageMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case image.FieldImageID:
		return m.ImageID()
	case image.FieldName:
		return m.Name()
	case image.FieldCreatedAt:
		return m.CreatedAt()
	case image.FieldUpdatedAt:
		return m.UpdatedAt()
	case image.FieldTenantID:
		return m.TenantID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ImageMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case image.FieldImageID:
		return m.OldImageID(ctx)
	case image.FieldName:
		return m.OldName(ctx)
	case image.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case image.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case image.FieldTenantID:
		return m.OldTenantID(ctx)
	}
	return nil, fmt.Errorf("unknown Image field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ImageMutation) SetField(name string, value ent.Value) error {
	switch name {
	case image.FieldImageID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetImageID(v)
		return nil
	case image.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case image.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case image.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case image.FieldTenantID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTenantID(v)
		return nil
	}
	return fmt.Errorf("unknown Image field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ImageMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ImageMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ImageMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Image numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ImageMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(image.FieldTenantID) {
		fields = append(fields, image.FieldTenantID)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ImageMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ImageMutation) ClearField(name string) error {
	switch name {
	case image.FieldTenantID:
		m.ClearTenantID()
		return nil
	}
	return fmt.Errorf("unknown Image nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ImageMutation) ResetField(name string) error {
	switch name {
	case image.FieldImageID:
		m.ResetImageID()
		return nil
	case image.FieldName:
		m.ResetName()
		return nil
	case image.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case image.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case image.FieldTenantID:
		m.ResetTenantID()
		return nil
	}
	return fmt.Errorf("unknown Image field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ImageMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.tenant != nil {
		edges = append(edges, image.EdgeTenant)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ImageMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case image.EdgeTenant:
		if id := m.tenant; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ImageMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ImageMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ImageMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtenant {
		edges = append(edges, image.EdgeTenant)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ImageMutation) EdgeCleared(name string) bool {
	switch name {
	case image.EdgeTenant:
		return m.clearedtenant
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ImageMutation) ClearEdge(name string) error {
	switch name {
	case image.EdgeTenant:
		m.ClearTenant()
		return nil
	}
	return fmt.Errorf("unknown Image unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ImageMutation) ResetEdge(name string) error {
	switch name {
	case image.EdgeTenant:
		m.ResetTenant()
		return nil
	}
	return fmt.Errorf("unknown Image edge %s", name)
}

// TenantMutation represents an operation that mutates the Tenant nodes in the graph.
type TenantMutation struct {
	config
	op              Op
	typ             string
	id              *int
	tenant_id       *uuid.UUID
	name            *string
	created_at      *time.Time
	updated_at      *time.Time
	clearedFields   map[string]struct{}
	children        *int
	clearedchildren bool
	parent          map[int]struct{}
	removedparent   map[int]struct{}
	clearedparent   bool
	image           map[int]struct{}
	removedimage    map[int]struct{}
	clearedimage    bool
	done            bool
	oldValue        func(context.Context) (*Tenant, error)
	predicates      []predicate.Tenant
}

var _ ent.Mutation = (*TenantMutation)(nil)

// tenantOption allows management of the mutation configuration using functional options.
type tenantOption func(*TenantMutation)

// newTenantMutation creates new mutation for the Tenant entity.
func newTenantMutation(c config, op Op, opts ...tenantOption) *TenantMutation {
	m := &TenantMutation{
		config:        c,
		op:            op,
		typ:           TypeTenant,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTenantID sets the ID field of the mutation.
func withTenantID(id int) tenantOption {
	return func(m *TenantMutation) {
		var (
			err   error
			once  sync.Once
			value *Tenant
		)
		m.oldValue = func(ctx context.Context) (*Tenant, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Tenant.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTenant sets the old Tenant of the mutation.
func withTenant(node *Tenant) tenantOption {
	return func(m *TenantMutation) {
		m.oldValue = func(context.Context) (*Tenant, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TenantMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TenantMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TenantMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TenantMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Tenant.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTenantID sets the "tenant_id" field.
func (m *TenantMutation) SetTenantID(u uuid.UUID) {
	m.tenant_id = &u
}

// TenantID returns the value of the "tenant_id" field in the mutation.
func (m *TenantMutation) TenantID() (r uuid.UUID, exists bool) {
	v := m.tenant_id
	if v == nil {
		return
	}
	return *v, true
}

// OldTenantID returns the old "tenant_id" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldTenantID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTenantID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTenantID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTenantID: %w", err)
	}
	return oldValue.TenantID, nil
}

// ResetTenantID resets all changes to the "tenant_id" field.
func (m *TenantMutation) ResetTenantID() {
	m.tenant_id = nil
}

// SetName sets the "name" field.
func (m *TenantMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *TenantMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *TenantMutation) ResetName() {
	m.name = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *TenantMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TenantMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TenantMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TenantMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TenantMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Tenant entity.
// If the Tenant object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TenantMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TenantMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetChildrenID sets the "children" edge to the Tenant entity by id.
func (m *TenantMutation) SetChildrenID(id int) {
	m.children = &id
}

// ClearChildren clears the "children" edge to the Tenant entity.
func (m *TenantMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the Tenant entity was cleared.
func (m *TenantMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// ChildrenID returns the "children" edge ID in the mutation.
func (m *TenantMutation) ChildrenID() (id int, exists bool) {
	if m.children != nil {
		return *m.children, true
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ChildrenID instead. It exists only for internal usage by the builders.
func (m *TenantMutation) ChildrenIDs() (ids []int) {
	if id := m.children; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *TenantMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
}

// AddParentIDs adds the "parent" edge to the Tenant entity by ids.
func (m *TenantMutation) AddParentIDs(ids ...int) {
	if m.parent == nil {
		m.parent = make(map[int]struct{})
	}
	for i := range ids {
		m.parent[ids[i]] = struct{}{}
	}
}

// ClearParent clears the "parent" edge to the Tenant entity.
func (m *TenantMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared reports if the "parent" edge to the Tenant entity was cleared.
func (m *TenantMutation) ParentCleared() bool {
	return m.clearedparent
}

// RemoveParentIDs removes the "parent" edge to the Tenant entity by IDs.
func (m *TenantMutation) RemoveParentIDs(ids ...int) {
	if m.removedparent == nil {
		m.removedparent = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.parent, ids[i])
		m.removedparent[ids[i]] = struct{}{}
	}
}

// RemovedParent returns the removed IDs of the "parent" edge to the Tenant entity.
func (m *TenantMutation) RemovedParentIDs() (ids []int) {
	for id := range m.removedparent {
		ids = append(ids, id)
	}
	return
}

// ParentIDs returns the "parent" edge IDs in the mutation.
func (m *TenantMutation) ParentIDs() (ids []int) {
	for id := range m.parent {
		ids = append(ids, id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *TenantMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
	m.removedparent = nil
}

// AddImageIDs adds the "image" edge to the Image entity by ids.
func (m *TenantMutation) AddImageIDs(ids ...int) {
	if m.image == nil {
		m.image = make(map[int]struct{})
	}
	for i := range ids {
		m.image[ids[i]] = struct{}{}
	}
}

// ClearImage clears the "image" edge to the Image entity.
func (m *TenantMutation) ClearImage() {
	m.clearedimage = true
}

// ImageCleared reports if the "image" edge to the Image entity was cleared.
func (m *TenantMutation) ImageCleared() bool {
	return m.clearedimage
}

// RemoveImageIDs removes the "image" edge to the Image entity by IDs.
func (m *TenantMutation) RemoveImageIDs(ids ...int) {
	if m.removedimage == nil {
		m.removedimage = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.image, ids[i])
		m.removedimage[ids[i]] = struct{}{}
	}
}

// RemovedImage returns the removed IDs of the "image" edge to the Image entity.
func (m *TenantMutation) RemovedImageIDs() (ids []int) {
	for id := range m.removedimage {
		ids = append(ids, id)
	}
	return
}

// ImageIDs returns the "image" edge IDs in the mutation.
func (m *TenantMutation) ImageIDs() (ids []int) {
	for id := range m.image {
		ids = append(ids, id)
	}
	return
}

// ResetImage resets all changes to the "image" edge.
func (m *TenantMutation) ResetImage() {
	m.image = nil
	m.clearedimage = false
	m.removedimage = nil
}

// Where appends a list predicates to the TenantMutation builder.
func (m *TenantMutation) Where(ps ...predicate.Tenant) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TenantMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TenantMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Tenant, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TenantMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TenantMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Tenant).
func (m *TenantMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TenantMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.tenant_id != nil {
		fields = append(fields, tenant.FieldTenantID)
	}
	if m.name != nil {
		fields = append(fields, tenant.FieldName)
	}
	if m.created_at != nil {
		fields = append(fields, tenant.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, tenant.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TenantMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tenant.FieldTenantID:
		return m.TenantID()
	case tenant.FieldName:
		return m.Name()
	case tenant.FieldCreatedAt:
		return m.CreatedAt()
	case tenant.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TenantMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tenant.FieldTenantID:
		return m.OldTenantID(ctx)
	case tenant.FieldName:
		return m.OldName(ctx)
	case tenant.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case tenant.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Tenant field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TenantMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tenant.FieldTenantID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTenantID(v)
		return nil
	case tenant.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case tenant.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case tenant.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Tenant field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TenantMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TenantMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TenantMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Tenant numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TenantMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TenantMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TenantMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Tenant nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TenantMutation) ResetField(name string) error {
	switch name {
	case tenant.FieldTenantID:
		m.ResetTenantID()
		return nil
	case tenant.FieldName:
		m.ResetName()
		return nil
	case tenant.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case tenant.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Tenant field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TenantMutation) AddedEdges() []string {
	edges := make([]string, 0, 3)
	if m.children != nil {
		edges = append(edges, tenant.EdgeChildren)
	}
	if m.parent != nil {
		edges = append(edges, tenant.EdgeParent)
	}
	if m.image != nil {
		edges = append(edges, tenant.EdgeImage)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TenantMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case tenant.EdgeChildren:
		if id := m.children; id != nil {
			return []ent.Value{*id}
		}
	case tenant.EdgeParent:
		ids := make([]ent.Value, 0, len(m.parent))
		for id := range m.parent {
			ids = append(ids, id)
		}
		return ids
	case tenant.EdgeImage:
		ids := make([]ent.Value, 0, len(m.image))
		for id := range m.image {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TenantMutation) RemovedEdges() []string {
	edges := make([]string, 0, 3)
	if m.removedparent != nil {
		edges = append(edges, tenant.EdgeParent)
	}
	if m.removedimage != nil {
		edges = append(edges, tenant.EdgeImage)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TenantMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case tenant.EdgeParent:
		ids := make([]ent.Value, 0, len(m.removedparent))
		for id := range m.removedparent {
			ids = append(ids, id)
		}
		return ids
	case tenant.EdgeImage:
		ids := make([]ent.Value, 0, len(m.removedimage))
		for id := range m.removedimage {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TenantMutation) ClearedEdges() []string {
	edges := make([]string, 0, 3)
	if m.clearedchildren {
		edges = append(edges, tenant.EdgeChildren)
	}
	if m.clearedparent {
		edges = append(edges, tenant.EdgeParent)
	}
	if m.clearedimage {
		edges = append(edges, tenant.EdgeImage)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TenantMutation) EdgeCleared(name string) bool {
	switch name {
	case tenant.EdgeChildren:
		return m.clearedchildren
	case tenant.EdgeParent:
		return m.clearedparent
	case tenant.EdgeImage:
		return m.clearedimage
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TenantMutation) ClearEdge(name string) error {
	switch name {
	case tenant.EdgeChildren:
		m.ClearChildren()
		return nil
	}
	return fmt.Errorf("unknown Tenant unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TenantMutation) ResetEdge(name string) error {
	switch name {
	case tenant.EdgeChildren:
		m.ResetChildren()
		return nil
	case tenant.EdgeParent:
		m.ResetParent()
		return nil
	case tenant.EdgeImage:
		m.ResetImage()
		return nil
	}
	return fmt.Errorf("unknown Tenant edge %s", name)
}
