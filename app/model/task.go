// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"goframe_learn/app/model/internal"
)

// Task is the golang structure for table task.
type Task internal.Task

// Fill with you ideas below.
type TaskApiResponse struct {
	Id    uint64 `orm:"id,primary" json:"id"`    //
	Title string `orm:"title"      json:"Title"` //
	Done  bool   `orm:"done"       json:"Done"`  //
}

type TaskApiRequest struct {
	Title string `orm:"title"      json:"Title"` //
	Done  bool   `orm:"done"       json:"Done"`  //
}
