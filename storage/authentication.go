package storage

import (
	"fmt"
	"time"
)

// Get a new auth token for the storage client
func (c *StorageClient) getAuthenticationToken() error {
	var authHeaders map[string]string
	authHeaders = make(map[string]string)
	authHeaders["X-Storage-User"] = c.getUserName()
	authHeaders["X-Storage-Pass"] = *c.client.Password

	rsp, err := c.executeRequest("GET", "/auth/v1.0", authHeaders)
	if err != nil {
		return err
	}

	var authToken string
	if authToken = rsp.Header.Get("X-Auth-Token"); authToken == "" {
		return fmt.Errorf("No authentication token found in response %#v", rsp)
	}

	c.client.DebugLogString("Successfully authenticated to IaaS Storage")
	c.authToken = &authToken
	c.tokenIssued = time.Now()
	return nil
}
