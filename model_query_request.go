/*
 * FIWARE-NGSI v2 Specification
 *
 * TODO: Add a description
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type QueryRequest struct {
	// 
	Entities []interface{} `json:"entities"`
	// 
	Attrs []string `json:"attrs"`
	// 
	Expression *interface{} `json:"expression"`
	// 
	Metadata []string `json:"metadata"`
}
