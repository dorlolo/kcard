package service

import (
	"kcardDesgin/backend/internal/domain"
	"time"
)

type ReviewSessionService struct{ Scheduler ReviewScheduler }

func (s ReviewSessionService) Start(id domain.ID, workspaceID domain.ID, deckID domain.ID) domain.ReviewSession {
	return domain.ReviewSession{ID: id, LearnerWorkspaceID: workspaceID, DeckID: deckID, SessionType: "direct", Status: "active", StartedAt: time.Now().UTC()}
}
func (s ReviewSessionService) Answer(sessionID domain.ID, cardID domain.ID, result string) domain.ReviewResult {
	now := time.Now().UTC()
	return domain.ReviewResult{ID: sessionID + ":" + cardID, ReviewSessionID: sessionID, CardID: cardID, Result: result, ReviewedAt: now, NextDueAt: s.Scheduler.NextDue(result, now)}
}
