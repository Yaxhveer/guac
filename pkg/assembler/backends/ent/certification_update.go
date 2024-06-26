// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/artifact"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/certification"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/predicate"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/sourcename"
)

// CertificationUpdate is the builder for updating Certification entities.
type CertificationUpdate struct {
	config
	hooks    []Hook
	mutation *CertificationMutation
}

// Where appends a list predicates to the CertificationUpdate builder.
func (cu *CertificationUpdate) Where(ps ...predicate.Certification) *CertificationUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetSourceID sets the "source_id" field.
func (cu *CertificationUpdate) SetSourceID(u uuid.UUID) *CertificationUpdate {
	cu.mutation.SetSourceID(u)
	return cu
}

// SetNillableSourceID sets the "source_id" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableSourceID(u *uuid.UUID) *CertificationUpdate {
	if u != nil {
		cu.SetSourceID(*u)
	}
	return cu
}

// ClearSourceID clears the value of the "source_id" field.
func (cu *CertificationUpdate) ClearSourceID() *CertificationUpdate {
	cu.mutation.ClearSourceID()
	return cu
}

// SetPackageVersionID sets the "package_version_id" field.
func (cu *CertificationUpdate) SetPackageVersionID(u uuid.UUID) *CertificationUpdate {
	cu.mutation.SetPackageVersionID(u)
	return cu
}

// SetNillablePackageVersionID sets the "package_version_id" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillablePackageVersionID(u *uuid.UUID) *CertificationUpdate {
	if u != nil {
		cu.SetPackageVersionID(*u)
	}
	return cu
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (cu *CertificationUpdate) ClearPackageVersionID() *CertificationUpdate {
	cu.mutation.ClearPackageVersionID()
	return cu
}

// SetPackageNameID sets the "package_name_id" field.
func (cu *CertificationUpdate) SetPackageNameID(u uuid.UUID) *CertificationUpdate {
	cu.mutation.SetPackageNameID(u)
	return cu
}

// SetNillablePackageNameID sets the "package_name_id" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillablePackageNameID(u *uuid.UUID) *CertificationUpdate {
	if u != nil {
		cu.SetPackageNameID(*u)
	}
	return cu
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (cu *CertificationUpdate) ClearPackageNameID() *CertificationUpdate {
	cu.mutation.ClearPackageNameID()
	return cu
}

// SetArtifactID sets the "artifact_id" field.
func (cu *CertificationUpdate) SetArtifactID(u uuid.UUID) *CertificationUpdate {
	cu.mutation.SetArtifactID(u)
	return cu
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableArtifactID(u *uuid.UUID) *CertificationUpdate {
	if u != nil {
		cu.SetArtifactID(*u)
	}
	return cu
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (cu *CertificationUpdate) ClearArtifactID() *CertificationUpdate {
	cu.mutation.ClearArtifactID()
	return cu
}

// SetType sets the "type" field.
func (cu *CertificationUpdate) SetType(c certification.Type) *CertificationUpdate {
	cu.mutation.SetType(c)
	return cu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableType(c *certification.Type) *CertificationUpdate {
	if c != nil {
		cu.SetType(*c)
	}
	return cu
}

// SetJustification sets the "justification" field.
func (cu *CertificationUpdate) SetJustification(s string) *CertificationUpdate {
	cu.mutation.SetJustification(s)
	return cu
}

// SetNillableJustification sets the "justification" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableJustification(s *string) *CertificationUpdate {
	if s != nil {
		cu.SetJustification(*s)
	}
	return cu
}

// SetKnownSince sets the "known_since" field.
func (cu *CertificationUpdate) SetKnownSince(t time.Time) *CertificationUpdate {
	cu.mutation.SetKnownSince(t)
	return cu
}

// SetNillableKnownSince sets the "known_since" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableKnownSince(t *time.Time) *CertificationUpdate {
	if t != nil {
		cu.SetKnownSince(*t)
	}
	return cu
}

// SetOrigin sets the "origin" field.
func (cu *CertificationUpdate) SetOrigin(s string) *CertificationUpdate {
	cu.mutation.SetOrigin(s)
	return cu
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableOrigin(s *string) *CertificationUpdate {
	if s != nil {
		cu.SetOrigin(*s)
	}
	return cu
}

// SetCollector sets the "collector" field.
func (cu *CertificationUpdate) SetCollector(s string) *CertificationUpdate {
	cu.mutation.SetCollector(s)
	return cu
}

// SetNillableCollector sets the "collector" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableCollector(s *string) *CertificationUpdate {
	if s != nil {
		cu.SetCollector(*s)
	}
	return cu
}

// SetDocumentRef sets the "document_ref" field.
func (cu *CertificationUpdate) SetDocumentRef(s string) *CertificationUpdate {
	cu.mutation.SetDocumentRef(s)
	return cu
}

// SetNillableDocumentRef sets the "document_ref" field if the given value is not nil.
func (cu *CertificationUpdate) SetNillableDocumentRef(s *string) *CertificationUpdate {
	if s != nil {
		cu.SetDocumentRef(*s)
	}
	return cu
}

// SetSource sets the "source" edge to the SourceName entity.
func (cu *CertificationUpdate) SetSource(s *SourceName) *CertificationUpdate {
	return cu.SetSourceID(s.ID)
}

// SetPackageVersion sets the "package_version" edge to the PackageVersion entity.
func (cu *CertificationUpdate) SetPackageVersion(p *PackageVersion) *CertificationUpdate {
	return cu.SetPackageVersionID(p.ID)
}

// SetAllVersionsID sets the "all_versions" edge to the PackageName entity by ID.
func (cu *CertificationUpdate) SetAllVersionsID(id uuid.UUID) *CertificationUpdate {
	cu.mutation.SetAllVersionsID(id)
	return cu
}

// SetNillableAllVersionsID sets the "all_versions" edge to the PackageName entity by ID if the given value is not nil.
func (cu *CertificationUpdate) SetNillableAllVersionsID(id *uuid.UUID) *CertificationUpdate {
	if id != nil {
		cu = cu.SetAllVersionsID(*id)
	}
	return cu
}

// SetAllVersions sets the "all_versions" edge to the PackageName entity.
func (cu *CertificationUpdate) SetAllVersions(p *PackageName) *CertificationUpdate {
	return cu.SetAllVersionsID(p.ID)
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (cu *CertificationUpdate) SetArtifact(a *Artifact) *CertificationUpdate {
	return cu.SetArtifactID(a.ID)
}

// Mutation returns the CertificationMutation object of the builder.
func (cu *CertificationUpdate) Mutation() *CertificationMutation {
	return cu.mutation
}

// ClearSource clears the "source" edge to the SourceName entity.
func (cu *CertificationUpdate) ClearSource() *CertificationUpdate {
	cu.mutation.ClearSource()
	return cu
}

// ClearPackageVersion clears the "package_version" edge to the PackageVersion entity.
func (cu *CertificationUpdate) ClearPackageVersion() *CertificationUpdate {
	cu.mutation.ClearPackageVersion()
	return cu
}

// ClearAllVersions clears the "all_versions" edge to the PackageName entity.
func (cu *CertificationUpdate) ClearAllVersions() *CertificationUpdate {
	cu.mutation.ClearAllVersions()
	return cu
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (cu *CertificationUpdate) ClearArtifact() *CertificationUpdate {
	cu.mutation.ClearArtifact()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CertificationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CertificationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CertificationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CertificationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CertificationUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := certification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Certification.type": %w`, err)}
		}
	}
	return nil
}

func (cu *CertificationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(certification.Table, certification.Columns, sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(certification.FieldType, field.TypeEnum, value)
	}
	if value, ok := cu.mutation.Justification(); ok {
		_spec.SetField(certification.FieldJustification, field.TypeString, value)
	}
	if value, ok := cu.mutation.KnownSince(); ok {
		_spec.SetField(certification.FieldKnownSince, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Origin(); ok {
		_spec.SetField(certification.FieldOrigin, field.TypeString, value)
	}
	if value, ok := cu.mutation.Collector(); ok {
		_spec.SetField(certification.FieldCollector, field.TypeString, value)
	}
	if value, ok := cu.mutation.DocumentRef(); ok {
		_spec.SetField(certification.FieldDocumentRef, field.TypeString, value)
	}
	if cu.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.SourceTable,
			Columns: []string{certification.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.SourceTable,
			Columns: []string{certification.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.PackageVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.PackageVersionTable,
			Columns: []string{certification.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PackageVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.PackageVersionTable,
			Columns: []string{certification.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AllVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.AllVersionsTable,
			Columns: []string{certification.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AllVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.AllVersionsTable,
			Columns: []string{certification.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.ArtifactTable,
			Columns: []string{certification.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.ArtifactTable,
			Columns: []string{certification.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CertificationUpdateOne is the builder for updating a single Certification entity.
type CertificationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CertificationMutation
}

// SetSourceID sets the "source_id" field.
func (cuo *CertificationUpdateOne) SetSourceID(u uuid.UUID) *CertificationUpdateOne {
	cuo.mutation.SetSourceID(u)
	return cuo
}

// SetNillableSourceID sets the "source_id" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableSourceID(u *uuid.UUID) *CertificationUpdateOne {
	if u != nil {
		cuo.SetSourceID(*u)
	}
	return cuo
}

// ClearSourceID clears the value of the "source_id" field.
func (cuo *CertificationUpdateOne) ClearSourceID() *CertificationUpdateOne {
	cuo.mutation.ClearSourceID()
	return cuo
}

// SetPackageVersionID sets the "package_version_id" field.
func (cuo *CertificationUpdateOne) SetPackageVersionID(u uuid.UUID) *CertificationUpdateOne {
	cuo.mutation.SetPackageVersionID(u)
	return cuo
}

// SetNillablePackageVersionID sets the "package_version_id" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillablePackageVersionID(u *uuid.UUID) *CertificationUpdateOne {
	if u != nil {
		cuo.SetPackageVersionID(*u)
	}
	return cuo
}

// ClearPackageVersionID clears the value of the "package_version_id" field.
func (cuo *CertificationUpdateOne) ClearPackageVersionID() *CertificationUpdateOne {
	cuo.mutation.ClearPackageVersionID()
	return cuo
}

// SetPackageNameID sets the "package_name_id" field.
func (cuo *CertificationUpdateOne) SetPackageNameID(u uuid.UUID) *CertificationUpdateOne {
	cuo.mutation.SetPackageNameID(u)
	return cuo
}

// SetNillablePackageNameID sets the "package_name_id" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillablePackageNameID(u *uuid.UUID) *CertificationUpdateOne {
	if u != nil {
		cuo.SetPackageNameID(*u)
	}
	return cuo
}

// ClearPackageNameID clears the value of the "package_name_id" field.
func (cuo *CertificationUpdateOne) ClearPackageNameID() *CertificationUpdateOne {
	cuo.mutation.ClearPackageNameID()
	return cuo
}

// SetArtifactID sets the "artifact_id" field.
func (cuo *CertificationUpdateOne) SetArtifactID(u uuid.UUID) *CertificationUpdateOne {
	cuo.mutation.SetArtifactID(u)
	return cuo
}

// SetNillableArtifactID sets the "artifact_id" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableArtifactID(u *uuid.UUID) *CertificationUpdateOne {
	if u != nil {
		cuo.SetArtifactID(*u)
	}
	return cuo
}

// ClearArtifactID clears the value of the "artifact_id" field.
func (cuo *CertificationUpdateOne) ClearArtifactID() *CertificationUpdateOne {
	cuo.mutation.ClearArtifactID()
	return cuo
}

// SetType sets the "type" field.
func (cuo *CertificationUpdateOne) SetType(c certification.Type) *CertificationUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableType(c *certification.Type) *CertificationUpdateOne {
	if c != nil {
		cuo.SetType(*c)
	}
	return cuo
}

// SetJustification sets the "justification" field.
func (cuo *CertificationUpdateOne) SetJustification(s string) *CertificationUpdateOne {
	cuo.mutation.SetJustification(s)
	return cuo
}

// SetNillableJustification sets the "justification" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableJustification(s *string) *CertificationUpdateOne {
	if s != nil {
		cuo.SetJustification(*s)
	}
	return cuo
}

// SetKnownSince sets the "known_since" field.
func (cuo *CertificationUpdateOne) SetKnownSince(t time.Time) *CertificationUpdateOne {
	cuo.mutation.SetKnownSince(t)
	return cuo
}

// SetNillableKnownSince sets the "known_since" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableKnownSince(t *time.Time) *CertificationUpdateOne {
	if t != nil {
		cuo.SetKnownSince(*t)
	}
	return cuo
}

// SetOrigin sets the "origin" field.
func (cuo *CertificationUpdateOne) SetOrigin(s string) *CertificationUpdateOne {
	cuo.mutation.SetOrigin(s)
	return cuo
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableOrigin(s *string) *CertificationUpdateOne {
	if s != nil {
		cuo.SetOrigin(*s)
	}
	return cuo
}

// SetCollector sets the "collector" field.
func (cuo *CertificationUpdateOne) SetCollector(s string) *CertificationUpdateOne {
	cuo.mutation.SetCollector(s)
	return cuo
}

// SetNillableCollector sets the "collector" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableCollector(s *string) *CertificationUpdateOne {
	if s != nil {
		cuo.SetCollector(*s)
	}
	return cuo
}

// SetDocumentRef sets the "document_ref" field.
func (cuo *CertificationUpdateOne) SetDocumentRef(s string) *CertificationUpdateOne {
	cuo.mutation.SetDocumentRef(s)
	return cuo
}

// SetNillableDocumentRef sets the "document_ref" field if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableDocumentRef(s *string) *CertificationUpdateOne {
	if s != nil {
		cuo.SetDocumentRef(*s)
	}
	return cuo
}

// SetSource sets the "source" edge to the SourceName entity.
func (cuo *CertificationUpdateOne) SetSource(s *SourceName) *CertificationUpdateOne {
	return cuo.SetSourceID(s.ID)
}

// SetPackageVersion sets the "package_version" edge to the PackageVersion entity.
func (cuo *CertificationUpdateOne) SetPackageVersion(p *PackageVersion) *CertificationUpdateOne {
	return cuo.SetPackageVersionID(p.ID)
}

// SetAllVersionsID sets the "all_versions" edge to the PackageName entity by ID.
func (cuo *CertificationUpdateOne) SetAllVersionsID(id uuid.UUID) *CertificationUpdateOne {
	cuo.mutation.SetAllVersionsID(id)
	return cuo
}

// SetNillableAllVersionsID sets the "all_versions" edge to the PackageName entity by ID if the given value is not nil.
func (cuo *CertificationUpdateOne) SetNillableAllVersionsID(id *uuid.UUID) *CertificationUpdateOne {
	if id != nil {
		cuo = cuo.SetAllVersionsID(*id)
	}
	return cuo
}

// SetAllVersions sets the "all_versions" edge to the PackageName entity.
func (cuo *CertificationUpdateOne) SetAllVersions(p *PackageName) *CertificationUpdateOne {
	return cuo.SetAllVersionsID(p.ID)
}

// SetArtifact sets the "artifact" edge to the Artifact entity.
func (cuo *CertificationUpdateOne) SetArtifact(a *Artifact) *CertificationUpdateOne {
	return cuo.SetArtifactID(a.ID)
}

// Mutation returns the CertificationMutation object of the builder.
func (cuo *CertificationUpdateOne) Mutation() *CertificationMutation {
	return cuo.mutation
}

// ClearSource clears the "source" edge to the SourceName entity.
func (cuo *CertificationUpdateOne) ClearSource() *CertificationUpdateOne {
	cuo.mutation.ClearSource()
	return cuo
}

// ClearPackageVersion clears the "package_version" edge to the PackageVersion entity.
func (cuo *CertificationUpdateOne) ClearPackageVersion() *CertificationUpdateOne {
	cuo.mutation.ClearPackageVersion()
	return cuo
}

// ClearAllVersions clears the "all_versions" edge to the PackageName entity.
func (cuo *CertificationUpdateOne) ClearAllVersions() *CertificationUpdateOne {
	cuo.mutation.ClearAllVersions()
	return cuo
}

// ClearArtifact clears the "artifact" edge to the Artifact entity.
func (cuo *CertificationUpdateOne) ClearArtifact() *CertificationUpdateOne {
	cuo.mutation.ClearArtifact()
	return cuo
}

// Where appends a list predicates to the CertificationUpdate builder.
func (cuo *CertificationUpdateOne) Where(ps ...predicate.Certification) *CertificationUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CertificationUpdateOne) Select(field string, fields ...string) *CertificationUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Certification entity.
func (cuo *CertificationUpdateOne) Save(ctx context.Context) (*Certification, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CertificationUpdateOne) SaveX(ctx context.Context) *Certification {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CertificationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CertificationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CertificationUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := certification.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Certification.type": %w`, err)}
		}
	}
	return nil
}

func (cuo *CertificationUpdateOne) sqlSave(ctx context.Context) (_node *Certification, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(certification.Table, certification.Columns, sqlgraph.NewFieldSpec(certification.FieldID, field.TypeUUID))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Certification.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, certification.FieldID)
		for _, f := range fields {
			if !certification.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != certification.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(certification.FieldType, field.TypeEnum, value)
	}
	if value, ok := cuo.mutation.Justification(); ok {
		_spec.SetField(certification.FieldJustification, field.TypeString, value)
	}
	if value, ok := cuo.mutation.KnownSince(); ok {
		_spec.SetField(certification.FieldKnownSince, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Origin(); ok {
		_spec.SetField(certification.FieldOrigin, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Collector(); ok {
		_spec.SetField(certification.FieldCollector, field.TypeString, value)
	}
	if value, ok := cuo.mutation.DocumentRef(); ok {
		_spec.SetField(certification.FieldDocumentRef, field.TypeString, value)
	}
	if cuo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.SourceTable,
			Columns: []string{certification.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.SourceTable,
			Columns: []string{certification.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcename.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.PackageVersionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.PackageVersionTable,
			Columns: []string{certification.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PackageVersionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.PackageVersionTable,
			Columns: []string{certification.PackageVersionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageversion.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AllVersionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.AllVersionsTable,
			Columns: []string{certification.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AllVersionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.AllVersionsTable,
			Columns: []string{certification.AllVersionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packagename.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ArtifactCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.ArtifactTable,
			Columns: []string{certification.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ArtifactIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   certification.ArtifactTable,
			Columns: []string{certification.ArtifactColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artifact.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Certification{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{certification.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
