package server

import imageserver "go_blog/server/image_server"

type ServerGroup struct {
	ImageServer imageserver.ImageServer // 上传图片
}

var Server = new(ServerGroup)
