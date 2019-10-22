package graphqlbackend

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/backend"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/internal/pkg/search"
	"github.com/sourcegraph/sourcegraph/internal/api"
)

type CombyQueryArgs struct {
	RepositoryNames []string
	MatchTemplate   string
	Rule            *string
	RewriteTemplate *string
}

func (schemaResolver) Comby(ctx context.Context, arg *struct {
	CombyQueryArgs
}) (combyPayload, error) {
	var allResults []combyResult
	for _, repoName := range arg.RepositoryNames {
		results, err := runComby(ctx, &arg.CombyQueryArgs, repoName)
		if err != nil {
			return nil, err
		}
		allResults = append(allResults, results...)
	}

	return combyPayload(allResults), nil
}

func runComby(ctx context.Context, arg *CombyQueryArgs, repoName string) ([]combyResult, error) {
	repo, err := backend.Repos.GetByName(ctx, api.RepoName(repoName))
	if err != nil {
		return nil, err
	}
	s := func(s *string) string {
		if s == nil {
			return ""
		}
		return *s
	}
	results, err := callCodemodInRepo(ctx, &search.RepositoryRevisions{
		Repo: repo,
		Revs: []search.RevisionSpecifier{{RevSpec: "HEAD"}},
	}, &args{
		matchTemplate:     arg.MatchTemplate,
		rewriteTemplate:   s(arg.RewriteTemplate),
		includeFileFilter: ".test.ts,.ts,.tsx,.go,.java,.gradle",
	})
	if err != nil {
		return nil, err
	}

	rs := make([]combyResult, len(results))
	for i, r := range results {
		rs[i] = combyResult{
			file: &gitTreeEntryResolver{
				commit: r.commit,
				stat:   createFileInfo(r.path, false),
			},
			rawDiff: r.diff,
		}
	}
	return rs, nil
}

type combyPayload []combyResult

func (p combyPayload) Results() []combyResult {
	return []combyResult(p)
}

type combyResult struct {
	file    *gitTreeEntryResolver
	rawDiff string
}

func (r combyResult) File() *gitTreeEntryResolver { return r.file }

func (r combyResult) RawDiff() *string {
	if r.rawDiff == "" {
		return nil
	}
	return &r.rawDiff
}