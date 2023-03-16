/**
 * @Author: 蛋白质先生
 * @Description:
 * @File: tag
 * @Version: 1.0.0
 * @Date: 2023/3/15 07:17
 */

package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// Get
// @Summary get tag
// @Produce json
// @Success 200 {string} string "success"
// @Router /api/v1/tags [get]
//
func (t Tag) Get(ctx *gin.Context)    {}
func (t Tag) List(ctx *gin.Context)   {}
func (t Tag) Create(ctx *gin.Context) {}
func (t Tag) Update(ctx *gin.Context) {}
func (t Tag) Delete(ctx *gin.Context) {}
