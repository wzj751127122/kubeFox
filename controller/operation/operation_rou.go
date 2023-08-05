package operation

import "github.com/gin-gonic/gin"

func OperationRouter(ginEngine *gin.RouterGroup) {

	operaRoute := ginEngine.Group("/operation")
	{
		operaRoute.GET("/get_operations", Operation.GetOperationRecordList)
		operaRoute.DELETE("/:id/delete_operation", Operation.DeleteOperationRecord)
		operaRoute.POST("/delete_operations", Operation.DeleteOperationRecords)
	}

}
