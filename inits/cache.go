package inits

import (
	"blog/global"
	"github.com/beego/beego/v2/client/cache"
)

func init()  {
	global.Bm, _ = cache.NewCache("file", `{"CachePath":"./cache","FileSuffix":".cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
}
