package apprunner

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VpcIngressConnections() *schema.Table {
	tableName := "aws_apprunner_vpc_ingress_connections"
	return &schema.Table{
		Name: tableName,
		Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html

Notes:
- 'account_id' has been renamed to 'source_account_id' to avoid conflict with the 'account_id' column that indicates what account this was synced from.`,
		Resolver:            fetchApprunnerVpcIngressConnections,
		PreResourceResolver: getVpcIngressConnection,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "apprunner"),
		Transform:           transformers.TransformWithStruct(&types.VpcIngressConnection{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcIngressConnectionArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "source_account_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveApprunnerTags("VpcIngressConnectionArn"),
			},
		},
	}
}
