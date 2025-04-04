// Copyright (C) 2024  Morgan S Hein
//
// This program is subject to the terms
// of the GNU Affero General Public License, version 3.
// If a copy of the AGPL was not distributed with this file, You
// can obtain one at https://www.gnu.org/licenses/.

package gadgets

import (
	"go/format"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSubCommandsTemplate(t *testing.T) {
	// Define mock data for the template
	data := renderData{
		GoGoImportPath: GOGOIMPORTPATH,
		RootCmd: GoCmd{
			Name:  "PrintHello",
			Short: "A short description",
			Long:  "A much longer description. Much wow!",
			GoFlags: []GoFlag{
				{Type: "string", Name: "config", Short: 'c', Default: `""`, Help: "config file (default is ./config.yaml)"},
				{Type: "bool", Name: "verbose", Short: 'v', Default: false, Help: "enable verbose mode"},
			},
		},
		SubCommands: []GoCmd{
			{
				Name:  "SubCommandA",
				Short: "A short description for SubCmdA",
				Long:  "A much longer description for SubCmdA. Much wow!",
				GoFlags: []GoFlag{
					{Type: "bool", Name: "print", Short: 'p', Default: false, Help: "Print extra information on the result."},
					{Type: "string", Name: "shout", Default: `""`, Help: "Words to shout."},
				},
			},
		},
	}

	templateNames := []string{
		"templates/main.go.tmpl",
		"templates/subCmd.go.tmpl",
		"templates/function.go.tmpl",
	}

	//// find mod root
	//root, err := mod.FindModuleRoot()
	//require.NoError(t, err, "Failed to find module root: %v", err)
	//
	//// for each template name, pass in the full path
	//for i, name := range templateNames {
	//	templateNames[i] = path.Join(root, "pkg", "funcs", name)
	//}

	// Render the template with the mock data
	rendered, err := renderFromTemplates(data, defaultFuncMap(), templateNames)
	require.NoError(t, err, "Failed to execute template: %v", err)
	assert.Contains(t, rendered, "package main", "Expected package main")
	assert.Contains(t, rendered, "PrintHello", "Expected PrintHello function")
	formatted, err := format.Source([]byte(rendered))
	require.NoError(t, err, "Failed to format source: %v", err)
	cupaloy.SnapshotT(t, formatted)
}
