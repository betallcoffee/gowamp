// serialization.go
package serialization

type Serialization interface {
	Encode([]interface{}) ([]byte, error)
	Decode([]byte) ([]interface{}, error)
}
