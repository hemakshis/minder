// Copyright 2023 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package rule provides the CLI subcommand for managing rules

// Package interfaces provides necessary interfaces and implementations for
// implementing engine plugins
package interfaces

import (
	"context"
	"encoding/json"

	billy "github.com/go-git/go-billy/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/stacklok/minder/internal/db"
	evalerrors "github.com/stacklok/minder/internal/engine/errors"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
)

// Ingester is the interface for a rule type ingester
type Ingester interface {
	// Ingest does the actual data ingestion for a rule type
	Ingest(ctx context.Context, ent protoreflect.ProtoMessage, params map[string]any) (*Result, error)
	// GetType returns the type of the ingester
	GetType() string
	// GetConfig returns the config for the ingester
	GetConfig() protoreflect.ProtoMessage
}

// Evaluator is the interface for a rule type evaluator
type Evaluator interface {
	Eval(ctx context.Context, profile map[string]any, res *Result) error
}

// Result is the result of an ingester
type Result struct {
	// Object is the object that was ingested. Normally comes from an external
	// system like an HTTP server.
	Object any
	// Fs is the filesystem that was created as a result of the ingestion. This
	// is normally used by the evaluator to do rule evaluation. The filesystem
	// may be a git repo, or a memory filesystem.
	Fs billy.Filesystem
}

// ActionOpt is the type that defines what action to take when remediating
type ActionOpt int

const (
	// ActionOptOn means perform the remediation
	ActionOptOn ActionOpt = iota
	// ActionOptOff means do not perform the remediation
	ActionOptOff
	// ActionOptDryRun means perform a dry run of the remediation
	ActionOptDryRun
	// ActionOptUnknown means the action is unknown. This is a sentinel value.
	ActionOptUnknown
)

// ActionOptFromString returns the ActionOpt from a string representation
func ActionOptFromString(s *string, defAction ActionOpt) ActionOpt {
	var actionOptMap = map[string]ActionOpt{
		"on":      ActionOptOn,
		"off":     ActionOptOff,
		"dry_run": ActionOptDryRun,
	}

	if s == nil {
		return defAction
	}

	if v, ok := actionOptMap[*s]; ok {
		return v
	}

	return ActionOptUnknown
}

// ActionType represents the type of action, i.e., remediate, alert, etc.
type ActionType string

// Action is the interface for a rule type action
type Action interface {
	Class() ActionType
	Type() string
	GetOnOffState(*pb.Profile) ActionOpt
	Do(ctx context.Context, cmd ActionCmd, setting ActionOpt, entity protoreflect.ProtoMessage,
		params ActionsParams, metadata *json.RawMessage) (json.RawMessage, error)
}

// ActionCmd is the type that defines what effect an action should have
type ActionCmd string

const (
	// ActionCmdOff means turn off the action
	ActionCmdOff ActionCmd = "turn_off"
	// ActionCmdOn means turn on the action
	ActionCmdOn ActionCmd = "turn_on"
	// ActionCmdDoNothing means the action should do nothing
	ActionCmdDoNothing ActionCmd = "do_nothing"
)

// EvalStatusParams is a helper struct to pass parameters to createOrUpdateEvalStatus
// to avoid confusion with the parameters' order. Since at the moment, all our entities are bound to
// a repo and most profiles are expecting a repo, the RepoID parameter is mandatory. For entities
// other than artifacts, the ArtifactID should be 0 that is translated to NULL in the database.
type EvalStatusParams struct {
	Profile          *pb.Profile
	Rule             *pb.Profile_Rule
	RuleType         *pb.RuleType
	ProfileID        uuid.UUID
	RepoID           uuid.UUID
	ArtifactID       uuid.NullUUID
	PullRequestID    uuid.NullUUID
	EntityType       db.Entities
	RuleTypeID       uuid.UUID
	EvalStatusFromDb *db.ListRuleEvaluationsByProfileIdRow
	evalErr          error
	actionsErr       evalerrors.ActionsError
}

// Ensure EvalStatusParams implements the necessary interfaces
var _ ActionsParams = (*EvalStatusParams)(nil)
var _ EvalParams = (*EvalStatusParams)(nil)

// GetEvalErr returns the evaluation error
func (e *EvalStatusParams) GetEvalErr() error {
	return e.evalErr
}

// SetEvalErr sets the evaluation error
func (e *EvalStatusParams) SetEvalErr(err error) {
	e.evalErr = err
}

// SetActionsErr sets the actions' error
func (e *EvalStatusParams) SetActionsErr(ctx context.Context, actionErr evalerrors.ActionsError) {
	// Get logger
	logger := zerolog.Ctx(ctx)

	// Make sure we don't try to push a nil json.RawMessage accidentally
	if actionErr.AlertMeta == nil {
		// Default to an empty json struct if the action did not return anything
		m, err := json.Marshal(&map[string]any{})
		if err != nil {
			// This should never happen since we are marshaling an empty struct
			logger.Error().Err(err).Msg("error marshaling empty json.RawMessage")
		}
		actionErr.AlertMeta = m
	}
	if actionErr.RemediateMeta == nil {
		// Default to an empty json struct if the action did not return anything
		m, err := json.Marshal(&map[string]any{})
		if err != nil {
			// This should never happen since we are marshaling an empty struct
			logger.Error().Err(err).Msg("error marshaling empty json.RawMessage")
		}
		actionErr.RemediateMeta = m
	}
	// All okay
	e.actionsErr = actionErr
}

// GetActionsErr returns the actions' error
func (e *EvalStatusParams) GetActionsErr() evalerrors.ActionsError {
	return e.actionsErr
}

// GetRule returns the rule
func (e *EvalStatusParams) GetRule() *pb.Profile_Rule {
	return e.Rule
}

// GetEvalStatusFromDb returns the evaluation status from the database
func (e *EvalStatusParams) GetEvalStatusFromDb() *db.ListRuleEvaluationsByProfileIdRow {
	return e.EvalStatusFromDb
}

// GetRuleType returns the rule type
func (e *EvalStatusParams) GetRuleType() *pb.RuleType {
	return e.RuleType
}

// GetProfile returns the profile
func (e *EvalStatusParams) GetProfile() *pb.Profile {
	return e.Profile
}

// EvalParams is the interface used for a rule type evaluator
type EvalParams interface {
	GetRule() *pb.Profile_Rule
}

// ActionsParams is the interface used for processing a rule type action
type ActionsParams interface {
	GetEvalErr() error
	GetEvalStatusFromDb() *db.ListRuleEvaluationsByProfileIdRow
	GetRuleType() *pb.RuleType
	GetProfile() *pb.Profile
	GetRule() *pb.Profile_Rule
}
