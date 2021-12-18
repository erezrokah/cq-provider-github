package resources

import (
	"context"

	"github.com/cloudquery/cq-provider-github/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/google/go-github/v40/github"
)

func Repositories() *schema.Table {
	return &schema.Table{
		Name:     "repositories",
		Resolver: fetchRepositories,
		// Those are optional
		// DeleteFilter: nil,
		// Multiplex:    nil,
		// IgnoreError:  nil,
		//PostResourceResolver: nil,
		Options: schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "node_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("node_id"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "full_name",
				Type: schema.TypeString,
			},
			{
				Name: "default_branch",
				Type: schema.TypeString,
			},
			{
				Name: "description",
				Type: schema.TypeString,
			},
			{
				Name: "language",
				Type: schema.TypeString,
			},
			{
				Name: "forks_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "open_issues_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "stargazers_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "subscribe_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "watchers_count",
				Type: schema.TypeBigInt,
			},
			{
				Name: "size",
				Type: schema.TypeBigInt,
			},
			{
				Name: "allow_rebase_merge",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_squash_merge",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_merge_commit",
				Type: schema.TypeBool,
			},
			{
				Name: "allow_auto_merge",
				Type: schema.TypeBool,
			},
			{
				Name: "delete_branch_on_merge",
				Type: schema.TypeBool,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchRepositories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	opts := github.RepositoryListOptions{}
	repositories, response, err := c.GithubClient.Repositories.List(ctx, "cloudquery", &opts)
	if err != nil {
		return err
	}
	_ = response
	res <- repositories
	return nil
}
