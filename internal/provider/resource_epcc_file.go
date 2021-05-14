package provider

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccFile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccFileCreate,
		ReadContext:   resourceEpccFileRead,
		DeleteContext: resourceEpccFileDelete,
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
				Type:     schema.TypeBool,
				Description: "TBD - But remember the behaviour of this differs from the API, in that terraform ignores this setting if you specify it for file_location.",
				Optional: true,
				ForceNew: true,
			},
			"file_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}

}

func resourceEpccFileDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	fileID := d.Id()

	err := epcc.Files.Delete(client, fileID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccFileRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	fileId := d.Id()

	file, err := epcc.Files.Get(client, fileId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("file_name", file.Data.FileName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("mime_type", file.Data.MimeType); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("file_size", file.Data.FileSize); err != nil {
		return diag.FromErr(err)
	}

	if file.Data.Link != nil {
		if err := d.Set("file_link", file.Data.Link.Href); err != nil {
			return diag.FromErr(err)
		}
	}

	if _, fileLocationSet := d.GetOk("file_location"); !fileLocationSet {
		if err := d.Set("public", file.Data.Public); err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func resourceEpccFileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	_, fileSet := d.GetOk("file")
	_, fileLocationSet := d.GetOk("file_location")

	fileContentsBase64 := d.Get("file").(string)
	fileLocation := d.Get("file_location").(string)
	public := d.Get("public").(bool)

	if fileSet && fileLocationSet {
		return diag.FromErr(errors.New("Cannot specify file if file_location is specified"))
	}

	if fileSet {
		b, err := base64.StdEncoding.DecodeString(fileContentsBase64)
		if err != nil {
			return diag.FromErr(err)
		}

		createdFileData, apiError := epcc.Files.CreateFromFile(client, d.Get("file_name").(string), public, bytes.NewBuffer(b))

		if apiError != nil {
			return FromAPIError(apiError)
		}

		d.SetId(createdFileData.Data.Id)

		resourceEpccFileRead(ctx, d, m)

		return diags
	} else if fileLocationSet {
		createdFileData, apiError := epcc.Files.CreateFromFileLocation(client, fileLocation)

		if apiError != nil {
			return FromAPIError(apiError)
		}

		d.SetId(createdFileData.Data.Id)

		resourceEpccFileRead(ctx, d, m)

		return diags
	} else {
		return diag.FromErr(errors.New("You must specify a file location or a file"))
	}
}