package data

import (
	"context"
	"strconv"

	"heytom-auth/internal/biz"
)

type User struct {
	ID
	Name         string `json:"name" gorm:"size:30;not null;comment:用户名称"`
	Mobile       string `json:"mobile" gorm:"size:24;not null;index;comment:用户手机号"`
	Password     string `json:"-" gorm:"not null;default:'';comment:用户密码"`
	GithubOpenid int64  `json:"github_openid" gorm:"comment:github openid"`
	Avatar       string `json:"avatar" gorm:"size:255;not null;comment:头像"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.Id))
}

type userRepo struct {
	data *Data
}

// NewUserRepo .
func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (r *userRepo) PageList(ctx context.Context, pageIndex int64, pageSize int64) (int64, []*biz.User, error) {
	var users []*biz.User
	var total int64
	
	// 查询总数
	if err := r.data.db.Model(&User{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}
	
	// 分页查询
	offset := (pageIndex - 1) * pageSize
	if err := r.data.db.Offset(int(offset)).Limit(int(pageSize)).Find(&users).Error; err != nil {
		return 0, nil, err
	}
	
	return total, users, nil
}

func (r *userRepo) Get(ctx context.Context, id int64) (*biz.User, error) {
	var user biz.User
	if err := r.data.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) Create(ctx context.Context, req *biz.User) (*biz.User, error) {
	if err := r.data.db.Create(req).Error; err != nil {
		return nil, err
	}
	return req, nil
}

func (r *userRepo) Update(ctx context.Context, req *biz.User) error {
	if err := r.data.db.Save(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	if err := r.data.db.Where("id = ?", id).Delete(&biz.User{}).Error; err != nil {
		return err
	}
	return nil
}