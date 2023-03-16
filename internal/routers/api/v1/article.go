/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: article
 * @Version: 1.0.0
 * @Date: 2023/3/15 07:17
 */

package v1

import (
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(ctx *gin.Context)    {}
func (a Article) List(ctx *gin.Context)   {}
func (a Article) Create(ctx *gin.Context) {}
func (a Article) Update(ctx *gin.Context) {}
func (a Article) Delete(ctx *gin.Context) {}
