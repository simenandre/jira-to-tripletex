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

package jira

import (
	"github.com/cobraz/jira-to-tripletex/internal/pkg/config"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// New represents new Jira Config
func New() (*jira.Client, error) {
	cnf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	tp := jira.BasicAuthTransport{
		Username: cnf.JiraEmail,
		Password: cnf.JiraToken,
	}

	return jira.NewClient(tp.Client(), cnf.Host)
}
