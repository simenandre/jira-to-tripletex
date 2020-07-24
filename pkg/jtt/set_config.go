// Copyright 2020 Simen A. W. Olsen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jtt

import (
	"github.com/cobraz/jira-to-tripletex/internal/pkg/config"
	"github.com/urfave/cli/v2"
)

// SetConfig is used to set config
func SetConfig(c *cli.Context) error {
	err := config.SetConfig(config.Config{
		Host:          c.String("host"),
		JiraToken:     c.String("jiraToken"),
		JiraEmail:     c.String("jiraEmail"),
		EmployeeToken: c.String("employeeToken"),
		ConsumerToken: c.String("consumerToken"),
		ActivityCode:  c.String("activityCode"),
	})
	if err != nil {
		return err
	}
	return nil
}
