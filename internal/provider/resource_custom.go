package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceCustom is a function that defines a custom resource for Terraform
func ResourceCustom() *schema.Resource {
	return &schema.Resource{
		//Create: resourceCustomCreate,
		Read: resourceCustomRead,
		//Update: resourceCustomUpdate,
		//Delete: resourceCustomDelete,
		Schema: TicketSchema(),
	}
}

/*
func resourceCustomCreate(d *schema.ResourceData, m interface{}) error {
	// Resource creation logic
	d.SetId("custom_" + d.Get("name").(string))
	return resourceCustomRead(d, m)
}

// Read function for the custom resource
func resourceCustomRead(d *schema.ResourceData, m interface{}) error {
	// Read data from the external resource or state
	log.Printf("[DEBUG] Reading custom resource: %s", d.Id())

	// For demonstration, setting the name and value from state
	d.Set("name", "custom_name_from_state")
	d.Set("value", "custom_value_from_state")

	return nil
}
*/

// resourceExampleRead reads the state of a resource from the custom API
func resourceCustomRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*CloudportalAPIClient)
	resourceID := d.Id()

	// Example API call to read a resource
	url := fmt.Sprintf("%s/ticket/%s", client.BaseURL, resourceID)
	resp, err := client.Client.Get(url)
	if err != nil {
		return fmt.Errorf("error reading resource: %s", err)
	}
	defer resp.Body.Close()

	// Example: If the resource exists, update the state
	//d.Set("name", "Example Resource Name") // Set the actual resource fields from the response
	return nil
}

/*
func resourceCustomUpdate(d *schema.ResourceData, m interface{}) error {
	// Resource update logic
	return resourceCustomRead(d, m)
}

func resourceCustomDelete(d *schema.ResourceData, m interface{}) error {
	// Resource delete logic
	d.SetId("")
	return nil
}
*/
