// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"time"
)

type CreateTodoBadRequest TodoOpeError

func (*CreateTodoBadRequest) createTodoRes() {}

type CreateTodoInternalServerError TodoOpeError

func (*CreateTodoInternalServerError) createTodoRes() {}

type DeleteTodoBadRequest TodoOpeError

func (*DeleteTodoBadRequest) deleteTodoRes() {}

type DeleteTodoInternalServerError TodoOpeError

func (*DeleteTodoInternalServerError) deleteTodoRes() {}

type DeleteTodoNotFound TodoOpeError

func (*DeleteTodoNotFound) deleteTodoRes() {}

// DeleteTodoOK is response for DeleteTodo operation.
type DeleteTodoOK struct{}

func (*DeleteTodoOK) deleteTodoRes() {}

type GetTodoBadRequest TodoOpeError

func (*GetTodoBadRequest) getTodoRes() {}

type GetTodoInternalServerError TodoOpeError

func (*GetTodoInternalServerError) getTodoRes() {}

type GetTodoOKApplicationJSON []TodoInformation

func (*GetTodoOKApplicationJSON) getTodoRes() {}

// NewOptDate returns new OptDate with value set to v.
func NewOptDate(v time.Time) OptDate {
	return OptDate{
		Value: v,
		Set:   true,
	}
}

// OptDate is optional time.Time.
type OptDate struct {
	Value time.Time
	Set   bool
}

// IsSet returns true if OptDate was set.
func (o OptDate) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptDate) Reset() {
	var v time.Time
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptDate) SetTo(v time.Time) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptDate) Get() (v time.Time, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptDate) Or(d time.Time) time.Time {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptInt32 returns new OptInt32 with value set to v.
func NewOptInt32(v int32) OptInt32 {
	return OptInt32{
		Value: v,
		Set:   true,
	}
}

// OptInt32 is optional int32.
type OptInt32 struct {
	Value int32
	Set   bool
}

// IsSet returns true if OptInt32 was set.
func (o OptInt32) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt32) Reset() {
	var v int32
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt32) SetTo(v int32) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt32) Get() (v int32, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt32) Or(d int32) int32 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTodoInformation returns new OptTodoInformation with value set to v.
func NewOptTodoInformation(v TodoInformation) OptTodoInformation {
	return OptTodoInformation{
		Value: v,
		Set:   true,
	}
}

// OptTodoInformation is optional TodoInformation.
type OptTodoInformation struct {
	Value TodoInformation
	Set   bool
}

// IsSet returns true if OptTodoInformation was set.
func (o OptTodoInformation) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTodoInformation) Reset() {
	var v TodoInformation
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTodoInformation) SetTo(v TodoInformation) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTodoInformation) Get() (v TodoInformation, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTodoInformation) Or(d TodoInformation) TodoInformation {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/TodoId
type TodoId struct {
	// TODOの登録番号.
	TodoID OptInt32 `json:"todo_id"`
}

// GetTodoID returns the value of TodoID.
func (s *TodoId) GetTodoID() OptInt32 {
	return s.TodoID
}

// SetTodoID sets the value of TodoID.
func (s *TodoId) SetTodoID(val OptInt32) {
	s.TodoID = val
}

func (*TodoId) createTodoRes() {}

// Ref: #/components/schemas/TodoInformation
type TodoInformation struct {
	// TODOの件名.
	Title string `json:"title"`
	// TODOの詳細.
	Detail OptString `json:"detail"`
	// TODOの進捗率.
	Progress int32 `json:"progress"`
	// 開始予定日.
	StartDate OptDate `json:"start_date"`
	// 終了予定日.
	EndDate OptDate `json:"end_date"`
}

// GetTitle returns the value of Title.
func (s *TodoInformation) GetTitle() string {
	return s.Title
}

// GetDetail returns the value of Detail.
func (s *TodoInformation) GetDetail() OptString {
	return s.Detail
}

// GetProgress returns the value of Progress.
func (s *TodoInformation) GetProgress() int32 {
	return s.Progress
}

// GetStartDate returns the value of StartDate.
func (s *TodoInformation) GetStartDate() OptDate {
	return s.StartDate
}

// GetEndDate returns the value of EndDate.
func (s *TodoInformation) GetEndDate() OptDate {
	return s.EndDate
}

// SetTitle sets the value of Title.
func (s *TodoInformation) SetTitle(val string) {
	s.Title = val
}

// SetDetail sets the value of Detail.
func (s *TodoInformation) SetDetail(val OptString) {
	s.Detail = val
}

// SetProgress sets the value of Progress.
func (s *TodoInformation) SetProgress(val int32) {
	s.Progress = val
}

// SetStartDate sets the value of StartDate.
func (s *TodoInformation) SetStartDate(val OptDate) {
	s.StartDate = val
}

// SetEndDate sets the value of EndDate.
func (s *TodoInformation) SetEndDate(val OptDate) {
	s.EndDate = val
}

// Ref: #/components/schemas/TodoOpeError
type TodoOpeError struct {
	// エラーコード.
	ErrorCode OptFloat64 `json:"error_code"`
	// エラーメッセージ.
	ErrorMessage OptString `json:"error_message"`
}

// GetErrorCode returns the value of ErrorCode.
func (s *TodoOpeError) GetErrorCode() OptFloat64 {
	return s.ErrorCode
}

// GetErrorMessage returns the value of ErrorMessage.
func (s *TodoOpeError) GetErrorMessage() OptString {
	return s.ErrorMessage
}

// SetErrorCode sets the value of ErrorCode.
func (s *TodoOpeError) SetErrorCode(val OptFloat64) {
	s.ErrorCode = val
}

// SetErrorMessage sets the value of ErrorMessage.
func (s *TodoOpeError) SetErrorMessage(val OptString) {
	s.ErrorMessage = val
}

type UpdateTodoBadRequest TodoOpeError

func (*UpdateTodoBadRequest) updateTodoRes() {}

type UpdateTodoInternalServerError TodoOpeError

func (*UpdateTodoInternalServerError) updateTodoRes() {}

// UpdateTodoNoContent is response for UpdateTodo operation.
type UpdateTodoNoContent struct{}

func (*UpdateTodoNoContent) updateTodoRes() {}

type UpdateTodoNotFound TodoOpeError

func (*UpdateTodoNotFound) updateTodoRes() {}
