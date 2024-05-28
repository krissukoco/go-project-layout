package transaction

import (
	"context"
	"database/sql"
	"log"
)

type sqlController struct {
	db *sql.DB
}

type sqlExecutor struct {
	tx *sql.Tx
}

func NewSqlController(db *sql.DB) Controller {
	return &sqlController{db}
}

func (c *sqlController) WithTransaction(ctx context.Context, execFunc ExecFunction) (err error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ex := &sqlExecutor{tx}
	defer func() {
		// Handle panic
		if r := recover(); r != nil {
			if rberr := tx.Rollback(); rberr != nil {
				log.Printf("rollback error: %v", rberr)
			}
			// Re-throw panic
			panic(r)
		}
		// Handle error
		if err != nil {
			if rberr := tx.Rollback(); rberr != nil {
				log.Printf("rollback error: %v", rberr)
			}
			return
		}
		// Finally, commit transaction
		err = tx.Commit()
	}()

	err = execFunc(ctx, ex)
	return err
}

func (e *sqlExecutor) Get(key interface{}) interface{} {
	if key == Key_Sql {
		return e.tx
	}
	return nil
}
