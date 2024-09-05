package api

type Vin struct {
	Txid         string   `json:"txid"`
	Vout         int      `json:"vout"`
	Prevout      Prevout  `json:"prevout"`
	Scriptsig    string   `json:"scriptsig"`
	ScriptsigAsm string   `json:"scriptsig_asm"`
	Witness      []string `json:"witness"`
	IsCoinbase   bool     `json:"is_coinbase"`
	Sequence     int64    `json:"sequence"`
}

type Prevout struct {
	Scriptpubkey        string `json:"scriptpubkey"`
	ScriptpubkeyAsm     string `json:"scriptpubkey_asm"`
	ScriptpubkeyType    string `json:"scriptpubkey_type"`
	ScriptpubkeyAddress string `json:"scriptpubkey_address"`
	Value               int    `json:"value"`
}
