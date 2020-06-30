package linode

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/linode/linodego"
)

func setStackScriptUserDefinedFields(d *schema.ResourceData, ss *linodego.Stackscript) {
	if ss.UserDefinedFields == nil {
		d.Set("user_defined_fields", nil)
		return
	}

	udfs := []map[string]string{}
	for _, udf := range *ss.UserDefinedFields {
		udfs = append(udfs, map[string]string{
			"default": udf.Default,
			"example": udf.Example,
			"many_of": udf.ManyOf,
			"one_of":  udf.OneOf,
			"label":   udf.Label,
			"name":    udf.Name,
		})
	}
	d.Set("user_defined_fields", udfs)
}
