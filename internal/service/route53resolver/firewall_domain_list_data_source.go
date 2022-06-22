package route53resolver

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)

func DataSourceFirewallDomainList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallDomainListRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creator_request_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_owner_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modification_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallDomainListRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).Route53ResolverConn

	input := &route53resolver.GetFirewallDomainListInput{
		FirewallDomainListId: aws.String(d.Get("id").(string)),
	}

	output, err := conn.GetFirewallDomainList(input)

	if err != nil {
		return fmt.Errorf("error getting Route53 Firewall Domain List: %w", err)
	}

	if output == nil {
		return fmt.Errorf("no Route53 Firewall Domain List found matching criteria; try different search")
	}

	d.SetId(aws.StringValue(output.FirewallDomainList.Id))
	d.Set("arn", output.FirewallDomainList.Arn)
	d.Set("creation_time", output.FirewallDomainList.CreationTime)
	d.Set("creator_request_id", output.FirewallDomainList.CreatorRequestId)
	d.Set("domain_count", output.FirewallDomainList.DomainCount)
	d.Set("name", output.FirewallDomainList.Name)
	d.Set("managed_owner_name", output.FirewallDomainList.ManagedOwnerName)
	d.Set("modification_time", output.FirewallDomainList.ModificationTime)
	d.Set("status", output.FirewallDomainList.Status)
	d.Set("status_message", output.FirewallDomainList.StatusMessage)

	return nil
}
