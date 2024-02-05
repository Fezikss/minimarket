package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"main.go/config"
	"main.go/storage"

	_ "github.com/golang-migrate/migrate/v4/database"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, cfg config.Config) (storage.IStorage, error) {
	url := fmt.Sprintf(
		`postgres://%s:%s@%s:%s/%s?sslmode=disable`,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	poolconfig, err := pgxpool.ParseConfig(url)
	if err != nil {
		fmt.Println("error while parsing config")
		return nil, err
	}
	poolconfig.MaxConns = 100

	pool, err := pgxpool.NewWithConfig(ctx, poolconfig)
	if err != nil {
		fmt.Println("error while connecting to db", err.Error())
		return nil, err
	}

	// migration
	fmt.Println("test1")
	m, err := migrate.New("file://migration/postgres/", url)

	if err != nil {
		fmt.Println("error while migrating", err.Error())
		return nil, err
	}

	if err = m.Up(); err != nil {

		if !strings.Contains(err.Error(), "no change") {
			version, dirty, err := m.Version()
			if err != nil {
				fmt.Println("err in checking version and dirty", err.Error())
				return nil, err

			}

			if dirty {
				version--
				if err = m.Force(int(version)); err != nil {
					fmt.Println("error in making force", err.Error())
					return nil, err
				}
			}
			fmt.Println("ERROR in migrating", err.Error())
			return nil, err
		}
	}

	return &Store{
		pool: pool,
	}, nil
}

func (s *Store) Close() {
	s.pool.Close()
}
func (s *Store) Branch() storage.IBranchStorage {
	return NewBranchRepo(s.pool)
}
func (s *Store) Sale() storage.ISaleStorage {
	return NewSaleRepo(s.pool)

}
func (s *Store) Category() storage.ICategoryStorage {
	return NewCategoryRepo(s.pool)
}
func (s *Store) Product() storage.IProductStorage {
	return NewProductRepo(s.pool)

}
func (s *Store) Storage() storage.IStorageStorage {
	return NewStorageRepo(s.pool)

}
