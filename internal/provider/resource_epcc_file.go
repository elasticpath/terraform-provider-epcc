package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccFile() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [File Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/files/index.html#the-file-object).",
		CreateContext: addDiagToContext(resourceEpccFileCreate),
		ReadContext:   addDiagToContext(resourceEpccFileRead),
		DeleteContext: addDiagToContext(resourceEpccFileDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"file_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mime_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"file": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"public": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "TBD - But remember the behaviour of this differs from the API, in that terraform ignores this setting if you specify it for file_location.",
				Optional:    true,
				ForceNew:    true,
			},
			"file_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}

}

func resourceEpccFileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	fileID := d.Id()

	err := epcc.Files.Delete(&ctx, client, fileID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccFileRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	fileId := d.Id()

	file, err := epcc.Files.Get(&ctx, client, fileId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("file_name", file.Data.FileName); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("mime_type", file.Data.MimeType); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("file_size", file.Data.FileSize); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if file.Data.Link != nil {
		if err := d.Set("file_link", file.Data.Link.Href); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	}

	if _, fileLocationSet := d.GetOk("file_location"); !fileLocationSet {
		if err := d.Set("public", file.Data.Public); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	}
}

func resourceEpccFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	_, fileSet := d.GetOk("file")
	_, fileLocationSet := d.GetOk("file_location")

	fileContentsBase64 := d.Get("file").(string)
	fileLocation := d.Get("file_location").(string)
	public := d.Get("public").(bool)

	if fileSet && fileLocationSet {
		addToDiag(ctx, diag.FromErr(errors.New("Cannot specify file if file_location is specified")))
		return
	}

	if fileSet {
		b, err := base64.StdEncoding.DecodeString(fileContentsBase64)
		if err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		createdFileData, apiError := epcc.Files.CreateFromFile(&ctx, client, d.Get("file_name").(string), public, bytes.NewBuffer(b))

		if apiError != nil {
			ReportAPIError(ctx, apiError)
			return
		}

		d.SetId(createdFileData.Data.Id)

		resourceEpccFileRead(ctx, d, m)

	} else if fileLocationSet {
		createdFileData, apiError := epcc.Files.CreateFromFileLocation(&ctx, client, fileLocation)

		if apiError != nil {
			ReportAPIError(ctx, apiError)
			return
		}

		d.SetId(createdFileData.Data.Id)

		resourceEpccFileRead(ctx, d, m)

	} else {
		addToDiag(ctx, diag.FromErr(errors.New("You must specify a file location or a file")))
	}
}
