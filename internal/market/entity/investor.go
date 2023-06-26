package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosstion []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosstion: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(AssetPosition *InvestorAssetPosition) {
	i.AssetPosstion = append(i.AssetPosstion, AssetPosition)
}

func (i *Investor) UpdateAssetPosition(AssetID string, qtdShares int) {
	assetPosition := i.GetAssetPosition(AssetID)
	if assetPosition == nil {
		i.AssetPosstion = append(i.AssetPosstion, NewInvestorAssetPosition(AssetID, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPosition := range i.AssetPosstion {
		if assetPosition.AssetID == assetID {
			return assetPosition
		} else {
			return nil
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
