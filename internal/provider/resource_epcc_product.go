package provider

import (
	"context"
	"fmt"

	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccProduct() *schema.Resource {
	return &schema.Resource{
		Description:   "Allows the caller to create, update or delete an Elastic Path Commerce Cloud PCM [product](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html).",
		CreateContext: resourceEpccProductCreate,
		ReadContext:   resourceEpccProductRead,
		UpdateContext: resourceEpccProductUpdate,
		DeleteContext: resourceEpccProductDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The unique identifier of the product.",
				Computed:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The product name to display to customers.",
				Required:    true,
			},
			"commodity_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Valid values: `physical` or `digital`.",
				Required:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					valueString := val.(string)
					if valueString != "physical" && valueString != "digital" {
						errs = append(errs, fmt.Errorf("%q must be either \"physical\" or \"digital\", but was set to: %q", key, valueString))
					}
					return
				},
			},
			"sku": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The unique _stock keeping unit_ of the product.",
				Required:    true,
			},
			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The unique slug of the product.",
				Optional:    true,
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The product description to display to customers.",
				Optional:    true,
			},
			"mpn": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The _manufacturer part number_ of the product.",
				Optional:    true,
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Valid values: `draft` or `live`. Default is `draft`.",
				Optional:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					valueString := val.(string)
					if valueString != "draft" && valueString != "live" {
						errs = append(errs, fmt.Errorf("%q must be either \"draft\" or \"live\", but was set to: %q", key, valueString))
					}
					return
				},
			},
			"upc_ean": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The _universal product code_ or _european article number_ of the product.",
				Optional:    true,
			},
		},
	}

}

func resourceEpccProductDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	productId := d.Id()

	err := epcc.Products.Delete(client, productId)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccProductUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	productId := d.Id()

	product := &epcc.Product{
		Id:   productId,
		Type: "product",
		Attributes: epcc.ProductAttributes{
			Name:          d.Get("name").(string),
			CommodityType: d.Get("commodity_type").(string),
			Sku:           d.Get("sku").(string),
			Slug:          d.Get("slug").(string),
			Description:   d.Get("description").(string),
			Mpn:           d.Get("mpn").(string),
			Status:        d.Get("status").(string),
			UpcEan:        d.Get("upc_ean").(string),
		},
	}

	updatedProductData, apiError := epcc.Products.Update(client, productId, product)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedProductData.Data.Id)

	return resourceEpccProductRead(ctx, d, m)
}

func resourceEpccProductRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	productId := d.Id()

	product, err := epcc.Products.Get(client, productId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", product.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", product.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEpccProductCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	product := &epcc.Product{
		Type: "product",
		Attributes: epcc.ProductAttributes{
			Name:          d.Get("name").(string),
			CommodityType: d.Get("commodity_type").(string),
			Sku:           d.Get("sku").(string),
			Slug:          d.Get("slug").(string),
			Description:   d.Get("description").(string),
			Mpn:           d.Get("mpn").(string),
			Status:        d.Get("status").(string),
			UpcEan:        d.Get("upc_ean").(string),
		},
	}

	createdProductData, apiError := epcc.Products.Create(client, product)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdProductData.Data.Id)

	resourceEpccProductRead(ctx, d, m)

	return diags
}
