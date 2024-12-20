// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"context"
	"net/http"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/otelogen"
)

// handleCreateTodoRequest handles create-todo operation.
//
// TODOを登録する.
//
// POST /todo
func (s *Server) handleCreateTodoRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("create-todo"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.HTTPRouteKey.String("/todo"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "CreateTodo",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "CreateTodo",
			ID:   "create-todo",
		}
	)
	request, close, err := s.decodeCreateTodoRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response CreateTodoRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "CreateTodo",
			OperationSummary: "",
			OperationID:      "create-todo",
			Body:             request,
			Params:           middleware.Parameters{},
			Raw:              r,
		}

		type (
			Request  = OptTodoInformation
			Params   = struct{}
			Response = CreateTodoRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			nil,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.CreateTodo(ctx, request)
				return response, err
			},
		)
	} else {
		response, err = s.h.CreateTodo(ctx, request)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeCreateTodoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleDeleteTodoRequest handles delete-todo operation.
//
// TODOを削除する.
//
// DELETE /todo/{todo_id}
func (s *Server) handleDeleteTodoRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("delete-todo"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.HTTPRouteKey.String("/todo/{todo_id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "DeleteTodo",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "DeleteTodo",
			ID:   "delete-todo",
		}
	)
	params, err := decodeDeleteTodoParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response DeleteTodoRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "DeleteTodo",
			OperationSummary: "",
			OperationID:      "delete-todo",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "todo_id",
					In:   "path",
				}: params.TodoID,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = DeleteTodoParams
			Response = DeleteTodoRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackDeleteTodoParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.DeleteTodo(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.DeleteTodo(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeDeleteTodoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleGetTodoRequest handles get-todo operation.
//
// TODOの一覧を取得する.
//
// GET /todo
func (s *Server) handleGetTodoRequest(args [0]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("get-todo"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.HTTPRouteKey.String("/todo"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "GetTodo",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "GetTodo",
			ID:   "get-todo",
		}
	)
	params, err := decodeGetTodoParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	var response GetTodoRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "GetTodo",
			OperationSummary: "",
			OperationID:      "get-todo",
			Body:             nil,
			Params: middleware.Parameters{
				{
					Name: "todo_id",
					In:   "query",
				}: params.TodoID,
				{
					Name: "offset",
					In:   "query",
				}: params.Offset,
				{
					Name: "limit",
					In:   "query",
				}: params.Limit,
			},
			Raw: r,
		}

		type (
			Request  = struct{}
			Params   = GetTodoParams
			Response = GetTodoRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackGetTodoParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.GetTodo(ctx, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.GetTodo(ctx, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeGetTodoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}

// handleUpdateTodoRequest handles update-todo operation.
//
// TODOを更新する.
//
// PATCH /todo/{todo_id}
func (s *Server) handleUpdateTodoRequest(args [1]string, argsEscaped bool, w http.ResponseWriter, r *http.Request) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("update-todo"),
		semconv.HTTPRequestMethodKey.String("PATCH"),
		semconv.HTTPRouteKey.String("/todo/{todo_id}"),
	}

	// Start a span for this request.
	ctx, span := s.cfg.Tracer.Start(r.Context(), "UpdateTodo",
		trace.WithAttributes(otelAttrs...),
		serverSpanKind,
	)
	defer span.End()

	// Add Labeler to context.
	labeler := &Labeler{attrs: otelAttrs}
	ctx = contextWithLabeler(ctx, labeler)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		attrOpt := metric.WithAttributeSet(labeler.AttributeSet())

		// Increment request counter.
		s.requests.Add(ctx, 1, attrOpt)

		// Use floating point division here for higher precision (instead of Millisecond method).
		s.duration.Record(ctx, float64(float64(elapsedDuration)/float64(time.Millisecond)), attrOpt)
	}()

	var (
		recordError = func(stage string, err error) {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			s.errors.Add(ctx, 1, metric.WithAttributeSet(labeler.AttributeSet()))
		}
		err          error
		opErrContext = ogenerrors.OperationContext{
			Name: "UpdateTodo",
			ID:   "update-todo",
		}
	)
	params, err := decodeUpdateTodoParams(args, argsEscaped, r)
	if err != nil {
		err = &ogenerrors.DecodeParamsError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeParams", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	request, close, err := s.decodeUpdateTodoRequest(r)
	if err != nil {
		err = &ogenerrors.DecodeRequestError{
			OperationContext: opErrContext,
			Err:              err,
		}
		defer recordError("DecodeRequest", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}
	defer func() {
		if err := close(); err != nil {
			recordError("CloseRequest", err)
		}
	}()

	var response UpdateTodoRes
	if m := s.cfg.Middleware; m != nil {
		mreq := middleware.Request{
			Context:          ctx,
			OperationName:    "UpdateTodo",
			OperationSummary: "Update a Todo",
			OperationID:      "update-todo",
			Body:             request,
			Params: middleware.Parameters{
				{
					Name: "todo_id",
					In:   "path",
				}: params.TodoID,
			},
			Raw: r,
		}

		type (
			Request  = OptTodoInformation
			Params   = UpdateTodoParams
			Response = UpdateTodoRes
		)
		response, err = middleware.HookMiddleware[
			Request,
			Params,
			Response,
		](
			m,
			mreq,
			unpackUpdateTodoParams,
			func(ctx context.Context, request Request, params Params) (response Response, err error) {
				response, err = s.h.UpdateTodo(ctx, request, params)
				return response, err
			},
		)
	} else {
		response, err = s.h.UpdateTodo(ctx, request, params)
	}
	if err != nil {
		defer recordError("Internal", err)
		s.cfg.ErrorHandler(ctx, w, r, err)
		return
	}

	if err := encodeUpdateTodoResponse(response, w, span); err != nil {
		defer recordError("EncodeResponse", err)
		if !errors.Is(err, ht.ErrInternalServerErrorResponse) {
			s.cfg.ErrorHandler(ctx, w, r, err)
		}
		return
	}
}
