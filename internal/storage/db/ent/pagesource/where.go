// Code generated by ent, DO NOT EDIT.

package pagesource

import (
	"entgo.io/ent/dialect/sql"
	"github.com/drakejin/place/internal/storage/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSource), v))
	})
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSource), v))
	})
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.PageSource {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSource), v...))
	})
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.PageSource {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSource), v...))
	})
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSource), v))
	})
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSource), v))
	})
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSource), v))
	})
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSource), v))
	})
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSource), v))
	})
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSource), v))
	})
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSource), v))
	})
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSource), v))
	})
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSource), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PageSource) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PageSource) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PageSource) predicate.PageSource {
	return predicate.PageSource(func(s *sql.Selector) {
		p(s.Not())
	})
}