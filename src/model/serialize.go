package model

import (
	"encoding/json"
	"fmt"
)

func (f Field) MarshalJSON() ([]byte, error) {
	type UnresolvedField Field
	return json.Marshal(&struct {
		UnresolvedField
		Type UnresolvedType `json:"type"`
	}{
		UnresolvedField: (UnresolvedField)(f),
		Type: NewUnresolvedFromIType(f.Type),
	})
}

func (f *Field) UnmarshalJSON(data []byte) error {
	type UnresolvedField Field
	temp := &struct {
		UnresolvedField
		Type UnresolvedType `json:"type"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*f = Field(temp.UnresolvedField)
	f.Type = &temp.Type
	return nil
}

func prepareTypes() (map[string]interface{}, error) {
	tempMap := make(map[string]interface{})

	for key, itype := range customTypes {
		var serializedType interface{}

		switch t := itype.(type) {
		case *PrimitiveType:
			serializedType = struct {
				Type string `json:"type"`
				*PrimitiveType
			}{
				Type:          "Primitive",
				PrimitiveType: t,
			}
		case *CompositeType:
			serializedType = struct {
				Type string `json:"type"`
				*CompositeType
			}{
				Type:         "Composite",
				CompositeType: t,
			}
		default:
			return nil, fmt.Errorf("Unknown IType implementation for key: %s", key)
		}

		tempMap[key] = serializedType
	}

	return tempMap, nil
}

func SerializeTypes() ([]byte, error) {
	types, err := prepareTypes()
	if err != nil {
		return []byte{}, err
	}
	withVersion := struct {
		Version string               `json:"version"`
		Types map[string]interface{} `json:"types"`
	}{
		Version: LATEST_SERIALIZE_VERSION,
		Types: types,
	}

	return json.Marshal(withVersion)
}

func resolveTypes(types map[string]IType) error {
	for _, t := range types {
		if comp, ok := t.(*CompositeType); ok {
			for i := range comp.GetFields() {
				if err := comp.Fields[i].resolve(types); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func deserializeTypesAfterVersion(unparsedTypes map[string]json.RawMessage) (map[string]IType, error) {
	types := make(map[string]IType)
	for key, rawValue := range unparsedTypes {
		var typeFieldExtract struct {
			Type string `json:"type"`
		}

		if err := json.Unmarshal(rawValue, &typeFieldExtract); err != nil {
			return nil, fmt.Errorf("\"%s\": %w", key, err)
		}
		var itype IType
		switch typeFieldExtract.Type {
		case "Primitive":
			itype = &PrimitiveType{}
			json.Unmarshal(rawValue, itype)
		case "Composite":
			itype = &CompositeType{}
			json.Unmarshal(rawValue, itype)
		default:
			return nil, fmt.Errorf("Failed to parse type (%s) due to unknown type type: \"%s\"", key, typeFieldExtract.Type)
		}
		if key != itype.GetName() {
			return nil, fmt.Errorf("Key for types map (%s) does not match type name (%s)", key, itype.GetName())
		}
		types[key] = itype
	}

	err := resolveTypes(types)
	if err != nil {
		return nil, err
	}
	return types, nil
}

func DeserializeTypes(data []byte) (map[string]IType, error) {
	withVersion := struct {
		Version string                   `json:"version"`
		Types map[string]json.RawMessage `json:"types"`
	}{}

	err := json.Unmarshal(data, &withVersion)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal to version struct: %w", err)
	}

	if withVersion.Version != LATEST_SERIALIZE_VERSION {
		return nil, fmt.Errorf("This types file seems to have a different Version (%s) than the one supported (%s)", withVersion.Version, LATEST_SERIALIZE_VERSION)
	}

	return deserializeTypesAfterVersion(withVersion.Types)
}
