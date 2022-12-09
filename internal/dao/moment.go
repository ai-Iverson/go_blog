// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go_blog/internal/dao/internal"
)

// internalMomentDao is internal type for wrapping internal DAO implements.
type internalMomentDao = *internal.MomentDao

// momentDao is the data access object for table moment.
// You can define custom methods on it to extend its functionality as you wish.
type momentDao struct {
	internalMomentDao
}

var (
	// Moment is globally public accessible object for table moment operations.
	Moment = momentDao{
		internal.NewMomentDao(),
	}
)

// Fill with you ideas below.
