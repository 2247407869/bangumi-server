// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package query

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"

	"github.com/bangumi/server/internal/app/internal/cachekey"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/errgo"
	"github.com/bangumi/server/internal/pkg/generic/gmap"
)

func (q Query) GetUser(ctx context.Context, userID model.UserID) (model.User, error) {
	q.metricUserQueryCount.Inc(1)
	var key = cachekey.User(userID)

	// try to read from cache
	var r model.User
	ok, err := q.cache.Get(ctx, key, &r)
	if err != nil {
		return r, errgo.Wrap(err, "cache.Get")
	}

	if ok {
		q.metricUserQueryCached.Inc(1)
		return r, nil
	}

	r, err = q.user.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return r, domain.ErrUserNotFound
		}

		return r, errgo.Wrap(err, "userRepo.Get")
	}

	if e := q.cache.Set(ctx, key, r, time.Hour); e != nil {
		q.log.Error("can't set response to cache", zap.Error(e))
	}

	return r, nil
}

func (q Query) GetUsersByIDs(ctx context.Context, userIDs ...model.UserID) (map[model.UserID]model.User, error) {
	q.metricUserQueryCount.Inc(int64(len(userIDs)))
	var notCached = make([]model.UserID, 0, len(userIDs))

	var result = make(map[model.UserID]model.User, len(userIDs))
	for _, userID := range userIDs {
		key := cachekey.User(userID)
		var s model.User
		ok, err := q.cache.Get(ctx, key, &s)
		if err != nil {
			return nil, errgo.Wrap(err, "cache.Get")
		}

		if ok {
			result[userID] = s
		} else {
			notCached = append(notCached, userID)
		}
	}

	q.metricUserQueryCached.Inc(int64(len(result)))
	newUserMap, err := q.user.GetByIDs(ctx, notCached...)
	if err != nil {
		return nil, errgo.Wrap(err, "failed to get subjects")
	}

	for userID, user := range newUserMap {
		err := q.cache.Set(ctx, cachekey.User(userID), user, time.Hour)
		if err != nil {
			q.log.Error("failed to set user cache")
		}
	}

	gmap.Copy(result, newUserMap)

	return result, nil
}
