package harvest

import (
	"fmt"
	"time"
)

type TimesheetService struct {
	Service
}

type DayEntry struct {
	ProjectId                    string					`json:"project_id"`
	Project     				 string					`json:"project"`
	UserId      				 int					`json:"user_id"`
	SpentAt                		 HarvestDate			`json:"spent_at"`
	TaskId         				 string					`json:"task_id"`
	Task            			 string					`json:"task"`
	Client                   	 string					`json:"client"`
	Id                     		 int					`json:"id"`
	Notes                        string					`json:"notes"`
	StartedAt               	 string					`json:"started_at"`
	EndedAt         			 string					`json:"ended_at"`
	HoursWithoutTimer			 float64				`json:"hours_without_timer"`
	Hours						 float64				`json:"hours"`
	CreatedAt                    time.Time				`json:"created_at"`
	UpdatedAt                    time.Time				`json:"updated_at"`
}

type TimesheetResponse struct{
	DayEntries				[]DayEntry			`json:"day_entries"`
	ForDay					HarvestDate			`json:"for_day"`
	Projects				[]Project			`json:"projects"`
}

func (c *TimesheetService) Find(day string, year string, personId string) (timesheetResponse TimesheetResponse, err error) {
	resourceURL := fmt.Sprintf("/daily?of_user=%v", personId)
	if year != "" && day != ""{
		resourceURL = fmt.Sprintf("/daily/%v/%v?of_user=%v", day, year, personId)
	}

	err = c.find(resourceURL, &timesheetResponse)
	return
}

