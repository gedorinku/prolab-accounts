// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package record

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Blogs", testBlogs)
	t.Run("Departments", testDepartments)
	t.Run("Entries", testEntries)
	t.Run("Profiles", testProfiles)
	t.Run("Roles", testRoles)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Blogs", testBlogsDelete)
	t.Run("Departments", testDepartmentsDelete)
	t.Run("Entries", testEntriesDelete)
	t.Run("Profiles", testProfilesDelete)
	t.Run("Roles", testRolesDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Blogs", testBlogsQueryDeleteAll)
	t.Run("Departments", testDepartmentsQueryDeleteAll)
	t.Run("Entries", testEntriesQueryDeleteAll)
	t.Run("Profiles", testProfilesQueryDeleteAll)
	t.Run("Roles", testRolesQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Blogs", testBlogsSliceDeleteAll)
	t.Run("Departments", testDepartmentsSliceDeleteAll)
	t.Run("Entries", testEntriesSliceDeleteAll)
	t.Run("Profiles", testProfilesSliceDeleteAll)
	t.Run("Roles", testRolesSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Blogs", testBlogsExists)
	t.Run("Departments", testDepartmentsExists)
	t.Run("Entries", testEntriesExists)
	t.Run("Profiles", testProfilesExists)
	t.Run("Roles", testRolesExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Blogs", testBlogsFind)
	t.Run("Departments", testDepartmentsFind)
	t.Run("Entries", testEntriesFind)
	t.Run("Profiles", testProfilesFind)
	t.Run("Roles", testRolesFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Blogs", testBlogsBind)
	t.Run("Departments", testDepartmentsBind)
	t.Run("Entries", testEntriesBind)
	t.Run("Profiles", testProfilesBind)
	t.Run("Roles", testRolesBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Blogs", testBlogsOne)
	t.Run("Departments", testDepartmentsOne)
	t.Run("Entries", testEntriesOne)
	t.Run("Profiles", testProfilesOne)
	t.Run("Roles", testRolesOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Blogs", testBlogsAll)
	t.Run("Departments", testDepartmentsAll)
	t.Run("Entries", testEntriesAll)
	t.Run("Profiles", testProfilesAll)
	t.Run("Roles", testRolesAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Blogs", testBlogsCount)
	t.Run("Departments", testDepartmentsCount)
	t.Run("Entries", testEntriesCount)
	t.Run("Profiles", testProfilesCount)
	t.Run("Roles", testRolesCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Blogs", testBlogsHooks)
	t.Run("Departments", testDepartmentsHooks)
	t.Run("Entries", testEntriesHooks)
	t.Run("Profiles", testProfilesHooks)
	t.Run("Roles", testRolesHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Blogs", testBlogsInsert)
	t.Run("Blogs", testBlogsInsertWhitelist)
	t.Run("Departments", testDepartmentsInsert)
	t.Run("Departments", testDepartmentsInsertWhitelist)
	t.Run("Entries", testEntriesInsert)
	t.Run("Entries", testEntriesInsertWhitelist)
	t.Run("Profiles", testProfilesInsert)
	t.Run("Profiles", testProfilesInsertWhitelist)
	t.Run("Roles", testRolesInsert)
	t.Run("Roles", testRolesInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("BlogToUserUsingUser", testBlogToOneUserUsingUser)
	t.Run("EntryToBlogUsingBlog", testEntryToOneBlogUsingBlog)
	t.Run("EntryToUserUsingAuthor", testEntryToOneUserUsingAuthor)
	t.Run("ProfileToDepartmentUsingDepartment", testProfileToOneDepartmentUsingDepartment)
	t.Run("ProfileToRoleUsingRole", testProfileToOneRoleUsingRole)
	t.Run("UserToProfileUsingProfile", testUserToOneProfileUsingProfile)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("BlogToEntries", testBlogToManyEntries)
	t.Run("DepartmentToProfiles", testDepartmentToManyProfiles)
	t.Run("ProfileToUsers", testProfileToManyUsers)
	t.Run("RoleToProfiles", testRoleToManyProfiles)
	t.Run("UserToBlogs", testUserToManyBlogs)
	t.Run("UserToAuthorEntries", testUserToManyAuthorEntries)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("BlogToUserUsingBlogs", testBlogToOneSetOpUserUsingUser)
	t.Run("EntryToBlogUsingEntries", testEntryToOneSetOpBlogUsingBlog)
	t.Run("EntryToUserUsingAuthorEntries", testEntryToOneSetOpUserUsingAuthor)
	t.Run("ProfileToDepartmentUsingProfiles", testProfileToOneSetOpDepartmentUsingDepartment)
	t.Run("ProfileToRoleUsingProfiles", testProfileToOneSetOpRoleUsingRole)
	t.Run("UserToProfileUsingUsers", testUserToOneSetOpProfileUsingProfile)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ProfileToDepartmentUsingProfiles", testProfileToOneRemoveOpDepartmentUsingDepartment)
	t.Run("ProfileToRoleUsingProfiles", testProfileToOneRemoveOpRoleUsingRole)
	t.Run("UserToProfileUsingUsers", testUserToOneRemoveOpProfileUsingProfile)
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
	t.Run("BlogToEntries", testBlogToManyAddOpEntries)
	t.Run("DepartmentToProfiles", testDepartmentToManyAddOpProfiles)
	t.Run("ProfileToUsers", testProfileToManyAddOpUsers)
	t.Run("RoleToProfiles", testRoleToManyAddOpProfiles)
	t.Run("UserToBlogs", testUserToManyAddOpBlogs)
	t.Run("UserToAuthorEntries", testUserToManyAddOpAuthorEntries)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("DepartmentToProfiles", testDepartmentToManySetOpProfiles)
	t.Run("ProfileToUsers", testProfileToManySetOpUsers)
	t.Run("RoleToProfiles", testRoleToManySetOpProfiles)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("DepartmentToProfiles", testDepartmentToManyRemoveOpProfiles)
	t.Run("ProfileToUsers", testProfileToManyRemoveOpUsers)
	t.Run("RoleToProfiles", testRoleToManyRemoveOpProfiles)
}

func TestReload(t *testing.T) {
	t.Run("Blogs", testBlogsReload)
	t.Run("Departments", testDepartmentsReload)
	t.Run("Entries", testEntriesReload)
	t.Run("Profiles", testProfilesReload)
	t.Run("Roles", testRolesReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Blogs", testBlogsReloadAll)
	t.Run("Departments", testDepartmentsReloadAll)
	t.Run("Entries", testEntriesReloadAll)
	t.Run("Profiles", testProfilesReloadAll)
	t.Run("Roles", testRolesReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Blogs", testBlogsSelect)
	t.Run("Departments", testDepartmentsSelect)
	t.Run("Entries", testEntriesSelect)
	t.Run("Profiles", testProfilesSelect)
	t.Run("Roles", testRolesSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Blogs", testBlogsUpdate)
	t.Run("Departments", testDepartmentsUpdate)
	t.Run("Entries", testEntriesUpdate)
	t.Run("Profiles", testProfilesUpdate)
	t.Run("Roles", testRolesUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Blogs", testBlogsSliceUpdateAll)
	t.Run("Departments", testDepartmentsSliceUpdateAll)
	t.Run("Entries", testEntriesSliceUpdateAll)
	t.Run("Profiles", testProfilesSliceUpdateAll)
	t.Run("Roles", testRolesSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
