// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package scheduler

import (
	"time"

	"github.com/mattermost/mattermost-server/server/v8/channels/jobs"
	"github.com/mattermost/mattermost-server/server/v8/public/model"
)

const schedFreq = 24 * time.Hour

func MakeScheduler(jobServer *jobs.JobServer) model.Scheduler {
	isEnabled := func(cfg *model.Config) bool {
		return true
	}
	return jobs.NewPeriodicScheduler(jobServer, model.JobTypePlugins, schedFreq, isEnabled)
}
