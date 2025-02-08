package products

import (
	"context"
	"fmt"
	"ordering/util"
)

func (product *product) DeleteProduct(ctx context.Context, id int64) error {
	_, err := product.store.GetProduct(ctx, id)
	if err != nil {
		return fmt.Errorf("%w: %v", util.ErrRecordNotFound, err)
	}

	err = product.store.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
