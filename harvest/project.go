package harvest

import (
	"fmt"
)

type ProjectService struct {
	Service
}

type Project struct {
	Id                               int         `json:"id"`
	ClientId                         int         `json:"client_id"`
	Name                             string      `json:"name"`
	Code                             string      `json:"code"`
	Notes                            string      `json:"notes"`
	BillBy                           string      `json:"bill_by"`
	BudgetBy                         string      `json:"budget_by"`
	Active                           bool        `json:"active"`
	CostBudgetIncludeExpenses        bool        `json:"cost_budget_include_expenses"`
	Billable                         bool        `json:"billable"`
	ShowBudgetToAll                  bool        `json:"show_budget_to_all"`
	CostBudget                       float32     `json:"cost_budget"`
	HourlyRate                       float32     `json:"hourly_rate"`
	Budget                           float32     `json:"budget"`
	NotifyWhenOverBudget             float32     `json:"notify_when_overbudget"`
	OverBudgetNotificationPercentage float32     `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             HarvestDate `json:"over_budget_notified_at"`
	CreatedAt                        HarvestDate `json:"created_at"`
	UpdateAt                         HarvestDate `json:"updated_at"`
	HintEarliestRecordAt             HarvestDate `json:"hint_earliest_record_at"`
	HintLatestRecordAt               HarvestDate `json:"hint_latest_record_at"`
}

type ProjectResponse struct {
	Project Project
}

func (c *ProjectService) List() (projects []Project, err error) {
	resourceURL := "/projects"
	var projectResponse []ProjectResponse
	err = c.list(resourceURL, &projectResponse)
	if err != nil {
		return
	}

	for _, element := range projectResponse {
		projects = append(projects, element.Project)
	}
	return
}

func (c *ProjectService) Find(projectID int) (project Project, err error) {
	resourceURL := fmt.Sprintf("/projects/%v", projectID)
	var projectResponse ProjectResponse
	err = c.find(resourceURL, &projectResponse)
	if err != nil {
		return
	}
	project = projectResponse.Project
	return
}
