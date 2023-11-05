package codesdk

import (
	"slices"
)

type CurrencyCode string

const (
	KIN CurrencyCode = "kin"
	AED CurrencyCode = "aed"
	AFN CurrencyCode = "afn"
	ALL CurrencyCode = "all"
	AMD CurrencyCode = "amd"
	ANG CurrencyCode = "ang"
	AOA CurrencyCode = "aoa"
	ARS CurrencyCode = "ars"
	AUD CurrencyCode = "aud"
	AWG CurrencyCode = "awg"
	AZN CurrencyCode = "azn"
	BAM CurrencyCode = "bam"
	BBD CurrencyCode = "bbd"
	BDT CurrencyCode = "bdt"
	BGN CurrencyCode = "bgn"
	BHD CurrencyCode = "bhd"
	BIF CurrencyCode = "bif"
	BMD CurrencyCode = "bmd"
	BND CurrencyCode = "bnd"
	BOB CurrencyCode = "bob"
	BRL CurrencyCode = "brl"
	BSD CurrencyCode = "bsd"
	BTN CurrencyCode = "btn"
	BWP CurrencyCode = "bwp"
	BYN CurrencyCode = "byn"
	BZD CurrencyCode = "bzd"
	CAD CurrencyCode = "cad"
	CDF CurrencyCode = "cdf"
	CHF CurrencyCode = "chf"
	CLP CurrencyCode = "clp"
	CNY CurrencyCode = "cny"
	COP CurrencyCode = "cop"
	CRC CurrencyCode = "crc"
	CUP CurrencyCode = "cup"
	CVE CurrencyCode = "cve"
	CZK CurrencyCode = "czk"
	DJF CurrencyCode = "djf"
	DKK CurrencyCode = "dkk"
	DOP CurrencyCode = "dop"
	DZD CurrencyCode = "dzd"
	EGP CurrencyCode = "egp"
	ERN CurrencyCode = "ern"
	ETB CurrencyCode = "etb"
	EUR CurrencyCode = "eur"
	FJD CurrencyCode = "fjd"
	FKP CurrencyCode = "fkp"
	GBP CurrencyCode = "gbp"
	GEL CurrencyCode = "gel"
	GHS CurrencyCode = "ghs"
	GIP CurrencyCode = "gip"
	GMD CurrencyCode = "gmd"
	GNF CurrencyCode = "gnf"
	GTQ CurrencyCode = "gtq"
	GYD CurrencyCode = "gyd"
	HKD CurrencyCode = "hkd"
	HNL CurrencyCode = "hnl"
	HRK CurrencyCode = "hrk"
	HTG CurrencyCode = "htg"
	HUF CurrencyCode = "huf"
	IDR CurrencyCode = "idr"
	ILS CurrencyCode = "ils"
	INR CurrencyCode = "inr"
	IQD CurrencyCode = "iqd"
	IRR CurrencyCode = "irr"
	ISK CurrencyCode = "isk"
	JMD CurrencyCode = "jmd"
	JOD CurrencyCode = "jod"
	JPY CurrencyCode = "jpy"
	KES CurrencyCode = "kes"
	KGS CurrencyCode = "kgs"
	KHR CurrencyCode = "khr"
	KMF CurrencyCode = "kmf"
	KPW CurrencyCode = "kpw"
	KRW CurrencyCode = "krw"
	KWD CurrencyCode = "kwd"
	KYD CurrencyCode = "kyd"
	KZT CurrencyCode = "kzt"
	LAK CurrencyCode = "lak"
	LBP CurrencyCode = "lbp"
	LKR CurrencyCode = "lkr"
	LRD CurrencyCode = "lrd"
	LYD CurrencyCode = "lyd"
	MAD CurrencyCode = "mad"
	MDL CurrencyCode = "mdl"
	MGA CurrencyCode = "mga"
	MKD CurrencyCode = "mkd"
	MMK CurrencyCode = "mmk"
	MNT CurrencyCode = "mnt"
	MOP CurrencyCode = "mop"
	MRU CurrencyCode = "mru"
	MUR CurrencyCode = "mur"
	MVR CurrencyCode = "mvr"
	MWK CurrencyCode = "mwk"
	MXN CurrencyCode = "mxn"
	MYR CurrencyCode = "myr"
	MZN CurrencyCode = "mzn"
	NAD CurrencyCode = "nad"
	NGN CurrencyCode = "ngn"
	NIO CurrencyCode = "nio"
	NOK CurrencyCode = "nok"
	NPR CurrencyCode = "npr"
	NZD CurrencyCode = "nzd"
	OMR CurrencyCode = "omr"
	PAB CurrencyCode = "pab"
	PEN CurrencyCode = "pen"
	PGK CurrencyCode = "pgk"
	PHP CurrencyCode = "php"
	PKR CurrencyCode = "pkr"
	PLN CurrencyCode = "pln"
	PYG CurrencyCode = "pyg"
	QAR CurrencyCode = "qar"
	RON CurrencyCode = "ron"
	RSD CurrencyCode = "rsd"
	RUB CurrencyCode = "rub"
	RWF CurrencyCode = "rwf"
	SAR CurrencyCode = "sar"
	SBD CurrencyCode = "sbd"
	SCR CurrencyCode = "scr"
	SDG CurrencyCode = "sdg"
	SEK CurrencyCode = "sek"
	SGD CurrencyCode = "sgd"
	SHP CurrencyCode = "shp"
	SLL CurrencyCode = "sll"
	SOS CurrencyCode = "sos"
	SRD CurrencyCode = "srd"
	SSP CurrencyCode = "ssp"
	STN CurrencyCode = "stn"
	SYP CurrencyCode = "syp"
	SZL CurrencyCode = "szl"
	THB CurrencyCode = "thb"
	TJS CurrencyCode = "tjs"
	TMT CurrencyCode = "tmt"
	TND CurrencyCode = "tnd"
	TOP CurrencyCode = "top"
	TRY CurrencyCode = "try"
	TTD CurrencyCode = "ttd"
	TWD CurrencyCode = "twd"
	TZS CurrencyCode = "tzs"
	UAH CurrencyCode = "uah"
	UGX CurrencyCode = "ugx"
	USD CurrencyCode = "usd"
	UYU CurrencyCode = "uyu"
	UZS CurrencyCode = "uzs"
	VES CurrencyCode = "ves"
	VND CurrencyCode = "vnd"
	VUV CurrencyCode = "vuv"
	WST CurrencyCode = "wst"
	XAF CurrencyCode = "xaf"
	XCD CurrencyCode = "xcd"
	XOF CurrencyCode = "xof"
	XPF CurrencyCode = "xpf"
	YER CurrencyCode = "yer"
	ZAR CurrencyCode = "zar"
	ZMW CurrencyCode = "zmw"
)

var (
	currenciesAtIndex = []CurrencyCode{
		KIN,
		AED,
		AFN,
		ALL,
		AMD,
		ANG,
		AOA,
		ARS,
		AUD,
		AWG,
		AZN,
		BAM,
		BBD,
		BDT,
		BGN,
		BHD,
		BIF,
		BMD,
		BND,
		BOB,
		BRL,
		BSD,
		BTN,
		BWP,
		BYN,
		BZD,
		CAD,
		CDF,
		CHF,
		CLP,
		CNY,
		COP,
		CRC,
		CUP,
		CVE,
		CZK,
		DJF,
		DKK,
		DOP,
		DZD,
		EGP,
		ERN,
		ETB,
		EUR,
		FJD,
		FKP,
		GBP,
		GEL,
		GHS,
		GIP,
		GMD,
		GNF,
		GTQ,
		GYD,
		HKD,
		HNL,
		HRK,
		HTG,
		HUF,
		IDR,
		ILS,
		INR,
		IQD,
		IRR,
		ISK,
		JMD,
		JOD,
		JPY,
		KES,
		KGS,
		KHR,
		KMF,
		KPW,
		KRW,
		KWD,
		KYD,
		KZT,
		LAK,
		LBP,
		LKR,
		LRD,
		LYD,
		MAD,
		MDL,
		MGA,
		MKD,
		MMK,
		MNT,
		MOP,
		MRU,
		MUR,
		MVR,
		MWK,
		MXN,
		MYR,
		MZN,
		NAD,
		NGN,
		NIO,
		NOK,
		NPR,
		NZD,
		OMR,
		PAB,
		PEN,
		PGK,
		PHP,
		PKR,
		PLN,
		PYG,
		QAR,
		RON,
		RSD,
		RUB,
		RWF,
		SAR,
		SBD,
		SCR,
		SDG,
		SEK,
		SGD,
		SHP,
		SLL,
		SOS,
		SRD,
		SSP,
		STN,
		SYP,
		SZL,
		THB,
		TJS,
		TMT,
		TND,
		TOP,
		TRY,
		TTD,
		TWD,
		TZS,
		UAH,
		UGX,
		USD,
		UYU,
		UZS,
		VES,
		VND,
		VUV,
		WST,
		XAF,
		XCD,
		XOF,
		XPF,
		YER,
		ZAR,
		ZMW,
	}
)

func IsValidCurrency(currency string) bool {
	_, err := CurrencyCode(currency).toIndex()
	return err == nil
}

func (c CurrencyCode) toIndex() (int, error) {
	index := slices.Index(currenciesAtIndex, c)
	if index < 0 {
		return -1, ErrInvalidCurrency
	}
	return index, nil
}
