package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

// Define the schema for the custom ticket
func TicketSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Unique identifier for the ticket",
		},
		"ticketno": {
			Type:        schema.TypeInt,
			Optional:    true,
			Description: "Ticket number",
		},
		"title": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Ticket title",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Ticket description",
		},
		"status": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Current status of the ticket",
		},
		"substatus": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Sub-status of the ticket",
		},
		"statuschangedat": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Timestamp when the status was last changed",
		},
		"createdat": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Timestamp when the ticket was created",
		},
		"createdby": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Details of the user who created the ticket",
			Elem:        userschema(),
		},
		"changedby": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Details of the user who last changed the ticket",
			Elem:        userschema(),
		},
		"claritycode": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Clarity code details",
			Elem:        claritycodeschema(),
		},
		"participants": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of participants in the ticket",
			Elem:        participantschema(),
		},
		"comments": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of comments on the ticket",
			Elem:        commentschema(),
		},
		"attachments": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of attachments for the ticket",
			Elem:        attachmentschema(),
		},
		"billingitems": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of billing items related to the ticket",
			Elem:        billingitemschema(),
		},
		"historyitems": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "History of changes to the ticket",
			Elem:        historyitemschema(),
		},
		"validactions": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "List of valid actions that can be performed on the ticket",
			Elem:        actionschema(),
		},
		"editableproperties": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of editable properties of the ticket",
			Elem:        schema.TypeString,
		},
		"mandatoryproperties": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "List of mandatory properties for the ticket",
			Elem:        schema.TypeString,
		},
		"etag": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ETag for the ticket",
		},
		"type": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Ticket type",
		},
		"serviceprovider": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Service provider name",
		},
		"cloudplatform": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Cloud platform for the ticket",
		},
		"catalogitems": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Catalog items associated with the ticket",
			Elem:        catalogitemschema(),
		},
	}
}

// Define the schema for the user object
func userschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's email address",
			},
			"userprincipalname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User principal name",
			},
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User ID",
			},
			"displayname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's display name",
			},
			"roles": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Roles of the user",
				Elem:        schema.TypeString,
			},
		},
	}
}

// Define the schema for the clarityCode object
func claritycodeschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Clarity code",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the clarity code",
			},
			"costcenter": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cost center for the clarity code",
			},
			"emails": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of emails related to the clarity code",
				Elem:        schema.TypeString,
			},
			"tower": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Tower associated with the clarity code",
			},
		},
	}
}

func participantschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"userinfo": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     userschema(),
			},
			"role": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Role of the participant",
			},
		},
	}
}

func commentschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Comment ID",
			},
			"createdat": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Comment creation timestamp",
			},
			"modifiedat": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Comment modification timestamp",
			},
			"author": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     userschema(),
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Comment content",
			},
			"loginuser": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     userschema(),
			},
			"iseditable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the comment is editable",
			},
			"iseditmode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the comment is in edit mode",
			},
			"contentcopy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A copy of the content",
			},
		},
	}
}

func attachmentschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of the attachment",
			},
			"uploaddatetime": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Upload timestamp of the attachment",
			},
			"uploadedby": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     userschema(),
			},
			"filename": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the attachment file",
			},
		},
	}
}

func billingitemschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Billing item ID",
			},
			"partitionkey": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Partition key for billing",
			},
			"subscriptionname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Subscription name",
			},
			"invoiceperiods": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Invoice periods for the billing item",
				Elem:        invoiceperiodschema(),
			},
		},
	}
}

func historyitemschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"date": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Date of history item",
			},
			"author": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     userschema(),
			},
			"changes": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of changes made to the ticket",
				Elem:        changeschema(),
			},
		},
	}
}

func changeschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"propertyname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the changed property",
			},
			"oldvalue": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Old value of the changed property",
				Elem:        schema.TypeString,
			},
			"newvalue": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "New value of the changed property",
				Elem:        schema.TypeString,
			},
		},
	}
}

func actionschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"actionname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the action",
			},
			"requiredproperties": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of required properties for the action",
				Elem:        schema.TypeString,
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of action",
			},
			"minnumofcatalogitems": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Minimum number of catalog items for the action",
			},
		},
	}
}

// InvoicePeriodSchema defines the schema for each invoice period in a billing item
func invoiceperiodschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"invoiceperiod": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Invoice period for the item",
			},
			"actualcost": {
				Type:        schema.TypeFloat,
				Required:    true,
				Description: "Actual cost of the billing item for the period",
			},
			"startdate": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Start date of the billing period",
			},
			"enddate": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "End date of the billing period",
			},
		},
	}
}

func catalogitemschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the catalog item",
			},
			"resourcename": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource name associated with the catalog item",
			},
			"label": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Label for the catalog item",
			},
			"catalogitemdisclaimer": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Disclaimer associated with the catalog item",
			},
			"catalogitemcloudplatform": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cloud platform associated with the catalog item (e.g., Azure)",
			},
			"tickettypes": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of ticket types for this catalog item",
				Elem:        schema.TypeString,
			},
			"active": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Indicates if the catalog item is active",
			},
			"catalogitemversion": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Version of the catalog item",
			},
			"catalogitemcreated": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Creation timestamp of the catalog item",
			},
			"catalogitemapproved": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Approval timestamp of the catalog item",
			},
			"catalogitemapprovedby": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User who approved the catalog item",
			},
			"catalogitemicon": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Icon associated with the catalog item",
			},
			"catalogfields": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of catalog fields for the item",
				Elem:        catalogfieldschema(),
			},
			"variables": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Variables associated with the catalog item",
				Elem:        schema.TypeString,
			},
			"resourcecontractname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the resource contract associated with the catalog item",
			},
			"resourcecontainername": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the resource container for the catalog item",
			},
		},
	}
}

// CatalogFieldSchema defines the schema for each catalog field in a catalog item
func catalogfieldschema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Key for the catalog field",
			},
			"label": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Label for the catalog field",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value of the catalog field",
			},
			"ismandatory": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Indicates if the catalog field is mandatory",
			},
			"lookupfunction": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Look-up function for the catalog field",
			},
			"lookupvalues": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of possible look-up values for the catalog field",
				Elem:        schema.TypeString,
			},
			"hintvalue": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hint value for the catalog field",
			},
			"inputtype": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Input type for the catalog field (e.g., text, number)",
			},
			"inputformat": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Input format for the catalog field",
			},
			"enabletoggleby": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Field that enables or disables this catalog field based on its value",
			},
			"disabled": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Indicates if the field is disabled",
			},
		},
	}
}
