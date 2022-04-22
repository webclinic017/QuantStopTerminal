package edgar

/*
This API returns all the company concepts data for a company into a single API call:

https://data.sec.gov/api/xbrl/companyfacts/CIK##########.json
*/

type CompanyFacts struct {
	Cik        int    `json:"cik"`
	EntityName string `json:"entityName"`
	Facts      struct {
		Dei struct {
			EntityCommonStockSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"EntityCommonStockSharesOutstanding"`
			EntityPublicFloat struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EntityPublicFloat"`
		} `json:"dei"`
		UsGaap struct {
			AccountsPayableCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccountsPayableCurrent"`
			AccrualForTaxesOtherThanIncomeTaxesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccrualForTaxesOtherThanIncomeTaxesCurrent"`
			AccruedIncomeTaxesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedIncomeTaxesCurrent"`
			AccruedLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedLiabilitiesCurrent"`
			AccruedRentCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedRentCurrent"`
			AccumulatedAmortizationDeferredFinanceCosts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedAmortizationDeferredFinanceCosts"`
			AccumulatedDepreciationDepletionAndAmortizationPropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedDepreciationDepletionAndAmortizationPropertyPlantAndEquipment"`
			AccumulatedOtherComprehensiveIncomeLossNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedOtherComprehensiveIncomeLossNetOfTax"`
			AdditionalPaidInCapitalCommonStock struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdditionalPaidInCapitalCommonStock"`
			AdjustmentsToAdditionalPaidInCapitalSharebasedCompensationRequisiteServicePeriodRecognitionValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdjustmentsToAdditionalPaidInCapitalSharebasedCompensationRequisiteServicePeriodRecognitionValue"`
			AdjustmentToAdditionalPaidInCapitalIncomeTaxEffectFromShareBasedCompensationNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdjustmentToAdditionalPaidInCapitalIncomeTaxEffectFromShareBasedCompensationNet"`
			AdvertisingExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdvertisingExpense"`
			AllowanceForDoubtfulAccountsReceivableCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AllowanceForDoubtfulAccountsReceivableCurrent"`
			AmortizationOfFinancingCosts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AmortizationOfFinancingCosts"`
			AmortizationOfFinancingCostsAndDiscounts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AmortizationOfFinancingCostsAndDiscounts"`
			AmortizationOfIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AmortizationOfIntangibleAssets"`
			AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount"`
			AssetImpairmentCharges struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetImpairmentCharges"`
			Assets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Assets"`
			AssetsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetsCurrent"`
			AssetsHeldForSaleLongLivedFairValueDisclosure struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetsHeldForSaleLongLivedFairValueDisclosure"`
			AssetsNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetsNoncurrent"`
			AssetsOfDisposalGroupIncludingDiscontinuedOperationCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetsOfDisposalGroupIncludingDiscontinuedOperationCurrent"`
			BusinessAcquisitionCostOfAcquiredEntityPurchasePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessAcquisitionCostOfAcquiredEntityPurchasePrice"`
			BusinessAcquisitionPercentageOfVotingInterestsAcquired struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"BusinessAcquisitionPercentageOfVotingInterestsAcquired"`
			BusinessAcquisitionPurchasePriceAllocationGoodwillAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessAcquisitionPurchasePriceAllocationGoodwillAmount"`
			BusinessCombinationAcquisitionRelatedCosts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationAcquisitionRelatedCosts"`
			BusinessCombinationContingentConsiderationArrangementsChangeInAmountOfContingentConsiderationLiability1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationContingentConsiderationArrangementsChangeInAmountOfContingentConsiderationLiability1"`
			BusinessCombinationContingentConsiderationArrangementsRangeOfOutcomesValueHigh struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationContingentConsiderationArrangementsRangeOfOutcomesValueHigh"`
			BusinessCombinationContingentConsiderationArrangementsRangeOfOutcomesValueLow struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationContingentConsiderationArrangementsRangeOfOutcomesValueLow"`
			BusinessCombinationContingentConsiderationLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationContingentConsiderationLiability"`
			CashAcquiredFromAcquisition struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashAcquiredFromAcquisition"`
			CashAndCashEquivalentsAtCarryingValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashAndCashEquivalentsAtCarryingValue"`
			CashAndCashEquivalentsPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashAndCashEquivalentsPeriodIncreaseDecrease"`
			CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents"`
			CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalentsPeriodIncreaseDecreaseIncludingExchangeRateEffect struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalentsPeriodIncreaseDecreaseIncludingExchangeRateEffect"`
			ClassOfWarrantOrRightNumberOfSecuritiesCalledByEachWarrantOrRight struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ClassOfWarrantOrRightNumberOfSecuritiesCalledByEachWarrantOrRight"`
			ClassOfWarrantOrRightNumberOfSecuritiesCalledByWarrantsOrRights struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ClassOfWarrantOrRightNumberOfSecuritiesCalledByWarrantsOrRights"`
			CommitmentsAndContingencies struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CommitmentsAndContingencies"`
			CommonStockDividendsPerShareDeclared struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start,omitempty"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"CommonStockDividendsPerShareDeclared"`
			CommonStockParOrStatedValuePerShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"CommonStockParOrStatedValuePerShare"`
			CommonStockSharesAuthorized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"CommonStockSharesAuthorized"`
			CommonStockSharesIssued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"CommonStockSharesIssued"`
			CommonStockSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"CommonStockSharesOutstanding"`
			CommonStockValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CommonStockValue"`
			ComprehensiveIncomeNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ComprehensiveIncomeNetOfTax"`
			ComprehensiveIncomeNetOfTaxAttributableToNoncontrollingInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ComprehensiveIncomeNetOfTaxAttributableToNoncontrollingInterest"`
			ComprehensiveIncomeNetOfTaxIncludingPortionAttributableToNoncontrollingInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ComprehensiveIncomeNetOfTaxIncludingPortionAttributableToNoncontrollingInterest"`
			ContractWithCustomerAssetCumulativeCatchUpAdjustmentToRevenueChangeInMeasureOfProgress struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerAssetCumulativeCatchUpAdjustmentToRevenueChangeInMeasureOfProgress"`
			ContractWithCustomerLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerLiability"`
			ContractWithCustomerLiabilityCumulativeCatchUpAdjustmentToRevenueChangeInMeasureOfProgress struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerLiabilityCumulativeCatchUpAdjustmentToRevenueChangeInMeasureOfProgress"`
			ContractWithCustomerLiabilityCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerLiabilityCurrent"`
			CostOfGoodsSold struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CostOfGoodsSold"`
			CostOfRevenue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CostOfRevenue"`
			CreditCardReceivables struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CreditCardReceivables"`
			CumulativeEffectOfInitialAdoptionOfNewAccountingPrinciple struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CumulativeEffectOfInitialAdoptionOfNewAccountingPrinciple"`
			CumulativeEffectOnRetainedEarningsBeforeTax1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start,omitempty"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CumulativeEffectOnRetainedEarningsBeforeTax1"`
			CumulativeEffectOnRetainedEarningsNetOfTax1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
						Start string `json:"start,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CumulativeEffectOnRetainedEarningsNetOfTax1"`
			CurrentFederalTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentFederalTaxExpenseBenefit"`
			CurrentForeignTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentForeignTaxExpenseBenefit"`
			CurrentIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentIncomeTaxExpenseBenefit"`
			CurrentStateAndLocalTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentStateAndLocalTaxExpenseBenefit"`
			DebtCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtCurrent"`
			DebtInstrumentRepurchaseAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtInstrumentRepurchaseAmount"`
			DeferredFederalIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredFederalIncomeTaxExpenseBenefit"`
			DeferredFinanceCostsGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredFinanceCostsGross"`
			DeferredFinanceCostsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredFinanceCostsNet"`
			DeferredFinanceCostsNoncurrentNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredFinanceCostsNoncurrentNet"`
			DeferredForeignIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredForeignIncomeTaxExpenseBenefit"`
			DeferredIncomeTaxAssetsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredIncomeTaxAssetsNet"`
			DeferredIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredIncomeTaxExpenseBenefit"`
			DeferredRevenueCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredRevenueCurrent"`
			DeferredStateAndLocalIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredStateAndLocalIncomeTaxExpenseBenefit"`
			DeferredTaxAssetsDeferredIncome struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsDeferredIncome"`
			DeferredTaxAssetsGoodwillAndIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsGoodwillAndIntangibleAssets"`
			DeferredTaxAssetsGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsGross"`
			DeferredTaxAssetsInventory struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsInventory"`
			DeferredTaxAssetsLiabilitiesNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsLiabilitiesNet"`
			DeferredTaxAssetsLiabilitiesNetCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsLiabilitiesNetCurrent"`
			DeferredTaxAssetsLiabilitiesNetNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsLiabilitiesNetNoncurrent"`
			DeferredTaxAssetsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsNet"`
			DeferredTaxAssetsNetCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsNetCurrent"`
			DeferredTaxAssetsNetNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsNetNoncurrent"`
			DeferredTaxAssetsOperatingLossCarryforwards struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsOperatingLossCarryforwards"`
			DeferredTaxAssetsOther struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsOther"`
			DeferredTaxAssetsPropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsPropertyPlantAndEquipment"`
			DeferredTaxAssetsTaxCreditCarryforwardsForeign struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxCreditCarryforwardsForeign"`
			DeferredTaxAssetsTaxDeferredExpenseCompensationAndBenefitsEmployeeCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxDeferredExpenseCompensationAndBenefitsEmployeeCompensation"`
			DeferredTaxAssetsTaxDeferredExpenseCompensationAndBenefitsShareBasedCompensationCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxDeferredExpenseCompensationAndBenefitsShareBasedCompensationCost"`
			DeferredTaxAssetsTaxDeferredExpenseReservesAndAccrualsDeferredRent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxDeferredExpenseReservesAndAccrualsDeferredRent"`
			DeferredTaxAssetsValuationAllowance struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsValuationAllowance"`
			DeferredTaxLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilities"`
			DeferredTaxLiabilitiesGoodwillAndIntangibleAssetsGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesGoodwillAndIntangibleAssetsGoodwill"`
			DeferredTaxLiabilitiesGoodwillAndIntangibleAssetsIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesGoodwillAndIntangibleAssetsIntangibleAssets"`
			DeferredTaxLiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesNoncurrent"`
			DeferredTaxLiabilitiesOther struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesOther"`
			DeferredTaxLiabilitiesPrepaidExpenses struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesPrepaidExpenses"`
			DeferredTaxLiabilitiesPropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesPropertyPlantAndEquipment"`
			DefinedBenefitPlanContributionsByEmployer struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DefinedBenefitPlanContributionsByEmployer"`
			DefinedContributionPlanCostRecognized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DefinedContributionPlanCostRecognized"`
			DefinedContributionPlanMaximumAnnualContributionsPerEmployeeAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DefinedContributionPlanMaximumAnnualContributionsPerEmployeeAmount"`
			DefinedContributionPlanMaximumAnnualContributionsPerEmployeePercent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"DefinedContributionPlanMaximumAnnualContributionsPerEmployeePercent"`
			Depreciation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Depreciation"`
			DepreciationAndAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DepreciationAndAmortization"`
			DerivativeNotionalAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeNotionalAmount"`
			DiscontinuedOperationIncomeLossFromDiscontinuedOperationNetOfTaxPerBasicShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"DiscontinuedOperationIncomeLossFromDiscontinuedOperationNetOfTaxPerBasicShare"`
			DiscontinuedOperationIncomeLossFromDiscontinuedOperationNetOfTaxPerDilutedShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"DiscontinuedOperationIncomeLossFromDiscontinuedOperationNetOfTaxPerDilutedShare"`
			DiscontinuedOperationTaxEffectOfDiscontinuedOperation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DiscontinuedOperationTaxEffectOfDiscontinuedOperation"`
			DisposalGroupIncludingDiscontinuedOperationGeneralAndAdministrativeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DisposalGroupIncludingDiscontinuedOperationGeneralAndAdministrativeExpense"`
			Dividends struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Dividends"`
			EarningsPerShareBasic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"EarningsPerShareBasic"`
			EarningsPerShareDiluted struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"EarningsPerShareDiluted"`
			EffectiveIncomeTaxRateContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateContinuingOperations"`
			EffectiveIncomeTaxRateReconciliationAtFederalStatutoryIncomeTaxRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationAtFederalStatutoryIncomeTaxRate"`
			EffectiveIncomeTaxRateReconciliationChangeInDeferredTaxAssetsValuationAllowance struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationChangeInDeferredTaxAssetsValuationAllowance"`
			EffectiveIncomeTaxRateReconciliationDeductionsQualifiedProductionActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationDeductionsQualifiedProductionActivities"`
			EffectiveIncomeTaxRateReconciliationDispositionOfBusiness struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationDispositionOfBusiness"`
			EffectiveIncomeTaxRateReconciliationForeignIncomeTaxRateDifferential struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationForeignIncomeTaxRateDifferential"`
			EffectiveIncomeTaxRateReconciliationNondeductibleExpenseImpairmentLosses struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationNondeductibleExpenseImpairmentLosses"`
			EffectiveIncomeTaxRateReconciliationNondeductibleExpenseOther struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationNondeductibleExpenseOther"`
			EffectiveIncomeTaxRateReconciliationNondeductibleExpenseShareBasedCompensationCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationNondeductibleExpenseShareBasedCompensationCost"`
			EffectiveIncomeTaxRateReconciliationOtherAdjustments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationOtherAdjustments"`
			EffectiveIncomeTaxRateReconciliationStateAndLocalIncomeTaxes struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationStateAndLocalIncomeTaxes"`
			EffectiveIncomeTaxRateReconciliationTaxCreditsForeign struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationTaxCreditsForeign"`
			EffectiveIncomeTaxRateReconciliationTaxExemptIncome struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationTaxExemptIncome"`
			EffectOfExchangeRateOnCashAndCashEquivalents struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EffectOfExchangeRateOnCashAndCashEquivalents"`
			EffectOfExchangeRateOnCashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EffectOfExchangeRateOnCashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents"`
			EmployeeRelatedLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeRelatedLiabilitiesCurrent"`
			EmployeeServiceShareBasedCompensationTaxBenefitFromCompensationExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationTaxBenefitFromCompensationExpense"`
			ExcessTaxBenefitFromShareBasedCompensationFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ExcessTaxBenefitFromShareBasedCompensationFinancingActivities"`
			ExcessTaxBenefitFromShareBasedCompensationOperatingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ExcessTaxBenefitFromShareBasedCompensationOperatingActivities"`
			FiniteLivedIntangibleAssetsAccumulatedAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAccumulatedAmortization"`
			FiniteLivedIntangibleAssetsAmortizationExpenseNextTwelveMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseNextTwelveMonths"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearFive"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearFour"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearThree"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearTwo"`
			FiniteLivedIntangibleAssetsWeightedAverageUsefulLife struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Year []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsWeightedAverageUsefulLife"`
			FixturesAndEquipmentGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FixturesAndEquipmentGross"`
			ForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstrumentsAtFairValueNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstrumentsAtFairValueNet"`
			ForeignCurrencyTransactionGainLossUnrealized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignCurrencyTransactionGainLossUnrealized"`
			GainLossOnDispositionOfAssets1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainLossOnDispositionOfAssets1"`
			GainLossOnSaleOfBusiness struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainLossOnSaleOfBusiness"`
			GainLossOnSaleOfOtherAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainLossOnSaleOfOtherAssets"`
			GainLossOnSaleOfPropertyPlantEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainLossOnSaleOfPropertyPlantEquipment"`
			GainsLossesOnExtinguishmentOfDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainsLossesOnExtinguishmentOfDebt"`
			Goodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Goodwill"`
			GoodwillAcquiredDuringPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillAcquiredDuringPeriod"`
			GoodwillForeignCurrencyTranslationGainLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillForeignCurrencyTranslationGainLoss"`
			GoodwillImpairedAccumulatedImpairmentLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillImpairedAccumulatedImpairmentLoss"`
			GoodwillImpairmentLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillImpairmentLoss"`
			GoodwillTransfers struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillTransfers"`
			GoodwillTranslationAdjustments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillTranslationAdjustments"`
			GoodwillWrittenOffRelatedToSaleOfBusinessUnit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillWrittenOffRelatedToSaleOfBusinessUnit"`
			GrossProfit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GrossProfit"`
			ImpairmentOfIntangibleAssetsExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfIntangibleAssetsExcludingGoodwill"`
			ImpairmentOfIntangibleAssetsFinitelived struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfIntangibleAssetsFinitelived"`
			ImpairmentOfIntangibleAssetsIndefinitelivedExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfIntangibleAssetsIndefinitelivedExcludingGoodwill"`
			ImpairmentOfInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfInvestments"`
			ImpairmentOfLongLivedAssetsHeldForUse struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfLongLivedAssetsHeldForUse"`
			ImpairmentOfLongLivedAssetsToBeDisposedOf struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfLongLivedAssetsToBeDisposedOf"`
			IncomeLossFromContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperations"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesDomestic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesDomestic"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesExtraordinaryItemsNoncontrollingInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesExtraordinaryItemsNoncontrollingInterest"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesForeign struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesForeign"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments"`
			IncomeLossFromContinuingOperationsPerBasicShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsPerBasicShare"`
			IncomeLossFromContinuingOperationsPerDilutedShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsPerDilutedShare"`
			IncomeLossFromDiscontinuedOperationsNetOfTaxAttributableToReportingEntity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromDiscontinuedOperationsNetOfTaxAttributableToReportingEntity"`
			IncomeTaxExaminationEstimateOfPossibleLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Eur []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"EUR"`
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxExaminationEstimateOfPossibleLoss"`
			IncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxExpenseBenefit"`
			IncomeTaxesPaid struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxesPaid"`
			IncomeTaxesPaidNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxesPaidNet"`
			IncomeTaxReceivable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReceivable"`
			IncomeTaxReconciliationChangeInEnactedTaxRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationChangeInEnactedTaxRate"`
			IncreaseDecreaseInAccountsPayableAndAccruedLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInAccountsPayableAndAccruedLiabilities"`
			IncreaseDecreaseInDeferredIncomeTaxes struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInDeferredIncomeTaxes"`
			IncreaseDecreaseInInventories struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInInventories"`
			IncreaseDecreaseInOtherNoncurrentLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInOtherNoncurrentLiabilities"`
			IncreaseDecreaseInOtherOperatingCapitalNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInOtherOperatingCapitalNet"`
			IncreaseDecreaseInPrepaidDeferredExpenseAndOtherAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInPrepaidDeferredExpenseAndOtherAssets"`
			IncreaseDecreaseInReceivables struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInReceivables"`
			IndefinitelivedIntangibleAssetsAcquired struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IndefinitelivedIntangibleAssetsAcquired"`
			IntangibleAssetsNetExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IntangibleAssetsNetExcludingGoodwill"`
			InterestExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestExpense"`
			InterestPaid struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestPaid"`
			InventoryNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InventoryNet"`
			InventoryValuationReserves struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InventoryValuationReserves"`
			InventoryWriteDown struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InventoryWriteDown"`
			InvestmentIncomeInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InvestmentIncomeInterest"`
			Land struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Land"`
			LeaseCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LeaseCost"`
			LesseeOperatingLeaseLiabilityPaymentsDue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDue"`
			LesseeOperatingLeaseLiabilityPaymentsDueAfterYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueAfterYearFive"`
			LesseeOperatingLeaseLiabilityPaymentsDueNextTwelveMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueNextTwelveMonths"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearFive"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearFour"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearThree"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearTwo"`
			LesseeOperatingLeaseLiabilityPaymentsRemainderOfFiscalYear struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsRemainderOfFiscalYear"`
			LesseeOperatingLeaseLiabilityUndiscountedExcessAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityUndiscountedExcessAmount"`
			LettersOfCreditOutstandingAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LettersOfCreditOutstandingAmount"`
			Liabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Liabilities"`
			LiabilitiesAndStockholdersEquity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesAndStockholdersEquity"`
			LiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesCurrent"`
			LiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesNoncurrent"`
			LiabilitiesOfDisposalGroupIncludingDiscontinuedOperationCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesOfDisposalGroupIncludingDiscontinuedOperationCurrent"`
			LineOfCredit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCredit"`
			LineOfCreditFacilityAmountOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityAmountOutstanding"`
			LineOfCreditFacilityAverageOutstandingAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityAverageOutstandingAmount"`
			LineOfCreditFacilityCurrentBorrowingCapacity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityCurrentBorrowingCapacity"`
			LineOfCreditFacilityFairValueOfAmountOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityFairValueOfAmountOutstanding"`
			LineOfCreditFacilityInterestRateDuringPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityInterestRateDuringPeriod"`
			LineOfCreditFacilityMaximumAmountOutstandingDuringPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityMaximumAmountOutstandingDuringPeriod"`
			LineOfCreditFacilityRemainingBorrowingCapacity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LineOfCreditFacilityRemainingBorrowingCapacity"`
			LinesOfCreditCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LinesOfCreditCurrent"`
			LongTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebt"`
			LongTermDebtCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtCurrent"`
			LongTermDebtNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtNoncurrent"`
			LongTermNotesPayable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermNotesPayable"`
			LossContingencyDamagesSoughtValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Eur []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"EUR"`
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LossContingencyDamagesSoughtValue"`
			LossContingencyLossInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LossContingencyLossInPeriod"`
			MarketingAndAdvertisingExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MarketingAndAdvertisingExpense"`
			MinorityInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MinorityInterest"`
			MinorityInterestDecreaseFromRedemptions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MinorityInterestDecreaseFromRedemptions"`
			NetCashProvidedByUsedInFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInFinancingActivities"`
			NetCashProvidedByUsedInInvestingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInInvestingActivities"`
			NetCashProvidedByUsedInOperatingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInOperatingActivities"`
			NetIncomeLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetIncomeLoss"`
			NetIncomeLossAttributableToNoncontrollingInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetIncomeLossAttributableToNoncontrollingInterest"`
			NewAccountingPronouncementOrChangeInAccountingPrincipleEffectOfAdoptionQuantification struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NewAccountingPronouncementOrChangeInAccountingPrincipleEffectOfAdoptionQuantification"`
			NoncashOrPartNoncashAcquisitionIntangibleAssetsAcquired1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NoncashOrPartNoncashAcquisitionIntangibleAssetsAcquired1"`
			NoncashOrPartNoncashAcquisitionNetNonmonetaryAssetsAcquiredLiabilitiesAssumed1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NoncashOrPartNoncashAcquisitionNetNonmonetaryAssetsAcquiredLiabilitiesAssumed1"`
			NotesPayable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NotesPayable"`
			NotesPayableCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NotesPayableCurrent"`
			NotesPayableFairValueDisclosure struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NotesPayableFairValueDisclosure"`
			NumberOfCountriesInWhichEntityOperates struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Country []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"country"`
					Pure []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"NumberOfCountriesInWhichEntityOperates"`
			NumberOfOperatingSegments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
					Segment []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"Segment"`
					Segment2 []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"segment"`
				} `json:"units"`
			} `json:"NumberOfOperatingSegments"`
			NumberOfReportableSegments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"pure"`
					Segment []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"segment"`
				} `json:"units"`
			} `json:"NumberOfReportableSegments"`
			NumberOfStores struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
					Store []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"Store"`
					Store2 []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"store"`
				} `json:"units"`
			} `json:"NumberOfStores"`
			OperatingIncomeLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingIncomeLoss"`
			OperatingLeaseCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseCost"`
			OperatingLeaseImpairmentLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseImpairmentLoss"`
			OperatingLeaseLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseLiability"`
			OperatingLeaseLiabilityCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseLiabilityCurrent"`
			OperatingLeaseLiabilityNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseLiabilityNoncurrent"`
			OperatingLeasePayments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasePayments"`
			OperatingLeaseRightOfUseAsset struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseRightOfUseAsset"`
			OperatingLeasesFutureMinimumPaymentsDue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDue"`
			OperatingLeasesFutureMinimumPaymentsDueCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueCurrent"`
			OperatingLeasesFutureMinimumPaymentsDueInFiveYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInFiveYears"`
			OperatingLeasesFutureMinimumPaymentsDueInFourYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInFourYears"`
			OperatingLeasesFutureMinimumPaymentsDueInThreeYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInThreeYears"`
			OperatingLeasesFutureMinimumPaymentsDueInTwoYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInTwoYears"`
			OperatingLeasesFutureMinimumPaymentsDueThereafter struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueThereafter"`
			OperatingLeasesRentExpenseContingentRentals struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesRentExpenseContingentRentals"`
			OperatingLeasesRentExpenseMinimumRentals struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesRentExpenseMinimumRentals"`
			OperatingLeasesRentExpenseNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesRentExpenseNet"`
			OperatingLeaseWeightedAverageDiscountRatePercent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"OperatingLeaseWeightedAverageDiscountRatePercent"`
			OtherAccruedLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAccruedLiabilitiesCurrent"`
			OtherAssetImpairmentCharges struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAssetImpairmentCharges"`
			OtherAssetsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAssetsCurrent"`
			OtherAssetsFairValueDisclosure struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAssetsFairValueDisclosure"`
			OtherAssetsNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAssetsNoncurrent"`
			OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationAdjustmentNetOfTaxPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationAdjustmentNetOfTaxPeriodIncreaseDecrease"`
			OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationAdjustmentNetOfTaxPortionAttributableToParent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationAdjustmentNetOfTaxPortionAttributableToParent"`
			OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationGainLossArisingDuringPeriodNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationGainLossArisingDuringPeriodNetOfTax"`
			OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationAdjustmentNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationAdjustmentNetOfTax"`
			OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationReclassificationAdjustmentFromAOCIRealizedUponSaleOrLiquidationNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationReclassificationAdjustmentFromAOCIRealizedUponSaleOrLiquidationNetOfTax"`
			OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationReclassificationAdjustmentFromAOCIRealizedUponSaleOrLiquidationTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationReclassificationAdjustmentFromAOCIRealizedUponSaleOrLiquidationTax"`
			OtherComprehensiveIncomeLossTaxPortionAttributableToParent1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossTaxPortionAttributableToParent1"`
			OtherDepreciationAndAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherDepreciationAndAmortization"`
			OtherLiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherLiabilitiesNoncurrent"`
			OtherOperatingActivitiesCashFlowStatement struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherOperatingActivitiesCashFlowStatement"`
			OtherReceivablesGrossCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherReceivablesGrossCurrent"`
			OtherSellingGeneralAndAdministrativeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherSellingGeneralAndAdministrativeExpense"`
			PaymentsForFees struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsForFees"`
			PaymentsForProceedsFromInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsForProceedsFromInvestments"`
			PaymentsForProceedsFromOtherInvestingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsForProceedsFromOtherInvestingActivities"`
			PaymentsForRepurchaseOfCommonStock struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsForRepurchaseOfCommonStock"`
			PaymentsOfDividends struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsOfDividends"`
			PaymentsOfFinancingCosts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsOfFinancingCosts"`
			PaymentsOfMergerRelatedCostsFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsOfMergerRelatedCostsFinancingActivities"`
			PaymentsOfStockIssuanceCosts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsOfStockIssuanceCosts"`
			PaymentsRelatedToTaxWithholdingForShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsRelatedToTaxWithholdingForShareBasedCompensation"`
			PaymentsToAcquireBusinessesGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireBusinessesGross"`
			PaymentsToAcquireBusinessesNetOfCashAcquired struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireBusinessesNetOfCashAcquired"`
			PaymentsToAcquireProductiveAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireProductiveAssets"`
			PaymentsToAcquirePropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquirePropertyPlantAndEquipment"`
			PreferredStockDividendRatePerDollarAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"PreferredStockDividendRatePerDollarAmount"`
			PreferredStockParOrStatedValuePerShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"PreferredStockParOrStatedValuePerShare"`
			PreferredStockSharesAuthorized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"PreferredStockSharesAuthorized"`
			PreferredStockSharesIssued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"PreferredStockSharesIssued"`
			PreferredStockSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"PreferredStockSharesOutstanding"`
			PreferredStockValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PreferredStockValue"`
			PrepaidExpenseAndOtherAssetsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PrepaidExpenseAndOtherAssetsCurrent"`
			PrepaidExpenseCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PrepaidExpenseCurrent"`
			PrepaidTaxes struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PrepaidTaxes"`
			ProceedsFromDivestitureOfBusinesses struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromDivestitureOfBusinesses"`
			ProceedsFromIncomeTaxRefunds struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromIncomeTaxRefunds"`
			ProceedsFromInterestReceived struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromInterestReceived"`
			ProceedsFromIssuanceOfCommonStock struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromIssuanceOfCommonStock"`
			ProceedsFromIssuanceOfLongTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromIssuanceOfLongTermDebt"`
			ProceedsFromIssuanceOfSharesUnderIncentiveAndShareBasedCompensationPlansIncludingStockOptions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromIssuanceOfSharesUnderIncentiveAndShareBasedCompensationPlansIncludingStockOptions"`
			ProceedsFromLifeInsurancePolicies struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromLifeInsurancePolicies"`
			ProceedsFromLinesOfCredit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromLinesOfCredit"`
			ProceedsFromOtherShortTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromOtherShortTermDebt"`
			ProceedsFromPaymentsForOtherFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromPaymentsForOtherFinancingActivities"`
			ProceedsFromSaleOfPropertyHeldForSale struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
						Start string `json:"start,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromSaleOfPropertyHeldForSale"`
			ProceedsFromSaleOfPropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromSaleOfPropertyPlantAndEquipment"`
			ProfitLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProfitLoss"`
			PropertyPlantAndEquipmentFairValueDisclosure struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PropertyPlantAndEquipmentFairValueDisclosure"`
			PropertyPlantAndEquipmentGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PropertyPlantAndEquipmentGross"`
			PropertyPlantAndEquipmentNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PropertyPlantAndEquipmentNet"`
			ReceivablesNetCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ReceivablesNetCurrent"`
			RepaymentsOfLinesOfCredit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start,omitempty"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfLinesOfCredit"`
			RepaymentsOfLongTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfLongTermDebt"`
			RepaymentsOfNotesPayable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfNotesPayable"`
			RepaymentsOfSeniorDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfSeniorDebt"`
			RestrictedCashAndCashEquivalents struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashAndCashEquivalents"`
			RestrictedCashAndCashEquivalentsAtCarryingValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashAndCashEquivalentsAtCarryingValue"`
			RestrictedCashAndCashEquivalentsNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashAndCashEquivalentsNoncurrent"`
			RestrictedCashAndInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashAndInvestments"`
			RestrictedCashCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashCurrent"`
			RestrictedCashNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashNoncurrent"`
			RestructuringCharges struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringCharges"`
			RestructuringCostsAndAssetImpairmentCharges struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringCostsAndAssetImpairmentCharges"`
			RestructuringReserve struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringReserve"`
			RestructuringReserveCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringReserveCurrent"`
			RestructuringReserveSettledWithCash struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringReserveSettledWithCash"`
			RestructuringReserveTranslationAdjustment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringReserveTranslationAdjustment"`
			RestructuringSettlementAndImpairmentProvisions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestructuringSettlementAndImpairmentProvisions"`
			RetainedEarningsAccumulatedDeficit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RetainedEarningsAccumulatedDeficit"`
			RevenueFromContractWithCustomerExcludingAssessedTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RevenueFromContractWithCustomerExcludingAssessedTax"`
			Revenues struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Revenues"`
			RightOfUseAssetObtainedInExchangeForOperatingLeaseLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RightOfUseAssetObtainedInExchangeForOperatingLeaseLiability"`
			SaleLeasebackTransactionNetProceedsInvestingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SaleLeasebackTransactionNetProceedsInvestingActivities"`
			SaleOfStockConsiderationReceivedOnTransaction struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SaleOfStockConsiderationReceivedOnTransaction"`
			SaleOfStockNumberOfSharesIssuedInTransaction struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"SaleOfStockNumberOfSharesIssuedInTransaction"`
			SalesRevenueGoodsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SalesRevenueGoodsNet"`
			SellingGeneralAndAdministrativeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SellingGeneralAndAdministrativeExpense"`
			SeniorLongTermNotes struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SeniorLongTermNotes"`
			SeniorNotesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SeniorNotesCurrent"`
			SettlementLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SettlementLiabilitiesCurrent"`
			ShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensation"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedDividendRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedDividendRate"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedTerm struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Year []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedTerm"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedVolatilityRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedVolatilityRate"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsRiskFreeInterestRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsRiskFreeInterestRate"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardNumberOfSharesAvailableForGrant struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardNumberOfSharesAvailableForGrant"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodTotalIntrinsicValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodTotalIntrinsicValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodGross"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingIntrinsicValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingIntrinsicValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingNumber struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingNumber"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeLowerRangeLimit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeLowerRangeLimit"`
			ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeNumberOfExercisableOptions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeNumberOfExercisableOptions"`
			ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeNumberOfOutstandingOptions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeNumberOfOutstandingOptions"`
			ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeUpperRangeLimit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeUpperRangeLimit"`
			SharebasedCompensationArrangementBySharebasedPaymentAwardOptionsExercisableIntrinsicValue1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SharebasedCompensationArrangementBySharebasedPaymentAwardOptionsExercisableIntrinsicValue1"`
			SharebasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeExercisableOptionsWeightedAverageExercisePrice1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"SharebasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeExercisableOptionsWeightedAverageExercisePrice1"`
			SharebasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeOutstandingOptionsWeightedAverageExercisePriceBeginningBalance1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"SharebasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeOutstandingOptionsWeightedAverageExercisePriceBeginningBalance1"`
			SharebasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeOutstandingOptionsWeightedAverageRemainingContractualTerm1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Year []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"SharebasedCompensationSharesAuthorizedUnderStockOptionPlansExercisePriceRangeOutstandingOptionsWeightedAverageRemainingContractualTerm1"`
			StockholdersEquity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockholdersEquity"`
			StockholdersEquityIncludingPortionAttributableToNoncontrollingInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockholdersEquityIncludingPortionAttributableToNoncontrollingInterest"`
			StockIssuedDuringPeriodSharesStockOptionsExercised struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodSharesStockOptionsExercised"`
			StockIssuedDuringPeriodValueNewIssues struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodValueNewIssues"`
			StockIssuedDuringPeriodValueShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodValueShareBasedCompensation"`
			StockIssuedDuringPeriodValueShareBasedCompensationGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodValueShareBasedCompensationGross"`
			StockRepurchasedAndRetiredDuringPeriodShares struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
						Start string `json:"start,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"StockRepurchasedAndRetiredDuringPeriodShares"`
			StockRepurchasedAndRetiredDuringPeriodValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start,omitempty"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchasedAndRetiredDuringPeriodValue"`
			StockRepurchasedDuringPeriodShares struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"StockRepurchasedDuringPeriodShares"`
			StockRepurchasedDuringPeriodValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchasedDuringPeriodValue"`
			StockRepurchaseProgramAuthorizedAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchaseProgramAuthorizedAmount"`
			StockRepurchaseProgramAuthorizedAmount1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchaseProgramAuthorizedAmount1"`
			StockRepurchaseProgramRemainingAuthorizedRepurchaseAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchaseProgramRemainingAuthorizedRepurchaseAmount"`
			StockRepurchaseProgramRemainingAuthorizedRepurchaseAmount1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchaseProgramRemainingAuthorizedRepurchaseAmount1"`
			TangibleAssetImpairmentCharges struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"TangibleAssetImpairmentCharges"`
			TaxCreditCarryforwardAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"TaxCreditCarryforwardAmount"`
			TreasuryStockAcquiredAverageCostPerShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"TreasuryStockAcquiredAverageCostPerShare"`
			TreasuryStockSharesAcquired struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"TreasuryStockSharesAcquired"`
			UndistributedEarningsOfForeignSubsidiaries struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UndistributedEarningsOfForeignSubsidiaries"`
			UnrecognizedTaxBenefits struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefits"`
			UnrecognizedTaxBenefitsDecreasesResultingFromPriorPeriodTaxPositions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsDecreasesResultingFromPriorPeriodTaxPositions"`
			UnrecognizedTaxBenefitsDecreasesResultingFromSettlementsWithTaxingAuthorities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsDecreasesResultingFromSettlementsWithTaxingAuthorities"`
			UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestAccrued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestAccrued"`
			UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestExpense"`
			UnrecognizedTaxBenefitsIncreasesResultingFromCurrentPeriodTaxPositions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncreasesResultingFromCurrentPeriodTaxPositions"`
			UnrecognizedTaxBenefitsIncreasesResultingFromPriorPeriodTaxPositions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncreasesResultingFromPriorPeriodTaxPositions"`
			UnrecognizedTaxBenefitsIncreasesResultingFromSettlementsWithTaxingAuthorities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncreasesResultingFromSettlementsWithTaxingAuthorities"`
			UnrecognizedTaxBenefitsReductionsResultingFromLapseOfApplicableStatuteOfLimitations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsReductionsResultingFromLapseOfApplicableStatuteOfLimitations"`
			UnrecognizedTaxBenefitsThatWouldImpactEffectiveTaxRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsThatWouldImpactEffectiveTaxRate"`
			ValuationAllowanceDeferredTaxAssetChangeInAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ValuationAllowanceDeferredTaxAssetChangeInAmount"`
			ValuationAllowancesAndReservesBalance struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ValuationAllowancesAndReservesBalance"`
			ValuationAllowancesAndReservesChargedToCostAndExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ValuationAllowancesAndReservesChargedToCostAndExpense"`
			ValuationAllowancesAndReservesChargedToOtherAccounts struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ValuationAllowancesAndReservesChargedToOtherAccounts"`
			ValuationAllowancesAndReservesDeductions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ValuationAllowancesAndReservesDeductions"`
			VariableLeaseCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"VariableLeaseCost"`
			WeightedAverageNumberDilutedSharesOutstandingAdjustment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"WeightedAverageNumberDilutedSharesOutstandingAdjustment"`
			WeightedAverageNumberOfDilutedSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"WeightedAverageNumberOfDilutedSharesOutstanding"`
			WeightedAverageNumberOfSharesOutstandingBasic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"WeightedAverageNumberOfSharesOutstandingBasic"`
			WriteOffOfDeferredDebtIssuanceCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"WriteOffOfDeferredDebtIssuanceCost"`
			CollateralizedFinancings struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CollateralizedFinancings"`
			DisposalGroupIncludingDiscontinuedOperationOperatingIncomeLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DisposalGroupIncludingDiscontinuedOperationOperatingIncomeLoss"`
			EffectiveIncomeTaxRateReconciliationShareBasedCompensationExcessTaxBenefitPercent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationShareBasedCompensationExcessTaxBenefitPercent"`
			PurchaseObligation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PurchaseObligation"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableNumber struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableNumber"`
		} `json:"us-gaap"`
	} `json:"facts"`
}
