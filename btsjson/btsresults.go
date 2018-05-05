package btsjson

import (
	"encoding/json"
	"github.com/juju/errors"
)

type AccountHistoryOperationDetail struct {
	TotalCount  int                `json:"total_count"`
	ResultCount int                `json:"result_count"`
	Details     []*OperationDetail `json:"details"`
}

type OperationDetail struct {
	Memo        string
	Description string
	Op          OperationArray
}

type OperationArray struct {
	Id string
	Op OperationContents
}

type OperationContent interface{}

type OperationContents []OperationContent

type OperationEnvelope struct {
	Type      int
	Operation interface{}
}

type TransferOperation struct {
	From   string      `json:"from"`
	To     string      `json:"to"`
	Amount AssetAmount `json:"amount"`
	Fee    AssetAmount `json:"fee"`
}

type AssetAmount struct {
	Asset  string `json:"asset_id"`
	Amount int64  `json:"amount"`
}

func (p *OperationEnvelope) UnmarshalJSON(data []byte) error {

	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "Unmarshal raw object")
	}

	if len(raw) != 2 {
		return errors.Errorf("Invalid operation data: %v", string(data))
	}

	if err := json.Unmarshal(raw[0], &p.Type); err != nil {
		return errors.Annotate(err, "Unmarshal OperationType")
	}

	if err := json.Unmarshal(raw[1], &p.Operation); err != nil {
		return errors.Annotate(err, "Unmarshal OperationDetail")
	}

	return nil
}

func ToBytes(in interface{}) []byte {
	b, err := json.Marshal(in)
	if err != nil {
		panic("toBytes is unable to marshal input")
	}
	return b
}

func (content *OperationContents) UnmarshalJSON(data []byte) error {

	var env OperationEnvelope
	if err := json.Unmarshal(data, &env); err != nil {
		return err
	}
	tops := make(OperationContents, 1)
	switch env.Type {
	case 0:
		var top TransferOperation
		if err := json.Unmarshal(ToBytes(env.Operation), &top); err != nil {
			return errors.Annotate(err, "unmarshal TransferOperation")
		}
		tops[0] = &top
	}

	*content = tops

	return nil

}

type Transfer2Result struct {
	TxHash      string
	Operation interface{}
}

func (p *Transfer2Result) UnmarshalJSON(data []byte) error {

	raw := make([]json.RawMessage, 2)
	if err := json.Unmarshal(data, &raw); err != nil {
		return errors.Annotate(err, "Unmarshal raw object")
	}

	if len(raw) != 2 {
		return errors.Errorf("Invalid operation data: %v", string(data))
	}

	if err := json.Unmarshal(raw[0], &p.TxHash); err != nil {
		return errors.Annotate(err, "Unmarshal OperationType")
	}

	if err := json.Unmarshal(raw[1], &p.Operation); err != nil {
		return errors.Annotate(err, "Unmarshal OperationDetail")
	}

	return nil
}
