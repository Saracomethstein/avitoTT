package openapi

import (
	"context"
	"errors"
	"log/slog"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *DefaultAPIService) ConvertIntoUUID(someId string) (uuid.UUID, error) {
	return uuid.Parse(someId)
}

func (s *DefaultAPIService) ConvertFromUUID(someId uuid.UUID) string {
	return someId.String()
}

func (s *DefaultAPIService) getTenderById(ctx context.Context, tenderId uuid.UUID) error {
	const op = "DefaultAPIService.getTenderById"
	log := s.log.With(slog.String("op", op))

	sql, args, err := s.builder.
		Select("1").
		From("tenders").
		Where(squirrel.Eq{"id": tenderId}).
		ToSql()

	var result int
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("tender does not exist", slog.Any("err", err))
			return ErrOrgNoRightsTender
		}
		log.Error("failed to build SQL query", slog.Any("err", err))
		return ErrSQLQuery
	}
	return nil

}

func (s *DefaultAPIService) getOrganizationById(ctx context.Context, id uuid.UUID) (*Organization, error) {
	const op = "DefaultAPIService.getOrganizationById"
	log := s.log.With(slog.String("op", op))

	sql, args, err := squirrel.Select("id", "name", "description", "type", "created_at", "updated_at").
		From("organization").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		log.Error("failed to build SQL query", slog.Any("err", err))
		return nil, ErrSQLQuery
	}
	row := s.pg.Pool.QueryRow(ctx, sql, args...)
	var organization Organization

	err = row.Scan(&organization.ID, &organization.Name, &organization.Description, &organization.Type, &organization.CreatedAt, &organization.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("organization not found", slog.Any("err", err))
			return nil, ErrNoOrganization
		}
		log.Error("failed to scan organization", "error", err)
		return nil, ErrSQLQuery
	}

	return &organization, nil
}

func (b *DefaultAPIService) organizationHasRights(ctx context.Context, orgId uuid.UUID, tenderId uuid.UUID) error {
	const op = "DefaultAPIService.organizationHasRights"
	log := b.log.With(slog.String("op", op))

	sql, args, err := squirrel.Select("1").
		From("tenders").
		Where(squirrel.Eq{
			"organization_id": orgId,
			"id":              tenderId,
		}).
		Limit(1).
		ToSql()

	if err != nil {
		log.Error("failed to build SQL query", slog.Any("err", err))
		return ErrSQLQuery
	}
	var result int
	err = b.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("organization has no rights for the tender", slog.Any("err", err))
			return ErrOrgNoRightsTender
		}
		log.Error("failed to build SQL query", slog.Any("err", err))
		return ErrSQLQuery
	}

	log.Info("organization has rights for the tender")
	return nil
}

func (b *DefaultAPIService) getUserById(ctx context.Context, userId uuid.UUID) (*User, error) {
	log := b.log.With(
		slog.String("op", "getUserById"),
	)

	sql, args, err := b.builder.
		Select("id", "username", "first_name", "last_name", "created_at", "updated_at").
		From("employee").
		Where(squirrel.Eq{"id": userId}).
		ToSql()

	if err != nil {
		log.Error("failed to build SQL query", slog.Any("err", err))
		return nil, ErrSQLQuery
	}

	resp := &User{}

	// Executing the query and scanning the result into the response struct
	err = b.pg.Pool.QueryRow(ctx, sql, args...).Scan(
		&resp.Id,
		&resp.Username,
		&resp.FirstName,
		&resp.LastName,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("user not found", slog.Any("err", err))
			return nil, ErrNoUser
		}
		log.Error("failed to query user", slog.Any("err", err))
		return nil, ErrSQLQuery
	}

	log.Info("user found", slog.Any("user_id", userId))
	return resp, nil
}

func (b *DefaultAPIService) userHasRights(ctx context.Context, userId uuid.UUID, tenderId uuid.UUID) error {
	const op = "DefaultAPIService.userHasRights"

	log := b.log.With(slog.String("op", op))

	sql, args, err := squirrel.Select("1").
		From("organization_responsible").
		Join("tenders ON tenders.organization_id = organization_responsible.organization_id").
		Where(squirrel.Eq{
			"organization_responsible.user_id": userId,
			"tenders.id":                       tenderId,
		}).
		Limit(1).
		ToSql()

	if err != nil {
		log.Error("failed to build query", slog.Any("error", err))
		return ErrSQLQuery
	}

	// Выполнение запроса
	var exists int
	err = b.pg.Pool.QueryRow(ctx, sql, args...).Scan(&exists)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("пользователь не имеет прав", slog.Any("err", err))
			return ErrUserNoRightsTender
		}
		log.Error("failed to execute query", slog.Any("error", err))
		return ErrSQLQuery
	}
	return nil
}

func (s *DefaultAPIService) userBelongsToOrganization(ctx context.Context, username string, organizationId uuid.UUID) error {
	const op = "DefaultAPIService.userBelongsToOrganization"
	log := s.log.With(slog.String("op", "userBelongsToOrganization"))

	sql, args, err := s.builder.
		Select("1").
		From("organization_responsible").
		Join("employee ON employee.id = organization_responsible.user_id").
		Where(squirrel.Eq{"employee.username": username}).
		Where(squirrel.Eq{"organization_responsible.organization_id": organizationId}).
		Limit(1).
		ToSql()

	if err != nil {
		log.Error("failed to build query", slog.Any("error", err))
		return ErrSQLQuery
	}

	var result int
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("user does not belong to organization", slog.Any("err", err))
			return ErrNotFound
		}
		return ErrSQLQuery
	}
	return nil
}

func (s *DefaultAPIService) getUserByName(ctx context.Context, username string) (*User, error) {
	const op = "getUserByName"
	log := s.log.With("operation", op)

	// Build the SQL query to select the user by username
	sql, args, err := s.builder.
		Select("id", "username", "first_name", "last_name", "created_at", "updated_at").
		From("employee").
		Where(squirrel.Eq{"username": username}).
		Limit(1).
		ToSql()

	if err != nil {
		log.Error("Failed to build SQL", slog.Any("error", err))
		return nil, ErrSQLQuery
	}
	var user User

	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("UserNotFound", slog.Any("err", err))
			return nil, ErrNoUser
		}
		log.Error("Failed to execute query", slog.Any("error", err))
		return nil, ErrSQLQuery
	}

	return &user, nil
}

func (s *DefaultAPIService) getBidById(ctx context.Context, bidId uuid.UUID) (*Bid, error) {
	const op = "getBidById"
	log := s.log.With(
		slog.String("op", op),
		slog.String("bid_id", bidId.String()))

	sql, args, err := s.builder.
		Select("bid_id", "name", "description", "status", "tender_id", "author_type", "author_id", "version", "created_at").
		From("bids").
		Where(squirrel.Eq{"bid_id": bidId}).
		ToSql()

	if err != nil {
		log.Error("failed to build SQL query", slog.Any("error", err))
		return nil, ErrSQLQuery
	}

	bid := &Bid{}

	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(
		&bid.Id,
		&bid.Name,
		&bid.Description,
		&bid.Status,
		&bid.TenderId,
		&bid.AuthorType,
		&bid.AuthorId,
		&bid.Version,
		&bid.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Error("bid not found", slog.String("bid_id", bidId.String()))
			return nil, ErrNotFound
		}
		log.Error("failed to execute SQL query", slog.Any("error", err))
		return nil, ErrSQLQuery
	}

	return bid, nil
}

func (s *DefaultAPIService) addVersionTable(ctx context.Context, tender *Tender) error {
	const op = "addVersionTable"
	log := s.log.With(slog.String("op", op))

	columns := []string{"tender_id", "status", "version", "updated_at"}
	values := []interface{}{tender.Id, tender.Status, tender.Version, tender.CreatedAt}

	if tender.Name != "" {
		columns = append(columns, "name")
		values = append(values, tender.Name)
	}

	if tender.Description != "" {
		columns = append(columns, "description")
		values = append(values, tender.Description)
	}

	if tender.ServiceType != "" {
		columns = append(columns, "service_type")
		values = append(values, tender.ServiceType)
	}

	if tender.OrganizationId != "" {
		newId, _ := s.ConvertIntoUUID(tender.OrganizationId)
		columns = append(columns, "organization_id")
		values = append(values, newId)
	}

	sql, args, err := s.builder.
		Insert("tender_versions").
		Columns(columns...).
		Values(values...).
		ToSql()

	if err != nil {
		log.Error("failed to execute version table insert", slog.Any("err", err))
		return ErrSQLQuery
	}

	_, err = s.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		log.Error("Failed to execute query", slog.Any("error", err))
		return ErrSQLQuery
	}

	return nil
}

//func (s *DefaultAPIService) userCanEditBid(ctx context.Context, username string, bidId uuid.UUID) error {
//	const op = "userCanEditBid"
//	log := s.log.With(
//		slog.String("op", op),
//		slog.String("user_id", username),
//		slog.String("bid_id", bidId.String()))
//
//	sql, args, err := s.builder.
//		Select("1").
//		From("bids").
//		Join("employee ON bids.author_id = employee.id").
//		Where(squirrel.Eq{"bid_id": bidId}).
//		Where(squirrel.Eq{"username": username}).
//		ToSql()
//
//	if err != nil {
//		log.Error("failed to build SQL query", slog.Any("error", err))
//		return ErrSQLQuery
//	}
//
//	var result int
//	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&result)
//
//	if err != nil {
//		if errors.Is(err, pgx.ErrNoRows) {
//			log.Info("user is not the author of the bid", slog.String("bid_id", bidId.String()))
//			return ErrUserNoRightsBid
//		}
//		log.Error("failed to execute SQL query", slog.Any("error", err))
//		return ErrSQLQuery
//	}
//	return nil
//}
