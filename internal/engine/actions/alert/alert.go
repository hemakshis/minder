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

// Package alert provides necessary interfaces and implementations for
// processing alerts.
package alert

import (
	"fmt"

	"github.com/stacklok/minder/internal/engine/actions/alert/noop"
	"github.com/stacklok/minder/internal/engine/actions/alert/security_advisory"
	engif "github.com/stacklok/minder/internal/engine/interfaces"
	"github.com/stacklok/minder/internal/providers"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
)

// ActionType is the type of the alert engine
const ActionType engif.ActionType = "alert"

// NewRuleAlert creates a new rule alert engine
func NewRuleAlert(rt *pb.RuleType, pbuild *providers.ProviderBuilder) (engif.Action, error) {
	alertCfg := rt.Def.GetAlert()
	if alertCfg == nil {
		return noop.NewNoopAlert(ActionType)
	}

	// nolint:revive // let's keep the switch here, it would be nicer to extend a switch in the future
	switch alertCfg.GetType() {
	case security_advisory.AlertType:
		if alertCfg.GetSecurityAdvisory() == nil {
			return nil, fmt.Errorf("alert engine missing security-advisory configuration")
		}
		return security_advisory.NewSecurityAdvisoryAlert(ActionType, alertCfg.GetSecurityAdvisory(), pbuild)
	}

	return nil, fmt.Errorf("unknown alert type: %s", alertCfg.GetType())
}
