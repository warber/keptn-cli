package keptn

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"github.com/keptn/cli2/pkg/printer"
	"github.com/keptn/go-utils/pkg/api/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetProject(t *testing.T) {
	const projectJSON = `{
   "creationDate": "1643878177026773540",
   "projectName": "echoproject",
   "shipyardVersion": "spec.keptn.sh/0.2.0",
   "stages": [
      {
         "services": [
            {
               "creationDate": "1643878187124666966",
               "lastEventTypes": {
                  "sh.keptn.event.echo.finished": {
                     "eventId": "ef3a0296-bd7d-4a27-b377-a740eec4ff9c",
                     "keptnContext": "2e11c227-da72-47d8-926b-e761683038c5",
                     "time": "1643903429716629859"
                  },
                  "sh.keptn.event.echo.started": {
                     "eventId": "a7e5fc6c-a495-4dc9-a277-deee9121e1c8",
                     "keptnContext": "2e11c227-da72-47d8-926b-e761683038c5",
                     "time": "1643903427721251404"
                  },
                  "sh.keptn.event.echo.triggered": {
                     "eventId": "c1c46913-cb24-43c5-b2a0-b072481e0c85",
                     "keptnContext": "2e11c227-da72-47d8-926b-e761683038c5",
                     "time": "1643903427621153484"
                  }
               },
               "openApprovals": null,
               "serviceName": "echoservice"
            }
         ],
         "stageName": "firststage"
      }
   ]
}
`

	t.Run("TestGetProject", func(t *testing.T) {
		p1 := &models.Project{}
		err := json.Unmarshal([]byte(projectJSON), p1)
		require.Nil(t, err)
		p1.ProjectName = "project1"

		p2 := &models.Project{}
		err = yaml.Unmarshal([]byte(projectJSON), p2)
		require.Nil(t, err)
		p2.ProjectName = "project2"

		k := NewTestKeptn(t)
		k.KeptnProjectInterfaceMock.EXPECT().GetAllProjects().Return([]*models.Project{p1, p2}, nil)
		projects, err := k.GetProjects(GetProjectsOptions{})
		require.Nil(t, err)
		assert.Equal(t, 2, len(projects))
	})

	t.Run("TestGetProject - filtered by project name", func(t *testing.T) {
		p1 := &models.Project{}
		err := json.Unmarshal([]byte(projectJSON), p1)
		require.Nil(t, err)
		p1.ProjectName = "project1"

		p2 := &models.Project{}
		err = yaml.Unmarshal([]byte(projectJSON), p2)
		require.Nil(t, err)
		p2.ProjectName = "project2"

		k := NewTestKeptn(t)
		k.KeptnProjectInterfaceMock.EXPECT().GetAllProjects().Return([]*models.Project{p1, p2}, nil)
		projects, err := k.GetProjects(GetProjectsOptions{
			ProjectName: "project1",
		})
		require.Nil(t, err)
		assert.Equal(t, 1, len(projects))
	})

	t.Run("TestGetProject - filtered by project name - project not found", func(t *testing.T) {
		p1 := &models.Project{}
		err := json.Unmarshal([]byte(projectJSON), p1)
		require.Nil(t, err)
		p1.ProjectName = "project1"

		p2 := &models.Project{}
		err = yaml.Unmarshal([]byte(projectJSON), p2)
		require.Nil(t, err)
		p2.ProjectName = "project2"

		k := NewTestKeptn(t)
		k.KeptnProjectInterfaceMock.EXPECT().GetAllProjects().Return([]*models.Project{p1, p2}, nil)
		projects, err := k.GetProjects(GetProjectsOptions{
			ProjectName: "echoproject",
		})
		require.Nil(t, err)
		assert.Equal(t, 0, len(projects))
	})
}

func TestPrintProjects(t *testing.T) {

	const projectYAML = `shipyardversion: spec.keptn.sh/0.2.3
stages:
    - services:
        - creationdate: "1642695659989253963"
          deployedimage: ""
          lasteventtypes:
            sh.keptn.event.evaluation.finished:
                eventid: 71d9a574-2fad-4f74-a5a7-b43f4860c517
                keptncontext: f0b0f2f4-1a44-4717-9d83-c3f37c63b735
                time: "1642695664120575247"
            sh.keptn.event.evaluation.started:
                eventid: 2742cbe7-9dd5-4919-82cb-0b38e8c46d8a
                keptncontext: f0b0f2f4-1a44-4717-9d83-c3f37c63b735
                time: "1642695663924583122"
            sh.keptn.event.evaluation.triggered:
                eventid: 6847ad11-6166-4273-8935-a18db1c235dc
                keptncontext: f0b0f2f4-1a44-4717-9d83-c3f37c63b735
                time: "1642695661820867900"
          openapprovals: []
          servicename: my-service
      stagename: dev
`

	const projectJSON = `{
   "creationDate": "1643878177026773540",
   "projectName": "echoproject",
   "shipyardVersion": "spec.keptn.sh/0.2.0",
   "stages": [
      {
         "services": [
            {
               "creationDate": "1643878187124666966",
               "lastEventTypes": {
                  "sh.keptn.event.echo.finished": {
                     "eventId": "ef3a0296-bd7d-4a27-b377-a740eec4ff9c",
                     "keptnContext": "2e11c227-da72-47d8-926b-e761683038c5",
                     "time": "1643903429716629859"
                  },
                  "sh.keptn.event.echo.started": {
                     "eventId": "a7e5fc6c-a495-4dc9-a277-deee9121e1c8",
                     "keptnContext": "2e11c227-da72-47d8-926b-e761683038c5",
                     "time": "1643903427721251404"
                  },
                  "sh.keptn.event.echo.triggered": {
                     "eventId": "c1c46913-cb24-43c5-b2a0-b072481e0c85",
                     "keptnContext": "2e11c227-da72-47d8-926b-e761683038c5",
                     "time": "1643903427621153484"
                  }
               },
               "openApprovals": null,
               "serviceName": "echoservice"
            }
         ],
         "stageName": "firststage"
      }
   ]
}
`
	t.Run("TestGetProject -o yaml", func(t *testing.T) {
		p := &models.Project{}
		err := yaml.Unmarshal([]byte(projectYAML), p)
		require.Nil(t, err)
		k := NewTestKeptn(t)
		k.PrintProjects([]*models.Project{p}, &printer.PrintOptions{Format: printer.FormatYAML})

		require.Nil(t, err)
		printedProject := &models.Project{}
		err = yaml.Unmarshal([]byte(k.TestCmdContext.Output()), printedProject)

		require.Nil(t, err)
		assert.Equal(t, p, printedProject)
	})

	t.Run("TestGetProject -o json", func(t *testing.T) {
		p := &models.Project{}
		err := json.Unmarshal([]byte(projectJSON), p)
		require.Nil(t, err)
		k := NewTestKeptn(t)
		k.PrintProjects([]*models.Project{p}, &printer.PrintOptions{Format: printer.FormatJSON})

		require.Nil(t, err)
		printedProject := &models.Project{}
		err = json.Unmarshal([]byte(k.TestCmdContext.Output()), printedProject)

		require.Nil(t, err)
		assert.Equal(t, p, printedProject)
	})
}
