package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func addDiagToContext(f func(context.Context, *schema.ResourceData, interface{})) func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var contextDiags = new(diag.Diagnostics)
		f(context.WithValue(ctx, "diags", contextDiags), d, m)
		return *contextDiags
	}
}
