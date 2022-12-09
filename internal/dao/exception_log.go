// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go_blog/internal/dao/internal"
)

// internalExceptionLogDao is internal type for wrapping internal DAO implements.
type internalExceptionLogDao = *internal.ExceptionLogDao

// exceptionLogDao is the data access object for table exception_log.
// You can define custom methods on it to extend its functionality as you wish.
type exceptionLogDao struct {
	internalExceptionLogDao
}

var (
	// ExceptionLog is globally public accessible object for table exception_log operations.
	ExceptionLog = exceptionLogDao{
		internal.NewExceptionLogDao(),
	}
)

// Fill with you ideas below.
