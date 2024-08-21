package keeper

import "context"

func (k Keeper) ExtraEIPs(ctx context.Context) ([]int, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	extraEIPs := make([]int, len(params.ExtraEIPs))
	for i, eip := range params.ExtraEIPs {
		extraEIPs[i] = int(eip)
	}

	return extraEIPs, nil
}

func (k Keeper) GetFeeDenom(ctx context.Context) (string, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return "", err
	}

	return params.FeeDenom, nil
}
