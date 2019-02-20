// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDepartments(t *testing.T) {
	t.Parallel()

	query := Departments()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDepartmentsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDepartmentsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Departments().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDepartmentsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DepartmentSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDepartmentsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DepartmentExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Department exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DepartmentExists to return true, but got false.")
	}
}

func testDepartmentsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	departmentFound, err := FindDepartment(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if departmentFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDepartmentsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Departments().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDepartmentsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Departments().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDepartmentsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	departmentOne := &Department{}
	departmentTwo := &Department{}
	if err = randomize.Struct(seed, departmentOne, departmentDBTypes, false, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}
	if err = randomize.Struct(seed, departmentTwo, departmentDBTypes, false, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = departmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = departmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Departments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDepartmentsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	departmentOne := &Department{}
	departmentTwo := &Department{}
	if err = randomize.Struct(seed, departmentOne, departmentDBTypes, false, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}
	if err = randomize.Struct(seed, departmentTwo, departmentDBTypes, false, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = departmentOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = departmentTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func departmentBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func departmentAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Department) error {
	*o = Department{}
	return nil
}

func testDepartmentsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Department{}
	o := &Department{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, departmentDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Department object: %s", err)
	}

	AddDepartmentHook(boil.BeforeInsertHook, departmentBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	departmentBeforeInsertHooks = []DepartmentHook{}

	AddDepartmentHook(boil.AfterInsertHook, departmentAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	departmentAfterInsertHooks = []DepartmentHook{}

	AddDepartmentHook(boil.AfterSelectHook, departmentAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	departmentAfterSelectHooks = []DepartmentHook{}

	AddDepartmentHook(boil.BeforeUpdateHook, departmentBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	departmentBeforeUpdateHooks = []DepartmentHook{}

	AddDepartmentHook(boil.AfterUpdateHook, departmentAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	departmentAfterUpdateHooks = []DepartmentHook{}

	AddDepartmentHook(boil.BeforeDeleteHook, departmentBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	departmentBeforeDeleteHooks = []DepartmentHook{}

	AddDepartmentHook(boil.AfterDeleteHook, departmentAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	departmentAfterDeleteHooks = []DepartmentHook{}

	AddDepartmentHook(boil.BeforeUpsertHook, departmentBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	departmentBeforeUpsertHooks = []DepartmentHook{}

	AddDepartmentHook(boil.AfterUpsertHook, departmentAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	departmentAfterUpsertHooks = []DepartmentHook{}
}

func testDepartmentsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDepartmentsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(departmentColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDepartmentToManyProfiles(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Department
	var b, c Profile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, profileDBTypes, false, profileColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.DepartmentID, a.ID)
	queries.Assign(&c.DepartmentID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Profiles().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.DepartmentID, b.DepartmentID) {
			bFound = true
		}
		if queries.Equal(v.DepartmentID, c.DepartmentID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DepartmentSlice{&a}
	if err = a.L.LoadProfiles(ctx, tx, false, (*[]*Department)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Profiles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Profiles = nil
	if err = a.L.LoadProfiles(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Profiles); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDepartmentToManyAddOpProfiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Department
	var b, c, d, e Profile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, departmentDBTypes, false, strmangle.SetComplement(departmentPrimaryKeyColumns, departmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Profile{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Profile{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProfiles(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.DepartmentID) {
			t.Error("foreign key was wrong value", a.ID, first.DepartmentID)
		}
		if !queries.Equal(a.ID, second.DepartmentID) {
			t.Error("foreign key was wrong value", a.ID, second.DepartmentID)
		}

		if first.R.Department != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Department != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Profiles[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Profiles[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Profiles().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDepartmentToManySetOpProfiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Department
	var b, c, d, e Profile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, departmentDBTypes, false, strmangle.SetComplement(departmentPrimaryKeyColumns, departmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Profile{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetProfiles(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Profiles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetProfiles(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Profiles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DepartmentID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DepartmentID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.DepartmentID) {
		t.Error("foreign key was wrong value", a.ID, d.DepartmentID)
	}
	if !queries.Equal(a.ID, e.DepartmentID) {
		t.Error("foreign key was wrong value", a.ID, e.DepartmentID)
	}

	if b.R.Department != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Department != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Department != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Department != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Profiles[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Profiles[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDepartmentToManyRemoveOpProfiles(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Department
	var b, c, d, e Profile

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, departmentDBTypes, false, strmangle.SetComplement(departmentPrimaryKeyColumns, departmentColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Profile{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, profileDBTypes, false, strmangle.SetComplement(profilePrimaryKeyColumns, profileColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddProfiles(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Profiles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveProfiles(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Profiles().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.DepartmentID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.DepartmentID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Department != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Department != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Department != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Department != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Profiles) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Profiles[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Profiles[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDepartmentsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDepartmentsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DepartmentSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDepartmentsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Departments().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	departmentDBTypes = map[string]string{`ID`: `bigint`, `Name`: `character varying`, `ShortName`: `character varying`}
	_                 = bytes.MinRead
)

func testDepartmentsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(departmentPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(departmentColumns) == len(departmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDepartmentsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(departmentColumns) == len(departmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Department{}
	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, departmentDBTypes, true, departmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(departmentColumns, departmentPrimaryKeyColumns) {
		fields = departmentColumns
	} else {
		fields = strmangle.SetComplement(
			departmentColumns,
			departmentPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := DepartmentSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDepartmentsUpsert(t *testing.T) {
	t.Parallel()

	if len(departmentColumns) == len(departmentPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Department{}
	if err = randomize.Struct(seed, &o, departmentDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Department: %s", err)
	}

	count, err := Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, departmentDBTypes, false, departmentPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Department struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Department: %s", err)
	}

	count, err = Departments().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
