package repository

import (
	"database/sql"
	"time"
)

type DeclarationType string

const (
	DeclarationTypeImport DeclarationType = "import"
)

type Status string

const (
	StatusPending               Status = "pending"
	StatusDraft                 Status = "draft"
	StatusIssued                Status = "issued"
	StatusTestPassed            Status = "test_passed"
	StatusReleasedCustoms       Status = "released_customs"
	StatusTransferredToDelivery Status = "transferred_to_delivery"
	StatusReceived              Status = "received"
	StatusFinalized             Status = "finalized"
)

type Declaration struct {
	ID                      int             `db:"id"`
	DeclarationNumber       string          `db:"declaration_number"`
	DeclarationType         DeclarationType `db:"declaration_type"`
	TrackNumber             sql.NullString  `db:"track_number"`
	IIN                     sql.NullString  `db:"iin"`
	Document                sql.NullString  `db:"document"` // Assuming it's stored as a path or reference
	PriceOverallDuty        sql.NullFloat64 `db:"price_overall_duty"`
	Status                  Status          `db:"status"`
	CitizenshipID           sql.NullInt64   `db:"citizenship_id"`
	FromCountryID           sql.NullInt64   `db:"from_country_id"`
	ToCountryID             sql.NullInt64   `db:"to_country_id"`
	KedenUserID             int             `db:"keden_user_id"`
	KedenUserDocumentID     sql.NullInt64   `db:"keden_user_document_id"`
	KedenAddressPermanentID sql.NullInt64   `db:"keden_address_permanent_id"`
	KedenAddressTemporaryID sql.NullInt64   `db:"keden_address_temporary_id"`
	FIO                     sql.NullString  `db:"fio"`
	ReceiverPhone           sql.NullString  `db:"receiver_phone"`
	FullAddress             sql.NullString  `db:"full_address"`
	ReceiverAddress         sql.NullString  `db:"receiver_address"`
	ReceiverCity            sql.NullString  `db:"receiver_city"`
	CreatedAt               time.Time       `db:"created_at"`
	UpdatedAt               time.Time       `db:"updated_at"`
}

func (d *Declaration) ChangeStatus(newStatus Status) {
	d.Status = newStatus
	// Implement database update logic here
}
