package provider

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"

	"encoding/json"
	"net/http"

	"github.com/terraform-provider-cloudportal/cloudportal/internal/logger"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Ticket struct {
	ID                  string        `json:"id"`                  // Unique identifier for the ticket.
	TicketNo            int           `json:"ticketno"`            // Ticket number.
	Title               string        `json:"title"`               // Ticket title.
	Description         string        `json:"description"`         // Ticket description.
	Status              string        `json:"status"`              // Current status of the ticket.
	SubStatus           string        `json:"substatus"`           // Sub-status of the ticket.
	StatusChangedAt     string        `json:"statuschangedat"`     // Timestamp when the status was last changed.
	CreatedAt           string        `json:"createdat"`           // Timestamp when the ticket was created.
	CreatedBy           User          `json:"createdby"`           // Details of the user who created the ticket.
	ChangedBy           User          `json:"changedby"`           // Details of the user who last changed the ticket.
	ClarityCode         ClarityCode   `json:"claritycode"`         // Clarity code details.
	Participants        []Participant `json:"participants"`        // List of participants in the ticket.
	Comments            []Comment     `json:"comments"`            // List of comments on the ticket.
	Attachments         []Attachment  `json:"attachments"`         // List of attachments for the ticket.
	BillingItems        []BillingItem `json:"billingitems"`        // List of billing items related to the ticket.
	HistoryItems        []HistoryItem `json:"historyitems"`        // History of changes to the ticket.
	ValidActions        []Action      `json:"validactions"`        // List of valid actions that can be performed on the ticket.
	EditableProperties  []string      `json:"editableproperties"`  // List of editable properties of the ticket.
	MandatoryProperties []string      `json:"mandatoryproperties"` // List of mandatory properties for the ticket.
	ETag                string        `json:"etag"`                // ETag for the ticket.
	Type                string        `json:"type"`                // Ticket type.
	ServiceProvider     string        `json:"serviceprovider"`     // Service provider name.
	CloudPlatform       string        `json:"cloudplatform"`       // Cloud platform for the ticket.
	CatalogItems        []CatalogItem `json:"catalogitems"`        // Catalog items associated with the ticket.
}

type User struct {
	ID                string   `json:"id"`
	Email             string   `json:"email"`
	UserPrincipalName string   `json:"userprincipalname"`
	DisplayName       string   `json:"displayname"`
	Roles             []string `json:"roles"`
}

type Participant struct {
	UserInfo User   `json:"userinfo"`
	Role     string `json:"role"`
}

type Comment struct {
	ID          string `json:"id"`
	Createdat   string `json:"createdat"`
	Modifiedat  string `json:"modifiedat"`
	Author      User   `json:"author"`
	Content     string `json:"content"`
	Loginuser   User   `json:"loginuser"`
	Iseditable  bool   `json:"iseditable"`
	Iseditmode  bool   `json:"IsEditMode"`
	Contentcopy string `json:"contentcopy"`
}

type Action struct {
	ActionName           string   `json:"actionname"`
	RequiredProperties   []string `json:"requiredproperties"`
	Type                 string   `json:"type"`
	MinNumOfCatalogItems int      `json:"minnumofcatalogitems"`
}

// InvoicePeriod represents a billing period and related details.
type InvoicePeriod struct {
	InvoicePeriod string  `json:"invoiceperiod"`
	ActualCost    float64 `json:"actualcost"`
	StartDate     string  `json:"startdate"`
	EndDate       string  `json:"enddate"`
}

// BillingItem represents a billing item with associated metadata.
type BillingItem struct {
	ID               string                   `json:"id"`
	PartitionKey     string                   `json:"partitionkey"`
	SubscriptionName string                   `json:"subscriptionname"`
	InvoicePeriods   map[string]InvoicePeriod `json:"invoiceperiods"`
}

// Change represents a single modification or update made to a ticket.
type Change struct {
	PropertyName string            `json:"propertyname"` // The name of the property that was changed (e.g., "status").
	OldValue     map[string]string `json:"oldvalue"`     // The old value of the property (before change).
	NewValue     map[string]string `json:"newvalue"`     // The new value of the property (after change).
}

type HistoryItem struct {
	Date    string   `json:"date"`    // The date when the history item was created.
	Author  []User   `json:"author"`  // The user who made the change.
	Changes []Change `json:"changes"` // List of changes that were made in this history item.
}

// CatalogItem represents a catalog item with all associated metadata.
type CatalogItem struct {
	Name                     string            `json:"name"`
	ResourceName             string            `json:"resourcename"`
	Label                    string            `json:"label"`
	CatalogItemDisclaimer    *string           `json:"catalogitemdisclaimer,omitempty"`
	CatalogItemCloudPlatform string            `json:"catalogitemcloudplatform"`
	TicketTypes              []string          `json:"tickettypes"`
	Active                   bool              `json:"active"`
	CatalogItemVersion       int               `json:"catalogitemversion"`
	CatalogItemCreated       string            `json:"catalogitemcreated"`
	CatalogItemApproved      string            `json:"catalogitemapproved"`
	CatalogItemApprovedBy    string            `json:"catalogitemapprovedby"`
	CatalogItemIcon          *string           `json:"catalogitemicon,omitempty"`
	CatalogFields            []CatalogField    `json:"catalogfields"`
	Variables                map[string]string `json:"variables"`
	ResourceContractName     *string           `json:"resourcecontractname,omitempty"`
	ResourceContainerName    *string           `json:"resourcecontainername,omitempty"`
}

// CatalogField represents a field in a catalog item with various attributes.
type CatalogField struct {
	Key            string   `json:"key"`
	Label          string   `json:"label"`
	Value          string   `json:"value"`
	IsMandatory    bool     `json:"ismandatory"`
	LookupFunction *string  `json:"lookupfunction,omitempty"`
	LookupValues   []string `json:"lookupvalues,omitempty"`
	HintValue      *string  `json:"hintvalue,omitempty"`
	InputType      *string  `json:"inputType,omitempty"`
	InputFormat    *string  `json:"inputformat,omitempty"`
	EnableToggleBy *string  `json:"enabletoggleby,omitempty"`
	Disabled       *string  `json:"disabled,omitempty"`
}

// Attachment represents the details of an attachment with metadata.
type Attachment struct {
	URL            string `json:"url"`
	UploadDateTime string `json:"uploaddatetime"`
	UploadedBy     []User `json:"uploadedby"`
	Filename       string `json:"filename"`
}

type ClarityCode struct {
	Code        string   `json:"code"`        // Clarity code
	Description string   `json:"description"` // Description of the clarity code
	CostCenter  string   `json:"costcenter"`  // Cost center for the clarity code
	Emails      []string `json:"emails"`      // List of emails related to the clarity code
	Tower       string   `json:"tower"`       // Tower associated with the clarity code
}

func dataSourceTicket() *schema.Resource {
	return &schema.Resource{
		Read:   dataSourceTicketRead,
		Schema: TicketSchema(), // Reuse the Ticket schema defined earlier

		// Ensure the 'id' is the only required field for querying the data source
		// In this case, the `ticket_id` is the identifier to fetch the ticket.
	}
}

// dataSourceTicketRead function is responsible for reading the ticket from the API
func dataSourceTicketRead(d *schema.ResourceData, meta interface{}) error {
	cred := meta.(*CloudportalAPIClient)

	if cred.isdebug {
		// Create a new logger with debug enabled
		// Initialize the logger once, using debugEnabled=true
		_, err := logger.NewLogger(true)
		if err != nil {
			log.Fatal("Error initializing logger:", err)
		}
		defer logger.Close()
	}

	ticketID := d.Get("id").(string)
	//ticketurl := "https://demand-module-dev.azurewebsites.net/api"
	// Example: API call to fetch the ticket details by ID
	// Replace this with your actual API client logic

	// Construct the URL for fetching the ticket details
	url := fmt.Sprintf("%s/ticket/%s", cred.BaseURL, ticketID)

	logger.Debug(url)

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Error(err.Error())
		return fmt.Errorf("failed to create HTTP request: %s", err)
	}

	// Step 2: Prepare token request options

	tokenRequestOptions := policy.TokenRequestOptions{
		Scopes: []string{cred.tenantID + "/.default"}, // Use the required scope for Azure management API Global.Appl.GoogleCloudPlatform.X
	}

	// Step 3: Get the access token
	token, err := cred.aziclient.GetToken(context.Background(), tokenRequestOptions)
	if err != nil {
		logger.Error(err.Error())
		log.Fatalf("failed to obtain a token: %v", err)
	}
	logger.Debug("Token : " + token.Token)
	// Set custom headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-IN,en-GB;q=0.9,en;q=0.8,en-US;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	// Set the API key in the Authorization header (replace with your actual method of authentication)
	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.APIKey))

	// Create a new HTTP client
	//client := &http.Client{}

	// Send the request using the HTTP client
	resp, err := cred.Client.Do(req)
	if err != nil {
		logger.Error("Send request : " + err.Error())
		return err
	}
	defer resp.Body.Close()

	logger.Debug(resp.Status)

	// Check if the response status is OK (200)
	if resp.StatusCode != http.StatusOK {
		logger.Error("Response status : " + resp.Status)
		logger.Error("Response status code : " + string(resp.StatusCode))
		return fmt.Errorf("API call failed with status %d: %s", resp.StatusCode, resp.Status)
	}

	// Check if the response is gzip encoded
	var reader io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		// Create a new gzip reader to decompress the content
		gzipReader, err := gzip.NewReader(reader)
		if err != nil {
			logger.Error(err.Error())
		}
		defer gzipReader.Close()
		reader = gzipReader
	}

	// Read the decompressed body
	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		logger.Error(err.Error())
	}

	// Print the raw response for debugging (you can remove this in production)
	logger.Debug(string(bodyBytes))

	// If the response is JSON, we can unmarshal it into a Go struct
	var ticket Ticket //interface{} // You can replace `interface{}` with a custom struct based on the JSON structure
	err = json.Unmarshal(bodyBytes, &ticket)
	if err != nil {
		logger.Error(err.Error())
	}

	// Set values to the Terraform resource schema
	d.Set("id", ticket.ID)
	d.Set("ticketno", ticket.TicketNo)
	d.Set("title", ticket.Title)
	d.Set("description", ticket.Description)
	d.Set("status", ticket.Status)
	d.Set("substatus", ticket.SubStatus)
	d.Set("statuschangedat", ticket.StatusChangedAt)
	d.Set("createdat", ticket.CreatedAt)
	d.Set("createdby", ticket.CreatedBy)
	d.Set("changedby", ticket.ChangedBy)

	// Optional attributes
	d.Set("participants", ticket.Participants)
	d.Set("comments", flattenComments(ticket.Comments))
	d.Set("attachments", flattenAttachments(ticket.Attachments))
	d.Set("billingitems", flattenBillingItems(ticket.BillingItems))
	d.Set("historyitems", flattenHistoryItems(ticket.HistoryItems))
	d.Set("validactions", flattenActions(ticket.ValidActions))

	// Mark the resource as read and set its ID
	d.SetId(ticket.ID)

	return nil
}

// Helper function to flatten a list of user objects
func flattenUsers(users []User) []interface{} {
	var result []interface{}
	for _, user := range users {
		result = append(result, map[string]interface{}{
			"email":             user.Email,
			"userprincipalname": user.UserPrincipalName,
			"id":                user.ID,
			"displayname":       user.DisplayName,
			"roles":             user.Roles,
		})
	}
	return result
}

/*
// Helper function to flatten participants
func flattenParticipants(participants []Participant) []interface{} {
	var result []interface{}
	for _, participant := range participants {
		result = append(result, map[string]interface{}{
			"userinfo": participant.UserInfo,
			"role":     participant.Role,
		})
	}
	return result
}*/

// Helper function to flatten comments
func flattenComments(comments []Comment) []interface{} {
	var result []interface{}
	for _, comment := range comments {
		result = append(result, map[string]interface{}{
			"id":          comment.ID,
			"createdat":   comment.Createdat,
			"modifiedat":  comment.Modifiedat,
			"author":      comment.Author,
			"content":     comment.Content,
			"loginuser":   comment.Loginuser,
			"iseditable":  comment.Iseditable,
			"iseditmode":  comment.Iseditmode,
			"contentcopy": comment.Contentcopy,
		})
	}
	return result
}

// Helper function to flatten attachments
func flattenAttachments(attachments []Attachment) []interface{} {
	var result []interface{}
	for _, attachment := range attachments {
		result = append(result, map[string]interface{}{
			"url":            attachment.URL,
			"uploaddatetime": attachment.UploadDateTime,
			"uploadedby":     flattenUsers(attachment.UploadedBy),
			"filename":       attachment.Filename,
		})
	}
	return result
}

// Helper function to flatten billing items
func flattenBillingItems(billingItems []BillingItem) []interface{} {
	var result []interface{}
	for _, item := range billingItems {
		result = append(result, map[string]interface{}{
			"id":               item.ID,
			"partitionkey":     item.PartitionKey,
			"subscriptionname": item.SubscriptionName,
			"invoiceperiods":   flattenInvoicePeriods(item.InvoicePeriods),
		})
	}
	return result
}

// Helper function to flatten invoice periods
func flattenInvoicePeriods(invoicePeriods map[string]InvoicePeriod) []interface{} {
	var result []interface{}
	for key, period := range invoicePeriods {
		result = append(result, map[string]interface{}{
			"invoiceperiod": key,
			"actualcost":    period.ActualCost,
			"startdate":     period.StartDate,
			"enddate":       period.EndDate,
		})
	}
	return result
}

// Helper function to flatten history items
func flattenHistoryItems(historyItems []HistoryItem) []interface{} {
	var result []interface{}
	for _, historyItem := range historyItems {
		result = append(result, map[string]interface{}{
			"date":    historyItem.Date,
			"author":  flattenUsers(historyItem.Author),
			"changes": flattenChanges(historyItem.Changes),
		})
	}
	return result
}

// Helper function to flatten changes
func flattenChanges(changes []Change) []interface{} {
	var result []interface{}
	for _, change := range changes {
		result = append(result, map[string]interface{}{
			"propertyname": change.PropertyName,
			"oldvalue":     flattenStringMap(change.OldValue),
			"newvalue":     flattenStringMap(change.NewValue),
		})
	}
	return result
}

// Helper function to flatten string maps (oldvalue/newvalue can be maps of strings)
func flattenStringMap(m map[string]string) map[string]interface{} {
	flattened := make(map[string]interface{})
	for k, v := range m {
		flattened[k] = v
	}
	return flattened
}

// Helper function to flatten actions
func flattenActions(actions []Action) []interface{} {
	var result []interface{}
	for _, action := range actions {
		result = append(result, map[string]interface{}{
			"actionname":           action.ActionName,
			"requiredproperties":   flattenStringList(action.RequiredProperties),
			"type":                 action.Type,
			"minnumofcatalogitems": action.MinNumOfCatalogItems,
		})
	}
	return result
}

// Helper function to flatten string lists (for requiredproperties field)
func flattenStringList(list []string) []interface{} {
	var result []interface{}
	for _, item := range list {
		result = append(result, item)
	}
	return result
}
