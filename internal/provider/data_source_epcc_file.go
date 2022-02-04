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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for the file.",
			},
			"file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of the file",
			},
			"file_link": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A link to the file",
			},
			"mime_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The MIME type of the file",
			},
			"file_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size of the file",
			},
			"public": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether the file is public.",
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
