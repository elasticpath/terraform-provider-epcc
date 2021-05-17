package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func addDiagToContext(a func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics) func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var contextDiags = new(diag.Diagnostics)
		returnedDiag := a(context.WithValue(ctx, "diags", contextDiags), d, m)
		return append(returnedDiag, *contextDiags...)
	}
}
