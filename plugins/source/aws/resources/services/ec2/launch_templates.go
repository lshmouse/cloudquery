package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LaunchTemplates() *schema.Table {
	tableName := "aws_ec2_launch_templates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplate.html`,
		Resolver:    fetchEc2LaunchTemplates,
		Transform:   transformers.TransformWithStruct(&types.LaunchTemplate{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ec2"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveEc2LaunchTemplateArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},

		Relations: []*schema.Table{
			LaunchTemplateVersions(),
		},
	}
}
