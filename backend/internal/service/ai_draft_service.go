package service

import "kcardDesgin/backend/internal/domain"

type AIDraftService struct{}

func (AIDraftService) Approve(draft domain.AIDraft, recordID domain.ID) domain.AIDraft {
	draft.Status = "approved"
	draft.Payload = map[string]any{"approvedRecordId": recordID}
	return draft
}
func (AIDraftService) Discard(draft domain.AIDraft) domain.AIDraft {
	draft.Status = "discarded"
	return draft
}
