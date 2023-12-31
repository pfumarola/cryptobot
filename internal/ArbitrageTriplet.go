package internal

type ArbitrageTriplet struct {
	FirstPair  string
	SecondPair string
	ThirdPair  string
}

var ArbitrageTriplets = []ArbitrageTriplet{
	{"BTC-USDT", "BTC-TUSD", "TUSD-USDT"},
	{"BTC-USDT", "BTC-USDC", "USDC-USDT"},
	{"BTC-USDT", "BTC-EUR", "EUR-USDT"},
	{"BTC-USDT", "BTC-FDUSD", "FDUSD-USDT"},
	{"BTC-USDT", "BTC-AEUR", "AEUR-USDT"},
	{"ETH-USDT", "ETH-BTC", "BTC-USDT"},
	{"ETH-USDT", "ETH-TUSD", "TUSD-USDT"},
	{"ETH-USDT", "ETH-USDC", "USDC-USDT"},
	{"ETH-USDT", "ETH-EUR", "EUR-USDT"},
	{"ETH-USDT", "ETH-FDUSD", "FDUSD-USDT"},
	{"ETH-USDT", "ETH-AEUR", "AEUR-USDT"},
	{"BNB-USDT", "BNB-BTC", "BTC-USDT"},
	{"BNB-USDT", "BNB-ETH", "ETH-USDT"},
	{"BNB-USDT", "BNB-TUSD", "TUSD-USDT"},
	{"BNB-USDT", "BNB-USDC", "USDC-USDT"},
	{"BNB-USDT", "BNB-EUR", "EUR-USDT"},
	{"BNB-USDT", "BNB-FDUSD", "FDUSD-USDT"},
	{"NEO-USDT", "NEO-BTC", "BTC-USDT"},
	{"NEO-USDT", "NEO-ETH", "ETH-USDT"},
	{"LTC-USDT", "LTC-BTC", "BTC-USDT"},
	{"LTC-USDT", "LTC-ETH", "ETH-USDT"},
	{"LTC-USDT", "LTC-BNB", "BNB-USDT"},
	{"LTC-USDT", "LTC-TUSD", "TUSD-USDT"},
	{"LTC-USDT", "LTC-EUR", "EUR-USDT"},
	{"LTC-USDT", "LTC-FDUSD", "FDUSD-USDT"},
	{"QTUM-USDT", "QTUM-BTC", "BTC-USDT"},
	{"QTUM-USDT", "QTUM-ETH", "ETH-USDT"},
	{"ADA-USDT", "ADA-BTC", "BTC-USDT"},
	{"ADA-USDT", "ADA-ETH", "ETH-USDT"},
	{"ADA-USDT", "ADA-BNB", "BNB-USDT"},
	{"ADA-USDT", "ADA-TUSD", "TUSD-USDT"},
	{"ADA-USDT", "ADA-USDC", "USDC-USDT"},
	{"ADA-USDT", "ADA-EUR", "EUR-USDT"},
	{"ADA-USDT", "ADA-FDUSD", "FDUSD-USDT"},
	{"XRP-USDT", "XRP-BTC", "BTC-USDT"},
	{"XRP-USDT", "XRP-ETH", "ETH-USDT"},
	{"XRP-USDT", "XRP-BNB", "BNB-USDT"},
	{"XRP-USDT", "XRP-TUSD", "TUSD-USDT"},
	{"XRP-USDT", "XRP-USDC", "USDC-USDT"},
	{"XRP-USDT", "XRP-EUR", "EUR-USDT"},
	{"XRP-USDT", "XRP-FDUSD", "FDUSD-USDT"},
	{"EOS-USDT", "EOS-BTC", "BTC-USDT"},
	{"EOS-USDT", "EOS-ETH", "ETH-USDT"},
	{"IOTA-USDT", "IOTA-BTC", "BTC-USDT"},
	{"IOTA-USDT", "IOTA-ETH", "ETH-USDT"},
	{"IOTA-USDT", "IOTA-FDUSD", "FDUSD-USDT"},
	{"XLM-USDT", "XLM-BTC", "BTC-USDT"},
	{"XLM-USDT", "XLM-ETH", "ETH-USDT"},
	{"XLM-USDT", "XLM-BNB", "BNB-USDT"},
	{"XLM-USDT", "XLM-EUR", "EUR-USDT"},
	{"ONT-USDT", "ONT-BTC", "BTC-USDT"},
	{"TRX-USDT", "TRX-BTC", "BTC-USDT"},
	{"TRX-USDT", "TRX-ETH", "ETH-USDT"},
	{"TRX-USDT", "TRX-BNB", "BNB-USDT"},
	{"TRX-USDT", "TRX-XRP", "XRP-USDT"},
	{"TRX-USDT", "TRX-EUR", "EUR-USDT"},
	{"ETC-USDT", "ETC-BTC", "BTC-USDT"},
	{"ETC-USDT", "ETC-ETH", "ETH-USDT"},
	{"ETC-USDT", "ETC-BNB", "BNB-USDT"},
	{"ICX-USDT", "ICX-BTC", "BTC-USDT"},
	{"NULS-USDT", "NULS-BTC", "BTC-USDT"},
	{"VET-USDT", "VET-BTC", "BTC-USDT"},
	{"VET-USDT", "VET-ETH", "ETH-USDT"},
	{"VET-USDT", "VET-BNB", "BNB-USDT"},
	{"VET-USDT", "VET-EUR", "EUR-USDT"},
	{"LINK-USDT", "LINK-BTC", "BTC-USDT"},
	{"LINK-USDT", "LINK-ETH", "ETH-USDT"},
	{"LINK-USDT", "LINK-BNB", "BNB-USDT"},
	{"LINK-USDT", "LINK-TUSD", "TUSD-USDT"},
	{"LINK-USDT", "LINK-EUR", "EUR-USDT"},
	{"LINK-USDT", "LINK-FDUSD", "FDUSD-USDT"},
	{"WAVES-USDT", "WAVES-BTC", "BTC-USDT"},
	{"WAVES-USDT", "WAVES-ETH", "ETH-USDT"},
	{"ONG-USDT", "ONG-BTC", "BTC-USDT"},
	{"HOT-USDT", "HOT-ETH", "ETH-USDT"},
	{"ZIL-USDT", "ZIL-BTC", "BTC-USDT"},
	{"ZIL-USDT", "ZIL-ETH", "ETH-USDT"},
	{"ZRX-USDT", "ZRX-BTC", "BTC-USDT"},
	{"FET-USDT", "FET-BTC", "BTC-USDT"},
	{"FET-USDT", "FET-BNB", "BNB-USDT"},
	{"FET-USDT", "FET-FDUSD", "FDUSD-USDT"},
	{"BAT-USDT", "BAT-BTC", "BTC-USDT"},
	{"XMR-USDT", "XMR-BTC", "BTC-USDT"},
	{"XMR-USDT", "XMR-ETH", "ETH-USDT"},
	{"XMR-USDT", "XMR-BNB", "BNB-USDT"},
	{"ZEC-USDT", "ZEC-BTC", "BTC-USDT"},
	{"ZEC-USDT", "ZEC-ETH", "ETH-USDT"},
	{"IOST-USDT", "IOST-BTC", "BTC-USDT"},
	{"IOST-USDT", "IOST-ETH", "ETH-USDT"},
	{"CELR-USDT", "CELR-BTC", "BTC-USDT"},
	{"DASH-USDT", "DASH-BTC", "BTC-USDT"},
	{"DASH-USDT", "DASH-ETH", "ETH-USDT"},
	{"THETA-USDT", "THETA-BTC", "BTC-USDT"},
	{"THETA-USDT", "THETA-ETH", "ETH-USDT"},
	{"ENJ-USDT", "ENJ-BTC", "BTC-USDT"},
	{"ENJ-USDT", "ENJ-ETH", "ETH-USDT"},
	{"MATIC-USDT", "MATIC-BTC", "BTC-USDT"},
	{"MATIC-USDT", "MATIC-ETH", "ETH-USDT"},
	{"MATIC-USDT", "MATIC-BNB", "BNB-USDT"},
	{"MATIC-USDT", "MATIC-TUSD", "TUSD-USDT"},
	{"MATIC-USDT", "MATIC-USDC", "USDC-USDT"},
	{"MATIC-USDT", "MATIC-EUR", "EUR-USDT"},
	{"MATIC-USDT", "MATIC-FDUSD", "FDUSD-USDT"},
	{"ATOM-USDT", "ATOM-BTC", "BTC-USDT"},
	{"ATOM-USDT", "ATOM-ETH", "ETH-USDT"},
	{"ATOM-USDT", "ATOM-BNB", "BNB-USDT"},
	{"ATOM-USDT", "ATOM-EUR", "EUR-USDT"},
	{"ATOM-USDT", "ATOM-FDUSD", "FDUSD-USDT"},
	{"TFUEL-USDT", "TFUEL-BTC", "BTC-USDT"},
	{"ONE-USDT", "ONE-BTC", "BTC-USDT"},
	{"ONE-USDT", "ONE-BNB", "BNB-USDT"},
	{"FTM-USDT", "FTM-BTC", "BTC-USDT"},
	{"FTM-USDT", "FTM-ETH", "ETH-USDT"},
	{"FTM-USDT", "FTM-BNB", "BNB-USDT"},
	{"FTM-USDT", "FTM-EUR", "EUR-USDT"},
	{"FTM-USDT", "FTM-FDUSD", "FDUSD-USDT"},
	{"ALGO-USDT", "ALGO-BTC", "BTC-USDT"},
	{"ALGO-USDT", "ALGO-ETH", "ETH-USDT"},
	{"ALGO-USDT", "ALGO-BNB", "BNB-USDT"},
	{"ALGO-USDT", "ALGO-FDUSD", "FDUSD-USDT"},
	{"DOGE-USDT", "DOGE-BTC", "BTC-USDT"},
	{"DOGE-USDT", "DOGE-TUSD", "TUSD-USDT"},
	{"DOGE-USDT", "DOGE-EUR", "EUR-USDT"},
	{"DOGE-USDT", "DOGE-FDUSD", "FDUSD-USDT"},
	{"DUSK-USDT", "DUSK-BTC", "BTC-USDT"},
	{"ANKR-USDT", "ANKR-BTC", "BTC-USDT"},
	{"WIN-USDT", "WIN-BNB", "BNB-USDT"},
	{"WIN-USDT", "WIN-TRX", "TRX-USDT"},
	{"WIN-USDT", "WIN-EUR", "EUR-USDT"},
	{"COS-USDT", "COS-BTC", "BTC-USDT"},
	{"COS-USDT", "COS-BNB", "BNB-USDT"},
	{"MTL-USDT", "MTL-BTC", "BTC-USDT"},
	{"DENT-USDT", "DENT-ETH", "ETH-USDT"},
	{"DOCK-USDT", "DOCK-BTC", "BTC-USDT"},
	{"WAN-USDT", "WAN-BTC", "BTC-USDT"},
	{"WAN-USDT", "WAN-ETH", "ETH-USDT"},
	{"FUN-USDT", "FUN-ETH", "ETH-USDT"},
	{"FUN-USDT", "FUN-BNB", "BNB-USDT"},
	{"CHZ-USDT", "CHZ-BTC", "BTC-USDT"},
	{"CHZ-USDT", "CHZ-BNB", "BNB-USDT"},
	{"CHZ-USDT", "CHZ-EUR", "EUR-USDT"},
	{"BAND-USDT", "BAND-BTC", "BTC-USDT"},
	{"XTZ-USDT", "XTZ-BTC", "BTC-USDT"},
	{"REN-USDT", "REN-BTC", "BTC-USDT"},
	{"RVN-USDT", "RVN-BTC", "BTC-USDT"},
	{"HBAR-USDT", "HBAR-BTC", "BTC-USDT"},
	{"HBAR-USDT", "HBAR-BNB", "BNB-USDT"},
	{"NKN-USDT", "NKN-BTC", "BTC-USDT"},
	{"STX-USDT", "STX-BTC", "BTC-USDT"},
	{"STX-USDT", "STX-BNB", "BNB-USDT"},
	{"KAVA-USDT", "KAVA-BTC", "BTC-USDT"},
	{"KAVA-USDT", "KAVA-ETH", "ETH-USDT"},
	{"ARPA-USDT", "ARPA-BTC", "BTC-USDT"},
	{"ARPA-USDT", "ARPA-BNB", "BNB-USDT"},
	{"IOTX-USDT", "IOTX-BTC", "BTC-USDT"},
	{"IOTX-USDT", "IOTX-ETH", "ETH-USDT"},
	{"RLC-USDT", "RLC-BTC", "BTC-USDT"},
	{"RLC-USDT", "RLC-ETH", "ETH-USDT"},
	{"CTXC-USDT", "CTXC-BTC", "BTC-USDT"},
	{"BCH-USDT", "BCH-BTC", "BTC-USDT"},
	{"BCH-USDT", "BCH-BNB", "BNB-USDT"},
	{"BCH-USDT", "BCH-TUSD", "TUSD-USDT"},
	{"BCH-USDT", "BCH-EUR", "EUR-USDT"},
	{"BCH-USDT", "BCH-FDUSD", "FDUSD-USDT"},
	{"VITE-USDT", "VITE-BTC", "BTC-USDT"},
	{"EUR-USDT", "EUR-AEUR", "AEUR-USDT"},
	{"OGN-USDT", "OGN-BTC", "BTC-USDT"},
	{"DREP-USDT", "DREP-BTC", "BTC-USDT"},
	{"LSK-USDT", "LSK-BTC", "BTC-USDT"},
	{"LSK-USDT", "LSK-ETH", "ETH-USDT"},
	{"BNT-USDT", "BNT-BTC", "BTC-USDT"},
	{"BNT-USDT", "BNT-ETH", "ETH-USDT"},
	{"LTO-USDT", "LTO-BTC", "BTC-USDT"},
	{"COTI-USDT", "COTI-BTC", "BTC-USDT"},
	{"COTI-USDT", "COTI-BNB", "BNB-USDT"},
	{"STPT-USDT", "STPT-BTC", "BTC-USDT"},
	{"DATA-USDT", "DATA-BTC", "BTC-USDT"},
	{"SOL-USDT", "SOL-BTC", "BTC-USDT"},
	{"SOL-USDT", "SOL-ETH", "ETH-USDT"},
	{"SOL-USDT", "SOL-BNB", "BNB-USDT"},
	{"SOL-USDT", "SOL-TUSD", "TUSD-USDT"},
	{"SOL-USDT", "SOL-USDC", "USDC-USDT"},
	{"SOL-USDT", "SOL-EUR", "EUR-USDT"},
	{"SOL-USDT", "SOL-FDUSD", "FDUSD-USDT"},
	{"CTSI-USDT", "CTSI-BTC", "BTC-USDT"},
	{"CTSI-USDT", "CTSI-BNB", "BNB-USDT"},
	{"HIVE-USDT", "HIVE-BTC", "BTC-USDT"},
	{"CHR-USDT", "CHR-BTC", "BTC-USDT"},
	{"CHR-USDT", "CHR-ETH", "ETH-USDT"},
	{"CHR-USDT", "CHR-BNB", "BNB-USDT"},
	{"ARDR-USDT", "ARDR-BTC", "BTC-USDT"},
	{"MDT-USDT", "MDT-BTC", "BTC-USDT"},
	{"KNC-USDT", "KNC-BTC", "BTC-USDT"},
	{"LRC-USDT", "LRC-BTC", "BTC-USDT"},
	{"LRC-USDT", "LRC-ETH", "ETH-USDT"},
	{"COMP-USDT", "COMP-BTC", "BTC-USDT"},
	{"COMP-USDT", "COMP-TUSD", "TUSD-USDT"},
	{"SC-USDT", "SC-ETH", "ETH-USDT"},
	{"ZEN-USDT", "ZEN-BTC", "BTC-USDT"},
	{"ZEN-USDT", "ZEN-ETH", "ETH-USDT"},
	{"SNX-USDT", "SNX-BTC", "BTC-USDT"},
	{"SNX-USDT", "SNX-ETH", "ETH-USDT"},
	{"SNX-USDT", "SNX-BNB", "BNB-USDT"},
	{"DGB-USDT", "DGB-BTC", "BTC-USDT"},
	{"SXP-USDT", "SXP-BTC", "BTC-USDT"},
	{"SXP-USDT", "SXP-BNB", "BNB-USDT"},
	{"MKR-USDT", "MKR-BTC", "BTC-USDT"},
	{"DCR-USDT", "DCR-BTC", "BTC-USDT"},
	{"STORJ-USDT", "STORJ-BTC", "BTC-USDT"},
	{"MANA-USDT", "MANA-BTC", "BTC-USDT"},
	{"MANA-USDT", "MANA-ETH", "ETH-USDT"},
	{"YFI-USDT", "YFI-BTC", "BTC-USDT"},
	{"BAL-USDT", "BAL-BTC", "BTC-USDT"},
	{"BLZ-USDT", "BLZ-BTC", "BTC-USDT"},
	{"BLZ-USDT", "BLZ-FDUSD", "FDUSD-USDT"},
	{"IRIS-USDT", "IRIS-BTC", "BTC-USDT"},
	{"KMD-USDT", "KMD-BTC", "BTC-USDT"},
	{"JST-USDT", "JST-BTC", "BTC-USDT"},
	{"ANT-USDT", "ANT-BTC", "BTC-USDT"},
	{"CRV-USDT", "CRV-BTC", "BTC-USDT"},
	{"CRV-USDT", "CRV-ETH", "ETH-USDT"},
	{"SAND-USDT", "SAND-BTC", "BTC-USDT"},
	{"SAND-USDT", "SAND-ETH", "ETH-USDT"},
	{"SAND-USDT", "SAND-BNB", "BNB-USDT"},
	{"SAND-USDT", "SAND-FDUSD", "FDUSD-USDT"},
	{"OCEAN-USDT", "OCEAN-BTC", "BTC-USDT"},
	{"OCEAN-USDT", "OCEAN-BNB", "BNB-USDT"},
	{"NMR-USDT", "NMR-BTC", "BTC-USDT"},
	{"DOT-USDT", "DOT-BTC", "BTC-USDT"},
	{"DOT-USDT", "DOT-ETH", "ETH-USDT"},
	{"DOT-USDT", "DOT-BNB", "BNB-USDT"},
	{"DOT-USDT", "DOT-TUSD", "TUSD-USDT"},
	{"DOT-USDT", "DOT-USDC", "USDC-USDT"},
	{"DOT-USDT", "DOT-EUR", "EUR-USDT"},
	{"DOT-USDT", "DOT-FDUSD", "FDUSD-USDT"},
	{"PAXG-USDT", "PAXG-BTC", "BTC-USDT"},
	{"TRB-USDT", "TRB-BTC", "BTC-USDT"},
	{"SUSHI-USDT", "SUSHI-BTC", "BTC-USDT"},
	{"KSM-USDT", "KSM-BTC", "BTC-USDT"},
	{"EGLD-USDT", "EGLD-BTC", "BTC-USDT"},
	{"EGLD-USDT", "EGLD-ETH", "ETH-USDT"},
	{"EGLD-USDT", "EGLD-BNB", "BNB-USDT"},
	{"EGLD-USDT", "EGLD-EUR", "EUR-USDT"},
	{"EGLD-USDT", "EGLD-FDUSD", "FDUSD-USDT"},
	{"DIA-USDT", "DIA-BTC", "BTC-USDT"},
	{"RUNE-USDT", "RUNE-BTC", "BTC-USDT"},
	{"RUNE-USDT", "RUNE-ETH", "ETH-USDT"},
	{"RUNE-USDT", "RUNE-BNB", "BNB-USDT"},
	{"RUNE-USDT", "RUNE-TUSD", "TUSD-USDT"},
	{"RUNE-USDT", "RUNE-EUR", "EUR-USDT"},
	{"RUNE-USDT", "RUNE-FDUSD", "FDUSD-USDT"},
	{"FIO-USDT", "FIO-BTC", "BTC-USDT"},
	{"UMA-USDT", "UMA-BTC", "BTC-USDT"},
	{"BEL-USDT", "BEL-BTC", "BTC-USDT"},
	{"WING-USDT", "WING-BTC", "BTC-USDT"},
	{"UNI-USDT", "UNI-BTC", "BTC-USDT"},
	{"UNI-USDT", "UNI-ETH", "ETH-USDT"},
	{"UNI-USDT", "UNI-BNB", "BNB-USDT"},
	{"OXT-USDT", "OXT-BTC", "BTC-USDT"},
	{"AVAX-USDT", "AVAX-BTC", "BTC-USDT"},
	{"AVAX-USDT", "AVAX-ETH", "ETH-USDT"},
	{"AVAX-USDT", "AVAX-BNB", "BNB-USDT"},
	{"AVAX-USDT", "AVAX-TUSD", "TUSD-USDT"},
	{"AVAX-USDT", "AVAX-USDC", "USDC-USDT"},
	{"AVAX-USDT", "AVAX-EUR", "EUR-USDT"},
	{"AVAX-USDT", "AVAX-FDUSD", "FDUSD-USDT"},
	{"FLM-USDT", "FLM-BTC", "BTC-USDT"},
	{"ORN-USDT", "ORN-BTC", "BTC-USDT"},
	{"UTK-USDT", "UTK-BTC", "BTC-USDT"},
	{"XVS-USDT", "XVS-BTC", "BTC-USDT"},
	{"XVS-USDT", "XVS-BNB", "BNB-USDT"},
	{"ALPHA-USDT", "ALPHA-BTC", "BTC-USDT"},
	{"AAVE-USDT", "AAVE-BTC", "BTC-USDT"},
	{"AAVE-USDT", "AAVE-ETH", "ETH-USDT"},
	{"AAVE-USDT", "AAVE-BNB", "BNB-USDT"},
	{"NEAR-USDT", "NEAR-BTC", "BTC-USDT"},
	{"NEAR-USDT", "NEAR-ETH", "ETH-USDT"},
	{"NEAR-USDT", "NEAR-BNB", "BNB-USDT"},
	{"NEAR-USDT", "NEAR-EUR", "EUR-USDT"},
	{"NEAR-USDT", "NEAR-FDUSD", "FDUSD-USDT"},
	{"FIL-USDT", "FIL-BTC", "BTC-USDT"},
	{"FIL-USDT", "FIL-ETH", "ETH-USDT"},
	{"FIL-USDT", "FIL-BNB", "BNB-USDT"},
	{"FIL-USDT", "FIL-FDUSD", "FDUSD-USDT"},
	{"INJ-USDT", "INJ-BTC", "BTC-USDT"},
	{"INJ-USDT", "INJ-ETH", "ETH-USDT"},
	{"INJ-USDT", "INJ-BNB", "BNB-USDT"},
	{"INJ-USDT", "INJ-TUSD", "TUSD-USDT"},
	{"INJ-USDT", "INJ-USDC", "USDC-USDT"},
	{"INJ-USDT", "INJ-FDUSD", "FDUSD-USDT"},
	{"AUDIO-USDT", "AUDIO-BTC", "BTC-USDT"},
	{"CTK-USDT", "CTK-BTC", "BTC-USDT"},
	{"CTK-USDT", "CTK-BNB", "BNB-USDT"},
	{"AXS-USDT", "AXS-BTC", "BTC-USDT"},
	{"AXS-USDT", "AXS-ETH", "ETH-USDT"},
	{"AXS-USDT", "AXS-BNB", "BNB-USDT"},
	{"HARD-USDT", "HARD-BTC", "BTC-USDT"},
	{"STRAX-USDT", "STRAX-BTC", "BTC-USDT"},
	{"UNFI-USDT", "UNFI-BTC", "BTC-USDT"},
	{"ROSE-USDT", "ROSE-BTC", "BTC-USDT"},
	{"ROSE-USDT", "ROSE-ETH", "ETH-USDT"},
	{"AVA-USDT", "AVA-BTC", "BTC-USDT"},
	{"SKL-USDT", "SKL-BTC", "BTC-USDT"},
	{"GRT-USDT", "GRT-BTC", "BTC-USDT"},
	{"GRT-USDT", "GRT-ETH", "ETH-USDT"},
	{"GRT-USDT", "GRT-EUR", "EUR-USDT"},
	{"PSG-USDT", "PSG-BTC", "BTC-USDT"},
	{"1INCH-USDT", "1INCH-BTC", "BTC-USDT"},
	{"OG-USDT", "OG-BTC", "BTC-USDT"},
	{"CELO-USDT", "CELO-BTC", "BTC-USDT"},
	{"RIF-USDT", "RIF-BTC", "BTC-USDT"},
	{"TRU-USDT", "TRU-BTC", "BTC-USDT"},
	{"TWT-USDT", "TWT-BTC", "BTC-USDT"},
	{"FIRO-USDT", "FIRO-BTC", "BTC-USDT"},
	{"LIT-USDT", "LIT-BTC", "BTC-USDT"},
	{"SFP-USDT", "SFP-BTC", "BTC-USDT"},
	{"DODO-USDT", "DODO-BTC", "BTC-USDT"},
	{"CAKE-USDT", "CAKE-BTC", "BTC-USDT"},
	{"CAKE-USDT", "CAKE-BNB", "BNB-USDT"},
	{"CAKE-USDT", "CAKE-TUSD", "TUSD-USDT"},
	{"BADGER-USDT", "BADGER-BTC", "BTC-USDT"},
	{"FIS-USDT", "FIS-BTC", "BTC-USDT"},
	{"OM-USDT", "OM-BTC", "BTC-USDT"},
	{"POND-USDT", "POND-BTC", "BTC-USDT"},
	{"DEGO-USDT", "DEGO-BTC", "BTC-USDT"},
	{"ALICE-USDT", "ALICE-BTC", "BTC-USDT"},
	{"LINA-USDT", "LINA-BTC", "BTC-USDT"},
	{"PERP-USDT", "PERP-BTC", "BTC-USDT"},
	{"SUPER-USDT", "SUPER-BTC", "BTC-USDT"},
	{"SUPER-USDT", "SUPER-FDUSD", "FDUSD-USDT"},
	{"CFX-USDT", "CFX-BTC", "BTC-USDT"},
	{"CFX-USDT", "CFX-TUSD", "TUSD-USDT"},
	{"TKO-USDT", "TKO-BTC", "BTC-USDT"},
	{"PUNDIX-USDT", "PUNDIX-ETH", "ETH-USDT"},
	{"TLM-USDT", "TLM-BTC", "BTC-USDT"},
	{"FORTH-USDT", "FORTH-BTC", "BTC-USDT"},
	{"BAKE-USDT", "BAKE-BTC", "BTC-USDT"},
	{"BAKE-USDT", "BAKE-BNB", "BNB-USDT"},
	{"SLP-USDT", "SLP-ETH", "ETH-USDT"},
	{"SHIB-USDT", "SHIB-TUSD", "TUSD-USDT"},
	{"SHIB-USDT", "SHIB-DOGE", "DOGE-USDT"},
	{"SHIB-USDT", "SHIB-EUR", "EUR-USDT"},
	{"SHIB-USDT", "SHIB-FDUSD", "FDUSD-USDT"},
	{"ICP-USDT", "ICP-BTC", "BTC-USDT"},
	{"ICP-USDT", "ICP-ETH", "ETH-USDT"},
	{"ICP-USDT", "ICP-BNB", "BNB-USDT"},
	{"ICP-USDT", "ICP-EUR", "EUR-USDT"},
	{"ICP-USDT", "ICP-FDUSD", "FDUSD-USDT"},
	{"AR-USDT", "AR-BTC", "BTC-USDT"},
	{"POLS-USDT", "POLS-BTC", "BTC-USDT"},
	{"POLS-USDT", "POLS-BNB", "BNB-USDT"},
	{"MDX-USDT", "MDX-BTC", "BTC-USDT"},
	{"MASK-USDT", "MASK-BNB", "BNB-USDT"},
	{"LPT-USDT", "LPT-BTC", "BTC-USDT"},
	{"LPT-USDT", "LPT-BNB", "BNB-USDT"},
	{"XVG-USDT", "XVG-ETH", "ETH-USDT"},
	{"XVG-USDT", "XVG-TUSD", "TUSD-USDT"},
	{"ATA-USDT", "ATA-BTC", "BTC-USDT"},
	{"GTC-USDT", "GTC-BTC", "BTC-USDT"},
	{"KLAY-USDT", "KLAY-BTC", "BTC-USDT"},
	{"PHA-USDT", "PHA-BTC", "BTC-USDT"},
	{"BOND-USDT", "BOND-BTC", "BTC-USDT"},
	{"MLN-USDT", "MLN-BTC", "BTC-USDT"},
	{"DEXE-USDT", "DEXE-ETH", "ETH-USDT"},
	{"C98-USDT", "C98-BTC", "BTC-USDT"},
	{"CLV-USDT", "CLV-BTC", "BTC-USDT"},
	{"QNT-USDT", "QNT-BTC", "BTC-USDT"},
	{"FLOW-USDT", "FLOW-BTC", "BTC-USDT"},
	{"FLOW-USDT", "FLOW-BNB", "BNB-USDT"},
	{"MINA-USDT", "MINA-BTC", "BTC-USDT"},
	{"RAY-USDT", "RAY-BNB", "BNB-USDT"},
	{"FARM-USDT", "FARM-BTC", "BTC-USDT"},
	{"ALPACA-USDT", "ALPACA-BTC", "BTC-USDT"},
	{"QUICK-USDT", "QUICK-BTC", "BTC-USDT"},
	{"QUICK-USDT", "QUICK-TUSD", "TUSD-USDT"},
	{"MBOX-USDT", "MBOX-BTC", "BTC-USDT"},
	{"MBOX-USDT", "MBOX-BNB", "BNB-USDT"},
	{"FOR-USDT", "FOR-BTC", "BTC-USDT"},
	{"REQ-USDT", "REQ-BTC", "BTC-USDT"},
	{"WAXP-USDT", "WAXP-BTC", "BTC-USDT"},
	{"ELF-USDT", "ELF-BTC", "BTC-USDT"},
	{"ELF-USDT", "ELF-ETH", "ETH-USDT"},
	{"DYDX-USDT", "DYDX-BTC", "BTC-USDT"},
	{"DYDX-USDT", "DYDX-BNB", "BNB-USDT"},
	{"DYDX-USDT", "DYDX-FDUSD", "FDUSD-USDT"},
	{"IDEX-USDT", "IDEX-BTC", "BTC-USDT"},
	{"VIDT-USDT", "VIDT-BTC", "BTC-USDT"},
	{"GALA-USDT", "GALA-BTC", "BTC-USDT"},
	{"GALA-USDT", "GALA-ETH", "ETH-USDT"},
	{"GALA-USDT", "GALA-BNB", "BNB-USDT"},
	{"GALA-USDT", "GALA-EUR", "EUR-USDT"},
	{"GALA-USDT", "GALA-FDUSD", "FDUSD-USDT"},
	{"ILV-USDT", "ILV-BTC", "BTC-USDT"},
	{"YGG-USDT", "YGG-BTC", "BTC-USDT"},
	{"SYS-USDT", "SYS-BTC", "BTC-USDT"},
	{"FIDA-USDT", "FIDA-BTC", "BTC-USDT"},
	{"FRONT-USDT", "FRONT-BTC", "BTC-USDT"},
	{"FRONT-USDT", "FRONT-TUSD", "TUSD-USDT"},
	{"AGLD-USDT", "AGLD-BTC", "BTC-USDT"},
	{"RAD-USDT", "RAD-BTC", "BTC-USDT"},
	{"BETA-USDT", "BETA-BTC", "BTC-USDT"},
	{"RARE-USDT", "RARE-BTC", "BTC-USDT"},
	{"LAZIO-USDT", "LAZIO-BTC", "BTC-USDT"},
	{"LAZIO-USDT", "LAZIO-EUR", "EUR-USDT"},
	{"CHESS-USDT", "CHESS-BTC", "BTC-USDT"},
	{"ADX-USDT", "ADX-BTC", "BTC-USDT"},
	{"ADX-USDT", "ADX-ETH", "ETH-USDT"},
	{"AUCTION-USDT", "AUCTION-BTC", "BTC-USDT"},
	{"AUCTION-USDT", "AUCTION-FDUSD", "FDUSD-USDT"},
	{"DAR-USDT", "DAR-BTC", "BTC-USDT"},
	{"DAR-USDT", "DAR-BNB", "BNB-USDT"},
	{"BNX-USDT", "BNX-BTC", "BTC-USDT"},
	{"MOVR-USDT", "MOVR-BTC", "BTC-USDT"},
	{"CITY-USDT", "CITY-BTC", "BTC-USDT"},
	{"ENS-USDT", "ENS-BTC", "BTC-USDT"},
	{"QI-USDT", "QI-BTC", "BTC-USDT"},
	{"PORTO-USDT", "PORTO-BTC", "BTC-USDT"},
	{"PORTO-USDT", "PORTO-EUR", "EUR-USDT"},
	{"POWR-USDT", "POWR-BTC", "BTC-USDT"},
	{"POWR-USDT", "POWR-ETH", "ETH-USDT"},
	{"PLA-USDT", "PLA-BTC", "BTC-USDT"},
	{"PYR-USDT", "PYR-BTC", "BTC-USDT"},
	{"RNDR-USDT", "RNDR-BTC", "BTC-USDT"},
	{"RNDR-USDT", "RNDR-FDUSD", "FDUSD-USDT"},
	{"ALCX-USDT", "ALCX-BTC", "BTC-USDT"},
	{"SANTOS-USDT", "SANTOS-BTC", "BTC-USDT"},
	{"BICO-USDT", "BICO-BTC", "BTC-USDT"},
	{"FLUX-USDT", "FLUX-BTC", "BTC-USDT"},
	{"FXS-USDT", "FXS-BTC", "BTC-USDT"},
	{"VOXEL-USDT", "VOXEL-BTC", "BTC-USDT"},
	{"HIGH-USDT", "HIGH-BTC", "BTC-USDT"},
	{"PEOPLE-USDT", "PEOPLE-BTC", "BTC-USDT"},
	{"JOE-USDT", "JOE-BTC", "BTC-USDT"},
	{"ACH-USDT", "ACH-BTC", "BTC-USDT"},
	{"IMX-USDT", "IMX-BTC", "BTC-USDT"},
	{"GLMR-USDT", "GLMR-BTC", "BTC-USDT"},
	{"LOKA-USDT", "LOKA-BTC", "BTC-USDT"},
	{"SCRT-USDT", "SCRT-BTC", "BTC-USDT"},
	{"SCRT-USDT", "SCRT-ETH", "ETH-USDT"},
	{"API3-USDT", "API3-BTC", "BTC-USDT"},
	{"ACA-USDT", "ACA-BTC", "BTC-USDT"},
	{"XNO-USDT", "XNO-BTC", "BTC-USDT"},
	{"WOO-USDT", "WOO-BTC", "BTC-USDT"},
	{"WOO-USDT", "WOO-BNB", "BNB-USDT"},
	{"ALPINE-USDT", "ALPINE-BTC", "BTC-USDT"},
	{"ALPINE-USDT", "ALPINE-EUR", "EUR-USDT"},
	{"ASTR-USDT", "ASTR-BTC", "BTC-USDT"},
	{"GMT-USDT", "GMT-BTC", "BTC-USDT"},
	{"GMT-USDT", "GMT-ETH", "ETH-USDT"},
	{"GMT-USDT", "GMT-BNB", "BNB-USDT"},
	{"GMT-USDT", "GMT-EUR", "EUR-USDT"},
	{"KDA-USDT", "KDA-BTC", "BTC-USDT"},
	{"APE-USDT", "APE-BTC", "BTC-USDT"},
	{"APE-USDT", "APE-ETH", "ETH-USDT"},
	{"BSW-USDT", "BSW-BNB", "BNB-USDT"},
	{"MULTI-USDT", "MULTI-BTC", "BTC-USDT"},
	{"STEEM-USDT", "STEEM-BTC", "BTC-USDT"},
	{"STEEM-USDT", "STEEM-ETH", "ETH-USDT"},
	{"MOB-USDT", "MOB-BTC", "BTC-USDT"},
	{"NEXO-USDT", "NEXO-BTC", "BTC-USDT"},
	{"GAL-USDT", "GAL-BTC", "BTC-USDT"},
	{"LDO-USDT", "LDO-BTC", "BTC-USDT"},
	{"LDO-USDT", "LDO-TUSD", "TUSD-USDT"},
	{"LDO-USDT", "LDO-FDUSD", "FDUSD-USDT"},
	{"OP-USDT", "OP-BTC", "BTC-USDT"},
	{"OP-USDT", "OP-ETH", "ETH-USDT"},
	{"OP-USDT", "OP-BNB", "BNB-USDT"},
	{"OP-USDT", "OP-TUSD", "TUSD-USDT"},
	{"OP-USDT", "OP-USDC", "USDC-USDT"},
	{"OP-USDT", "OP-EUR", "EUR-USDT"},
	{"OP-USDT", "OP-FDUSD", "FDUSD-USDT"},
	{"STG-USDT", "STG-BTC", "BTC-USDT"},
	{"GMX-USDT", "GMX-BTC", "BTC-USDT"},
	{"POLYX-USDT", "POLYX-BTC", "BTC-USDT"},
	{"APT-USDT", "APT-BTC", "BTC-USDT"},
	{"APT-USDT", "APT-ETH", "ETH-USDT"},
	{"APT-USDT", "APT-EUR", "EUR-USDT"},
	{"OSMO-USDT", "OSMO-BTC", "BTC-USDT"},
	{"HFT-USDT", "HFT-BTC", "BTC-USDT"},
	{"PHB-USDT", "PHB-BTC", "BTC-USDT"},
	{"HOOK-USDT", "HOOK-BTC", "BTC-USDT"},
	{"MAGIC-USDT", "MAGIC-BTC", "BTC-USDT"},
	{"HIFI-USDT", "HIFI-ETH", "ETH-USDT"},
	{"AGIX-USDT", "AGIX-BTC", "BTC-USDT"},
	{"GNS-USDT", "GNS-BTC", "BTC-USDT"},
	{"SYN-USDT", "SYN-BTC", "BTC-USDT"},
	{"VIB-USDT", "VIB-BTC", "BTC-USDT"},
	{"SSV-USDT", "SSV-BTC", "BTC-USDT"},
	{"SSV-USDT", "SSV-ETH", "ETH-USDT"},
	{"SSV-USDT", "SSV-TUSD", "TUSD-USDT"},
	{"LQTY-USDT", "LQTY-BTC", "BTC-USDT"},
	{"AMB-USDT", "AMB-BTC", "BTC-USDT"},
	{"USTC-USDT", "USTC-FDUSD", "FDUSD-USDT"},
	{"GAS-USDT", "GAS-BTC", "BTC-USDT"},
	{"GAS-USDT", "GAS-FDUSD", "FDUSD-USDT"},
	{"GLM-USDT", "GLM-BTC", "BTC-USDT"},
	{"PROM-USDT", "PROM-BTC", "BTC-USDT"},
	{"QKC-USDT", "QKC-BTC", "BTC-USDT"},
	{"QKC-USDT", "QKC-ETH", "ETH-USDT"},
	{"UFT-USDT", "UFT-ETH", "ETH-USDT"},
	{"ID-USDT", "ID-BTC", "BTC-USDT"},
	{"ID-USDT", "ID-BNB", "BNB-USDT"},
	{"ID-USDT", "ID-TUSD", "TUSD-USDT"},
	{"ARB-USDT", "ARB-BTC", "BTC-USDT"},
	{"ARB-USDT", "ARB-ETH", "ETH-USDT"},
	{"ARB-USDT", "ARB-TUSD", "TUSD-USDT"},
	{"ARB-USDT", "ARB-USDC", "USDC-USDT"},
	{"ARB-USDT", "ARB-EUR", "EUR-USDT"},
	{"ARB-USDT", "ARB-FDUSD", "FDUSD-USDT"},
	{"LOOM-USDT", "LOOM-BTC", "BTC-USDT"},
	{"OAX-USDT", "OAX-BTC", "BTC-USDT"},
	{"RDNT-USDT", "RDNT-BTC", "BTC-USDT"},
	{"RDNT-USDT", "RDNT-TUSD", "TUSD-USDT"},
	{"WBTC-USDT", "WBTC-BTC", "BTC-USDT"},
	{"WBTC-USDT", "WBTC-ETH", "ETH-USDT"},
	{"EDU-USDT", "EDU-BTC", "BTC-USDT"},
	{"EDU-USDT", "EDU-BNB", "BNB-USDT"},
	{"EDU-USDT", "EDU-TUSD", "TUSD-USDT"},
	{"SUI-USDT", "SUI-BTC", "BTC-USDT"},
	{"SUI-USDT", "SUI-BNB", "BNB-USDT"},
	{"SUI-USDT", "SUI-TUSD", "TUSD-USDT"},
	{"SUI-USDT", "SUI-EUR", "EUR-USDT"},
	{"SUI-USDT", "SUI-FDUSD", "FDUSD-USDT"},
	{"AERGO-USDT", "AERGO-BTC", "BTC-USDT"},
	{"PEPE-USDT", "PEPE-TUSD", "TUSD-USDT"},
	{"FLOKI-USDT", "FLOKI-TUSD", "TUSD-USDT"},
	{"AST-USDT", "AST-BTC", "BTC-USDT"},
	{"SNT-USDT", "SNT-BTC", "BTC-USDT"},
	{"SNT-USDT", "SNT-ETH", "ETH-USDT"},
	{"COMBO-USDT", "COMBO-BNB", "BNB-USDT"},
	{"MAV-USDT", "MAV-BTC", "BTC-USDT"},
	{"MAV-USDT", "MAV-TUSD", "TUSD-USDT"},
	{"PENDLE-USDT", "PENDLE-BTC", "BTC-USDT"},
	{"PENDLE-USDT", "PENDLE-TUSD", "TUSD-USDT"},
	{"ARKM-USDT", "ARKM-BTC", "BTC-USDT"},
	{"ARKM-USDT", "ARKM-BNB", "BNB-USDT"},
	{"ARKM-USDT", "ARKM-TUSD", "TUSD-USDT"},
	{"WBETH-USDT", "WBETH-ETH", "ETH-USDT"},
	{"WLD-USDT", "WLD-BTC", "BTC-USDT"},
	{"WLD-USDT", "WLD-FDUSD", "FDUSD-USDT"},
	{"SEI-USDT", "SEI-BTC", "BTC-USDT"},
	{"SEI-USDT", "SEI-BNB", "BNB-USDT"},
	{"SEI-USDT", "SEI-TUSD", "TUSD-USDT"},
	{"SEI-USDT", "SEI-FDUSD", "FDUSD-USDT"},
	{"CYBER-USDT", "CYBER-BTC", "BTC-USDT"},
	{"CYBER-USDT", "CYBER-ETH", "ETH-USDT"},
	{"CYBER-USDT", "CYBER-BNB", "BNB-USDT"},
	{"CYBER-USDT", "CYBER-TUSD", "TUSD-USDT"},
	{"CYBER-USDT", "CYBER-FDUSD", "FDUSD-USDT"},
	{"NTRN-USDT", "NTRN-BTC", "BTC-USDT"},
	{"NTRN-USDT", "NTRN-BNB", "BNB-USDT"},
	{"TIA-USDT", "TIA-BTC", "BTC-USDT"},
	{"TIA-USDT", "TIA-TUSD", "TUSD-USDT"},
	{"MEME-USDT", "MEME-ETH", "ETH-USDT"},
	{"MEME-USDT", "MEME-BNB", "BNB-USDT"},
	{"MEME-USDT", "MEME-TUSD", "TUSD-USDT"},
	{"MEME-USDT", "MEME-FDUSD", "FDUSD-USDT"},
	{"ORDI-USDT", "ORDI-BTC", "BTC-USDT"},
	{"ORDI-USDT", "ORDI-TUSD", "TUSD-USDT"},
	{"ORDI-USDT", "ORDI-USDC", "USDC-USDT"},
	{"ORDI-USDT", "ORDI-FDUSD", "FDUSD-USDT"},
	{"PIVX-USDT", "PIVX-BTC", "BTC-USDT"},
	{"VIC-USDT", "VIC-BTC", "BTC-USDT"},
	{"BLUR-USDT", "BLUR-BTC", "BTC-USDT"},
	{"BLUR-USDT", "BLUR-FDUSD", "FDUSD-USDT"},
	{"VANRY-USDT", "VANRY-BTC", "BTC-USDT"},
	{"JTO-USDT", "JTO-FDUSD", "FDUSD-USDT"},
	{"1000SATS-USDT", "1000SATS-FDUSD", "FDUSD-USDT"},
	{"BONK-USDT", "BONK-FDUSD", "FDUSD-USDT"},
	{"ACE-USDT", "ACE-BTC", "BTC-USDT"},
	{"ACE-USDT", "ACE-BNB", "BNB-USDT"},
	{"ACE-USDT", "ACE-FDUSD", "FDUSD-USDT"},
	{"NFP-USDT", "NFP-BTC", "BTC-USDT"},
	{"NFP-USDT", "NFP-BNB", "BNB-USDT"},
	{"NFP-USDT", "NFP-TUSD", "TUSD-USDT"},
	{"NFP-USDT", "NFP-FDUSD", "FDUSD-USDT"},
	{"AI-USDT", "AI-BTC", "BTC-USDT"},
	{"AI-USDT", "AI-BNB", "BNB-USDT"},
	{"AI-USDT", "AI-TUSD", "TUSD-USDT"},
	{"AI-USDT", "AI-FDUSD", "FDUSD-USDT"},
}

// var ArbitrageTriplets = []ArbitrageTriplet{
// 	{"BNB-FDUSD", "BNB-BTC", "BTC-FDUSD"},
// 	{"BNB-FDUSD", "BNB-ETH", "ETH-FDUSD"},
// 	{"ETH-FDUSD", "ETH-BTC", "BTC-FDUSD"},
// 	{"SEI-FDUSD", "SEI-BNB", "BNB-FDUSD"},
// 	{"SEI-FDUSD", "SEI-BTC", "BTC-FDUSD"},
// 	{"CYBER-FDUSD", "CYBER-BNB", "BNB-FDUSD"},
// 	{"CYBER-FDUSD", "CYBER-BTC", "BTC-FDUSD"},
// 	{"CYBER-FDUSD", "CYBER-ETH", "ETH-FDUSD"},
// 	{"SOL-FDUSD", "SOL-BNB", "BNB-FDUSD"},
// 	{"SOL-FDUSD", "SOL-BTC", "BTC-FDUSD"},
// 	{"SOL-FDUSD", "SOL-ETH", "ETH-FDUSD"},
// 	{"XRP-FDUSD", "XRP-BNB", "BNB-FDUSD"},
// 	{"XRP-FDUSD", "XRP-BTC", "BTC-FDUSD"},
// 	{"XRP-FDUSD", "XRP-ETH", "ETH-FDUSD"},
// 	{"DOGE-FDUSD", "DOGE-BTC", "BTC-FDUSD"},
// 	{"ARB-FDUSD", "ARB-BTC", "BTC-FDUSD"},
// 	{"ARB-FDUSD", "ARB-ETH", "ETH-FDUSD"},
// 	{"SUI-FDUSD", "SUI-BNB", "BNB-FDUSD"},
// 	{"SUI-FDUSD", "SUI-BTC", "BTC-FDUSD"},
// 	{"FIL-FDUSD", "FIL-BNB", "BNB-FDUSD"},
// 	{"FIL-FDUSD", "FIL-BTC", "BTC-FDUSD"},
// 	{"FIL-FDUSD", "FIL-ETH", "ETH-FDUSD"},
// 	{"LTC-FDUSD", "LTC-BNB", "BNB-FDUSD"},
// 	{"LTC-FDUSD", "LTC-BTC", "BTC-FDUSD"},
// 	{"LTC-FDUSD", "LTC-ETH", "ETH-FDUSD"},
// 	{"ADA-FDUSD", "ADA-BNB", "BNB-FDUSD"},
// 	{"ADA-FDUSD", "ADA-BTC", "BTC-FDUSD"},
// 	{"ADA-FDUSD", "ADA-ETH", "ETH-FDUSD"},
// 	{"ATOM-FDUSD", "ATOM-BNB", "BNB-FDUSD"},
// 	{"ATOM-FDUSD", "ATOM-BTC", "BTC-FDUSD"},
// 	{"ATOM-FDUSD", "ATOM-ETH", "ETH-FDUSD"},
// 	{"AVAX-FDUSD", "AVAX-BNB", "BNB-FDUSD"},
// 	{"AVAX-FDUSD", "AVAX-BTC", "BTC-FDUSD"},
// 	{"AVAX-FDUSD", "AVAX-ETH", "ETH-FDUSD"},
// 	{"BCH-FDUSD", "BCH-BNB", "BNB-FDUSD"},
// 	{"BCH-FDUSD", "BCH-BTC", "BTC-FDUSD"},
// 	{"MATIC-FDUSD", "MATIC-BNB", "BNB-FDUSD"},
// 	{"MATIC-FDUSD", "MATIC-BTC", "BTC-FDUSD"},
// 	{"MATIC-FDUSD", "MATIC-ETH", "ETH-FDUSD"},
// 	{"ALGO-FDUSD", "ALGO-BNB", "BNB-FDUSD"},
// 	{"ALGO-FDUSD", "ALGO-BTC", "BTC-FDUSD"},
// 	{"ALGO-FDUSD", "ALGO-ETH", "ETH-FDUSD"},
// 	{"DOT-FDUSD", "DOT-BNB", "BNB-FDUSD"},
// 	{"DOT-FDUSD", "DOT-BTC", "BTC-FDUSD"},
// 	{"DOT-FDUSD", "DOT-ETH", "ETH-FDUSD"},
// 	{"FTM-FDUSD", "FTM-BNB", "BNB-FDUSD"},
// 	{"FTM-FDUSD", "FTM-BTC", "BTC-FDUSD"},
// 	{"FTM-FDUSD", "FTM-ETH", "ETH-FDUSD"},
// 	{"LINK-FDUSD", "LINK-BNB", "BNB-FDUSD"},
// 	{"LINK-FDUSD", "LINK-BTC", "BTC-FDUSD"},
// 	{"LINK-FDUSD", "LINK-ETH", "ETH-FDUSD"},
// 	{"NEAR-FDUSD", "NEAR-BNB", "BNB-FDUSD"},
// 	{"NEAR-FDUSD", "NEAR-BTC", "BTC-FDUSD"},
// 	{"NEAR-FDUSD", "NEAR-ETH", "ETH-FDUSD"},
// 	{"MEME-FDUSD", "MEME-BNB", "BNB-FDUSD"},
// 	{"MEME-FDUSD", "MEME-ETH", "ETH-FDUSD"},
// 	{"EGLD-FDUSD", "EGLD-BNB", "BNB-FDUSD"},
// 	{"EGLD-FDUSD", "EGLD-BTC", "BTC-FDUSD"},
// 	{"EGLD-FDUSD", "EGLD-ETH", "ETH-FDUSD"},
// 	{"FET-FDUSD", "FET-BNB", "BNB-FDUSD"},
// 	{"FET-FDUSD", "FET-BTC", "BTC-FDUSD"},
// 	{"GAS-FDUSD", "GAS-BTC", "BTC-FDUSD"},
// 	{"OP-FDUSD", "OP-BNB", "BNB-FDUSD"},
// 	{"OP-FDUSD", "OP-BTC", "BTC-FDUSD"},
// 	{"OP-FDUSD", "OP-ETH", "ETH-FDUSD"},
// 	{"ORDI-FDUSD", "ORDI-BTC", "BTC-FDUSD"},
// 	{"RNDR-FDUSD", "RNDR-BTC", "BTC-FDUSD"},
// 	{"DYDX-FDUSD", "DYDX-BNB", "BNB-FDUSD"},
// 	{"DYDX-FDUSD", "DYDX-BTC", "BTC-FDUSD"},
// 	{"RUNE-FDUSD", "RUNE-BNB", "BNB-FDUSD"},
// 	{"RUNE-FDUSD", "RUNE-BTC", "BTC-FDUSD"},
// 	{"RUNE-FDUSD", "RUNE-ETH", "ETH-FDUSD"},
// 	{"GALA-FDUSD", "GALA-BNB", "BNB-FDUSD"},
// 	{"GALA-FDUSD", "GALA-BTC", "BTC-FDUSD"},
// 	{"GALA-FDUSD", "GALA-ETH", "ETH-FDUSD"},
// 	{"WLD-FDUSD", "WLD-BTC", "BTC-FDUSD"},
// 	{"BLUR-FDUSD", "BLUR-BTC", "BTC-FDUSD"},
// 	{"SUPER-FDUSD", "SUPER-BTC", "BTC-FDUSD"},
// 	{"AUCTION-FDUSD", "AUCTION-BTC", "BTC-FDUSD"},
// 	{"IOTA-FDUSD", "IOTA-BTC", "BTC-FDUSD"},
// 	{"IOTA-FDUSD", "IOTA-ETH", "ETH-FDUSD"},
// 	{"SHIB-FDUSD", "SHIB-DOGE", "DOGE-FDUSD"},
// 	{"SAND-FDUSD", "SAND-BNB", "BNB-FDUSD"},
// 	{"SAND-FDUSD", "SAND-BTC", "BTC-FDUSD"},
// 	{"SAND-FDUSD", "SAND-ETH", "ETH-FDUSD"},
// 	{"INJ-FDUSD", "INJ-BNB", "BNB-FDUSD"},
// 	{"INJ-FDUSD", "INJ-BTC", "BTC-FDUSD"},
// 	{"INJ-FDUSD", "INJ-ETH", "ETH-FDUSD"},
// 	{"ACE-FDUSD", "ACE-BNB", "BNB-FDUSD"},
// 	{"ACE-FDUSD", "ACE-BTC", "BTC-FDUSD"},
// 	{"BLZ-FDUSD", "BLZ-BTC", "BTC-FDUSD"},
// 	{"NFP-FDUSD", "NFP-BNB", "BNB-FDUSD"},
// 	{"NFP-FDUSD", "NFP-BTC", "BTC-FDUSD"},
// 	{"AI-FDUSD", "AI-BNB", "BNB-FDUSD"},
// 	{"AI-FDUSD", "AI-BTC", "BTC-FDUSD"},
// }

// var ArbitrageTriplets = []ArbitrageTriplet{
// 	{"ETH-TUSD", "ETH-BTC", "BTC-TUSD"},
// 	{"BNB-TUSD", "BNB-BTC", "BTC-TUSD"},
// 	{"BNB-TUSD", "BNB-ETH", "ETH-TUSD"},
// 	{"XRP-TUSD", "XRP-BTC", "BTC-TUSD"},
// 	{"XRP-TUSD", "XRP-ETH", "ETH-TUSD"},
// 	{"XRP-TUSD", "XRP-BNB", "BNB-TUSD"},
// 	{"ADA-TUSD", "ADA-BTC", "BTC-TUSD"},
// 	{"ADA-TUSD", "ADA-ETH", "ETH-TUSD"},
// 	{"ADA-TUSD", "ADA-BNB", "BNB-TUSD"},
// 	{"LINK-TUSD", "LINK-BTC", "BTC-TUSD"},
// 	{"LINK-TUSD", "LINK-ETH", "ETH-TUSD"},
// 	{"LINK-TUSD", "LINK-BNB", "BNB-TUSD"},
// 	{"LTC-TUSD", "LTC-BTC", "BTC-TUSD"},
// 	{"LTC-TUSD", "LTC-ETH", "ETH-TUSD"},
// 	{"LTC-TUSD", "LTC-BNB", "BNB-TUSD"},
// 	{"BCH-TUSD", "BCH-BTC", "BTC-TUSD"},
// 	{"BCH-TUSD", "BCH-BNB", "BNB-TUSD"},
// 	{"ARB-TUSD", "ARB-BTC", "BTC-TUSD"},
// 	{"ARB-TUSD", "ARB-ETH", "ETH-TUSD"},
// 	{"ID-TUSD", "ID-BTC", "BTC-TUSD"},
// 	{"ID-TUSD", "ID-BNB", "BNB-TUSD"},
// 	{"LDO-TUSD", "LDO-BTC", "BTC-TUSD"},
// 	{"MATIC-TUSD", "MATIC-BTC", "BTC-TUSD"},
// 	{"MATIC-TUSD", "MATIC-ETH", "ETH-TUSD"},
// 	{"MATIC-TUSD", "MATIC-BNB", "BNB-TUSD"},
// 	{"OP-TUSD", "OP-BTC", "BTC-TUSD"},
// 	{"OP-TUSD", "OP-ETH", "ETH-TUSD"},
// 	{"OP-TUSD", "OP-BNB", "BNB-TUSD"},
// 	{"SOL-TUSD", "SOL-BTC", "BTC-TUSD"},
// 	{"SOL-TUSD", "SOL-ETH", "ETH-TUSD"},
// 	{"SOL-TUSD", "SOL-BNB", "BNB-TUSD"},
// 	{"SSV-TUSD", "SSV-BTC", "BTC-TUSD"},
// 	{"SSV-TUSD", "SSV-ETH", "ETH-TUSD"},
// 	{"RDNT-TUSD", "RDNT-BTC", "BTC-TUSD"},
// 	{"DOGE-TUSD", "DOGE-BTC", "BTC-TUSD"},
// 	{"EDU-TUSD", "EDU-BTC", "BTC-TUSD"},
// 	{"EDU-TUSD", "EDU-BNB", "BNB-TUSD"},
// 	{"SUI-TUSD", "SUI-BTC", "BTC-TUSD"},
// 	{"SUI-TUSD", "SUI-BNB", "BNB-TUSD"},
// 	{"MAV-TUSD", "MAV-BTC", "BTC-TUSD"},
// 	{"CFX-TUSD", "CFX-BTC", "BTC-TUSD"},
// 	{"PENDLE-TUSD", "PENDLE-BTC", "BTC-TUSD"},
// 	{"XVG-TUSD", "XVG-ETH", "ETH-TUSD"},
// 	{"ARKM-TUSD", "ARKM-BTC", "BTC-TUSD"},
// 	{"ARKM-TUSD", "ARKM-BNB", "BNB-TUSD"},
// 	{"AVAX-TUSD", "AVAX-BTC", "BTC-TUSD"},
// 	{"AVAX-TUSD", "AVAX-ETH", "ETH-TUSD"},
// 	{"AVAX-TUSD", "AVAX-BNB", "BNB-TUSD"},
// 	{"COMP-TUSD", "COMP-BTC", "BTC-TUSD"},
// 	{"QUICK-TUSD", "QUICK-BTC", "BTC-TUSD"},
// 	{"CYBER-TUSD", "CYBER-BTC", "BTC-TUSD"},
// 	{"CYBER-TUSD", "CYBER-ETH", "ETH-TUSD"},
// 	{"CYBER-TUSD", "CYBER-BNB", "BNB-TUSD"},
// 	{"SEI-TUSD", "SEI-BTC", "BTC-TUSD"},
// 	{"SEI-TUSD", "SEI-BNB", "BNB-TUSD"},
// 	{"FRONT-TUSD", "FRONT-BTC", "BTC-TUSD"},
// 	{"RUNE-TUSD", "RUNE-BTC", "BTC-TUSD"},
// 	{"RUNE-TUSD", "RUNE-ETH", "ETH-TUSD"},
// 	{"RUNE-TUSD", "RUNE-BNB", "BNB-TUSD"},
// 	{"MEME-TUSD", "MEME-ETH", "ETH-TUSD"},
// 	{"MEME-TUSD", "MEME-BNB", "BNB-TUSD"},
// 	{"INJ-TUSD", "INJ-BTC", "BTC-TUSD"},
// 	{"INJ-TUSD", "INJ-ETH", "ETH-TUSD"},
// 	{"INJ-TUSD", "INJ-BNB", "BNB-TUSD"},
// 	{"ORDI-TUSD", "ORDI-BTC", "BTC-TUSD"},
// 	{"SHIB-TUSD", "SHIB-DOGE", "DOGE-TUSD"},
// 	{"CAKE-TUSD", "CAKE-BTC", "BTC-TUSD"},
// 	{"CAKE-TUSD", "CAKE-BNB", "BNB-TUSD"},
// 	{"TIA-TUSD", "TIA-BTC", "BTC-TUSD"},
// 	{"DOT-TUSD", "DOT-BTC", "BTC-TUSD"},
// 	{"DOT-TUSD", "DOT-ETH", "ETH-TUSD"},
// 	{"DOT-TUSD", "DOT-BNB", "BNB-TUSD"},
// 	{"NFP-TUSD", "NFP-BTC", "BTC-TUSD"},
// 	{"NFP-TUSD", "NFP-BNB", "BNB-TUSD"},
// 	{"AI-TUSD", "AI-BTC", "BTC-TUSD"},
// 	{"AI-TUSD", "AI-BNB", "BNB-TUSD"},
// }

// var ArbitrageTriplets = []ArbitrageTriplet{
// 	{"BNB-USDC", "BNB-BTC", "BTC-USDC"},
// 	{"BNB-USDC", "BNB-ETH", "ETH-USDC"},
// 	{"ETH-USDC", "ETH-BTC", "BTC-USDC"},
// 	{"XRP-USDC", "XRP-BNB", "BNB-USDC"},
// 	{"XRP-USDC", "XRP-BTC", "BTC-USDC"},
// 	{"XRP-USDC", "XRP-ETH", "ETH-USDC"},
// 	{"ADA-USDC", "ADA-BNB", "BNB-USDC"},
// 	{"ADA-USDC", "ADA-BTC", "BTC-USDC"},
// 	{"ADA-USDC", "ADA-ETH", "ETH-USDC"},
// 	{"SOL-USDC", "SOL-BNB", "BNB-USDC"},
// 	{"SOL-USDC", "SOL-BTC", "BTC-USDC"},
// 	{"SOL-USDC", "SOL-ETH", "ETH-USDC"},
// 	{"ARB-USDC", "ARB-BTC", "BTC-USDC"},
// 	{"ARB-USDC", "ARB-ETH", "ETH-USDC"},
// 	{"AVAX-USDC", "AVAX-BNB", "BNB-USDC"},
// 	{"AVAX-USDC", "AVAX-BTC", "BTC-USDC"},
// 	{"AVAX-USDC", "AVAX-ETH", "ETH-USDC"},
// 	{"DOT-USDC", "DOT-BNB", "BNB-USDC"},
// 	{"DOT-USDC", "DOT-BTC", "BTC-USDC"},
// 	{"DOT-USDC", "DOT-ETH", "ETH-USDC"},
// 	{"INJ-USDC", "INJ-BNB", "BNB-USDC"},
// 	{"INJ-USDC", "INJ-BTC", "BTC-USDC"},
// 	{"INJ-USDC", "INJ-ETH", "ETH-USDC"},
// 	{"MATIC-USDC", "MATIC-BNB", "BNB-USDC"},
// 	{"MATIC-USDC", "MATIC-BTC", "BTC-USDC"},
// 	{"MATIC-USDC", "MATIC-ETH", "ETH-USDC"},
// 	{"OP-USDC", "OP-BNB", "BNB-USDC"},
// 	{"OP-USDC", "OP-BTC", "BTC-USDC"},
// 	{"OP-USDC", "OP-ETH", "ETH-USDC"},
// 	{"ORDI-USDC", "ORDI-BTC", "BTC-USDC"},
// }

// var ArbitrageTriplets = []ArbitrageTriplet{
// 	{"BTC-EUR", "BTC-EUR", "EUR-AEUR"},
// 	{"ETH-EUR", "ETH-BTC", "BTC-EUR"},
// 	{"ETH-EUR", "ETH-BTC", "BTC-AEUR"},
// 	{"ETH-EUR", "ETH-EUR", "EUR-AEUR"},
// 	{"BNB-EUR", "BNB-BTC", "BTC-EUR"},
// 	{"BNB-EUR", "BNB-ETH", "ETH-EUR"},
// 	{"BNB-EUR", "BNB-BTC", "BTC-AEUR"},
// 	{"BNB-EUR", "BNB-ETH", "ETH-AEUR"},
// 	{"BNB-EUR", "BNB-EUR", "EUR-AEUR"},
// 	{"XRP-EUR", "XRP-BTC", "BTC-EUR"},
// 	{"XRP-EUR", "XRP-ETH", "ETH-EUR"},
// 	{"XRP-EUR", "XRP-BNB", "BNB-EUR"},
// 	{"XRP-EUR", "XRP-BTC", "BTC-AEUR"},
// 	{"XRP-EUR", "XRP-ETH", "ETH-AEUR"},
// 	{"XRP-EUR", "XRP-EUR", "EUR-AEUR"},
// 	{"LINK-EUR", "LINK-BTC", "BTC-EUR"},
// 	{"LINK-EUR", "LINK-ETH", "ETH-EUR"},
// 	{"LINK-EUR", "LINK-BNB", "BNB-EUR"},
// 	{"LINK-EUR", "LINK-BTC", "BTC-AEUR"},
// 	{"LINK-EUR", "LINK-ETH", "ETH-AEUR"},
// 	{"LINK-EUR", "LINK-EUR", "EUR-AEUR"},
// 	{"DOT-EUR", "DOT-BTC", "BTC-EUR"},
// 	{"DOT-EUR", "DOT-ETH", "ETH-EUR"},
// 	{"DOT-EUR", "DOT-BNB", "BNB-EUR"},
// 	{"DOT-EUR", "DOT-BTC", "BTC-AEUR"},
// 	{"DOT-EUR", "DOT-ETH", "ETH-AEUR"},
// 	{"DOT-EUR", "DOT-EUR", "EUR-AEUR"},
// 	{"LTC-EUR", "LTC-BTC", "BTC-EUR"},
// 	{"LTC-EUR", "LTC-ETH", "ETH-EUR"},
// 	{"LTC-EUR", "LTC-BNB", "BNB-EUR"},
// 	{"LTC-EUR", "LTC-BTC", "BTC-AEUR"},
// 	{"LTC-EUR", "LTC-ETH", "ETH-AEUR"},
// 	{"LTC-EUR", "LTC-EUR", "EUR-AEUR"},
// 	{"ADA-EUR", "ADA-BTC", "BTC-EUR"},
// 	{"ADA-EUR", "ADA-ETH", "ETH-EUR"},
// 	{"ADA-EUR", "ADA-BNB", "BNB-EUR"},
// 	{"ADA-EUR", "ADA-BTC", "BTC-AEUR"},
// 	{"ADA-EUR", "ADA-ETH", "ETH-AEUR"},
// 	{"ADA-EUR", "ADA-EUR", "EUR-AEUR"},
// 	{"BCH-EUR", "BCH-BTC", "BTC-EUR"},
// 	{"BCH-EUR", "BCH-BNB", "BNB-EUR"},
// 	{"BCH-EUR", "BCH-BTC", "BTC-AEUR"},
// 	{"BCH-EUR", "BCH-EUR", "EUR-AEUR"},
// 	{"XLM-EUR", "XLM-BTC", "BTC-EUR"},
// 	{"XLM-EUR", "XLM-ETH", "ETH-EUR"},
// 	{"XLM-EUR", "XLM-BNB", "BNB-EUR"},
// 	{"XLM-EUR", "XLM-BTC", "BTC-AEUR"},
// 	{"XLM-EUR", "XLM-ETH", "ETH-AEUR"},
// 	{"XLM-EUR", "XLM-EUR", "EUR-AEUR"},
// 	{"GRT-EUR", "GRT-BTC", "BTC-EUR"},
// 	{"GRT-EUR", "GRT-ETH", "ETH-EUR"},
// 	{"GRT-EUR", "GRT-BTC", "BTC-AEUR"},
// 	{"GRT-EUR", "GRT-ETH", "ETH-AEUR"},
// 	{"GRT-EUR", "GRT-EUR", "EUR-AEUR"},
// 	{"DOGE-EUR", "DOGE-BTC", "BTC-EUR"},
// 	{"DOGE-EUR", "DOGE-BTC", "BTC-AEUR"},
// 	{"DOGE-EUR", "DOGE-EUR", "EUR-AEUR"},
// 	{"EGLD-EUR", "EGLD-BTC", "BTC-EUR"},
// 	{"EGLD-EUR", "EGLD-ETH", "ETH-EUR"},
// 	{"EGLD-EUR", "EGLD-BNB", "BNB-EUR"},
// 	{"EGLD-EUR", "EGLD-BTC", "BTC-AEUR"},
// 	{"EGLD-EUR", "EGLD-ETH", "ETH-AEUR"},
// 	{"EGLD-EUR", "EGLD-EUR", "EUR-AEUR"},
// 	{"AVAX-EUR", "AVAX-BTC", "BTC-EUR"},
// 	{"AVAX-EUR", "AVAX-ETH", "ETH-EUR"},
// 	{"AVAX-EUR", "AVAX-BNB", "BNB-EUR"},
// 	{"AVAX-EUR", "AVAX-BTC", "BTC-AEUR"},
// 	{"AVAX-EUR", "AVAX-ETH", "ETH-AEUR"},
// 	{"AVAX-EUR", "AVAX-EUR", "EUR-AEUR"},
// 	{"CHZ-EUR", "CHZ-BTC", "BTC-EUR"},
// 	{"CHZ-EUR", "CHZ-BNB", "BNB-EUR"},
// 	{"CHZ-EUR", "CHZ-BTC", "BTC-AEUR"},
// 	{"CHZ-EUR", "CHZ-EUR", "EUR-AEUR"},
// 	{"MATIC-EUR", "MATIC-BTC", "BTC-EUR"},
// 	{"MATIC-EUR", "MATIC-ETH", "ETH-EUR"},
// 	{"MATIC-EUR", "MATIC-BNB", "BNB-EUR"},
// 	{"MATIC-EUR", "MATIC-BTC", "BTC-AEUR"},
// 	{"MATIC-EUR", "MATIC-ETH", "ETH-AEUR"},
// 	{"MATIC-EUR", "MATIC-EUR", "EUR-AEUR"},
// 	{"WIN-EUR", "WIN-BNB", "BNB-EUR"},
// 	{"WIN-EUR", "WIN-TRX", "TRX-EUR"},
// 	{"WIN-EUR", "WIN-EUR", "EUR-AEUR"},
// 	{"VET-EUR", "VET-BTC", "BTC-EUR"},
// 	{"VET-EUR", "VET-ETH", "ETH-EUR"},
// 	{"VET-EUR", "VET-BNB", "BNB-EUR"},
// 	{"VET-EUR", "VET-BTC", "BTC-AEUR"},
// 	{"VET-EUR", "VET-ETH", "ETH-AEUR"},
// 	{"VET-EUR", "VET-EUR", "EUR-AEUR"},
// 	{"TRX-EUR", "TRX-BTC", "BTC-EUR"},
// 	{"TRX-EUR", "TRX-ETH", "ETH-EUR"},
// 	{"TRX-EUR", "TRX-BNB", "BNB-EUR"},
// 	{"TRX-EUR", "TRX-XRP", "XRP-EUR"},
// 	{"TRX-EUR", "TRX-BTC", "BTC-AEUR"},
// 	{"TRX-EUR", "TRX-ETH", "ETH-AEUR"},
// 	{"TRX-EUR", "TRX-EUR", "EUR-AEUR"},
// 	{"SHIB-EUR", "SHIB-DOGE", "DOGE-EUR"},
// 	{"SHIB-EUR", "SHIB-EUR", "EUR-AEUR"},
// 	{"SOL-EUR", "SOL-BTC", "BTC-EUR"},
// 	{"SOL-EUR", "SOL-ETH", "ETH-EUR"},
// 	{"SOL-EUR", "SOL-BNB", "BNB-EUR"},
// 	{"SOL-EUR", "SOL-BTC", "BTC-AEUR"},
// 	{"SOL-EUR", "SOL-ETH", "ETH-AEUR"},
// 	{"SOL-EUR", "SOL-EUR", "EUR-AEUR"},
// 	{"ICP-EUR", "ICP-BTC", "BTC-EUR"},
// 	{"ICP-EUR", "ICP-ETH", "ETH-EUR"},
// 	{"ICP-EUR", "ICP-BNB", "BNB-EUR"},
// 	{"ICP-EUR", "ICP-BTC", "BTC-AEUR"},
// 	{"ICP-EUR", "ICP-ETH", "ETH-AEUR"},
// 	{"ICP-EUR", "ICP-EUR", "EUR-AEUR"},
// 	{"RUNE-EUR", "RUNE-BTC", "BTC-EUR"},
// 	{"RUNE-EUR", "RUNE-ETH", "ETH-EUR"},
// 	{"RUNE-EUR", "RUNE-BNB", "BNB-EUR"},
// 	{"RUNE-EUR", "RUNE-BTC", "BTC-AEUR"},
// 	{"RUNE-EUR", "RUNE-ETH", "ETH-AEUR"},
// 	{"RUNE-EUR", "RUNE-EUR", "EUR-AEUR"},
// 	{"LAZIO-EUR", "LAZIO-BTC", "BTC-EUR"},
// 	{"LAZIO-EUR", "LAZIO-BTC", "BTC-AEUR"},
// 	{"LAZIO-EUR", "LAZIO-EUR", "EUR-AEUR"},
// 	{"PORTO-EUR", "PORTO-BTC", "BTC-EUR"},
// 	{"PORTO-EUR", "PORTO-BTC", "BTC-AEUR"},
// 	{"PORTO-EUR", "PORTO-EUR", "EUR-AEUR"},
// 	{"ALPINE-EUR", "ALPINE-BTC", "BTC-EUR"},
// 	{"ALPINE-EUR", "ALPINE-BTC", "BTC-AEUR"},
// 	{"ALPINE-EUR", "ALPINE-EUR", "EUR-AEUR"},
// 	{"ATOM-EUR", "ATOM-BTC", "BTC-EUR"},
// 	{"ATOM-EUR", "ATOM-ETH", "ETH-EUR"},
// 	{"ATOM-EUR", "ATOM-BNB", "BNB-EUR"},
// 	{"ATOM-EUR", "ATOM-BTC", "BTC-AEUR"},
// 	{"ATOM-EUR", "ATOM-ETH", "ETH-AEUR"},
// 	{"ATOM-EUR", "ATOM-EUR", "EUR-AEUR"},
// 	{"GALA-EUR", "GALA-BTC", "BTC-EUR"},
// 	{"GALA-EUR", "GALA-ETH", "ETH-EUR"},
// 	{"GALA-EUR", "GALA-BNB", "BNB-EUR"},
// 	{"GALA-EUR", "GALA-BTC", "BTC-AEUR"},
// 	{"GALA-EUR", "GALA-ETH", "ETH-AEUR"},
// 	{"GALA-EUR", "GALA-EUR", "EUR-AEUR"},
// 	{"NEAR-EUR", "NEAR-BTC", "BTC-EUR"},
// 	{"NEAR-EUR", "NEAR-ETH", "ETH-EUR"},
// 	{"NEAR-EUR", "NEAR-BNB", "BNB-EUR"},
// 	{"NEAR-EUR", "NEAR-BTC", "BTC-AEUR"},
// 	{"NEAR-EUR", "NEAR-ETH", "ETH-AEUR"},
// 	{"NEAR-EUR", "NEAR-EUR", "EUR-AEUR"},
// 	{"GMT-EUR", "GMT-BTC", "BTC-EUR"},
// 	{"GMT-EUR", "GMT-ETH", "ETH-EUR"},
// 	{"GMT-EUR", "GMT-BNB", "BNB-EUR"},
// 	{"GMT-EUR", "GMT-BTC", "BTC-AEUR"},
// 	{"GMT-EUR", "GMT-ETH", "ETH-AEUR"},
// 	{"GMT-EUR", "GMT-EUR", "EUR-AEUR"},
// 	{"FTM-EUR", "FTM-BTC", "BTC-EUR"},
// 	{"FTM-EUR", "FTM-ETH", "ETH-EUR"},
// 	{"FTM-EUR", "FTM-BNB", "BNB-EUR"},
// 	{"FTM-EUR", "FTM-BTC", "BTC-AEUR"},
// 	{"FTM-EUR", "FTM-ETH", "ETH-AEUR"},
// 	{"FTM-EUR", "FTM-EUR", "EUR-AEUR"},
// 	{"DAR-EUR", "DAR-BTC", "BTC-EUR"},
// 	{"DAR-EUR", "DAR-BNB", "BNB-EUR"},
// 	{"DAR-EUR", "DAR-BTC", "BTC-AEUR"},
// 	{"DAR-EUR", "DAR-EUR", "EUR-AEUR"},
// 	{"OP-EUR", "OP-BTC", "BTC-EUR"},
// 	{"OP-EUR", "OP-ETH", "ETH-EUR"},
// 	{"OP-EUR", "OP-BNB", "BNB-EUR"},
// 	{"OP-EUR", "OP-BTC", "BTC-AEUR"},
// 	{"OP-EUR", "OP-ETH", "ETH-AEUR"},
// 	{"OP-EUR", "OP-EUR", "EUR-AEUR"},
// 	{"APT-EUR", "APT-BTC", "BTC-EUR"},
// 	{"APT-EUR", "APT-ETH", "ETH-EUR"},
// 	{"APT-EUR", "APT-BTC", "BTC-AEUR"},
// 	{"APT-EUR", "APT-ETH", "ETH-AEUR"},
// 	{"APT-EUR", "APT-EUR", "EUR-AEUR"},
// 	{"ARB-EUR", "ARB-BTC", "BTC-EUR"},
// 	{"ARB-EUR", "ARB-ETH", "ETH-EUR"},
// 	{"ARB-EUR", "ARB-BTC", "BTC-AEUR"},
// 	{"ARB-EUR", "ARB-ETH", "ETH-AEUR"},
// 	{"ARB-EUR", "ARB-EUR", "EUR-AEUR"},
// 	{"SUI-EUR", "SUI-BTC", "BTC-EUR"},
// 	{"SUI-EUR", "SUI-BNB", "BNB-EUR"},
// 	{"SUI-EUR", "SUI-BTC", "BTC-AEUR"},
// 	{"SUI-EUR", "SUI-EUR", "EUR-AEUR"},
// 	{"BTC-AEUR", "BTC-EUR", "EUR-AEUR"},
// 	{"ETH-AEUR", "ETH-BTC", "BTC-EUR"},
// 	{"ETH-AEUR", "ETH-BTC", "BTC-AEUR"},
// 	{"ETH-AEUR", "ETH-EUR", "EUR-AEUR"},
// }
