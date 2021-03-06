package provider

import (
	"context"

	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccProduct() *schema.Resource {
	return &schema.Resource{
		Description: "Allows the caller to look up details of an Elastic Path Commerce Cloud PCM [product](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html).",
		ReadContext: addDiagToContext(dataSourceEpccProductRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Description: "The unique identifier of the product.",
				Required:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The product name to display to customers.",
				Computed:    true,
			},
			"commodity_type": {
				Type:        schema.TypeString,
				Description: "The type of the product; either `physical` or `digital`.",
				Computed:    true,
			},
			"sku": {
				Type:        schema.TypeString,
				Description: "The unique _stock keeping unit_ of the product.",
				Computed:    true,
			},
			"slug": {
				Type:        schema.TypeString,
				Description: "The unique slug of the product.",
				Computed:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Description: "The product description to display to customers.",
				Computed:    true,
			},
			"mpn": {
				Type:        schema.TypeString,
				Description: "The _manufacturer part number_ of the product.",
				Computed:    true,
			},
			"status": {
				Type:        schema.TypeString,
				Description: "The status of the product; either `draft` or `live`. Default is `draft`.",
				Computed:    true,
			},
			"upc_ean": {
				Type:        schema.TypeString,
				Description: "The _universal product code_ or _european article number_ of the product.",
				Computed:    true,
			},
			"files": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceEpccProductRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	productId := d.Get("id").(string)

	product, err := epcc.Products.Get(&ctx, client, productId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	productFiles, err := epcc.Products.GetProductFiles(&ctx, client, productId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("name", product.Data.Attributes.Name)
	d.Set("commodity_type", product.Data.Attributes.CommodityType)
	d.Set("sku", product.Data.Attributes.Sku)
	d.Set("slug", product.Data.Attributes.Slug)
	d.Set("description", product.Data.Attributes.Description)
	d.Set("mpn", product.Data.Attributes.Mpn)
	d.Set("status", product.Data.Attributes.Status)
	d.Set("upc_ean", product.Data.Attributes.UpcEan)

	d.SetId(product.Data.Id)

	if productFiles != nil && productFiles.Data != nil {
		fileIds := convertJsonTypesToIds(productFiles.Data)

		if err := d.Set("files", fileIds); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	} else {
		if err := d.Set("files", [0]string{}); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	}
}
