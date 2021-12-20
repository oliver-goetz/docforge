// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package reactor

import (
	"github.com/gardener/docforge/pkg/api"
	"github.com/gardener/docforge/pkg/resourcehandlers"
	githubHandler "github.com/gardener/docforge/pkg/resourcehandlers/github"
	"github.com/gardener/docforge/pkg/resourcehandlers/testhandler"
	"github.com/gardener/docforge/pkg/util/tests"
	"github.com/gardener/docforge/pkg/util/urls"
	"github.com/google/go-github/v32/github"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

// TODO: This is a flaky test. In the future the ResourceHandler should be mocked.
func Test_processLink(t *testing.T) {
	nodeA := &api.Node{
		Name:   "node_A.md",
		Source: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
	}
	nodeB := &api.Node{
		Name:   "node_B.md",
		Source: "https://github.com/gardener/gardener/blob/v1.10.0/docs/extensions/overview.md",
	}
	nodeA.Nodes = []*api.Node{nodeB}
	nodeA.SetParentsDownwards()

	testCases := []struct {
		name              string
		node              *api.Node
		destination       string
		contentSourcePath string
		wantDestination   *string
		//wantText          *string
		//wantTitle         *string
		//wantDownload      *DownloadTask
		wantErr error
		mutate  func(c *nodeContentProcessor)
	}{
		// skipped links
		{
			name:              "Internal document links are not processed",
			destination:       "#internal-link",
			contentSourcePath: "",
			wantDestination:   tests.StrPtr("#internal-link"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name:              "mailto protocol is not processed",
			destination:       "mailto:a@b.com",
			contentSourcePath: "",
			wantDestination:   tests.StrPtr("mailto:a@b.com"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name:              "Absolute links to releases not processed",
			destination:       "https://github.com/gardener/gardener/releases/tag/v1.4.0",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/gardener/gardener/releases/tag/v1.4.0"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name:              "Relative links to releases not processed",
			destination:       "../../../releases/tag/v1.4.0",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/gardener/gardener/releases/tag/v1.4.0"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		// links to resources
		{
			name:              "Relative link to resource NOT in download scope",
			destination:       "./image.png",
			contentSourcePath: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			wantDestination:   tests.StrPtr("https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name: "Relative link to resource in download scope",
			node: &api.Node{
				Source: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			},
			destination:       "./image.png",
			contentSourcePath: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			wantDestination:   tests.StrPtr("/__resources"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload: &DownloadTask{
			//	Source:    "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png",
			//	Target:    "",
			//	Referer:   "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			//	Reference: "./image.png",
			//},
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				h := testhandler.NewTestResourceHandler().WithAccept(func(uri string) bool {
					return true
				}).WithBuildAbsLink(func(source, link string) (string, error) {
					return "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png", nil
				}).WithSetVersion(func(absLink, version string) (string, error) {
					return absLink, nil
				})
				c.resourceHandlers = resourcehandlers.NewRegistry(h)
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
					Downloads: &api.Downloads{
						Scope: map[string]api.ResourceRenameRules{
							"/gardener/gardener/(blob|raw|wiki)/v1.10.0/docs": nil,
						},
					},
				}
			},
		},
		{
			name: "Relative link to resource NOT in download scope",
			node: &api.Node{
				Source: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			},
			destination:       "../image.png",
			contentSourcePath: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			wantDestination:   tests.StrPtr("https://github.com/gardener/gardener/blob/v1.10.0/image.png"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name:              "Absolute link to resource NOT in download scope",
			node:              nodeA,
			destination:       "https://github.com/owner/repo/blob/master/docs/image.png",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/owner/repo/blob/master/docs/image.png"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name: "Absolute link to resource in download scope",
			node: &api.Node{
				Source: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			},
			destination:       "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png",
			contentSourcePath: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			wantDestination:   tests.StrPtr("/__resources"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload: &DownloadTask{
			//	Source:    "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png",
			//	Target:    "",
			//	Referer:   "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			//	Reference: "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png",
			//},
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
					Downloads: &api.Downloads{
						Scope: map[string]api.ResourceRenameRules{
							"/gardener/gardener/(blob|raw|wiki)/v1.10.0/docs": nil,
						},
					},
				}
			},
		},
		// links to documents
		{
			name:              "Absolute link to document NOT in download scope and NOT from structure",
			node:              nodeA,
			destination:       "https://github.com/owner/repo/blob/master/docs/doc.md",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/owner/repo/blob/master/docs/doc.md"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name:              "Absolute link to document in download scope and from structure",
			node:              nodeA,
			destination:       nodeB.Source,
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("./node_B.md"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				c.sourceLocations = map[string][]*api.Node{nodeB.Source: {nodeB}}
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
					Downloads: &api.Downloads{
						Scope: map[string]api.ResourceRenameRules{
							"/gardener/gardener/(blob|raw|wiki)/v1.10.0/docs": nil,
						},
					},
				}
			},
		},
		{
			name:              "Relative link to document NOT in download scope and NOT from structure",
			node:              nodeA,
			destination:       "https://github.com/owner/repo/blob/master/docs/doc.md",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/owner/repo/blob/master/docs/doc.md"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
		},
		{
			name:              "Relative link to document in download scope and NOT from structure",
			node:              nodeA,
			destination:       "./another.md",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/gardener/gardener/blob/v1.10.0/docs/another.md"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
					Downloads: &api.Downloads{
						Scope: map[string]api.ResourceRenameRules{
							"/gardener/gardener/(blob|raw|wiki)/v1.10.0/docs": nil,
						},
					},
				}
			},
		},
		// Version rewrite
		{
			name:              "Absolute link to document not in download scope version rewrite",
			node:              nodeA,
			destination:       "https://github.com/gardener/gardener/blob/master/sample.md",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("https://github.com/gardener/gardener/blob/v1.10.0/sample.md"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload:      nil,
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/(blob)": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
				}
			},
		},
		{
			name:              "Absolute link to resource version rewrite",
			node:              nodeA,
			destination:       "https://github.com/gardener/gardener/blob/master/docs/image.png",
			contentSourcePath: nodeA.Source,
			wantDestination:   tests.StrPtr("/__resources"),
			//wantText:          nil,
			//wantTitle:         nil,
			//wantDownload: &DownloadTask{
			//	Source:    "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png",
			//	Target:    "",
			//	Referer:   "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			//	Reference: "https://github.com/gardener/gardener/blob/master/docs/image.png",
			//},
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
					Downloads: &api.Downloads{
						Scope: map[string]api.ResourceRenameRules{
							"/gardener/gardener/(blob|raw|wiki)/v1.10.0/docs": nil,
						},
					},
				}
			},
		},
		{
			name: "Relative link to resource in download scope with rewrites",
			node: &api.Node{
				Source: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
				Links: &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png": {
							Text:  tests.StrPtr("Test text"),
							Title: tests.StrPtr("Test title"),
						},
					},
				},
			},
			destination:       "./image.png",
			contentSourcePath: "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			wantDestination:   tests.StrPtr("/__resources"),
			//wantText:          tests.StrPtr("Test text"),
			//wantTitle:         tests.StrPtr("Test title"),
			//wantDownload: &DownloadTask{
			//	Source:    "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png",
			//	Target:    "",
			//	Referer:   "https://github.com/gardener/gardener/blob/v1.10.0/docs/README.md",
			//	Reference: "./image.png",
			//},
			wantErr: nil,
			mutate: func(c *nodeContentProcessor) {
				h := testhandler.NewTestResourceHandler().WithAccept(func(uri string) bool {
					return true
				}).WithBuildAbsLink(func(source, link string) (string, error) {
					return "https://github.com/gardener/gardener/blob/v1.10.0/docs/image.png", nil
				}).WithSetVersion(func(absLink, version string) (string, error) {
					return absLink, nil
				})
				c.resourceHandlers = resourcehandlers.NewRegistry(h)
				c.globalLinksConfig = &api.Links{
					Rewrites: map[string]*api.LinkRewriteRule{
						"/gardener/gardener/": {
							Version: tests.StrPtr("v1.10.0"),
						},
					},
					Downloads: &api.Downloads{
						Scope: map[string]api.ResourceRenameRules{
							"/gardener/gardener/(blob|raw|wiki)/v1.10.0/docs": nil,
						},
					},
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			c := &nodeContentProcessor{
				resourceAbsLinks: make(map[string]string),
				resourcesRoot:    "/__resources",
				resourceHandlers: resourcehandlers.NewRegistry(githubHandler.NewResourceHandler(github.NewClient(nil), http.DefaultClient, []string{"github.com"})),
				rewriteEmbedded:  true,
				validator:        &fakeValidator{},
				downloader:       &fakeDownload{},
			}
			if tc.mutate != nil {
				tc.mutate(c)
			}
			//ctx, cancel := context.WithCancel(context.Background())
			//defer cancel()

			lr := &linkResolver{
				nodeContentProcessor: c,
				node:                 tc.node,
				source:               tc.contentSourcePath,
			}
			link, gotErr := lr.resolveBaseLink(tc.destination)

			assert.Equal(t, tc.wantErr, gotErr)
			//if gotDownload != nil {
			//	if tc.wantDownload != nil {
			//		tc.wantDownload.Target = gotDownload.Target
			//	}
			//}
			//assert.Equal(t, tc.wantDownload, gotDownload)
			var destination /*, text, title*/ string
			if link.Destination != nil {
				destination = *link.Destination
			}
			if tc.wantDestination != nil {
				if !strings.HasPrefix(destination, *tc.wantDestination) {
					t.Errorf("expected destination starting with %s, was %s", *tc.wantDestination, destination)
					return
				}
			} else {
				assert.Equal(t, tc.wantDestination, link.Destination)
			}
			//if link.Text != nil {
			//	text = *link.Text
			//}
			//if tc.wantText != nil {
			//	assert.Equal(t, *tc.wantText, text)
			//} else {
			//	assert.Equal(t, tc.wantText, link.Text)
			//}
			//if link.Title != nil {
			//	title = *link.Title
			//}
			//if tc.wantText != nil {
			//	assert.Equal(t, *tc.wantTitle, title)
			//} else {
			//	assert.Equal(t, tc.wantTitle, link.Title)
			//}
		})
	}
}

type fakeValidator struct{}

func (f *fakeValidator) ValidateLink(linkUrl *urls.URL, linkDestination, contentSourcePath string) bool {
	return true
}

type fakeDownload struct{}

func (f fakeDownload) Schedule(task *DownloadTask) error {
	return nil
}
