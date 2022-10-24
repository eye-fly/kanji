package json

import (
	"fmt"
	"sql_filler/subjects/assignment"
	"strconv"
)

func (bec *backEnd) saveProgress(userID int, progress map[string][]interface{}) error {
	for key, val := range progress {
		subjectID, err := strconv.Atoi(key)
		if err != nil {
			return err
		}

		if len(val) != 2 {
			return fmt.Errorf("bad input format list, should have 2 values")
		}
		l1, ok := (val[0]).(float64)
		if !ok {
			return fmt.Errorf("atoi[0] err original S: %v", val[0])
		}

		l2, ok := (val[1]).(float64)
		if !ok {
			//return fmt.Errorf("atoi[1] err original S: %v", val[1])
			l2 = 0
		}

		var stageDelta assignment.AssignmentDStage
		if int(l1)+int(l2) > 0 {
			stageDelta = assignment.AssignmentDDown
		} else {
			stageDelta = assignment.AssignmentDUp
		}

		err = assignment.UpdateAssignment(bec.db, userID, subjectID, stageDelta)
		if err != nil {
			return err
		}
	}
	return nil
}
