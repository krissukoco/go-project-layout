package transaction

import "context"

type execKey struct {
	key string
}

var (
	Key_Sql = &execKey{"database/sql"}
)

type ExecFunction func(ctx context.Context, ex Executor) error

type Controller interface {
	// WithTransaction runs execFunc with database transaction.
	// If execFunc returns non-nil error, then the transaction will be rollback.
	//
	// ExecFunction contains Executor which can be used for obtaining underlying database transaction such as *sql.Tx
	//
	// Usage example:
	//
	// 	// Example creating company and user. If user failed to be created, then creating company will be rolled back
	// 	func createCompanyAndUser(ctx context.Context, company Company, user User) (companyId int64, userId int64, err error) {
	// 		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//  	defer cancel()
	// 		var companyId int64
	// 		var userId int64
	//
	// 		err = WithTransaction(ctx, func(ctx context.Context, ex Executor) error {
	// 			// Create Company
	// 			cid, err := companyRepo.TxCreate(ex, company)
	// 			if err != nil {
	// 				return err
	//			}
	// 			// Create User
	//			uid, err := userRepo.TxCreate(ex, user)
	// 			if err != nil {
	// 				return err
	//			}
	// 			// Publish event (example)
	// 			if err := publisher.Publish(ctx, "user_created", uid); err != nil {
	// 				return err
	// 			}
	// 			companyId = cid
	// 			userId = uid
	// 			return nil
	//		})
	//		if err != nil {
	// 			return 0, 0, err
	// 		}
	// 		return companyId, userId, nil
	//	}
	//
	// And inside companyRepo:
	//
	// 	func (r *companyRepo) TxCreate(ex transaction.Executor, company Company) (int64, error) {
	// 		tx, ok := ex.Get(transaction.Key_Sql).(*sql.Tx)
	// 		if !ok {
	// 			return 0, errors.New("unknown database transaction")
	//		}
	// 		var id int64
	// 		err := tx.Query(`INSERT INTO .....`).Scan(&id)
	// 		if err != nil {
	// 			return 0, err
	//		}
	// 		return id, nil
	//	}
	WithTransaction(ctx context.Context, execFunc ExecFunction) error
}

type Executor interface {
	Get(key interface{}) interface{}
}
