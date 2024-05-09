package awx
import (
"bytes"
"encoding/json"
"fmt"
)

// InventoryScriptsService implements awx inventory scripts apis.
type InventoryScriptsService struct {
    client *Client
}

// ListInventoryScriptsResponse represents `ListInventoryScripts` endpoint response.
type ListInventoryScriptsResponse struct {
    Pagination
    Results []*InventoryScript `json:"results"`
}

const inventoryScriptsAPIEndpoint = "/api/v2/inventory_scripts/"

// GetInventoryScriptByID shows the details of a awx inventroy Scripts.
func (i *InventoryScriptsService) GetInventoryScriptByID(id int, params map[string]string) (*InventoryScript, error) {
    result := new(InventoryScript)
    endpoint := fmt.Sprintf("%s%d/", inventoryScriptsAPIEndpoint, id)
    resp, err := i.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// ListInventoryScripts shows list of awx inventories.
func (i *InventoryScriptsService) ListInventoryScripts(params map[string]string) ([]*InventoryScript, *ListInventoryScriptsResponse, error) {
    result := new(ListInventoryScriptsResponse)
    resp, err := i.client.Requester.GetJSON(inventoryScriptsAPIEndpoint, result, params)
    if err != nil {
        return nil, result, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, result, err
    }

    return result.Results, result, nil
}

// CreateInventoryScript creates an awx InventoryScript.
func (i *InventoryScriptsService) CreateInventoryScript(data map[string]interface{}, params map[string]string) (*InventoryScript, error) {
    mandatoryFields = []string{"name", "organization"}
    validate, status := ValidateParams(data, mandatoryFields)

    if !status {
        err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
        return nil, err
    }

    result := new(InventoryScript)
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    // Add check if InventoryScript exists and return proper error

    resp, err := i.client.Requester.PostJSON(inventoryScriptsAPIEndpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// UpdateInventoryScript update an awx InventoryScript
func (i *InventoryScriptsService) UpdateInventoryScript(id int, data map[string]interface{}, params map[string]string) (*InventoryScript, error) {
    result := new(InventoryScript)
    endpoint := fmt.Sprintf("%s%d", inventoryScriptsAPIEndpoint, id)
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }
    resp, err := i.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, nil)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// GetInventoryScript retrives the InventoryScript information from its ID or Name
func (i *InventoryScriptsService) GetInventoryScript(id int, params map[string]string) (*InventoryScript, error) {
    endpoint := fmt.Sprintf("%s%d", inventoryScriptsAPIEndpoint, id)
    result := new(InventoryScript)
    resp, err := i.client.Requester.GetJSON(endpoint, result, map[string]string{})
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

// DeleteInventoryScript delete an InventoryScript from AWX
func (i *InventoryScriptsService) DeleteInventoryScript(id int) (*InventoryScript, error) {
    result := new(InventoryScript)
    endpoint := fmt.Sprintf("%s%d", inventoryScriptsAPIEndpoint, id)

    resp, err := i.client.Requester.Delete(endpoint, result, nil)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result, nil
}

