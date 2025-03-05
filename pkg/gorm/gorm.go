package gorm

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CrudInterface[T any] interface {
	Scope(scope func(db *gorm.DB) *gorm.DB) *Crud[T]
	Unscoped() *Crud[T]
	FindByIDs(ctx context.Context, ids []int64, preloads ...[]string) ([]*T, error)
	Find(ctx context.Context, id int64, preloads ...[]string) (*T, error)
	FindWithLock(ctx context.Context, id int64, preloads ...[]string) (*T, error)
	FindByConds(ctx context.Context, conds any, preloads ...[]string) (*T, error)
	FindWithLockByConds(ctx context.Context, conds any, preloads ...[]string) (*T, error)
	ListByConds(ctx context.Context, conds any, order any, preloads ...[]string) ([]*T, error)
	PageByConds(ctx context.Context, page, pageSize int64, conds any, order any, preloads ...[]string) ([]*T, int64, error)
	LargePageByConds(ctx context.Context, page, pageSize int64, conds any, order any, preloads ...[]string) ([]*T, bool, error)
	Pluck(ctx context.Context, column string, conds any) ([]any, error)
	PluckInt64(ctx context.Context, column string, conds any) ([]int64, error)
	PluckString(ctx context.Context, column string, conds any) ([]string, error)
	CountByConds(ctx context.Context, conds any) (int64, error)
	ExistsByConds(ctx context.Context, conds any) (bool, error)
	SumByConds(ctx context.Context, conds any, key string) (int64, error)
	Create(ctx context.Context, item *T) error
	BatchCreate(ctx context.Context, items []*T) error
	BatchChunkCreate(ctx context.Context, items []*T, chunkSize int) error
	Update(ctx context.Context, id int64, data any) error
	BatchUpdate(ctx context.Context, ids []int64, data any) error
	UpdateByConds(ctx context.Context, conds any, data any) error
	Upsert(ctx context.Context, item *T, updateFields []string) error
	BatchUpsert(ctx context.Context, items []*T, updateFields []string) error
	Increment(ctx context.Context, id int64, column string, value int64) error
	IncrementByConds(ctx context.Context, conds any, column string, value int64) error
	Decrement(ctx context.Context, id int64, column string, value int64) error
	DecrementByConds(ctx context.Context, conds any, column string, value int64) error
	Delete(ctx context.Context, id int64) error
	BatchDelete(ctx context.Context, ids []int64) error
	DeleteByConds(ctx context.Context, conds any) error
	ForceDelete(ctx context.Context, id int64) error
	ForceDeleteByConds(ctx context.Context, conds any) error
}

type DB = gorm.DB

type Crud[T any] struct {
	*Model
	Scopes []func(db *gorm.DB) *gorm.DB
}

type Model struct {
	client *DB
}

// NewModel 创建一个 Model
func NewModel(c *DB) *Model {
	return &Model{
		client: c,
	}
}

// NewCrud 创建一个默认存储模型
func NewCrud[T any](m *Model) *Crud[T] {
	return &Crud[T]{
		Model:  m,
		Scopes: []func(db *gorm.DB) *gorm.DB{},
	}
}

var _ CrudInterface[any] = (*Crud[any])(nil)

// Paginate 数据分页
func Paginate(page int32, perPage int32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case perPage > 30:
			perPage = 30
		case perPage <= 0:
			perPage = 10
		}

		offset := (page - 1) * perPage

		return db.Offset(int(offset)).Limit(int(perPage))
	}
}

// Paginator 数据分页，新版，不影响用户期望
func Paginator(page, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}

// WithPreloads 预加载关联关系
func WithPreloads(preloads []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, preload := range preloads {
			db = db.Preload(preload)
		}
		return db
	}
}

// MapWhere 根据条件查询
func MapWhere(conds any) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch c := conds.(type) {
		case map[string]any:
			for k, v := range c {
				db = db.Where(k, v)
			}
		default:
			db = db.Where(c)
		}
		return db
	}
}

func FirstOrDefault[T any](l []T, d T) T {
	if len(l) > 0 {
		return l[0]
	}
	return d
}

// Scope 添加作用域
func (c *Crud[T]) Scope(scope func(db *gorm.DB) *gorm.DB) *Crud[T] {
	c.Scopes = append(c.Scopes, scope)
	return c
}

// Unscoped 取消软删除
func (c *Crud[T]) Unscoped() *Crud[T] {
	return c.Scope(func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	})
}

// WithLock 加锁
func (c *Crud[T]) WithLock() *Crud[T] {
	return c.Scope(func(db *gorm.DB) *gorm.DB {
		return db.Clauses(clause.Locking{Strength: "UPDATE"})
	})
}

// Find 查询单条记录
func (c *Crud[T]) Find(ctx context.Context, id int64, preloads ...[]string) (*T, error) {
	var m T

	err := c.client.WithContext(ctx).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Scopes(c.Scopes...).
		Find(&m, id).
		Error

	return &m, err
}

// FindWithLock 查询单条记录并加锁
func (c *Crud[T]) FindWithLock(ctx context.Context, id int64, preloads ...[]string) (*T, error) {
	var m T

	err := c.client.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Scopes(c.Scopes...).
		Find(&m, id).
		Error

	return &m, err
}

// FindByConds 根据条件查询单条记录
func (c *Crud[T]) FindByConds(ctx context.Context, conds any, preloads ...[]string) (*T, error) {
	var m T

	err := c.client.WithContext(ctx).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Find(&m).
		Error

	return &m, err
}

// FindByIDs 批量查询记录
func (c *Crud[T]) FindByIDs(ctx context.Context, ids []int64, preloads ...[]string) ([]*T, error) {
	var ms []*T
	err := c.client.WithContext(ctx).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Where("id IN ?", ids).
		Find(&ms).
		Error

	return ms, err
}

// FindWithLockByConds 根据条件查询单条记录并加锁
func (c *Crud[T]) FindWithLockByConds(ctx context.Context, conds any, preloads ...[]string) (*T, error) {
	var m T

	err := c.client.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Find(&m).
		Error

	return &m, err
}

// ListByConds 根据条件查询多条记录
func (c *Crud[T]) ListByConds(ctx context.Context, conds any, order any, preloads ...[]string) ([]*T, error) {
	var ml []*T

	err := c.client.WithContext(ctx).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Order(order).
		Find(&ml).
		Error

	return ml, err
}

// PageByConds 根据条件分页查询多条记录
func (c *Crud[T]) PageByConds(ctx context.Context, page, pageSize int64, conds any, order any, preloads ...[]string) ([]*T, int64, error) {
	var total int64
	var ml []*T
	var m T

	q := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...)

	if err := q.Count(&total).Error; err != nil {
		return ml, total, err
	}

	err := q.
		Scopes(Paginator(page, pageSize)).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Order(order).
		Find(&ml).
		Error

	return ml, total, err
}

// LargePageByConds 根据条件分页查询多条记录（适用于大数据量）
func (c *Crud[T]) LargePageByConds(ctx context.Context, page, pageSize int64, conds any, order any, preloads ...[]string) ([]*T, bool, error) {
	var ml []*T
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(Paginator(page, pageSize)).
		Scopes(WithPreloads(FirstOrDefault(preloads, nil))).
		Scopes(c.Scopes...).
		Order(order).
		Find(&ml).
		Error

	hasNext := len(ml) == int(pageSize)

	return ml, hasNext, err
}

// Pluck 获取单个列的值
func (c *Crud[T]) Pluck(ctx context.Context, column string, conds any) ([]any, error) {
	var results []any
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Pluck(column, &results).
		Error

	return results, err
}

// PluckInt64 获取单个列的 int64 值
func (c *Crud[T]) PluckInt64(ctx context.Context, column string, conds any) ([]int64, error) {
	var results []int64
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Pluck(column, &results).
		Error

	return results, err
}

// PluckString 获取单个列的 string 值
func (c *Crud[T]) PluckString(ctx context.Context, column string, conds any) ([]string, error) {
	var results []string
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Pluck(column, &results).
		Error

	return results, err
}

// CountByConds 根据条件统计记录数
func (c *Crud[T]) CountByConds(ctx context.Context, conds any) (int64, error) {
	var count int64
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Count(&count).
		Error

	return count, err
}

// ExistsByConds 根据条件判断记录是否存在
func (c *Crud[T]) ExistsByConds(ctx context.Context, conds any) (bool, error) {
	var count int64
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Count(&count).
		Error

	return count > 0, err
}

// SumByConds 根据条件统计某个字段的和
func (c *Crud[T]) SumByConds(ctx context.Context, conds any, key string) (int64, error) {
	var sum int64
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Select("ifnull(sum(" + key + "), 0) as sum").
		Scan(&sum).
		Error

	return sum, err
}

// Create 创建一条记录
func (c *Crud[T]) Create(ctx context.Context, item *T) error {
	err := c.client.WithContext(ctx).
		Create(item).
		Error

	return err
}

// BatchCreate 批量创建记录
func (c *Crud[T]) BatchCreate(ctx context.Context, items []*T) error {
	err := c.client.WithContext(ctx).
		Create(items).
		Error

	return err
}

// BatchChunkCreate 批量分块创建记录
func (c *Crud[T]) BatchChunkCreate(ctx context.Context, items []*T, chunkSize int) error {
	err := c.client.WithContext(ctx).
		CreateInBatches(items, chunkSize).
		Error

	return err
}

// Update 更新一条记录
func (c *Crud[T]) Update(ctx context.Context, id int64, data any) error {
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(c.Scopes...).
		Where("id = ?", id).
		Updates(data).
		Error

	return err
}

// BatchUpdate 批量更新记录
func (c *Crud[T]) BatchUpdate(ctx context.Context, ids []int64, data any) error {
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(c.Scopes...).
		Where("id IN (?)", ids).
		Updates(data).
		Error

	return err
}

// UpdateByConds 根据条件更新记录
func (c *Crud[T]) UpdateByConds(ctx context.Context, conds any, data any) error {
	var m T

	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Updates(data).
		Error

	return err
}

// Upsert 插入或更新一条记录
func (c *Crud[T]) Upsert(ctx context.Context, item *T, updateFields []string) error {
	err := c.client.WithContext(ctx).
		Clauses(clause.OnConflict{
			UpdateAll: len(updateFields) == 0,
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(updateFields),
		}).
		Create(item).
		Error

	return err
}

// BatchUpsert 批量插入或更新记录
func (c *Crud[T]) BatchUpsert(ctx context.Context, items []*T, updateFields []string) error {
	err := c.client.WithContext(ctx).
		Clauses(clause.OnConflict{
			UpdateAll: len(updateFields) == 0,
			Columns:   []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns(updateFields),
		}).
		Create(items).
		Error

	return err
}

// Increment 递增某个字段
func (c *Crud[T]) Increment(ctx context.Context, id int64, column string, value int64) error {
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Where("id = ?", id).
		Scopes(c.Scopes...).
		UpdateColumn(column, gorm.Expr(column+" + ?", value)).
		Error

	return err
}

// IncrementByConds 根据条件递增某个字段
func (c *Crud[T]) IncrementByConds(ctx context.Context, conds any, column string, value int64) error {
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		UpdateColumn(column, gorm.Expr(column+" + ?", value)).
		Error

	return err
}

// Decrement 递减某个字段
func (c *Crud[T]) Decrement(ctx context.Context, id int64, column string, value int64) error {
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Where("id = ?", id).
		Scopes(c.Scopes...).
		UpdateColumn(column, gorm.Expr(column+" - ?", value)).
		Error

	return err
}

// DecrementByConds 根据条件递减某个字段
func (c *Crud[T]) DecrementByConds(ctx context.Context, conds any, column string, value int64) error {
	var m T
	err := c.client.WithContext(ctx).
		Model(&m).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		UpdateColumn(column, gorm.Expr(column+" - ?", value)).
		Error

	return err
}

// Delete 删除一条记录
func (c *Crud[T]) Delete(ctx context.Context, id int64) error {
	var m T

	err := c.client.WithContext(ctx).
		Scopes(c.Scopes...).
		Delete(&m, id).
		Error

	return err
}

// BatchDelete 批量删除记录
func (c *Crud[T]) BatchDelete(ctx context.Context, ids []int64) error {
	var m T

	err := c.client.WithContext(ctx).
		Scopes(c.Scopes...).
		Delete(&m, ids).
		Error

	return err
}

// DeleteByConds 根据条件删除记录
func (c *Crud[T]) DeleteByConds(ctx context.Context, conds any) error {
	var m T

	err := c.client.WithContext(ctx).
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Delete(&m).
		Error

	return err
}

// ForceDelete 强制删除记录（包括软删除的记录）
func (c *Crud[T]) ForceDelete(ctx context.Context, id int64) error {
	var m T
	err := c.client.WithContext(ctx).
		Unscoped().
		Where("id = ?", id).
		Scopes(c.Scopes...).
		Delete(&m).
		Error

	return err
}

// ForceDeleteByConds 根据条件强制删除记录（包括软删除的记录）
func (c *Crud[T]) ForceDeleteByConds(ctx context.Context, conds any) error {
	var m T
	err := c.client.WithContext(ctx).
		Unscoped().
		Scopes(MapWhere(conds)).
		Scopes(c.Scopes...).
		Delete(&m).
		Error

	return err
}
