// Copyright 2018 The go-exactonline AUTHORS. All rights reserved.
//
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
// Code generated by gen-services.go; DO NOT EDIT.
package mailbox

import "github.com/stack11/go-exactonline/api"

type service struct {
	client *api.Client
}

// MailboxService is responsible for communication with the Mailbox
// endpoints of the Exact Online API.
type MailboxService struct {
	client *api.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Endpoints available under this service
	DefaultMailbox         *DefaultMailboxEndpoint
	Mailboxes              *MailboxesEndpoint
	MailMessageAttachments *MailMessageAttachmentsEndpoint
	MailMessagesReceived   *MailMessagesReceivedEndpoint
	MailMessagesSent       *MailMessagesSentEndpoint
	PreferredMailbox       *PreferredMailboxEndpoint
}

// NewMailboxService creates a new initialized instance of the
// MailboxService.
func NewMailboxService(apiClient *api.Client) *MailboxService {
	s := &MailboxService{client: apiClient}

	s.common.client = apiClient

	s.DefaultMailbox = (*DefaultMailboxEndpoint)(&s.common)
	s.Mailboxes = (*MailboxesEndpoint)(&s.common)
	s.MailMessageAttachments = (*MailMessageAttachmentsEndpoint)(&s.common)
	s.MailMessagesReceived = (*MailMessagesReceivedEndpoint)(&s.common)
	s.MailMessagesSent = (*MailMessagesSentEndpoint)(&s.common)
	s.PreferredMailbox = (*PreferredMailboxEndpoint)(&s.common)

	return s
}
