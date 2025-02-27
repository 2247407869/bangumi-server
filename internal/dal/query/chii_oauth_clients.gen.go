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

func newOAuthClient(db *gorm.DB, opts ...gen.DOOption) oAuthClient {
	_oAuthClient := oAuthClient{}

	_oAuthClient.oAuthClientDo.UseDB(db, opts...)
	_oAuthClient.oAuthClientDo.UseModel(&dao.OAuthClient{})

	tableName := _oAuthClient.oAuthClientDo.TableName()
	_oAuthClient.ALL = field.NewAsterisk(tableName)
	_oAuthClient.AppID = field.NewUint32(tableName, "app_id")
	_oAuthClient.ClientID = field.NewString(tableName, "client_id")
	_oAuthClient.ClientSecret = field.NewString(tableName, "client_secret")
	_oAuthClient.RedirectURI = field.NewString(tableName, "redirect_uri")
	_oAuthClient.GrantTypes = field.NewString(tableName, "grant_types")
	_oAuthClient.Scope = field.NewString(tableName, "scope")
	_oAuthClient.UserID = field.NewString(tableName, "user_id")
	_oAuthClient.App = oAuthClientBelongsToApp{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("App", "dao.App"),
	}

	_oAuthClient.fillFieldMap()

	return _oAuthClient
}

type oAuthClient struct {
	oAuthClientDo oAuthClientDo

	ALL          field.Asterisk
	AppID        field.Uint32
	ClientID     field.String
	ClientSecret field.String
	RedirectURI  field.String
	GrantTypes   field.String
	Scope        field.String
	UserID       field.String
	App          oAuthClientBelongsToApp

	fieldMap map[string]field.Expr
}

func (o oAuthClient) Table(newTableName string) *oAuthClient {
	o.oAuthClientDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o oAuthClient) As(alias string) *oAuthClient {
	o.oAuthClientDo.DO = *(o.oAuthClientDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *oAuthClient) updateTableName(table string) *oAuthClient {
	o.ALL = field.NewAsterisk(table)
	o.AppID = field.NewUint32(table, "app_id")
	o.ClientID = field.NewString(table, "client_id")
	o.ClientSecret = field.NewString(table, "client_secret")
	o.RedirectURI = field.NewString(table, "redirect_uri")
	o.GrantTypes = field.NewString(table, "grant_types")
	o.Scope = field.NewString(table, "scope")
	o.UserID = field.NewString(table, "user_id")

	o.fillFieldMap()

	return o
}

func (o *oAuthClient) WithContext(ctx context.Context) *oAuthClientDo {
	return o.oAuthClientDo.WithContext(ctx)
}

func (o oAuthClient) TableName() string { return o.oAuthClientDo.TableName() }

func (o oAuthClient) Alias() string { return o.oAuthClientDo.Alias() }

func (o *oAuthClient) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *oAuthClient) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 8)
	o.fieldMap["app_id"] = o.AppID
	o.fieldMap["client_id"] = o.ClientID
	o.fieldMap["client_secret"] = o.ClientSecret
	o.fieldMap["redirect_uri"] = o.RedirectURI
	o.fieldMap["grant_types"] = o.GrantTypes
	o.fieldMap["scope"] = o.Scope
	o.fieldMap["user_id"] = o.UserID

}

func (o oAuthClient) clone(db *gorm.DB) oAuthClient {
	o.oAuthClientDo.ReplaceConnPool(db.Statement.ConnPool)
	return o
}

func (o oAuthClient) replaceDB(db *gorm.DB) oAuthClient {
	o.oAuthClientDo.ReplaceDB(db)
	return o
}

type oAuthClientBelongsToApp struct {
	db *gorm.DB

	field.RelationField
}

func (a oAuthClientBelongsToApp) Where(conds ...field.Expr) *oAuthClientBelongsToApp {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a oAuthClientBelongsToApp) WithContext(ctx context.Context) *oAuthClientBelongsToApp {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a oAuthClientBelongsToApp) Model(m *dao.OAuthClient) *oAuthClientBelongsToAppTx {
	return &oAuthClientBelongsToAppTx{a.db.Model(m).Association(a.Name())}
}

type oAuthClientBelongsToAppTx struct{ tx *gorm.Association }

func (a oAuthClientBelongsToAppTx) Find() (result *dao.App, err error) {
	return result, a.tx.Find(&result)
}

func (a oAuthClientBelongsToAppTx) Append(values ...*dao.App) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a oAuthClientBelongsToAppTx) Replace(values ...*dao.App) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a oAuthClientBelongsToAppTx) Delete(values ...*dao.App) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a oAuthClientBelongsToAppTx) Clear() error {
	return a.tx.Clear()
}

func (a oAuthClientBelongsToAppTx) Count() int64 {
	return a.tx.Count()
}

type oAuthClientDo struct{ gen.DO }

func (o oAuthClientDo) Debug() *oAuthClientDo {
	return o.withDO(o.DO.Debug())
}

func (o oAuthClientDo) WithContext(ctx context.Context) *oAuthClientDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o oAuthClientDo) ReadDB() *oAuthClientDo {
	return o.Clauses(dbresolver.Read)
}

func (o oAuthClientDo) WriteDB() *oAuthClientDo {
	return o.Clauses(dbresolver.Write)
}

func (o oAuthClientDo) Session(config *gorm.Session) *oAuthClientDo {
	return o.withDO(o.DO.Session(config))
}

func (o oAuthClientDo) Clauses(conds ...clause.Expression) *oAuthClientDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o oAuthClientDo) Returning(value interface{}, columns ...string) *oAuthClientDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o oAuthClientDo) Not(conds ...gen.Condition) *oAuthClientDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o oAuthClientDo) Or(conds ...gen.Condition) *oAuthClientDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o oAuthClientDo) Select(conds ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o oAuthClientDo) Where(conds ...gen.Condition) *oAuthClientDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o oAuthClientDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *oAuthClientDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o oAuthClientDo) Order(conds ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o oAuthClientDo) Distinct(cols ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o oAuthClientDo) Omit(cols ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o oAuthClientDo) Join(table schema.Tabler, on ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o oAuthClientDo) LeftJoin(table schema.Tabler, on ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o oAuthClientDo) RightJoin(table schema.Tabler, on ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o oAuthClientDo) Group(cols ...field.Expr) *oAuthClientDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o oAuthClientDo) Having(conds ...gen.Condition) *oAuthClientDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o oAuthClientDo) Limit(limit int) *oAuthClientDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o oAuthClientDo) Offset(offset int) *oAuthClientDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o oAuthClientDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *oAuthClientDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o oAuthClientDo) Unscoped() *oAuthClientDo {
	return o.withDO(o.DO.Unscoped())
}

func (o oAuthClientDo) Create(values ...*dao.OAuthClient) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o oAuthClientDo) CreateInBatches(values []*dao.OAuthClient, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o oAuthClientDo) Save(values ...*dao.OAuthClient) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o oAuthClientDo) First() (*dao.OAuthClient, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthClient), nil
	}
}

func (o oAuthClientDo) Take() (*dao.OAuthClient, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthClient), nil
	}
}

func (o oAuthClientDo) Last() (*dao.OAuthClient, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthClient), nil
	}
}

func (o oAuthClientDo) Find() ([]*dao.OAuthClient, error) {
	result, err := o.DO.Find()
	return result.([]*dao.OAuthClient), err
}

func (o oAuthClientDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dao.OAuthClient, err error) {
	buf := make([]*dao.OAuthClient, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o oAuthClientDo) FindInBatches(result *[]*dao.OAuthClient, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o oAuthClientDo) Attrs(attrs ...field.AssignExpr) *oAuthClientDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o oAuthClientDo) Assign(attrs ...field.AssignExpr) *oAuthClientDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o oAuthClientDo) Joins(fields ...field.RelationField) *oAuthClientDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o oAuthClientDo) Preload(fields ...field.RelationField) *oAuthClientDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o oAuthClientDo) FirstOrInit() (*dao.OAuthClient, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthClient), nil
	}
}

func (o oAuthClientDo) FirstOrCreate() (*dao.OAuthClient, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dao.OAuthClient), nil
	}
}

func (o oAuthClientDo) FindByPage(offset int, limit int) (result []*dao.OAuthClient, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o oAuthClientDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o oAuthClientDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o oAuthClientDo) Delete(models ...*dao.OAuthClient) (result gen.ResultInfo, err error) {
	return o.DO.Delete(models)
}

func (o *oAuthClientDo) withDO(do gen.Dao) *oAuthClientDo {
	o.DO = *do.(*gen.DO)
	return o
}
