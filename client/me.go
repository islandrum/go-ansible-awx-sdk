package awx

type MeService struct {
	client *Client
}

type MeResponse struct {
	Pagination
	Results []*User `json:"results"`
}

const meAPIEndpoint = "/api/v2/me/"

// /api/v2/me is a special case of user that returns a single user in the results
// and that user is the logged in user
func (u *MeService) GetMe(params map[string]string) (*User, error) {
	result := new(MeResponse)
	resp, err := u.client.Requester.GetJSON(meAPIEndpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result.Results[0], nil
}
