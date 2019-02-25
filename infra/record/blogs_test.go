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

func testBlogs(t *testing.T) {
	t.Parallel()

	query := Blogs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBlogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
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

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Blogs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlogSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BlogExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Blog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlogExists to return true, but got false.")
	}
}

func testBlogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	blogFound, err := FindBlog(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if blogFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBlogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Blogs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBlogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Blogs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blogOne := &Blog{}
	blogTwo := &Blog{}
	if err = randomize.Struct(seed, blogOne, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}
	if err = randomize.Struct(seed, blogTwo, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blogOne := &Blog{}
	blogTwo := &Blog{}
	if err = randomize.Struct(seed, blogOne, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}
	if err = randomize.Struct(seed, blogTwo, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = blogOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = blogTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func blogBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func blogAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Blog) error {
	*o = Blog{}
	return nil
}

func testBlogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Blog{}
	o := &Blog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, blogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Blog object: %s", err)
	}

	AddBlogHook(boil.BeforeInsertHook, blogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	blogBeforeInsertHooks = []BlogHook{}

	AddBlogHook(boil.AfterInsertHook, blogAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	blogAfterInsertHooks = []BlogHook{}

	AddBlogHook(boil.AfterSelectHook, blogAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	blogAfterSelectHooks = []BlogHook{}

	AddBlogHook(boil.BeforeUpdateHook, blogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	blogBeforeUpdateHooks = []BlogHook{}

	AddBlogHook(boil.AfterUpdateHook, blogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	blogAfterUpdateHooks = []BlogHook{}

	AddBlogHook(boil.BeforeDeleteHook, blogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	blogBeforeDeleteHooks = []BlogHook{}

	AddBlogHook(boil.AfterDeleteHook, blogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	blogAfterDeleteHooks = []BlogHook{}

	AddBlogHook(boil.BeforeUpsertHook, blogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	blogBeforeUpsertHooks = []BlogHook{}

	AddBlogHook(boil.AfterUpsertHook, blogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	blogAfterUpsertHooks = []BlogHook{}
}

func testBlogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(blogColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlogToManyEntries(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Blog
	var b, c Entry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, entryDBTypes, false, entryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, entryDBTypes, false, entryColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.BlogID = a.ID
	c.BlogID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Entries().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.BlogID == b.BlogID {
			bFound = true
		}
		if v.BlogID == c.BlogID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := BlogSlice{&a}
	if err = a.L.LoadEntries(ctx, tx, false, (*[]*Blog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Entries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Entries = nil
	if err = a.L.LoadEntries(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Entries); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testBlogToManyAddOpEntries(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Blog
	var b, c, d, e Entry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blogDBTypes, false, strmangle.SetComplement(blogPrimaryKeyColumns, blogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Entry{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, entryDBTypes, false, strmangle.SetComplement(entryPrimaryKeyColumns, entryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Entry{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEntries(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.BlogID {
			t.Error("foreign key was wrong value", a.ID, first.BlogID)
		}
		if a.ID != second.BlogID {
			t.Error("foreign key was wrong value", a.ID, second.BlogID)
		}

		if first.R.Blog != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Blog != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Entries[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Entries[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Entries().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testBlogToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Blog
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blogDBTypes, false, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BlogSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Blog)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlogToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Blog
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blogDBTypes, false, strmangle.SetComplement(blogPrimaryKeyColumns, blogColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Blogs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testBlogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
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

func testBlogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BlogSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBlogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Blogs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blogDBTypes = map[string]string{`ID`: `bigint`, `URL`: `character varying`, `FeedURL`: `character varying`, `UserID`: `bigint`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`}
	_           = bytes.MinRead
)

func testBlogsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(blogColumns) == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blogDBTypes, true, blogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBlogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blogColumns) == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Blog{}
	if err = randomize.Struct(seed, o, blogDBTypes, true, blogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, blogDBTypes, true, blogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blogColumns, blogPrimaryKeyColumns) {
		fields = blogColumns
	} else {
		fields = strmangle.SetComplement(
			blogColumns,
			blogPrimaryKeyColumns,
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

	slice := BlogSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBlogsUpsert(t *testing.T) {
	t.Parallel()

	if len(blogColumns) == len(blogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Blog{}
	if err = randomize.Struct(seed, &o, blogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Blog: %s", err)
	}

	count, err := Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, blogDBTypes, false, blogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Blog struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Blog: %s", err)
	}

	count, err = Blogs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
