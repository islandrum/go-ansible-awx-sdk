package awx

import (
    "bytes"
    "encoding/json"
    "fmt"
)

type CredentialTypeService struct {
	client *Client
}

type ListCredentialTypeResponse struct {
	Pagination
	Results []*CredentialType `json:"results"`
}

const credentialTypesAPIEndpoint = "/api/v2/credential_types/"

func (cs *CredentialTypeService) ListCredentialsTypes(params map[string]string) ([]*CredentialType,
	*ListCredentialTypeResponse, error) {
	result := new(ListCredentialTypeResponse)
	resp, err := cs.client.Requester.GetJSON(credentialTypesAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	err = CheckResponse(resp)
	if err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}
func (cs *CredentialTypeService) CreateCredentialsTypes(data map[string]interface{}, params map[string]string) (*CredentialType, error) {
    result := new(CredentialType)
    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := cs.client.Requester.PostJSON(credentialTypesAPIEndpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    err = CheckResponse(resp)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (cs *CredentialTypeService) GetCredentialsTypesByID(id int, params map[string]string) (*CredentialType, error) {
    result := new(CredentialType)
    endpoint := fmt.Sprintf("%s%d", credentialTypesAPIEndpoint, id)
    resp, err := cs.client.Requester.GetJSON(endpoint, result, params)
    if err != nil {
        return nil, err
    }

    err = CheckResponse(resp)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (cs *CredentialTypeService) UpdateCredentialsTypesByID(id int, data map[string]interface{},
    params map[string]string) (*CredentialType, error) {
    result := new(CredentialType)
    endpoint := fmt.Sprintf("%s%d", credentialTypesAPIEndpoint, id)

    payload, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    resp, err := cs.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
    if err != nil {
        return nil, err
    }

    err = CheckResponse(resp)
    if err != nil {
        return nil, err
    }

    return result, nil
}

func (cs *CredentialTypeService) DeleteCredentialsTypesByID(id int, params map[string]string) error {
    endpoint := fmt.Sprintf("%s%d", credentialTypesAPIEndpoint, id)
    resp, err := cs.client.Requester.Delete(endpoint, nil, params)
    if err != nil {
        return err
    }

    err = CheckResponse(resp)
    if err != nil {
        return err
    }

    return nil
}
