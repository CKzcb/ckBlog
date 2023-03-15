/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: article
 * @Version: 1.0.0
 * @Date: 2023/3/15 07:17
 */

package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context)    {}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}
