// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/bangumi/server/internal/dal/dao"
)

func newFriend(db *gorm.DB, opts ...gen.DOOption) friend {
	_friend := friend{}

	_friend.friendDo.UseDB(db, opts...)
	_friend.friendDo.UseModel(&dao.Friend{})

	tableName := _friend.friendDo.TableName()
	_friend.ALL = field.NewAsterisk(tableName)
	_friend.UserID = field.NewField(tableName, "frd_uid")
	_friend.FriendID = field.NewField(tableName, "frd_fid")
	_friend.Grade = field.NewUint8(tableName, "frd_grade")
	_friend.CreatedTime = field.NewUint32(tableName, "frd_dateline")
	_friend.Description = field.NewString(tableName, "frd_description")

	_friend.fillFieldMap()

	return _friend
}

type friend struct {
	friendDo friendDo

	ALL         field.Asterisk
	UserID      field.Field
	FriendID    field.Field
	Grade       field.Uint8
	CreatedTime field.Uint32
	Description field.String

	fieldMap map[string]field.Expr
}

func (f friend) Table(newTableName string) *friend {
	f.friendDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f friend) As(alias string) *friend {
	f.friendDo.DO = *(f.friendDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *friend) updateTableName(table string) *friend {
	f.ALL = field.NewAsterisk(table)
	f.UserID = field.NewField(table, "frd_uid")
	f.FriendID = field.NewField(table, "frd_fid")
	f.Grade = field.NewUint8(table, "frd_grade")
	f.CreatedTime = field.NewUint32(table, "frd_dateline")
	f.Description = field.NewString(table, "frd_description")

	f.fillFieldMap()

	return f
}

func (f *friend) WithContext(ctx context.Context) *friendDo { return f.friendDo.WithContext(ctx) }

func (f friend) TableName() string { return f.friendDo.TableName() }

func (f friend) Alias() string { return f.friendDo.Alias() }

func (f *friend) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *friend) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 5)
	f.fieldMap["frd_uid"] = f.UserID
	f.fieldMap["frd_fid"] = f.FriendID
	f.fieldMap["frd_grade"] = f.Grade
	f.fieldMap["frd_dateline"] = f.CreatedTime
	f.fieldMap["frd_description"] = f.Description
}

func (f friend) clone(db *gorm.DB) friend {
	f.friendDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f friend) replaceDB(db *gorm.DB) friend {
	f.friendDo.ReplaceDB(db)
	return f
}

type friendDo struct{ gen.DO }

func (f friendDo) Debug() *friendDo {
	return f.withDO(f.DO.Debug())
}

func (f friendDo) WithContext(ctx context.Context) *friendDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f friendDo) ReadDB() *friendDo {
	return f.Clauses(dbresolver.Read)
}

func (f friendDo) WriteDB() *friendDo {
	return f.Clauses(dbresolver.Write)
}

func (f friendDo) Session(config *gorm.Session) *friendDo {
	return f.withDO(f.DO.Session(config))
}

func (f friendDo) Clauses(conds ...clause.Expression) *friendDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f friendDo) Returning(value interface{}, columns ...string) *friendDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f friendDo) Not(conds ...gen.Condition) *friendDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f friendDo) Or(conds ...gen.Condition) *friendDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f friendDo) Select(conds ...field.Expr) *friendDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f friendDo) Where(conds ...gen.Condition) *friendDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f friendDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *friendDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f friendDo) Order(conds ...field.Expr) *friendDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f friendDo) Distinct(cols ...field.Expr) *friendDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f friendDo) Omit(cols ...field.Expr) *friendDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f friendDo) Join(table schema.Tabler, on ...field.Expr) *friendDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f friendDo) LeftJoin(table schema.Tabler, on ...field.Expr) *friendDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f friendDo) RightJoin(table schema.Tabler, on ...field.Expr) *friendDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f friendDo) Group(cols ...field.Expr) *friendDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f friendDo) Having(conds ...gen.Condition) *friendDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f friendDo) Limit(limit int) *friendDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f friendDo) Offset(offset int) *friendDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f friendDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *friendDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f friendDo) Unscoped() *friendDo {
	return f.withDO(f.DO.Unscoped())
}

func (f friendDo) Create(values ...*dao.Friend) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f friendDo) CreateInBatches(values []*dao.Friend, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f friendDo) Save(values ...*dao.Friend) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f friendDo) First() (*dao.Friend, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Friend), nil
	}
}

func (f friendDo) Take() (*dao.Friend, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Friend), nil
	}
}

func (f friendDo) Last() (*dao.Friend, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Friend), nil
	}
}

func (f friendDo) Find() ([]*dao.Friend, error) {
	result, err := f.DO.Find()
	return result.([]*dao.Friend), err
}

func (f friendDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.Friend, err error) {
	buf := make([]*dao.Friend, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f friendDo) FindInBatches(result *[]*dao.Friend, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f friendDo) Attrs(attrs ...field.AssignExpr) *friendDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f friendDo) Assign(attrs ...field.AssignExpr) *friendDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f friendDo) Joins(fields ...field.RelationField) *friendDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f friendDo) Preload(fields ...field.RelationField) *friendDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f friendDo) FirstOrInit() (*dao.Friend, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Friend), nil
	}
}

func (f friendDo) FirstOrCreate() (*dao.Friend, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.Friend), nil
	}
}

func (f friendDo) FindByPage(offset int, limit int) (result []*dao.Friend, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f friendDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f friendDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f friendDo) Delete(models ...*dao.Friend) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *friendDo) withDO(do gen.Dao) *friendDo {
	f.DO = *do.(*gen.DO)
	return f
}
