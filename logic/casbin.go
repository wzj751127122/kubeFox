package logic

import (
	"k8s-platform/dao"
	"k8s-platform/model"
	"strconv"
	"sync"
	models "github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pkg/errors"

)

func UpdateCasbin(AuthorityID uint, casbinInfos []model.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range casbinInfos {
		rules = append(rules, []string{authorityId, v.Path, v.Method})
	}
	e := Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	err := e.InvalidateCache()
	if err != nil {
		return err
	}
	return nil
}

func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := dao.DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := Casbin()
	err = e.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}

func GetPolicyPathByAuthorityId(AuthorityID uint) (pathMaps []model.CasbinInfo) {
	e := Casbin()
	authorityId := strconv.Itoa(int(AuthorityID))
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, model.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

var (
	cachedEnforcer *casbin.CachedEnforcer
	once           sync.Once
)

func Casbin() *casbin.CachedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(dao.DB)
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub)  && keyMatch2(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "111"
		`
		m, err := models.NewModelFromString(text)
		if err != nil {
			return
		}
		cachedEnforcer, _ = casbin.NewCachedEnforcer(m, a)
		cachedEnforcer.SetExpireTime(60 * 60)
		_ = cachedEnforcer.LoadPolicy()
	})
	return cachedEnforcer
}
