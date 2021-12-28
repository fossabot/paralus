package gateway

import (
	"io"
	"io/ioutil"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/segmentio/encoding/json"
	"sigs.k8s.io/yaml"
)

const (
	yamlContentType string = "application/yaml"
)

// rafayJSON is the rafay object to YAML marshaller
type rafayYAML struct {
}

// NewRafayYAML returns new grpc gateway yaml marshaller
func NewRafayYAML() runtime.Marshaler {
	return &rafayYAML{}
}

// Marshal marshals "v" into byte sequence.
func (m *rafayYAML) Marshal(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	yb, err := yaml.JSONToYAML(b)
	if err != nil {
		return nil, err
	}

	return yb, nil
}

// Unmarshal unmarshals "data" into "v".
// "v" must be a pointer value.
func (m *rafayYAML) Unmarshal(yb []byte, v interface{}) error {
	jb, err := yaml.YAMLToJSON(yb)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jb, v)
	return err
}

// NewDecoder returns a Decoder which reads byte sequence from "r".
func (m *rafayYAML) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(v interface{}) error {
		yb, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		return m.Unmarshal(yb, v)
	})
}

// NewEncoder returns an Encoder which writes bytes sequence into "w".
func (m *rafayYAML) NewEncoder(w io.Writer) runtime.Encoder {
	return runtime.EncoderFunc(func(v interface{}) error {
		yb, err := m.Marshal(v)
		if err != nil {
			return err
		}
		_, err = w.Write(yb)

		return err
	})
}

// ContentType returns the Content-Type which this marshaler is responsible for.
func (m *rafayYAML) ContentType(v interface{}) string {
	return yamlContentType
}
