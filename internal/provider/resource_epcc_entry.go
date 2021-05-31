package provider

import (
	"context"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EntryResourceProvider struct {
}

func (p EntryResourceProvider) Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Entry Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/entries/index.html).",
		CreateContext: addDiagToContext(p.create),
		ReadContext:   addDiagToContext(p.read),
		UpdateContext: addDiagToContext(p.update),
		DeleteContext: addDiagToContext(p.delete),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_id": {
				Description: "Target core object identifier (can only be used for core flows)",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"strings": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"numbers": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"booleans": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
		},
	}
}

func (p EntryResourceProvider) create(ctx context.Context, d *schema.ResourceData, m interface{}) {
	if id, ok := d.GetOk("target_id"); ok {
		if err := checkCoreFlow(d); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		d.SetId(id.(string))
		p.update(ctx, d, m)
		return
	}
	client := m.(*epcc.Client)

	entry := &epcc.Entry{
		Type:     "entry",
		Strings:  convertMapToStringMap(d.Get("strings").(map[string]interface{})),
		Numbers:  convertMapToFloatMap(d.Get("numbers").(map[string]interface{})),
		Booleans: convertMapToBooleanMap(d.Get("booleans").(map[string]interface{})),
	}

	flowSlug := d.Get("slug").(string)

	created, apiError := epcc.Entries.Create(&ctx, client, flowSlug, entry)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(created.Data.Id)
}

func checkCoreFlow(d *schema.ResourceData) error {
	slug := d.Get("slug").(string)
	coreFlows := []string{
		"addresses",
		"products",
		"brands",
		"collections",
		"categories",
		"customers",
		"carts",
		"cart_items",
		"files",
		"orders",
		"order_items",
		"promotions",
	}
	for _, coreFlow := range coreFlows {
		if coreFlow == slug {
			return nil
		}
	}
	return fmt.Errorf("slug %v does not correspond to one of core flows: %v", slug, coreFlows)
}

func (p EntryResourceProvider) update(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	flowSlug := d.Get("slug").(string)
	entry := &epcc.Entry{
		Id:       d.Id(),
		Type:     "entry",
		Strings:  convertMapToStringMap(d.Get("strings").(map[string]interface{})),
		Numbers:  convertMapToFloatMap(d.Get("numbers").(map[string]interface{})),
		Booleans: convertMapToBooleanMap(d.Get("booleans").(map[string]interface{})),
	}

	_, err := epcc.Entries.Update(&ctx, client, flowSlug, entry)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	p.read(ctx, d, m)
}

func (p EntryResourceProvider) delete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	flowSlug := d.Get("slug").(string)
	entryID := d.Id()

	err := epcc.Entries.Delete(&ctx, client, flowSlug, entryID)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.SetId("")
}

func (p EntryResourceProvider) read(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	flowSlug := d.Get("slug").(string)
	entryID := d.Id()

	entry, err := epcc.Entries.Get(&ctx, client, flowSlug, entryID)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("strings", entry.Data.Strings); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("numbers", entry.Data.Numbers); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("booleans", entry.Data.Booleans); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}
