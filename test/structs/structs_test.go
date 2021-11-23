package structs_test

import (
	"encoding/json"
	"gitlab.idc.xiaozhu.com/xz-go/common/util"
	"testing"
)
type aa struct {
	Ver int `form:"ver" json:"ver"` // 版本
}
type bb struct {
	Ver int `form:"ver" json:"ver" json1:"v_er"` // 版本
}

func TestStructTag(t *testing.T) {
	a := aa{Ver: 2}
	b := bb{}
	_ = util.Bind(&b, a)
	bbb, _ := json.Marshal(b)
	c := util.FieldToMap(b, "json1")
	d, _ := json.Marshal(c)
	t.Logf("bbb %+v", string(bbb))
	//t.Logf("cc %+v %+v", c, err)
	t.Logf("dd %+v", string(d))
}
