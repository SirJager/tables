// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package core_repo

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	// -------------------------- ADD ONE TO -> CORE_SESSIONS --------------------------
	CreateSession(ctx context.Context, arg CreateSessionParams) (CoreSession, error)
	// -------------------------- ADD CORE_TABLES <-> CORE_TABLES --------------------------
	CreateTable(ctx context.Context, arg CreateTableParams) (CoreTable, error)
	// ------------------------------ ADD ONE CORE_USERS <-> CORE_USER  ------------------------------
	CreateUser(ctx context.Context, arg CreateUserParams) (CoreUser, error)
	// -------------------------- REMOVE CORE_TABLES <-> CORE_TABLES --------------------------
	DeleteTable(ctx context.Context, id int64) error
	DeleteTableWhereUserAndName(ctx context.Context, arg DeleteTableWhereUserAndNameParams) error
	DeleteTablesWhereUser(ctx context.Context, userID int64) error
	// ------------------------------ REMOVE ONE CORE_USERS -> nil  ------------------------------
	DeleteUser(ctx context.Context, id int64) error
	// ------------------------------ GET MULTIPLE CORE_TABLES <== [CORE_TABLES] ------------------------------
	GetAllTables(ctx context.Context) ([]CoreTable, error)
	// ------------------------------ GET MULTIPLE CORE_USERS <== [CORE_USERS] ------------------------------
	GetAllUsers(ctx context.Context) ([]CoreUser, error)
	// -------------------------- GET ONE FROM <- CORE_SESSIONS --------------------------
	GetSession(ctx context.Context, id uuid.UUID) (CoreSession, error)
	GetSomeTables(ctx context.Context, arg GetSomeTablesParams) ([]CoreTable, error)
	GetSomeTablesWhereUser(ctx context.Context, arg GetSomeTablesWhereUserParams) ([]CoreTable, error)
	GetSomeUsers(ctx context.Context, arg GetSomeUsersParams) ([]CoreUser, error)
	// -------------------------- GET ONE CORE_TABLES <- CORE_TABLES --------------------------
	GetTable(ctx context.Context, id int64) (CoreTable, error)
	GetTableWhereIDAndUser(ctx context.Context, arg GetTableWhereIDAndUserParams) (CoreTable, error)
	GetTableWhereName(ctx context.Context, name string) (CoreTable, error)
	// --------------------- GET MULTIPLE CORE_TABLES OF CORE_USERS.user_id <== [CORE_TABLES] ---------------------
	GetTablesWhereUser(ctx context.Context, userID int64) ([]CoreTable, error)
	// ------------------------------ GET ONE CORE_USERS <== CORE_USER  ------------------------------
	GetUser(ctx context.Context, id int64) (CoreUser, error)
	GetUserWhereEmail(ctx context.Context, email string) (CoreUser, error)
	GetUserWhereUsername(ctx context.Context, username string) (CoreUser, error)
	// -------------------------- UPDATE CORE_TABLES <-> CORE_TABLES --------------------------
	UpdateTableColumns(ctx context.Context, arg UpdateTableColumnsParams) (CoreTable, error)
	UpdateUserBlocked(ctx context.Context, arg UpdateUserBlockedParams) (CoreUser, error)
	// ------------------------------ UPDATE ONE CORE_USERS <-> CORE_USERS  ------------------------------
	UpdateUserFullName(ctx context.Context, arg UpdateUserFullNameParams) (CoreUser, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (CoreUser, error)
	UpdateUserPublic(ctx context.Context, arg UpdateUserPublicParams) (CoreUser, error)
	UpdateUserUsername(ctx context.Context, arg UpdateUserUsernameParams) (CoreUser, error)
	UpdateUserVerified(ctx context.Context, arg UpdateUserVerifiedParams) (CoreUser, error)
}

var _ Querier = (*Queries)(nil)
