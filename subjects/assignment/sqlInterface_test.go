package assignment

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/samonzeweb/godb"
	"github.com/samonzeweb/godb/adapters/sqlite"
)

type testCase struct {
	name        string
	ass         Json
	err         error
	returnedAss Json
	del         AssignmentDStage
}

var stagePasses = []testCase{
	{
		name: "correct update Started_at",
		ass: Json{Data: Data{
			SubjectID:   4,
			UnlockedAt:  time.Now().Add(-time.Hour),
			AvailableAt: time.Now().Add(-time.Hour),
			StartedAt:   time.Time{},
			SrsStage:    0,
		}},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(2 * time.Hour),
				StartedAt:   time.Now(),

				SrsStage: 1,
			},
		},
	},

	{
		name: "correct update passed_at",
		ass: Json{Data: Data{
			SubjectID:   4,
			UnlockedAt:  time.Now().Add(-time.Hour),
			AvailableAt: time.Now().Add(-time.Hour),
			StartedAt:   time.Now().Add(-time.Minute),
			SrsStage:    4,
		}},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(167 * time.Hour),
				StartedAt:   time.Now().Add(-time.Minute),
				PassedAt:    time.Now(),

				SrsStage: 5,
			},
		},
	},

	{
		name: "correct update burned_at",
		ass: Json{Data: Data{
			SubjectID:   4,
			UnlockedAt:  time.Now().Add(-time.Hour),
			AvailableAt: time.Now().Add(-time.Hour),
			StartedAt:   time.Now().Add(-2 * time.Minute),
			PassedAt:    time.Now().Add(-time.Minute),
			SrsStage:    8,
		}},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Time{},
				StartedAt:   time.Now().Add(-2 * time.Minute),
				PassedAt:    time.Now().Add(-time.Minute),
				BurnedAt:    time.Now(),
				SrsStage:    9,
			},
		},
	},
}

var deltaEgeCase = []testCase{
	{
		name: "normal up",
		ass: Json{Data: Data{
			SubjectID:   4,
			UnlockedAt:  time.Now().Add(-time.Hour),
			AvailableAt: time.Now().Add(-time.Hour),
			StartedAt:   time.Now().Add(-time.Minute),
			PassedAt:    time.Now().Add(-time.Minute),
			SrsStage:    5,
		}},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(335 * time.Hour),
				StartedAt:   time.Now().Add(-time.Minute),
				PassedAt:    time.Now().Add(-time.Minute),

				SrsStage: 6,
			},
		},
		del: AssignmentDUp,
	},

	{
		name: "normal down",
		ass: Json{Data: Data{
			SubjectID:   4,
			UnlockedAt:  time.Now().Add(-time.Hour),
			AvailableAt: time.Now().Add(-time.Hour),
			StartedAt:   time.Now().Add(-time.Minute),
			PassedAt:    time.Now().Add(-time.Minute),
			SrsStage:    7,
		}},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(335 * time.Hour),
				StartedAt:   time.Now().Add(-time.Minute),
				PassedAt:    time.Now().Add(-time.Minute),

				SrsStage: 6,
			},
		},
		del: AssignmentDDown,
	},

	{
		name: "under PassedAt",
		ass: Json{Data: Data{
			SubjectID:   4,
			UnlockedAt:  time.Now().Add(-time.Hour),
			AvailableAt: time.Now().Add(-time.Hour),
			StartedAt:   time.Now().Add(-time.Minute),
			PassedAt:    time.Now().Add(-time.Minute),
			SrsStage:    5,
		}},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(23 * time.Hour),
				StartedAt:   time.Now().Add(-time.Minute),
				PassedAt:    time.Now().Add(-time.Minute),

				SrsStage: 4,
			},
		},
		del: AssignmentDDown,
	},

	{
		name: "over burn stage",
		ass: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now(),
				StartedAt:   time.Now().Add(-2 * time.Minute),
				PassedAt:    time.Now().Add(-time.Minute),
				BurnedAt:    time.Now(),
				SrsStage:    9,
			},
		},
		err: fmt.Errorf("blond"),
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now(),
				StartedAt:   time.Now().Add(-2 * time.Minute),
				PassedAt:    time.Now().Add(-time.Minute),
				BurnedAt:    time.Now(),
				SrsStage:    9,
			},
		},
		del: AssignmentDUp,
	},
	{
		name: "under zero",
		ass: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(-time.Hour),
				SrsStage:    0,
			},
		},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(-time.Hour),
				SrsStage:    0,
			},
		},
		del: AssignmentDDown,
	},

	{
		name: "under starting stage",
		ass: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(-2 * time.Minute),
				StartedAt:   time.Now().Add(-2 * time.Minute),
				SrsStage:    1,
			},
		},
		err: nil,
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(2 * time.Hour),
				StartedAt:   time.Now().Add(-2 * time.Minute),
				SrsStage:    1,
			},
		},
		del: AssignmentDDown,
	},

	{
		name: "not unlocked",
		ass: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:  4,
				UnlockedAt: time.Time{},
				SrsStage:   0,
			},
		},
		err: fmt.Errorf("blond"),
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:  4,
				UnlockedAt: time.Time{},
				SrsStage:   0,
			},
		},
		del: AssignmentDUp,
	},

	{
		name: "unlocked in future",
		ass: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:  4,
				UnlockedAt: time.Now().Add(2 * time.Hour),
				SrsStage:   0,
			},
		},
		err: fmt.Errorf("blond"),
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:  4,
				UnlockedAt: time.Now().Add(2 * time.Hour),
				SrsStage:   0,
			},
		},
		del: AssignmentDUp,
	},

	{
		name: "not available",
		ass: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(2 * time.Hour),
				StartedAt:   time.Now().Add(-2 * time.Minute),
				SrsStage:    1,
			},
		},
		err: fmt.Errorf("blond"),
		returnedAss: Json{
			DataUpdatedAt: time.Now(),
			Data: Data{
				SubjectID:   4,
				UnlockedAt:  time.Now().Add(-time.Hour),
				AvailableAt: time.Now().Add(2 * time.Hour),
				StartedAt:   time.Now().Add(-2 * time.Minute),
				SrsStage:    1,
			},
		},
		del: AssignmentDUp,
	},
}

var opt = cmp.Comparer(func(x, y time.Time) bool {
	return x.Round(time.Second * 5).Equal(y.Round(time.Second * 5))
})

func TestProcessProgress(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	b = filepath.Dir(b)
	basepath := strings.TrimSuffix(b, "/subjects/assignment")
	db, err := godb.Open(sqlite.Adapter, basepath+"/subjects.db")
	if err != nil {
		t.Error(err)
	}

	for _, test := range stagePasses {
		t.Run("stagePasses/"+test.name, func(t *testing.T) {

			err := (&test.ass).processProgress(db, AssignmentDUp)
			if (err != nil) != (test.err != nil) {
				t.Errorf("difernt errors, got: %s", err)
			} else {
				diff := cmp.Diff(test.ass,
					test.returnedAss, opt)
				if len(diff) > 0 {
					t.Error(diff)
				}
			}
		})
	}

	for _, test := range deltaEgeCase {
		t.Run("deltaEgeCase/"+test.name, func(t *testing.T) {

			err := (&test.ass).processProgress(db, test.del)
			if (err != nil) != (test.err != nil) {
				t.Errorf("difernt errors, got: %s", err)
			} else {
				diff := cmp.Diff(test.ass,
					test.returnedAss, opt)
				if len(diff) > 0 {
					t.Error(diff)
				}
			}
		})
	}
}
