// Copyright (c) 2019, The Fonero developers
// See LICENSE for details.

package types

import piapi "github.com/fonero-project/politeia/politeiawww/api/www/v1"

// ProposalInfo holds the proposal details as document here
// https://github.comfonero-project/politeia/blob/master/politeiawww/api/www/v1/api.md#user-proposals.
// It also holds the votes status details. The ID field is auto incremented by
// the db.
type ProposalInfo struct {
	ID              int                `json:"id" storm:"id,increment"`
	Name            string             `json:"name"`
	State           ProposalStateType  `json:"state"`
	Status          ProposalStatusType `json:"status"`
	Timestamp       uint64             `json:"timestamp"`
	UserID          string             `json:"userid"`
	Username        string             `json:"username"`
	PublicKey       string             `json:"publickey"`
	Signature       string             `json:"signature"`
	Version         string             `json:"version"`
	NumComments     int32              `json:"numcomments"`
	StatusChangeMsg string             `json:"statuschangemessage"`
	PublishedDate   uint64             `json:"publishedat" storm:"index"`
	CensoredDate    uint64             `json:"censoredat"`
	AbandonedDate   uint64             `json:"abandonedat"`
	// RefID was added to create an easily readable part of the URL that helps
	// to reference the proposals details page. Storm db ignores entries with
	// duplicate pk but returns "ErrAlreadyExists" error if the field other than
	// the pk has the tag "unique".
	RefID            string `storm:"unique"`
	CensorshipRecord `json:"censorshiprecord"`
	ProposalVotes    `json:"votes"`
	// Files           []AttachmentFile   `json:"files"`
}

// Proposals defines an array of proposals payload as returned by RouteAllVetted route.
type Proposals struct {
	Data []*ProposalInfo `json:"proposals"`
}

// Proposal defines a proposal payload as returned by RouteProposalDetails route.
type Proposal struct {
	Data *ProposalInfo `json:"proposal"`
}

// CensorshipRecord is an entry that was created when the proposal was submitted.
// https://github.com/fonero-project/politeia/blob/master/politeiawww/api/www/v1/api.md#censorship-record
type CensorshipRecord struct {
	TokenVal   string `json:"token" storm:"unique"`
	MerkleRoot string `json:"merkle"`
	Signature  string `json:"signature"`
}

// AttachmentFile are files and documents submitted as proposal details.
// https://github.com/fonero-project/politeia/blob/master/politeiawww/api/www/v1/api.md#file
type AttachmentFile struct {
	Name      string `json:"name"`
	MimeType  string `json:"mime"`
	DigestKey string `json:"digest"`
	Payload   string `json:"payload"`
}

// ProposalVotes defines the proposal status(Votes infor for the public proposals).
// https://github.com/fonero-project/politeia/blob/master/politeiawww/api/www/v1/api.md#proposal-vote-status
type ProposalVotes struct {
	Token              string         `json:"token"`
	VoteStatus         VoteStatusType `json:"status"`
	VoteResults        []Results      `json:"optionsresult"`
	TotalVotes         int64          `json:"totalvotes"`
	Endheight          string         `json:"endheight"`
	NumOfEligibleVotes int64          `json:"numofeligiblevotes"`
	QuorumPercentage   uint32         `json:"quorumpercentage"`
	PassPercentage     uint32         `json:"passpercentage"`
}

// Votes defines a slice of VotesStatuses as returned by RouteAllVoteStatus.
type Votes struct {
	Data []ProposalVotes `json:"votesstatus"`
}

// Results defines the actual vote count info per the votes choices available.
type Results struct {
	Option        VoteOption `json:"option"`
	VotesReceived int64      `json:"votesreceived"`
}

// VoteOption defines the actual high level vote results for the specific agenda.
type VoteOption struct {
	OptionID    string `json:"id"`
	Description string `json:"description"`
	Bits        int32  `json:"bits"`
}

// ProposalStatusType defines the various proposal statuses available as referenced
// in https://github.com/fonero-project/politeia/blob/master/politeiawww/api/www/v1/v1.go
type ProposalStatusType piapi.PropStatusT

func (p ProposalStatusType) String() string {
	return piapi.PropStatus[piapi.PropStatusT(p)]
}

// VoteStatusType defines the various vote statuses available as referenced in
// https://github.com/fonero-project/politeia/blob/master/politeiawww/api/www/v1/v1.go
type VoteStatusType piapi.PropVoteStatusT

// ShorterDesc maps the short description to there respective vote status type.
var ShorterDesc = map[piapi.PropVoteStatusT]string{
	piapi.PropVoteStatusInvalid:       "Invalid",
	piapi.PropVoteStatusNotAuthorized: "Not Authorized",
	piapi.PropVoteStatusAuthorized:    "Authorized",
	piapi.PropVoteStatusStarted:       "Started",
	piapi.PropVoteStatusFinished:      "Finished",
	piapi.PropVoteStatusDoesntExist:   "Doesn't Exist",
}

// ShortDesc returns the shorter vote status description.
func (s VoteStatusType) ShortDesc() string {
	return ShorterDesc[piapi.PropVoteStatusT(s)]
}

// LongDesc returns the long vote status description.
func (s VoteStatusType) LongDesc() string {
	return piapi.PropVoteStatus[piapi.PropVoteStatusT(s)]
}

// ProposalStateType defines the proposal state entry.
type ProposalStateType int8

const (
	// InvalidState defines the invalid state proposals.
	InvalidState ProposalStateType = iota

	// UnvettedState defines the unvetted state proposals and includes proposals
	// with a status of:
	//   * PropStatusNotReviewed
	//   * PropStatusUnreviewedChanges
	//   * PropStatusCensored
	UnvettedState

	// VettedState defines the vetted state proposals and includes proposals
	// with a status of:
	//   * PropStatusPublic
	//   * PropStatusAbandoned
	VettedState

	// UnknownState indicates a proposal state that is unrecognized.
	UnknownState
)

func (f ProposalStateType) String() string {
	switch f {
	case InvalidState:
		return "invalid"
	case UnvettedState:
		return "unvetted"
	case VettedState:
		return "vetted"
	default:
		return "unknown"
	}
}

// ProposalStateFromStr converts the string into ProposalStateType value.
func ProposalStateFromStr(val string) ProposalStateType {
	switch val {
	case "invalid":
		return InvalidState
	case "unvetted":
		return UnvettedState
	case "vetted":
		return VettedState
	default:
		return UnknownState
	}
}

// VotesStatuses returns the ShorterDesc map contents exclusive of Invalid and
// Doesn't exist statuses.
func VotesStatuses() map[VoteStatusType]string {
	m := make(map[VoteStatusType]string)
	for k, val := range ShorterDesc {
		if k == piapi.PropVoteStatusInvalid ||
			k == piapi.PropVoteStatusDoesntExist {
			continue
		}
		m[VoteStatusType(k)] = val
	}
	return m
}

// IsEqual compares CensorshipRecord, Name, State, NumComments, StatusChangeMsg,
// Timestamp, CensoredDate, AbandonedDate, PublishedDate, Token, VoteStatus,
// TotalVotes and count of VoteResults between the two ProposalsInfo structs passed.
func (a *ProposalInfo) IsEqual(b *ProposalInfo) bool {
	if a.CensorshipRecord != b.CensorshipRecord || a.Name != b.Name || a.State != b.State ||
		a.NumComments != b.NumComments || a.StatusChangeMsg != b.StatusChangeMsg ||
		a.Status != b.Status || a.Timestamp != b.Timestamp || a.Token != b.Token ||
		a.CensoredDate != b.CensoredDate || a.AbandonedDate != b.AbandonedDate ||
		a.VoteStatus != b.VoteStatus || a.TotalVotes != b.TotalVotes ||
		a.PublishedDate != b.PublishedDate || len(a.VoteResults) != len(b.VoteResults) {
		return false
	}
	return true
}
