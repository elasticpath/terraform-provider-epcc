package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccFile() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccFileRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"file_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mime_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func dataSourceEpccFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)

	FileId := d.Get("id").(string)

	File, err := epcc.Files.Get(&ctx, client, FileId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("file_name", File.Data.FileName)

	if File.Data.Link != nil {
		d.Set("file_link", File.Data.Link.Href)
	}

	d.Set("mime_type", File.Data.MimeType)

	d.Set("file_size", File.Data.FileSize)

	d.Set("public", File.Data.Public)

	d.SetId(File.Data.Id)
}
