package tftime

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceTimeStatic() *schema.Resource {
	return &schema.Resource{
		Create: resourceTimeStaticCreate,
		Read:   resourceTimeStaticRead,
		Delete: schema.Noop,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"day": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hour": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"keepers": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"minute": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"month": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"rfc822": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc822z": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc850": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc1123": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc1123z": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rfc3339": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsRFC3339Time,
			},
			"second": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"unix": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"unixdate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"year": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceTimeStaticCreate(d *schema.ResourceData, m interface{}) error {
	timestamp := time.Now().UTC()

	if v, ok := d.GetOk("rfc3339"); ok {
		var err error
		timestamp, err = time.Parse(time.RFC3339, v.(string))

		if err != nil {
			return fmt.Errorf("error parsing rfc3339 (%s): %s", v.(string), err)
		}
	}

	d.SetId(timestamp.Format(time.RFC3339))

	return resourceTimeStaticRead(d, m)
}

func resourceTimeStaticRead(d *schema.ResourceData, m interface{}) error {
	timestamp, err := time.Parse(time.RFC3339, d.Id())

	if err != nil {
		return fmt.Errorf("error parsing timestamp (%s): %s", d.Id(), err)
	}

	if err := d.Set("day", timestamp.Day()); err != nil {
		return fmt.Errorf("error setting day: %s", err)
	}

	if err := d.Set("hour", timestamp.Hour()); err != nil {
		return fmt.Errorf("error setting hour: %s", err)
	}

	if err := d.Set("minute", timestamp.Minute()); err != nil {
		return fmt.Errorf("error setting minute: %s", err)
	}

	if err := d.Set("month", int(timestamp.Month())); err != nil {
		return fmt.Errorf("error setting month: %s", err)
	}

	if err := d.Set("rfc822", timestamp.Format(time.RFC822)); err != nil {
		return fmt.Errorf("error setting rfc822: %s", err)
	}

	if err := d.Set("rfc822z", timestamp.Format(time.RFC822Z)); err != nil {
		return fmt.Errorf("error setting rfc822z: %s", err)
	}

	if err := d.Set("rfc850", timestamp.Format(time.RFC850)); err != nil {
		return fmt.Errorf("error setting rfc850: %s", err)
	}

	if err := d.Set("rfc1123", timestamp.Format(time.RFC1123)); err != nil {
		return fmt.Errorf("error setting rfc1123: %s", err)
	}

	if err := d.Set("rfc1123z", timestamp.Format(time.RFC1123Z)); err != nil {
		return fmt.Errorf("error setting rfc1123z: %s", err)
	}

	if err := d.Set("rfc3339", timestamp.Format(time.RFC3339)); err != nil {
		return fmt.Errorf("error setting rfc3339: %s", err)
	}

	if err := d.Set("second", timestamp.Second()); err != nil {
		return fmt.Errorf("error setting second: %s", err)
	}

	if err := d.Set("unix", timestamp.Unix()); err != nil {
		return fmt.Errorf("error setting unix: %s", err)
	}

	if err := d.Set("unixdate", timestamp.Format(time.UnixDate)); err != nil {
		return fmt.Errorf("error setting unixdate: %s", err)
	}

	if err := d.Set("year", timestamp.Year()); err != nil {
		return fmt.Errorf("error setting year: %s", err)
	}

	return nil
}